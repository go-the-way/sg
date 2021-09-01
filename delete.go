package sgen

var (
	Delete = func(dgs []Ge, fgs ...Ge) Ge {
		if dgs == nil {
			return DeleteFrom(fgs...)
		}
		return NewJoiner(
			[]Ge{NewJoiner(dgs, ", ", "DELETE ", "", false),
				NewJoiner(fgs, ", ", "FROM ", "", false)}, " ", "", "", false,
		)
	}
	DeleteFrom = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "DELETE FROM ", "", false) }
)
