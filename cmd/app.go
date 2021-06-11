package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

// RunTestFlags is a configuration struct for RunTest command
type RunTestFlags struct {
	NumSimRead      int
	NumSimWrite     int
	FileSize        int
	SortDirectories bool
	Timeout         float64
	LoopCount       int
	MediaDirectory  string
}

func NewConfig() *RunTestFlags {
	return &RunTestFlags{}
}

func NewApp() *cli.App {
	config := NewConfig()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "sim-r",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous write.",
				Aliases:     []string{"r"},
				Destination: &config.NumSimRead,
			},
			&cli.IntFlag{
				Name:        "sim-w",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous read.",
				Aliases:     []string{"w"},
				Destination: &config.NumSimWrite,
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
				Destination: &config.FileSize,
			},
			&cli.IntFlag{
				Name:        "loop-count",
				Value:       1,
				Usage:       "total testing loop count.",
				Aliases:     []string{"l"},
				Destination: &config.LoopCount,
			},
			&cli.StringFlag{
				Name:        "root-dir",
				Value:       "/mnt/",
				Usage:       "media root directory to perform test on.",
				Aliases:     []string{"d"},
				Destination: &config.MediaDirectory,
			},
			&cli.BoolFlag{
				Name:        "sort-directories",
				Value:       false,
				Usage:       "sort directories by name.",
				Aliases:     []string{"n"},
				Destination: &config.SortDirectories,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("running test ...")
			return RunTest(c, config)
		},
	}
	return app
}
