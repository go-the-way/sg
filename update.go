package sgen

var (
	Update = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "UPDATE ", "", false) }
)
