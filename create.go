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
	// CreateView
	//
	// CREATE VIEW vm_hello_world AS
	// SELECT UUID() AS uuid
	CreateView = func(view Ge, viewDefinition Ge) Ge {
		return NewJoiner([]Ge{NewJoiner([]Ge{Create, View, view}, " ", "", "", false), viewDefinition}, " AS ", "", "", false)
	}
	// CreateIndex
	//
	// CREATE INDEX idx_name ON table(id, name)
	CreateIndex = func(unique bool, index Ge, table Ge, columns ...Ge) Ge {
		return NewJoiner([]Ge{
			NewJoiner([]Ge{
				Create, map[bool]Ge{true: Unique, false: P("")}[unique], Index, index, P("ON"), table,
			}, " ", "", "", false),
			NewJoiner(columns, ", ", " ", "", true),
		}, "", "", "", false)
	}
	// CreateUniqueIndex
	//
	// CREATE UNIQUE INDEX idx_name ON table(id, name)
	CreateUniqueIndex = func(index Ge, table Ge, columns ...Ge) Ge {
		return CreateIndex(true, index, table, columns...)
	}
	// IndexDefinition defines gen SQL like `unique? index idx_name (id, name)`
	IndexDefinition = func(unique bool, name Ge, columns ...Ge) Ge {
		return NewJoiner([]Ge{
			map[bool]Ge{true: Unique, false: P("")}[unique], Index, name,
			NewJoiner(columns, ", ", "", "", true),
		}, " ", "", "", false)
	}
	// ColumnDefinition defines gen SQL like `name varchar(50) not null default '000' comment 'Name'`
	ColumnDefinition = func(name Ge, typ Ge, null, autoIncrement, defaultSet bool, defaultVal, comment string) Ge {
		return NewJoiner([]Ge{
			name,
			typ,
			map[bool]Ge{true: Null, false: NotNull}[null],
			map[bool]Ge{true: Default(P(fmt.Sprintf("%s", defaultVal))), false: P("")}[defaultSet],
			map[bool]Ge{true: AutoIncrement, false: P("")}[autoIncrement],
			Comment(comment),
		}, " ", "", "", false)
	}
	// PrimaryKey defines gen SQL like `PRIMARY KEY(id, name)`
	PrimaryKey = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "PRIMARY KEY ", "", true) }
	// Default defines gen SQL like `DEFAULT '0'`
	Default = func(ge Ge) Ge { return NewJoiner([]Ge{P("DEFAULT "), ge}, "", "", "", false) }
	// Comment defines gen SQL like `COMMENT 'comment'`
	Comment = func(comment string) Ge { return P(fmt.Sprintf("COMMENT '%s'", comment)) }
)

// createTableBuilder
//
// 1.
// create table table_name(
// 	id int not null auto_increment comment 'ID',
// 	name varchar(200) not null default '0' comment 'NAME',
// 	primary key(id)
// );
//
// 2.
// create table table_name(
// 	id int not null auto_increment primary key comment 'ID',
// 	name varchar(200) not null default '0' comment 'NAME'
// );
type createTableBuilder struct {
	table            Ge
	ifNotExist       bool
	comment          string
	columnDefinition []Ge
	primaryKey       []Ge
	indexes          []Ge
}

func CreateTableBuilder() *createTableBuilder {
	return &createTableBuilder{
		columnDefinition: []Ge{},
		primaryKey:       []Ge{},
		indexes:          []Ge{}}
}

func (c *createTableBuilder) Table(table Ge) *createTableBuilder {
	c.table = table
	return c
}
func (c *createTableBuilder) IfNotExist() *createTableBuilder {
	c.ifNotExist = true
	return c
}

func (c *createTableBuilder) Comment(comment string) *createTableBuilder {
	c.comment = comment
	return c
}

func (c *createTableBuilder) ColumnDefinition(columnDefinition ...Ge) *createTableBuilder {
	c.columnDefinition = append(c.columnDefinition, columnDefinition...)
	return c
}

func (c *createTableBuilder) PrimaryKey(columns ...Ge) *createTableBuilder {
	c.primaryKey = append(c.primaryKey, columns...)
	return c
}

func (c *createTableBuilder) Index(index ...Ge) *createTableBuilder {
	c.indexes = append(c.indexes, index...)
	return c
}

func (c *createTableBuilder) Build() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		NewJoiner([]Ge{Create, Table, map[bool]Ge{true: P("IF NOT EXISTS"), false: C("")}[c.ifNotExist], c.table}, " ", "", "", false),
		NewJoiner([]Ge{
			NewJoiner(c.columnDefinition, ", ", "", "", false),
			PrimaryKey(c.primaryKey...),
			NewJoiner(c.indexes, ", ", "", "", false),
		}, ", ", "", "", true,
		), Comment(c.comment)}, "", "", "", false)
	return joiner.SQL()
}

func (c *createTableBuilder) Clear() *createTableBuilder {
	c.table = nil
	c.columnDefinition = []Ge{}
	c.primaryKey = []Ge{}
	c.indexes = []Ge{}
	return c
}

func (c *createTableBuilder) SQL() (string, []interface{}) {
	return c.Build()
}
