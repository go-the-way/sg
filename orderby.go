package sgen

var (
	OrderBy   = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "ORDER BY ", "", false) }
	Asc       = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "", " ASC", false) }
	Desc      = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "", " DESC", false) }
	AscGroup  = func(gs ...Ge) Ge { return NewJoinerWithAppend(gs, ", ", "", "", " ASC", false) }
	DescGroup = func(gs ...Ge) Ge { return NewJoinerWithAppend(gs, ", ", "", "", " DESC", false) }
)
