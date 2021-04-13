package main

import (
	"fmt"
	"image"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"go.viam.com/robotcore/artifact"
	"go.viam.com/robotcore/rimage"
	"go.viam.com/robotcore/vision/chess"

	"github.com/edaniels/golog"
	"github.com/tonyOreglia/glee/pkg/position"
)

/* TODO(erh): put back
func TestInit(t *testing.T) {
	state := boardStateGuesser{}

	fns, err := filepath.Glob(artifact.MustPath("samples/chess/init/board-*.png"))
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(fns)

	for idx, fn := range fns {
		fmt.Println(fn)
		depthDN := strings.Replace(fn, ".png", ".dat.gz", 1)

		board, err := chess.FindAndWarpBoardFromFiles(fn, depthDN)
		if err != nil {
			t.Fatal(err)
		}

		_, err = state.newData(board)
		if err != nil {
			t.Fatal(err)
		}

		pcs, err := state.game.GetSquaresWithPieces(board)
		if err != nil {
			err2 := board.WriteDebugImages("out/foo")
			if err2 != nil {
				panic(err2)
			}
			t.Fatal(err)
		}
		fmt.Printf("\t%s\n", pcs)
		if len(pcs) != 32 {
			temp := board.Annotate()
			tempfn := fmt.Sprintf("out/init-%d.png", idx)

			utils.WriteImageToFile(tempfn, temp)
			fmt.Printf("\t annotated -> %s\n", tempfn)
		}

		if state.Ready() {
			squares, err := state.GetSquaresWithPieces()
			if err != nil {
				t.Fatal(err)
			}

			if len(squares) != 32 {
				t.Errorf("wrong number of squares %d", len(squares))
			}

			for x := 'a'; x <= 'h'; x++ {
				for _, y := range []string{"1", "2", "7", "8"} {
					sq := string(x) + y
					if !squares[sq] {
						t.Errorf("missing %s", sq)
					}
				}
			}
		}
	}

	p := position.StartingPosition()
	m, err := state.GetPrevMove(p)
	if m != nil {
		t.Errorf("why is there a move!!!")
	}
	if err != nil {
		t.Fatal(err)
	}
}
*/
func TestOneMove(t *testing.T) {
	logger := golog.NewTestLogger(t)
	state := boardStateGuesser{}

	e2e4Path := artifact.MustPath("samples/chess/e2e4")
	fns, err := filepath.Glob(e2e4Path + "/board-*.png")
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(fns)

	for idx, fn := range fns {
		fmt.Println(fn)
		depthDN := strings.Replace(fn, ".png", ".dat.gz", 1)

		board, err := chess.FindAndWarpBoardFromFiles(fn, depthDN, logger)
		if err != nil {
			t.Fatal(err)
		}

		_, err = state.newData(board)
		if err != nil {
			t.Fatal(err)
		}

		board.WriteDebugImages(fmt.Sprintf("out/onemove-%d", idx))
	}

	bb, err := state.GetBitBoard()
	if err != nil {
		t.Fatal(err)
	}

	if bb.Value() != 18441959067825012735 {
		t.Errorf("TestOneMove initial value wrong %d", bb.Value())
	}

	p := position.StartingPosition()
	m, err := state.GetPrevMove(p)
	if m == nil {
		t.Errorf("why is there not a move!!!")
	}
	if err != nil {
		t.Fatal(err)
	}
	if m.String() != "e2e4" {
		t.Errorf("move is wrong: %s", m.String())
	}

}

func TestWristDepth1(t *testing.T) {
	dm, err := rimage.ParseDepthMap(artifact.MustPath("samples/chess/wristdepth1.dat.gz"))
	if err != nil {
		t.Fatal(err)
	}

	ppr := dm.ToPrettyPicture(0, 1000)
	pp := rimage.ConvertImage(ppr)

	center := image.Point{dm.Width() / 2, dm.Height() / 2}
	pp.Circle(center, 30, rimage.Red)

	lowest, lowestValue, highest, highestValue := findDepthPeaks(dm, center, 30)
	fmt.Printf("lowest: %v highest: %v\n", lowest, highest)
	fmt.Printf("lowestValue: %v highestValue: %v\n", lowestValue, highestValue)

	pp.Circle(lowest, 5, rimage.Blue)
	pp.Circle(highest, 5, rimage.Red)

	err = pp.WriteTo("/tmp/x.png")
	if err != nil {
		t.Fatal(err)
	}

}
