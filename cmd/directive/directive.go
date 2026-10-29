// directive.go is responsible for holding the shared data structures and types for gluster's internal
// data structure systems. These various data structures are implementations of generic queue and stack based systems.
//
// All of the source files within this directive are either the implementations of these data structures or helper modules
package directive

// CommandDirective is the object/node single that within the linked list
// A CommandDirective node holds an id, the next, the prev, the priority (which is optional), and the arg that node is responsible for
type CommandDirective struct {
	Id       string // this is a uuid representing the entire data structure
	Cost     int
	Arg      string
	Priority int
}
