package where

import (
	"fmt"
	"github.com/billcoding/sgen"
	"strings"
)

var (
	And      = func(g sgen.Generator) sgen.Generator { return &join{[]sgen.Generator{g}, "AND", ""} }
	Or       = func(g sgen.Generator) sgen.Generator { return &join{[]sgen.Generator{g}, "OR", ""} }
	AndGroup = func(gs ...sgen.Generator) sgen.Generator { return &join{gs, "", "AND"} }
	OrGroup  = func(gs ...sgen.Generator) sgen.Generator { return &join{gs, "", "OR"} }
)

type join struct {
	gs     []sgen.Generator
	prefix string
	sep    string
}

func (j *join) SQL() (string, []interface{}) {
	ss := make([]string, 0)
	params := make([]interface{}, 0)
	for _, g := range j.gs {
		sql, ps := g.SQL()
		if sql != "" {
			ss = append(ss, sql)
		}
		if ps != nil && len(ps) > 0 {
			params = append(params, ps...)
		}
	}
	if len(ss) == 0 {
		return "", params
	}
	return fmt.Sprintf(" %s (%s) ", j.prefix, strings.Join(ss, j.sep)), params
}
