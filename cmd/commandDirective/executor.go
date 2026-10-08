// The executor is the main system manager for the directives
// A directive node has the functionality to act as a node in multiple schemas: queue, linked list, two way linked list
// In the executor you get to choose what data structure you are working with for directives
// Executors can be mixed and matched between devices

package commandDirective

import (
	"github.com/google/uuid"
)

type ExecutorMode int

const (
	PriorityMode ExecutorMode = iota
	QueueMode
	SingleMode
	TwoMode
	StackMode
)

var modeName = map[ExecutorMode]string{
	PriorityMode: "priorityQueue",
	QueueMode:    "queue",
	SingleMode:   "singleDirection",
	TwoMode:      "twoDirection",
	StackMode:    "stack",
}

type DirectiveExecutor struct {
	id        string
	directive *CommandDirective
	mode      ExecutorMode
}

// Create a DirectiveExector for command directive collection
func createExector(directive *CommandDirective, mode ExecutorMode) *DirectiveExecutor {
	return &DirectiveExecutor{
		id:        uuid.NewString(),
		directive: directive,
		mode:      mode,
	}
}
