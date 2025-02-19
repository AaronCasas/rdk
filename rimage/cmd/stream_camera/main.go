// Package main streams a specific camera over WebRTC.
package main

import (
	"context"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	"github.com/edaniels/gostream/codec/x264"
	"go.uber.org/multierr"
	"go.viam.com/utils"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/components/camera/videosource"
	"go.viam.com/rdk/config"
)

func main() {
	utils.ContextualMain(mainWithArgs, logger)
}

var (
	defaultPort = 5555
	logger      = golog.NewDebugLogger("stream_camera")
)

// Arguments for the command.
type Arguments struct {
	Port        utils.NetPortFlag `flag:"0"`
	Debug       bool              `flag:"debug"`
	Dump        bool              `flag:"dump,usage=dump all camera info"`
	Format      string            `flag:"format"`
	Path        string            `flag:"path"`
	PathPattern string            `flag:"pathPattern"`
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) error {
	// both argesParsed and argsMap are similar, and should at some point be merged or refactored
	var argsParsed Arguments
	var argsMap videosource.WebcamAttrs
	if err := utils.ParseFlags(args, &argsParsed); err != nil {
		return err
	}
	if argsParsed.Port == 0 {
		argsParsed.Port = utils.NetPortFlag(defaultPort)
	}

	if argsParsed.Dump {
		all := gostream.QueryVideoDevices()
		for _, info := range all {
			logger.Debugf("%s", info.ID)
			logger.Debugf("\t labels: %v", info.Labels)
			for _, p := range info.Properties {
				logger.Debugf("\t %v %d x %d", p.FrameFormat, p.Width, p.Height)
			}
		}
		return nil
	}

	attrs := config.AttributeMap{}

	if argsParsed.Format != "" {
		attrs["format"] = argsParsed.Format
		argsMap.Format = argsParsed.Format
	}

	if argsParsed.Path != "" {
		attrs["path"] = argsParsed.Path
		argsMap.Path = argsParsed.Path
	}

	if argsParsed.PathPattern != "" {
		attrs["path_pattern"] = argsParsed.PathPattern
		argsMap.Format = argsParsed.PathPattern
	}

	if argsParsed.Debug {
		attrs["debug"] = true
		argsMap.Debug = true
	}

	if argsParsed.Debug {
		logger.Debugf("attrs: %v", attrs)
	}

	return viewCamera(ctx, argsMap, int(argsParsed.Port), argsParsed.Debug, logger)
}

func viewCamera(ctx context.Context, attrs videosource.WebcamAttrs, port int, debug bool, logger golog.Logger) error {
	webcam, err := videosource.NewWebcamSource(ctx, "camera", &attrs, logger)
	if err != nil {
		return err
	}

	if err := func() error {
		img, closer, err := camera.ReadImage(ctx, webcam)
		if err != nil {
			return err
		}
		defer closer()
		if debug {
			logger.Debugf("image type: %T dimensions: %v", img, img.Bounds())
		}
		return nil
	}(); err != nil {
		return err
	}

	remoteStream, err := gostream.NewStream(x264.DefaultStreamConfig)
	if err != nil {
		return err
	}

	server, err := gostream.NewStandaloneStreamServer(port, logger, nil, remoteStream)
	if err != nil {
		return err
	}
	if err := server.Start(ctx); err != nil {
		return err
	}

	utils.ContextMainReadyFunc(ctx)()
	return multierr.Combine(gostream.StreamVideoSource(ctx, webcam, remoteStream), server.Stop(ctx))
}
