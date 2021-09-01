package sgen

var (
	From = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "FROM ", "", false) }
)
