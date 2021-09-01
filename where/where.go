package where

import (
	"fmt"
	"github.com/billcoding/sgen"
	"strings"
)

type (
	wC   = whereCond
	wLC  = whereLikeCond
	wITC = whereInstrCond
	wIC  = whereInCond
	wBAC = whereBetweenAndCond
)

var (
	Gen        = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoiner(gs, "", " WHERE", "", false) }
	And        = func(g sgen.Ge) sgen.Ge { return sgen.NewJoiner([]sgen.Ge{g}, "", " AND", "", false) }
	Or         = func(g sgen.Ge) sgen.Ge { return sgen.NewJoiner([]sgen.Ge{g}, "", " OR", "", false) }
	AndGroup   = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoiner(gs, "AND", "", "", true) }
	OrGroup    = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoiner(gs, "OR", "", "", true) }
	Eq         = func(c sgen.C, v interface{}) sgen.Ge { return &wC{c, "=", v} }
	Gt         = func(c sgen.C, v interface{}) sgen.Ge { return &wC{c, ">", v} }
	GtEq       = func(c sgen.C, v interface{}) sgen.Ge { return &wC{c, ">=", v} }
	Lt         = func(c sgen.C, v interface{}) sgen.Ge { return &wC{c, "<", v} }
	LtEq       = func(c sgen.C, v interface{}) sgen.Ge { return &wC{c, "<=", v} }
	Like       = func(c sgen.C, v interface{}) sgen.Ge { return &wLC{c, "%", "%", v} }
	LeftLike   = func(c sgen.C, v interface{}) sgen.Ge { return &wLC{c, "%", "", v} }
	RightLike  = func(c sgen.C, v interface{}) sgen.Ge { return &wLC{c, "", "%", v} }
	Instr      = func(c sgen.C, v interface{}) sgen.Ge { return &wITC{c, v} }
	In         = func(c sgen.C, vs ...interface{}) sgen.Ge { return &wIC{c, vs} }
	BetweenAnd = func(c sgen.C, l, r interface{}) sgen.Ge { return &wBAC{c, l, r} }
)

// whereCond
// $c $op ?
// params: p
type whereCond struct {
	c  sgen.C
	op string
	p  interface{}
}

// whereLikeCond
// $c LIKE CONCAT(l, ? ,r)
// params : p
type whereLikeCond struct {
	c sgen.C
	l string
	r string
	p interface{}
}

// whereBetweenAndCond
// $c BETWEEN ? AND ?
// params : l, r
type whereBetweenAndCond struct {
	c sgen.C
	l interface{}
	r interface{}
}

// whereInCond
// $c IN (?...)
// params : vs
type whereInCond struct {
	c  sgen.C
	ps []interface{}
}

// whereInstrCond
// INSTR($c, ?) > 0
// params : vs
type whereInstrCond struct {
	c sgen.C
	p interface{}
}

func (w *whereCond) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (%s %s ?) ", w.c, w.op), []interface{}{w.p}
}

func (w *whereLikeCond) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (%s LIKE CONCAT('%s', ?, '%s')) ", w.c, w.l, w.r), []interface{}{w.p}
}

func (w *whereBetweenAndCond) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (%s BETWEEN ? AND ?) ", w.c), []interface{}{w.l, w.r}
}

func (w *whereInCond) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (%s IN (%s)) ", w.c, strings.TrimLeft(strings.Repeat(",?", len(w.ps)), ",")), w.ps
}

func (w *whereInstrCond) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (INSTR(%s, ?) > 0) ", w.c), []interface{}{w.p}
}
