package where

import (
	"fmt"
	"github.com/billcoding/sgen"
	"strings"
)

var (
	Eq         = func(c sgen.Column, v interface{}) sgen.Generator { return &whereCond{c, "=", v} }
	Gt         = func(c sgen.Column, v interface{}) sgen.Generator { return &whereCond{c, ">", v} }
	GtEq       = func(c sgen.Column, v interface{}) sgen.Generator { return &whereCond{c, ">=", v} }
	Lt         = func(c sgen.Column, v interface{}) sgen.Generator { return &whereCond{c, "<", v} }
	LtEq       = func(c sgen.Column, v interface{}) sgen.Generator { return &whereCond{c, "<=", v} }
	Like       = func(c sgen.Column, v interface{}) sgen.Generator { return &whereLikeCond{c, "%", "%", v} }
	LeftLike   = func(c sgen.Column, v interface{}) sgen.Generator { return &whereLikeCond{c, "%", "", v} }
	RightLike  = func(c sgen.Column, v interface{}) sgen.Generator { return &whereLikeCond{c, "", "%", v} }
	Instr      = func(c sgen.Column, v interface{}) sgen.Generator { return &whereInstr{c, v} }
	In         = func(c sgen.Column, vs ...interface{}) sgen.Generator { return &whereInCond{c, vs} }
	BetweenAnd = func(c sgen.Column, l, r interface{}) sgen.Generator { return &whereBetweenAndCond{c, l, r} }
)

// whereCond
// $c $op ?
// params: p
type whereCond struct {
	c  sgen.Column
	op string
	p  interface{}
}

// whereLikeCond
// $c LIKE CONCAT(l, ? ,r)
// params : p
type whereLikeCond struct {
	c sgen.Column
	l string
	r string
	p interface{}
}

// whereBetweenAndCond
// $c BETWEEN ? AND ?
// params : l, r
type whereBetweenAndCond struct {
	c sgen.Column
	l interface{}
	r interface{}
}

// whereInCond
// $c IN (?...)
// params : vs
type whereInCond struct {
	c  sgen.Column
	ps []interface{}
}

// whereInset
// INSTR($c, ?) > 0
// params : vs
type whereInstr struct {
	c sgen.Column
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

func (w *whereInstr) SQL() (string, []interface{}) {
	return fmt.Sprintf(" (INSTR(%s, ?) > 0) ", w.c), []interface{}{w.p}
}
