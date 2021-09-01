package sgen

import "fmt"

var (
	Alias = func(g Ge, as string) Ge { return &alias{g, as} }
	As    = Alias
)

type alias struct {
	g     Ge
	alias string
}

func (a *alias) SQL() (string, []interface{}) {
	s, ps := a.g.SQL()
	return fmt.Sprintf("%s AS %s", s, a.alias), ps
}
