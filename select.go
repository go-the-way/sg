package sgen

var (
	Select = func(gs ...Generator) Ge { return NewJoiner(gs, ", ", "SELECT ", "", false) }
)
