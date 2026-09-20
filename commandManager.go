package main

import (
	"bytes"
	"log"
	"os/exec"
)

//Run a command or something?

func commandExec(command string) int {
	cmd := exec.Command(command)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Printf("Exit Status: %d", exiterr.ExitCode())
			// if there is error -> return actual cmd exit code
			return exiterr.ExitCode()
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}

		//placeholder exit 0 since we cannot capture exit code on success
	}

	return 0

}
