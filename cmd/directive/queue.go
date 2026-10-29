package directive

type node struct {
	cmd  *CommandDirective
	next *node
	prev *node
}

type Queue struct {
	head *node
	tail *node
	size int
}

// Create a new Queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Push a command to the back of the queue
func (q *Queue) Push(cmd *CommandDirective) {
	n := &node{cmd: cmd}

	if q.tail == nil {
		//case 1 this node is a new tail (we are initializing a queue)
		q.head = n
		q.tail = n
	} else {
		//case 2 the queue has values in it, adding a tail
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}

	q.size++
}

// Pop the next node in the queue and return the command that is to be executed along with a success bool
func (q *Queue) Pop() (*CommandDirective, bool) {
	if q.head == nil {
		return nil, false
	}

	n := q.head

	//previous node node becomes the head
	q.head = n.next

	if q.head == nil {
		q.tail = nil //queue is now empty
	} else {
		q.head.prev = nil
	}

	q.size--

	return n.cmd, true
}
