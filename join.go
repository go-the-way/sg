package sgen

var (
	LeftJoin  = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "LEFT JOIN ", "", false) }
	RightJoin = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "RIGHT JOIN ", "", false) }
	InnerJoin = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "INNER JOIN ", "", false) }

	On = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "ON ", "", true) }
)
