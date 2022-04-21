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

// Package sg
/*

sg: A simple standard SQL generator written in Go.

Build a create table SQL:

	builder := CreateTableBuilder().
		Table(T("table_person")).
		ColumnDefinition(
			ColumnDefinition(P("id"), P("int"), false, true, false, "", "ID"),
			ColumnDefinition(P("name"), P("varchar(50)"), false, false, true, "000", "Name"),
		).
		PrimaryKey(C("id")).
		Index(IndexDefinition(false, P("idx_name"), C("name")))
	sql, _ := builder.Build()
	fmt.Println(sql)

Build a insert SQL:

	builder := InsertBuilder().
		Table(T("table_person")).
		Column(C("col1"), C("col2")).
		Value(Arg(100), Arg(200))
	sql, _ := builder.Build()
	fmt.Println(sql)

Build a delete SQL:

	builder := DeleteBuilder().
		Delete(T("t1.*")).
		From(As(C("table1"), "t1"), As(C("table2"), "t2")).
		Where(AndGroup(Gt("t1.col1", 100), Gt("t2.col2", 200)))
	sql, _ := builder.Build()
	fmt.Println(sql)

Build a update SQL:

	builder := UpdateBuilder().
		Update(As(T("table_person"), "t")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("t.col1 = ta.col1")))).
		Set(SetEq("col1", 100), SetEq("col2", 200)).
		Where(AndGroup(Eq("a", 100), Eq("b", 200)))
	sql, _ := builder.Build()
	fmt.Println(sql)

Build a select SQL:

	builder := SelectBuilder().
		Select(C("a"), C("b")).
		From(T("table_person")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1")))).
		Where(AndGroup(Eq("a", 100), Eq("b", 200))).
		OrderBy(DescGroup(C("a"), C("b")))
	sql, _ := builder.Build()
	fmt.Println(sql)

*/
package sg
