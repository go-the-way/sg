/*
 * Copyright 2021 sg(go-the-way) Author. All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sg

import "fmt"

var (
	// As alias for Alias
	As = Alias
	// Alias defines gen SQL like `column AS c` or `table AS t`
	Alias = func(g Ge, as string) Ge { return &alias{g, as} }
	// Arg defines gen SQL like `?` only means an argument
	Arg = func(p interface{}) Ge { return &arg{p} }
	// From defines gen SQL like `FROM ...`
	From = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "FROM ", "", false) }
)

type (
	// P type alias for C
	P = C
	// T type alias for C
	T = C
	// C type alias for Column
	C = Column
	// Ge type alias for Generator
	Ge = Generator
	// Column type def for string
	Column string
	// Generator defines Interface its
	Generator interface {
		SQL() (string, []interface{})
	}
	Builder interface {
		Build() (string, []interface{})
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

func (a *arg) SQL() (string, []interface{}) {
	return "?", []interface{}{a.p}
}

func (a *alias) SQL() (string, []interface{}) {
	s, ps := a.g.SQL()
	return fmt.Sprintf("%s AS %s", s, a.alias), ps
}
