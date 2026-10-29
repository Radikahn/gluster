// package commandDirective
//
//
//
// type TwoDirection interface {
// 	CreateNode()
// 	Remove()
// 	GetNext()
// 	GetPrev()
// 	IsHead()
// 	IsTail()
// }
//
// type SingleDirection interface {
// 	CreateNode(current *CommandDirective)
// 	Remove()
// 	GetNext()
// 	IsHead()
// 	IsTail()
// }
//
// type Queue interface {
// 	CreateQueue(arg string) (*CommandDirective, string)
// 	CreateNode(current *CommandDirective, arg string) *CommandDirective
// 	Remove(current *CommandDirective) error
// 	GetNext(current *CommandDirective) *CommandDirective
// 	IsHead(current *CommandDirective) bool
// 	IsTail(current *CommandDirective) bool
// 	Pop()
// }
//
// type PriorityQueue interface {
// 	CreateNode(tail *CommandDirective, arg string)
// 	Remove(current *CommandDirective)
// 	GetNext(current *CommandDirective)
// 	IsHead(current *CommandDirective)
// 	IsTail(current *CommandDirective)
// }
//
// type Stack interface {
// 	CreateNode()
// 	Remove()
// 	GetNext()
// 	IsHead()
// 	IsTail()
// }
