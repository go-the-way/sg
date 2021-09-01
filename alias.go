package sgen

import "fmt"

var (
	Alias = func(g Generator, as string) Generator { return &alias{g, as} }
	As    = Alias
)

type alias struct {
	g     Generator
	alias string
}

func (a *alias) SQL() (string, []interface{}) {
	s, ps := a.g.SQL()
	return fmt.Sprintf("%s AS %s", s, a.alias), ps
}
