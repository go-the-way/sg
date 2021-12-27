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

var (
	Select = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "SELECT ", "", false) }

	OrderBy   = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "ORDER BY ", "", false) }
	Asc       = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "", " ASC", false) }
	Desc      = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "", " DESC", false) }
	AscGroup  = func(gs ...Ge) Ge { return NewJoinerWithAppend(gs, ", ", "", "", " ASC", false) }
	DescGroup = func(gs ...Ge) Ge { return NewJoinerWithAppend(gs, ", ", "", "", " DESC", false) }

	GroupBy = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "GROUP BY ", "", false) }

	Having = func(gs ...Ge) Ge { return NewJoiner(gs, "", "HAVING ", "", false) }
)

// selectBuilder
/*
 A standard SELECT :
 SELECT t.a, t.b, t.c
 FROM table AS t
 WHERE t.a > 0 AND t.b > 0
 ORDER BY t.a DESC, t.b ASC
**/
type selectBuilder struct {
	selects []Ge
	from    []Ge
	join    []Ge
	where   []Ge
	orderBy []Ge
}

func SelectBuilder() *selectBuilder {
	return &selectBuilder{
		selects: []Ge{},
		from:    []Ge{},
		where:   []Ge{},
		orderBy: []Ge{},
	}
}

func (b *selectBuilder) Select(selects ...Ge) *selectBuilder {
	b.selects = append(b.selects, selects...)
	return b
}

func (b *selectBuilder) From(from ...Ge) *selectBuilder {
	b.from = append(b.from, from...)
	return b
}

func (b *selectBuilder) Join(joins ...Ge) *selectBuilder {
	b.join = append(b.join, joins...)
	return b
}

func (b *selectBuilder) Where(wheres ...Ge) *selectBuilder {
	b.where = append(b.where, wheres...)
	return b
}

func (b *selectBuilder) OrderBy(orderBys ...Ge) *selectBuilder {
	b.orderBy = append(b.orderBy, orderBys...)
	return b
}

func (b *selectBuilder) Clear() *selectBuilder {
	b.selects = []Ge{}
	b.from = []Ge{}
	b.where = []Ge{}
	b.orderBy = []Ge{}
	return b
}

func (b *selectBuilder) Build() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Select(b.selects...),
		From(b.from...),
		NewJoiner(b.join, "", "", "", false),
		Where(b.where...),
		OrderBy(b.orderBy...)}, " ", "", "", false)
	return joiner.SQL()
}

func (b *selectBuilder) SQL() (string, []interface{}) {
	return b.Build()
}
