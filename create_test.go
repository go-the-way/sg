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

import "testing"

func TestCreateView(t *testing.T) {
	testEqualWithoutPs(t, CreateView(P("vm_nowTime"), P(`select NOW() AS t`)),
		`CREATE VIEW vm_nowTime AS select NOW() AS t`)
}

func TestCreateIndex(t *testing.T) {
	testEqualWithoutPs(t, CreateIndex(false, P("idx_name"), T("table"), C("name")),
		`CREATE INDEX idx_name ON table (name)`)
	testEqualWithoutPs(t, CreateIndex(true, P("idx_name"), T("table"), C("name")),
		`CREATE UNIQUE INDEX idx_name ON table (name)`)
}

func TestCreateUniqueIndex(t *testing.T) {
	testEqualWithoutPs(t, CreateUniqueIndex(P("idx_name"), T("table"), C("name")),
		`CREATE UNIQUE INDEX idx_name ON table (name)`)
}

func TestIndexDefinition(t *testing.T) {
	testEqualWithoutPs(t, IndexDefinition(false, P("idx_name"), C("name")),
		`INDEX idx_name (name)`)
	testEqualWithoutPs(t, IndexDefinition(false, P("idx_name"), C("name"), C("name1"), C("name2")),
		`INDEX idx_name (name, name1, name2)`)
	testEqualWithoutPs(t, IndexDefinition(true, P("idx_name"), C("name")),
		`UNIQUE INDEX idx_name (name)`)
	testEqualWithoutPs(t, IndexDefinition(true, P("idx_name"), C("name"), C("name1"), C("name2")),
		`UNIQUE INDEX idx_name (name, name1, name2)`)
}

func TestColumnDefinition(t *testing.T) {
	testEqualWithoutPs(t, ColumnDefinition(P("id"), P("int"), false, true, false, "", "ID"),
		`id int NOT NULL AUTO_INCREMENT COMMENT 'ID'`)
	testEqualWithoutPs(t, ColumnDefinition(P("id"), P("bigint"), false, true, false, "", "ID"),
		`id bigint NOT NULL AUTO_INCREMENT COMMENT 'ID'`)
	testEqualWithoutPs(t, ColumnDefinition(P("name"), P("varchar(50)"), false, false, false, "", "Name"),
		`name varchar(50) NOT NULL COMMENT 'Name'`)
	testEqualWithoutPs(t, ColumnDefinition(P("name"), P("varchar(50)"), true, false, true, "'000'", "Name"),
		`name varchar(50) NULL DEFAULT '000' COMMENT 'Name'`)
}

func TestPrimaryKey(t *testing.T) {
	testEqualWithoutPs(t, PrimaryKey(C("id")), `PRIMARY KEY (id)`)
	testEqualWithoutPs(t, PrimaryKey(C("id"), C("name")), `PRIMARY KEY (id, name)`)
	testEqualWithoutPs(t, PrimaryKey(C("id"), C("name"), C("age")), `PRIMARY KEY (id, name, age)`)
}

func TestDefault(t *testing.T) {
	testEqualWithoutPs(t, Default(C("1")), `DEFAULT 1`)
	testEqualWithoutPs(t, Default(C("'HELP'")), `DEFAULT 'HELP'`)
}

func TestComment(t *testing.T) {
	testEqualWithoutPs(t, Comment("help for comment"), `COMMENT 'help for comment'`)
}

func TestCreateTableBuilder(t *testing.T) {
	builder := CreateTableBuilder().
		Table(T("table_person")).
		IfNotExist().
		Commented().
		Comment("the person table").
		ColumnDefinition(
			ColumnDefinition(P("id"), P("int"), false, true, false, "", "ID"),
			ColumnDefinition(P("name"), P("varchar(50)"), false, false, true, "'000'", "Name"),
		).
		PrimaryKey(C("id")).
		Index(IndexDefinition(false, P("idx_name"), C("name")))
	testEqualWithoutPs(t, builder, `CREATE TABLE IF NOT EXISTS table_person(id int NOT NULL AUTO_INCREMENT COMMENT 'ID', name varchar(50) NOT NULL DEFAULT '000' COMMENT 'Name', PRIMARY KEY (id), INDEX idx_name (name))COMMENT 'the person table'`)
}

func TestCreateTableBuilder_Clear(t *testing.T) {
	builder := CreateTableBuilder().Table(P("table_a"))
	builder.Clear()
	if builder.table != nil {
		t.Errorf("builder's table required nil")
	}
	if len(builder.columnDefinition) > 0 {
		t.Errorf("builder's columnDefinition required empty")
	}
	if len(builder.primaryKey) > 0 {
		t.Errorf("builder's primaryKey required empty")
	}
	if len(builder.indexes) > 0 {
		t.Errorf("builder's indexes required empty")
	}
}
