package sgen

var (
	Insert = func(table Ge, gs ...Ge) Ge {
		return NewJoiner([]Ge{P("INSERT INTO"), table,
			NewJoiner(gs, ", ", "", "", true)},
			" ", "", "", false)
	}

	Values = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "VALUES", "", true) }
)
