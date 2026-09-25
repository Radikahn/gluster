package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Radikahn/gluster/cmd/commandDirective"
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

	// At the moment the order for if a Device or CommanDirective comes first is
	// not determined
	var firstTask, currentDirectiveId = commandDirective.InitFromArray(tasks)
	fmt.Printf("current list id: %s \n", currentDirectiveId)

	if firstTask == nil {
		log.Fatal("No tasks provided")
		showHelpDialog()
	}

	var taskCount int = commandDirective.GetLength(firstTask)
	fmt.Printf("Number of tasks: %d \n", taskCount)

	//Register device
	var currentDevice, _, _ = device.RegisterDevice(
		"mainPC",
		"macos",
		"localhost",
		globalDevices,
	)
	device.ActivateDevice(currentDevice)

	device.AssignDirective(currentDirectiveId, currentDevice)

}

func showHelpDialog() {
	fmt.Println("Welcome to Gluster!")
}
