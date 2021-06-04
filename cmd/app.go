package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	var simReadWrite int
	var mediaDirectory string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "sim-rw",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous read and write.",
				Aliases:     []string{"s"},
				Destination: &simReadWrite,
			},
			&cli.StringFlag{
				Name:        "root-dir",
				Value:       "/mnt/",
				Usage:       "media root directory to perform test on.",
				Aliases:     []string{"r"},
				Destination: &mediaDirectory,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("running test ...")
			return RunTest(c, simReadWrite, mediaDirectory)
		},
	}
	return app
}
