package sgen

var (
	GroupBy = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "GROUP BY ", "", false) }
)
