package commandDirective

type TwoDirection interface {
	CreateNode()
	Remove()
	GetNext()
	GetPrev()
	IsHead()
	IsTail()
}

type SingleDirection interface {
	CreateNode()
	Remove()
	GetNext()
	IsHead()
	IsTail()
}

type Queue interface {
	CreateNode(current *CommandDirective, arg string)
	Remove(current *CommandDirective)
	GetNext(current *CommandDirective)
	IsHead(current *CommandDirective)
	IsTail(current *CommandDirective)
}

type PriorityQueue interface {
	CreateNode()
	Remove()
	GetNext()
	IsHead()
	IsTail()
}

type Stack interface {
	CreateNode()
	Remove()
	GetNext()
	IsHead()
	IsTail()
}
