package commandDirective

// todo implement queue executor

// Create a new node for the Queue data structure
// Requires the current tail
func (c CommandDirective) CreateNode(current *CommandDirective, arg string) *CommandDirective {
	//create the new node
	var new_node = CommandDirective{
		id:   current.id, //inherit the id of the list the tail belongs to
		prev: current,
		arg:  arg,
	}

	current.next = &new_node

	return &new_node
}
