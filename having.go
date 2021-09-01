package sgen

var (
	Having = func(gs ...Ge) Ge { return NewJoiner(gs, "", "HAVING ", "", false) }
)
