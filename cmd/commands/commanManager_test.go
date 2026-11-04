package commands

import (
	"testing"
)

func TestCommandExec(t *testing.T) {
	command := `go help`
	want := 0

	exitCode, err := CommandExec(command)

	if want != exitCode {
		t.Errorf(`commandExec("ls") = %d, %v, want match for %d, nil`, exitCode, err, want)
	}
}
