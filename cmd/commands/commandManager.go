package commands

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// Run command of given string
// Meant to be used for a CommandDirective.arg
func CommandExec(command string) (int, error) {
	app, commandArray, commandParseError := parseCommand(command)

	if commandParseError != nil {
		log.Fatalf("Command: %s failed to get parsed:\n%s", command, commandParseError)
		return 1, commandParseError
	}

	//exec command
	cmd := exec.Command(app, commandArray...)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Printf("Exit Status: %d", exiterr.ExitCode())
			// if there is error -> return actual cmd exit code
			return exiterr.ExitCode(), err
		} else {
			log.Fatalf("cmd.Wait: %v", err)
			return 1, err
		}
	}

	fmt.Printf("STDOUT: %s", out.String())
	//placeholder exit 0 since we cannot capture exit code on success
	return 0, nil

}

// Uses regex to split command by space char " "
// Returns the app and the list of args following
//
// The use of multiple apps in a command is experimental
// The use of && is experimental and not tested fully
func parseCommand(command string) (string, []string, error) {
	if len(command) <= 0 {
		log.Fatal("Cannot run command, given string argument is empty (length 0)")
		return "", []string{}, fmt.Errorf("String argument for command is empty")
	}

	if strings.Contains(command, "|") {
		return "", []string{}, fmt.Errorf("Pipe arguments are not supported in gluster")
	}

	if strings.Contains(command, ">") || strings.Contains(command, ">>") {
		return "", []string{}, fmt.Errorf("write to and append functions not supported in gluster")
	}

	if strings.Contains(command, "&&") {
		return "", []string{}, fmt.Errorf("AND argument not supported in Gluster")
	}

	commandArray := regexp.MustCompile(" ").Split(command, -1)

	return commandArray[0], commandArray[1:], nil
}

func evaluateCost() {
	// TODO: add logic to add cost for a given function
	// OR can we find a way to add costs when adding the node itself?
	// Cost table/directory?
}
