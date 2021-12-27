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
	Delete = func(dgs []Ge, fgs ...Ge) Ge {
		if dgs == nil || len(dgs) <= 0 {
			return DeleteFrom(fgs...)
		}
		return NewJoiner(
			[]Ge{NewJoiner(dgs, ", ", "DELETE ", "", false),
				NewJoiner(fgs, ", ", "FROM ", "", false)}, " ", "", "", false,
		)
	}
	DeleteFrom = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "DELETE FROM ", "", false) }
)

// deleteBuilder
/*
 A standard DELETE :
 DELETE FROM table WHERE col1 > 100
 DELETE t.* FROM table AS t WHERE t.col1 > 100
**/
type deleteBuilder struct {
	delete []Ge
	from   []Ge
	where  []Ge
}

func DeleteBuilder() *deleteBuilder {
	return &deleteBuilder{
		delete: []Ge{},
		from:   []Ge{},
		where:  []Ge{},
	}
}

func (b *deleteBuilder) Delete(deletes ...Ge) *deleteBuilder {
	b.delete = append(b.delete, deletes...)
	return b
}

func (b *deleteBuilder) From(from ...Ge) *deleteBuilder {
	b.from = append(b.from, from...)
	return b
}

func (b *deleteBuilder) Where(wheres ...Ge) *deleteBuilder {
	b.where = append(b.where, wheres...)
	return b
}

func (b *deleteBuilder) Clear() *deleteBuilder {
	b.delete = []Ge{}
	b.from = []Ge{}
	b.where = []Ge{}
	return b
}

func (b *deleteBuilder) Build() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Delete(b.delete, b.from...),
		Where(b.where...)}, " ", "", "", false)
	return joiner.SQL()
}

func (b *deleteBuilder) SQL() (string, []interface{}) {
	return b.Build()
}
