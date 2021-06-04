package main

import (
	"log"
	"os"

	"github.com/watercompany/usb-test/cmd"
)

func main() {
	app := cmd.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
