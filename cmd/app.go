package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	var simReadWrite int
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "lang",
				Value:       runtime.GOMAXPROCS(0),
				Usage:       "number of simultaneous read and write.",
				Destination: &simReadWrite,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("running test ...")
			return RunTest(c, simReadWrite)
		},
	}
	return app
}
