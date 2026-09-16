package main

type CommandDirective struct {
	id       int
	next     *CommandDirective
	prev     *CommandDirective
	priority int
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

func startList(
	id int,
	priority int,
	arg string,
) *CommandDirective {
	return &CommandDirective{id: id, priority: priority, arg: arg}
}

// Create a new node right after the current
func createNode(
	id int,
	prev *CommandDirective,
	priority int,
	arg string,
) *CommandDirective {
	newNode := &CommandDirective{
		id:       id,
		next:     prev.next,
		prev:     prev,
		priority: priority,
		arg:      arg,
	}

	if prev.next != nil {
		prev.next.prev = newNode
	}

	prev.next = newNode

	return newNode
}
