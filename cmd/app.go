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
	var fileSize int
	var sortDirectories bool
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
			&cli.IntFlag{
				Name:        "size",
				Value:       1024,
				Usage:       "total file size.",
				Aliases:     []string{"s"},
				Destination: &fileSize,
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
			return RunTest(c, simRead, simWrite, fileSize, mediaDirectory)
		},
	}
	return app
}
