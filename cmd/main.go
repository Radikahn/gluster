package main

import (
	"fmt"
	"os"

	"github.com/Radikahn/gluster/cmd/device"
)

// GLOBAL DEVICE MAP
var globalDevices map[string]string = make(map[string]string)

func main() {
	// command-line based arugments for testing
	tasks := os.Args[1:]

	//sanity check
	if len(tasks) == 0 {
		showHelpDialog()
		os.Exit(1)
	}
	//Register device
	var currentDevice, _, _ = device.RegisterDevice(
		"mainPC",
		"macos",
		"localhost",
		globalDevices,
	)
	device.ActivateDevice(currentDevice)

}

func showHelpDialog() {
	fmt.Println("Welcome to Gluster!")
}
