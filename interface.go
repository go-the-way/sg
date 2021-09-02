package sgen

import "fmt"

var (
	Alias = func(g Ge, as string) Ge { return &alias{g, as} }
	As    = Alias
	Arg   = func(p interface{}) Ge { return &arg{p} }
	From  = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "FROM ", "", false) }
)

type (
	P         = C
	C         = Column
	T         = Table
	Ge        = Generator
	Column    string
	Table     string
	Generator interface {
		SQL() (string, []interface{})
	}
	arg struct {
		p interface{}
	}
	alias struct {
		g     Ge
		alias string
	}
)

func (c Column) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s", c), nil
}

func (t Table) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s", t), nil
}

func (a *arg) SQL() (string, []interface{}) {
	return "?", []interface{}{a.p}
}

func (a *alias) SQL() (string, []interface{}) {
	s, ps := a.g.SQL()
	return fmt.Sprintf("%s AS %s", s, a.alias), ps
}
