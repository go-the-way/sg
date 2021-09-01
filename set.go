package sgen

var (
	Set   = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "SET ", "", false) }
	SetEq = func(c C, v interface{}) Ge { return &wC{c, "=", v, "", ""} }
)
