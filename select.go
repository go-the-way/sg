package sgen

var (
	Select = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "SELECT ", "", false) }
)
