package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	// config := config.NewConfig()
	var (
		simRead         int
		simWrite        int
		mediaDirectory  string
		fileSize        int
		sortDirectories bool
		// timeout         float64
		loopCount int
	)
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "sim-r",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous write.",
				Aliases:     []string{"r"},
				Destination: &simRead,
			},
			&cli.IntFlag{
				Name:        "sim-w",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous read.",
				Aliases:     []string{"w"},
				Destination: &simWrite,
			},
			// &cli.Float64Flag{
			// 	Name:        "timeout",
			// 	Value:       3600.0,
			// 	Usage:       "loop timeout.",
			// 	Aliases:     []string{"t"},
			// 	Destination: &timeout,
			// },
			&cli.IntFlag{
				Name:        "size",
				Value:       1024,
				Usage:       "total file size.",
				Aliases:     []string{"s"},
				Destination: &fileSize,
			},
			&cli.IntFlag{
				Name:        "loop-count",
				Value:       1,
				Usage:       "total testing loop count.",
				Aliases:     []string{"l"},
				Destination: &loopCount,
			},
			&cli.StringFlag{
				Name:        "root-dir",
				Value:       "/mnt/",
				Usage:       "media root directory to perform test on.",
				Aliases:     []string{"d"},
				Destination: &mediaDirectory,
			},
			&cli.BoolFlag{
				Name:        "sort-directories",
				Value:       false,
				Usage:       "sort directories by name.",
				Aliases:     []string{"n"},
				Destination: &sortDirectories,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("running test ...")
			return RunTest(c, simRead, simWrite, fileSize, sortDirectories, loopCount, mediaDirectory)
		},
	}
	return app
}
