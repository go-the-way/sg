// Copyright 2022 sg Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sg

import (
	"fmt"
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
	Where      = func(gs ...Ge) Ge { return NewJoiner(gs, "", "WHERE ", "", false) }
	And        = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "AND ", "", false) }
	Or         = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "OR ", "", false) }
	Not        = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "!", "", false) }
	AndGroup   = func(gs ...Ge) Ge { return NewJoiner(gs, " AND ", "", "", true) }
	OrGroup    = func(gs ...Ge) Ge { return NewJoiner(gs, " OR ", "", "", true) }
	Eq         = func(c C, v interface{}) Ge { return &wC{c, "=", v, "(", ")"} }
	NotEq      = func(c C, v interface{}) Ge { return &wC{c, "!=", v, "(", ")"} }
	Gt         = func(c C, v interface{}) Ge { return &wC{c, ">", v, "(", ")"} }
	GtEq       = func(c C, v interface{}) Ge { return &wC{c, ">=", v, "(", ")"} }
	Lt         = func(c C, v interface{}) Ge { return &wC{c, "<", v, "(", ")"} }
	LtEq       = func(c C, v interface{}) Ge { return &wC{c, "<=", v, "(", ")"} }
	Like       = func(c C, v interface{}) Ge { return &wLC{c, "%", "%", v} }
	LeftLike   = func(c C, v interface{}) Ge { return &wLC{c, "%", "", v} }
	RightLike  = func(c C, v interface{}) Ge { return &wLC{c, "", "%", v} }
	Instr      = func(c C, v interface{}) Ge { return &wITC{c, v} }
	In         = func(c C, vs ...interface{}) Ge { return &wIC{c, vs} }
	BetweenAnd = func(c C, l, r interface{}) Ge { return &wBAC{c, l, r} }
)

// whereCond
// $c $op ?
// params: p
type whereCond struct {
	c    C
	op   string
	p    interface{}
	l, r string
}

// whereLikeCond
// $c LIKE CONCAT(l, ? ,r)
// params : p
type whereLikeCond struct {
	c C
	l string
	r string
	p interface{}
}

// whereBetweenAndCond
// $c BETWEEN ? AND ?
// params : l, r
type whereBetweenAndCond struct {
	c C
	l interface{}
	r interface{}
}

// whereInCond
// $c IN (?...)
// params : vs
type whereInCond struct {
	c  C
	ps []interface{}
}

// whereInstrCond
// INSTR($c, ?) > 0
// params : vs
type whereInstrCond struct {
	c C
	p interface{}
}

func (w *whereCond) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s%s %s ?%s", w.l, w.c, w.op, w.r), []interface{}{w.p}
}

func (w *whereLikeCond) SQL() (string, []interface{}) {
	return fmt.Sprintf("(%s LIKE CONCAT('%s', ?, '%s'))", w.c, w.l, w.r), []interface{}{w.p}
}

func (w *whereBetweenAndCond) SQL() (string, []interface{}) {
	return fmt.Sprintf("(%s BETWEEN ? AND ?)", w.c), []interface{}{w.l, w.r}
}

func (w *whereInCond) SQL() (string, []interface{}) {
	return fmt.Sprintf("(%s IN (%s))", w.c, strings.TrimRight(strings.TrimLeft(strings.Repeat(", ?", len(w.ps)), ", "), " ")), w.ps
}

func (w *whereInstrCond) SQL() (string, []interface{}) {
	return fmt.Sprintf("(INSTR(%s, ?) > 0)", w.c), []interface{}{w.p}
}
