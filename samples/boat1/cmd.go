// Package main is the work-in-progress robotic boat from Viam.
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"go.viam.com/core/board"
	"go.viam.com/core/config"
	pb "go.viam.com/core/proto/api/v1"
	"go.viam.com/core/rlog"
	"go.viam.com/core/robot"
	robotimpl "go.viam.com/core/robot/impl"
	"go.viam.com/core/sensor"
	"go.viam.com/core/serial"
	"go.viam.com/core/utils"
	"go.viam.com/core/web"

	_ "go.viam.com/core/board/detector"

	"github.com/adrianmo/go-nmea"
	"github.com/edaniels/golog"
	"go.uber.org/multierr"
)

const (
	millisPerRotation = 200
	maxRPM            = 600.0
)

var logger = golog.NewDevelopmentLogger("boat1")

type Boat struct {
	theBoard        board.Board
	starboard, port board.Motor

	throttle, direction, mode, aSwitch board.DigitalInterrupt
	rightVertical, rightHorizontal     board.DigitalInterrupt
	activeBackgroundWorkers            sync.WaitGroup
}

func (b *Boat) MoveStraight(ctx context.Context, distanceMillis int, millisPerSec float64, block bool) (int, error) {
	dir := pb.DirectionRelative_DIRECTION_RELATIVE_FORWARD
	if distanceMillis < 0 {
		dir = pb.DirectionRelative_DIRECTION_RELATIVE_BACKWARD
		distanceMillis *= -1
	}

	if block {
		return 0, errors.New("boat can't block for move straight yet")
	}

	speed := (millisPerSec * 60.0) / float64(millisPerRotation)
	rotations := float64(distanceMillis) / millisPerRotation

	// TODO(erh): return how much it actually moved
	return distanceMillis, multierr.Combine(
		b.starboard.GoFor(ctx, dir, speed, rotations),
		b.port.GoFor(ctx, dir, speed, rotations),
	)

}

func (b *Boat) Spin(ctx context.Context, angleDeg float64, degsPerSec float64, block bool) (float64, error) {
	return math.NaN(), errors.New("boat can't spin yet")
}

func (b *Boat) WidthMillis(ctx context.Context) (int, error) {
	return 1, nil
}

func (b *Boat) Stop(ctx context.Context) error {
	return multierr.Combine(b.starboard.Off(ctx), b.port.Off(ctx))
}

func (b *Boat) Close() error {
	defer b.activeBackgroundWorkers.Wait()
	return b.Stop(context.Background())
}

func (b *Boat) StartRC(ctx context.Context) {
	b.activeBackgroundWorkers.Add(1)
	utils.ManagedGo(func() {
		for {
			if !utils.SelectContextOrWait(ctx, 10*time.Millisecond) {
				return
			}

			mode := b.mode.Value()
			if mode == 0 {
				continue
			}

			var port, starboard float64

			var portDirection = pb.DirectionRelative_DIRECTION_RELATIVE_FORWARD
			var starboardDirection = pb.DirectionRelative_DIRECTION_RELATIVE_FORWARD

			direction := 0.0

			if b.aSwitch.Value() >= 1600 {
				port = maxRPM * float64(b.rightVertical.Value()) / 100.0
				starboard = port

				if math.Abs(port) < 10 {
					// either not moving or spin mode
					port = maxRPM * float64(b.rightHorizontal.Value()) / 100.0
					starboard = -1 * port
				} else {
					// moving mostly forward or back, but turning a bit
					direction = float64(b.rightHorizontal.Value())
				}

				if port < 0 {
					portDirection = board.FlipDirection(portDirection)
					port = -1 * port
				}
				if starboard < 0 {
					starboardDirection = board.FlipDirection(starboardDirection)
					starboard = -1 * starboard
				}

			} else {
				if mode == 2 {
					portDirection = board.FlipDirection(portDirection)
					starboardDirection = board.FlipDirection(starboardDirection)
				}

				port = maxRPM * (float64(b.throttle.Value()) / 90)
				starboard = port

				direction = float64(b.direction.Value())

			}

			if direction > 0 {
				// we want to turn towards starboard
				// so we slow down the starboard motor
				starboard *= 1 - (direction / 100.0)
			} else if direction < 0 {
				port *= 1 - (direction / -100.0)
			}

			var err error

			if port < 8 && starboard < 8 {
				err = b.Stop(ctx)
			} else {
				err = multierr.Combine(
					b.starboard.GoFor(ctx, starboardDirection, starboard, 0),
					b.port.GoFor(ctx, portDirection, port, 0),
				)
			}

			if err != nil {
				log.Print(err)
			}

		}
	}, b.activeBackgroundWorkers.Done)
}

type SavedDetph struct {
	Longitude float64
	Latitude  float64
	Depth     float64
	Extra     interface{}
}

func storeAll(docs []SavedDetph) error {
	for _, doc := range docs {
		data, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		_, err = http.Post(
			"https://us-east-1.aws.webhooks.mongodb-realm.com/api/client/v2.0/app/boat1-lwcji/service/http1/incoming_webhook/depthRecord",
			"application/json",
			bytes.NewReader(data))
		if err != nil {
			return err
		}
	}

	return nil
}

var currentLocation nmea.GLL

func trackGPS() {
	dev, err := serial.Open("/dev/ttyAMA1")
	if err != nil {
		rlog.Logger.Fatalf("canot open gps serial %s", err)
	}
	defer dev.Close()

	r := bufio.NewReader(dev)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			rlog.Logger.Fatalf("can't read gps serial %s", err)
		}

		s, err := nmea.Parse(line)
		if err != nil {
			rlog.Logger.Debugf("can't parse nmea %s : %s", line, err)
			continue
		}

		gll, ok := s.(nmea.GLL)
		if ok {
			currentLocation = gll
		}
	}
}

var toStore []SavedDetph

func doRecordDepth(ctx context.Context, depthSensor sensor.Sensor) error {
	if currentLocation.Longitude == 0 {
		return errors.New("currentLocation is 0")
	}

	readings, err := depthSensor.Readings(ctx)
	if err != nil {
		return err
	}
	if len(readings) != 1 {
		return fmt.Errorf("readings is unexpected %v", readings)
	}

	m := readings[0].(map[string]interface{})

	confidence := m["confidence"].(float64)
	depth := m["distance"].(float64)

	if confidence < 90 {
		rlog.Logger.Debugf("confidence too low, skipping confidence: %v depth: %v", confidence, depth)
		return nil
	}

	d := SavedDetph{currentLocation.Longitude, currentLocation.Latitude, depth, m}

	toStore = append(toStore, d)

	err = storeAll(toStore)
	if err == nil {
		toStore = []SavedDetph{}
	}
	return err
}

func recordDepthWorker(ctx context.Context, depthSensor sensor.Sensor) {
	if depthSensor == nil {
		rlog.Logger.Fatalf("depthSensor cannot be nil")
	}

	for {
		if !utils.SelectContextOrWait(ctx, 5*time.Second) {
			return
		}

		err := doRecordDepth(ctx, depthSensor)
		if err != nil {
			rlog.Logger.Debugf("erorr recording depth %s", err)
		}
	}
}

func NewBoat(r robot.Robot) (*Boat, error) {
	b := &Boat{}
	b.theBoard = r.BoardByName("local")
	if b.theBoard == nil {
		return nil, errors.New("cannot find board")
	}

	b.starboard = b.theBoard.Motor("starboard")
	b.port = b.theBoard.Motor("port")

	if b.starboard == nil || b.port == nil {
		return nil, errors.New("need a starboard and port motor")
	}

	b.throttle = b.theBoard.DigitalInterrupt("throttle")
	b.direction = b.theBoard.DigitalInterrupt("direction")
	b.mode = b.theBoard.DigitalInterrupt("mode")
	b.aSwitch = b.theBoard.DigitalInterrupt("a")
	b.rightHorizontal = b.theBoard.DigitalInterrupt("right-horizontal")
	b.rightVertical = b.theBoard.DigitalInterrupt("right-vertical")

	if b.throttle == nil || b.direction == nil || b.mode == nil {
		return nil, errors.New("need a throttle and direction and mode")
	}

	return b, nil
}

func main() {
	utils.ContextualMain(mainWithArgs, logger)
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) (err error) {
	flag.Parse()

	cfg, err := config.Read(flag.Arg(0))
	if err != nil {
		return err
	}

	myRobot, err := robotimpl.New(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer myRobot.Close()

	boat, err := NewBoat(myRobot)
	if err != nil {
		return err
	}
	defer boat.Close()
	boat.StartRC(ctx)

	myRobot.AddBase(boat, config.Component{Name: "boatbot"})

	var activeBackgroundWorkers sync.WaitGroup
	activeBackgroundWorkers.Add(2)
	defer activeBackgroundWorkers.Wait()
	utils.ManagedGo(func() {
		trackGPS()
	}, activeBackgroundWorkers.Done)
	utils.ManagedGo(func() {
		recordDepthWorker(ctx, myRobot.SensorByName("depth1"))
	}, activeBackgroundWorkers.Done)

	return web.RunWeb(ctx, myRobot, web.NewOptions(), logger)
}
