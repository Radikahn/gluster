package commands

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// Run command of given string
// Meant to be used for a CommandDirective.arg
func commandExec(command string) int {
	cmd := exec.Command(command)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Printf("Exit Status: %d", exiterr.ExitCode())
			// if there is error -> return actual cmd exit code
			return exiterr.ExitCode()
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
	}

	fmt.Printf("STDOUT: %s", out.String())
	//placeholder exit 0 since we cannot capture exit code on success
	return 0

}

func evaluateCost() {
	// TODO: add logic to add cost for a given function
	// OR can we find a way to add costs when adding the node itself?
	// Cost table/directory?
}
