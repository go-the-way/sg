package sgen

import "fmt"

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
)

var (
	Arg = func(p interface{}) Ge { return &arg{p} }
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
