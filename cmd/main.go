package main

import (
	"flag"
	"fmt"

	"github.com/Radikahn/gluster/cmd/states"
	"github.com/Radikahn/gluster/runtime"
)

// GLOBAL DEVICE MAP
var globalDevices map[string]string = make(map[string]string)

func main() {

	daemonMode := flag.String("daemon", "{empty command}", "What daemon mode to activate")

	flag.Parse()

	mode := *daemonMode

	switch mode {
	case states.DaemonModeName[states.Orchestrator]:
		runtime.StartOrchestratorDaemon()
	case states.DaemonModeName[states.Worker]:
		fmt.Print("Worker")
	default:
		showHelpMenu()

	}
}

func showHelpMenu() {
	fmt.Println("Welcome to Gluster!")
	fmt.Println("To start the Gluster daemon run the following:")
	fmt.Println("\t--daemon=orchestrator (for starting the orchestrator daemon)")
	fmt.Println("\t--daemon=worker (for starting the worker daemon)")
}
