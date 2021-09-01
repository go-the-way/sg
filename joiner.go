package sgen

import (
	"fmt"
	"strings"
)

type Joiner struct {
	gs     []Generator
	sep    string
	prefix string
	suffix string
	append string
	group  bool
}

func NewJoiner(gs []Generator, sep, prefix, suffix string, group bool) *Joiner {
	return &Joiner{gs, sep, prefix, suffix, "", group}
}

func NewJoinerWithAppend(gs []Generator, sep, prefix, suffix, append string, group bool) *Joiner {
	return &Joiner{gs, sep, prefix, suffix, append, group}
}

func (j *Joiner) SQL() (string, []interface{}) {
	ss := make([]string, 0)
	params := make([]interface{}, 0)
	for _, g := range j.gs {
		sql, ps := g.SQL()
		if sql != "" {
			ss = append(ss, sql+j.append)
		}
		if ps != nil && len(ps) > 0 {
			params = append(params, ps...)
		}
	}
	if len(ss) == 0 {
		return "", params
	}
	gl, gr := "", ""
	if j.group {
		gl, gr = "(", ")"
	}
	return fmt.Sprintf("%s%s%s%s%s", j.prefix, gl, strings.Join(ss, j.sep), gr, j.suffix), params
}
