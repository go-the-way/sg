package sgen

var (
	Values = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "VALUES", "", true) }
)
