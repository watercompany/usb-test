package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	var simRead int
	var simWrite int
	var mediaDirectory string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "sim-r",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous write.",
				Aliases:     []string{"r"},
				Destination: &simWrite,
			},
			&cli.IntFlag{
				Name:        "sim-w",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous read.",
				Aliases:     []string{"w"},
				Destination: &simRead,
			},
			&cli.StringFlag{
				Name:        "root-dir",
				Value:       "/mnt/",
				Usage:       "media root directory to perform test on.",
				Aliases:     []string{"d"},
				Destination: &mediaDirectory,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("running test ...")
			return RunTest(c, simRead, simWrite, mediaDirectory)
		},
	}
	return app
}
