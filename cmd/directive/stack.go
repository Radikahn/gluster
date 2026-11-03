package directive

type stackNode struct {
	cmd  *CommandDirective
	next *stackNode
}

type Stack struct {
	head *stackNode
	size int
}

// Create a new Stack.
func NewStack() *Stack {
	return &Stack{}
}

// Add a commandDirective to the top of the stack
func (s *Stack) Push(cmd *CommandDirective) {
	n := &stackNode{cmd: cmd}

	if s.head == nil {
		//case 1 the stack is empty and we are adding a value for the first time
		s.head = n
	} else {
		//case 2 the stack has nodes in it, we have a new head
		var prev = s.head

		s.head = n
		s.head.next = prev
	}

	s.size++
}

// Pop the next node in the stack, return the commandDiretive and success boolean
func (s *Stack) Pop() (*CommandDirective, bool) {
	if s.head == nil {
		return nil, false
	}

	n := s.head

	s.head = n.next

	s.size--

	return n.cmd, true
}
