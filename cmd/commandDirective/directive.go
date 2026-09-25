// directive.go is responsible for the priority based linked list system that tasks can be placed into.
// `args` are the main content type in a CommandDirective, they are captures as a string but can be anything from a command
// to any piece of useful instruction.
package commandDirective

import (
	"github.com/google/uuid"
)

// CommandDirective is the object/node single that within the linked list
// A CommandDirective node holds an id, the next, the prev, the priority (which is optional), and the arg that node is responsible for
type CommandDirective struct {
	id       string // this is a uuid representing the entire list the node is a part of
	next     *CommandDirective
	prev     *CommandDirective
	priority int
	cost     int
	arg      string
}

func (curr *CommandDirective) getNext() *CommandDirective {
	return curr.next
}

func (curr *CommandDirective) getPrev() *CommandDirective {
	return curr.prev
}

func (curr *CommandDirective) isHead() bool {
	return curr.prev == nil
}

func (curr *CommandDirective) isTail() bool {
	return curr.next == nil
}

// remove unlinks curr from its list and clears its links.
func (curr *CommandDirective) remove() {
	next := curr.next
	prev := curr.prev

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	curr.next = nil
	curr.prev = nil
}

// Initialization function to start a CommandDirective linked list.
// This will return a *CommandDirective that is labeled as a head and the Id of the linked list
// `createNode` should be used for all other additions to nodes within the same linked list
func startList(
	priority int,
	cost int,
	arg string,
) (*CommandDirective, string) {
	var id = uuid.NewString()
	var head = CommandDirective{id: id, priority: priority, arg: arg}

	return &head, id
}

// Create a new node right after the current for
// the respective linked list the node belongs to
func createNode(
	prev *CommandDirective,
	priority int,
	cost int,
	arg string,
) *CommandDirective {
	newNode := &CommandDirective{
		id:       prev.id, //inheret the id of the previous node
		next:     prev.next,
		prev:     prev,
		priority: priority,
		cost:     cost,
		arg:      arg,
	}

	if prev.next != nil {
		prev.next.prev = newNode
	}

	prev.next = newNode

	return newNode
}

// Get length of linked list. Be cautious if your linked list is extremely long,
// this algorithim runs with O(n)
func GetLength(head *CommandDirective) int {
	var curr *CommandDirective = head
	var count int = 1 // Count starts at one to account for head

	for curr.next != nil {
		curr = curr.next
		count += 1
	}

	return count
}

// Loads a linked list into memory from an array object of strings.
// Each respective string in `tasks []string` will be the `arg` value of each CommandDirective node
// Returns a *CommandDirective that is the head of the linked list, or nil if `tasks` is empty
//
// All costs are set to 1 by default until cost eval is ran
func InitFromArray(tasks []string) (*CommandDirective, string) {
	if len(tasks) == 0 {
		return nil, ""
	}

	head, id := startList(1, 1, tasks[0])
	prev := head

	for i := 1; i < len(tasks); i++ {
		prev = createNode(prev, 1, 1, tasks[i])
	}

	return head, id
}
