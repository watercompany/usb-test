package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/watercompany/usb-test/utils"
)

func RunTest(ctx *cli.Context) error {

	// create a folder - name TestDir
	testDir := "TestDir"
	forceCreate := true
	if err := utils.CreateFolder(testDir, forceCreate); err != nil {
		return err
	}

	return nil
}
