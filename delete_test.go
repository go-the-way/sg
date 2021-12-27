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

import "testing"

func TestDelete(t *testing.T) {
	testEqual(t, Delete([]Ge{}, T("table_a")), `DELETE FROM table_a`, []interface{}{})
	testEqual(t, Delete([]Ge{T("ta.*")}, As(T("table_a"), "ta")), `DELETE ta.* FROM table_a AS ta`, []interface{}{})
	testEqual(t, Delete([]Ge{T("ta.*"), T("tb.*")}, As(T("table_a"), "ta"), As(T("table_b"), "tb")),
		`DELETE ta.*, tb.* FROM table_a AS ta, table_b AS tb`, []interface{}{})
	testEqual(t, Delete([]Ge{T("ta.*"), T("tb.*"), T("tc.*")}, As(T("table_a"), "ta"), As(T("table_b"), "tb"), As(T("table_c"), "tc")),
		`DELETE ta.*, tb.*, tc.* FROM table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestDeleteFrom(t *testing.T) {
	testEqual(t, DeleteFrom(T("table_a")), `DELETE FROM table_a`, []interface{}{})
	testEqual(t, DeleteFrom(T("table_a"), T("table_b")), `DELETE FROM table_a, table_b`, []interface{}{})
	testEqual(t, DeleteFrom(T("table_a"), T("table_b"), T("table_c")), `DELETE FROM table_a, table_b, table_c`, []interface{}{})
}

func TestDeleteBuilder(t *testing.T) {
	builder := DeleteBuilder().
		Delete(T("t1.*")).
		From(As(C("table1"), "t1"), As(C("table2"), "t2")).
		Where(AndGroup(Gt("t1.col1", 100), Gt("t2.col2", 200)))
	testEqual(t, builder, `DELETE t1.* FROM table1 AS t1, table2 AS t2 WHERE ((t1.col1 > ?) AND (t2.col2 > ?))`, []interface{}{100, 200})
}

func TestDeleteBuilder_Clear(t *testing.T) {
	builder := DeleteBuilder().Delete(C("col"))
	builder.Clear()
	if len(builder.delete) > 0 {
		t.Errorf("builder's delete required empty")
	}
	if len(builder.where) > 0 {
		t.Errorf("builder's where required empty")
	}
	if len(builder.from) > 0 {
		t.Errorf("builder's from required empty")
	}
}
