package cmd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/watercompany/usb-test/utils"
)

var (
	n = 1024
)

func RunTest(ctx *cli.Context) error {

	// if err := utils.CreateFolder(config.TestDir, config.ForceCreate); err != nil {
	// 	return err
	// }

	token := make([]byte, n*1024)
	rand.Read(token)
	// fmt.Println(token)
	lsblkJSON, err := ParseLsblk()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", lsblkJSON)

	return nil
}

// {
// 	"blockdevices": [
// 	   {"name":"vda", "maj:min":"254:0", "rm":false, "size":"59.6G", "ro":false, "type":"disk", "mountpoint":null,
// 		  "children": [
// 			 {"name":"vda1", "maj:min":"254:1", "rm":false, "size":"59.6G", "ro":false, "type":"part", "mountpoint":"/etc/hosts"}
// 		  ]
// 	   }
// 	]
//  }

// LSBLK is a wrapper for the lsblk output
type LSBLK struct {
	BlockDevices []Device `json:"blockdevices"`
}

// Device is a wrapper for device in LSBLK
type Device struct {
	Children []Child `json:"children"`
}

// Child is a wrapper for children in Device
type Child struct {
	Name       string  `json:"name"`
	MajMin     string  `json:"maj:min"`
	Rm         bool    `json:"rm"`
	Size       float64 `json:"size"`
	Ro         bool    `json:"ro"`
	Type       string  `json:"type"`
	MountPoint string  `json:"mountpoint"`
}

// ParseLsblk returns json format of `lsblk --json` command
func ParseLsblk() (LSBLK, error) {
	var parsedLSBLK LSBLK
	lsblkOutput, err := utils.RunCMD("lsblk", "--json")
	if err != nil {
		return LSBLK{}, err
	}
	json.Unmarshal([]byte(lsblkOutput), &parsedLSBLK)
	return parsedLSBLK, nil
}
