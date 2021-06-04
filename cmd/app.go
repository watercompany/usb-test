package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"a"},
				Usage:   "run usb read write tests",
				Action: func(c *cli.Context) error {
					fmt.Println("running test ...") //c.Args().First())
					return RunTest(c)
				},
			},
		},
		Action: func(c *cli.Context) error {
			// cli.ShowSubcommandHelp(c)
			fmt.Println("running test ...") //c.Args().First())
			return RunTest(c)
			// return nil
		},
	}
	return app
}
