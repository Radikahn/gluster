package runtime

import (
	"log"
	"os"

	"github.com/Radikahn/gluster/api"
	"github.com/sevlyar/go-daemon"
)

func StartOrchestratorDaemon() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        append([]string{"[gluster-orch-daemon]"}, os.Args[1:]...),
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}

	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("---------------")
	log.Print("[ORCHESTRATOR]: daemon started")

	api.StartOrchestratorAPI(*log.Default())
}

func StartWorkerDaemon() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        append([]string{"[gluster-worker-daemon]"}, os.Args[1:]...),
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}

	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("---------------")
	log.Print("[WORKER]: daemon started")

	api.StartWorkerAPI(*log.Default())
}
