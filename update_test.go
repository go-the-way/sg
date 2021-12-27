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

func TestUpdate(t *testing.T) {
	testEqual(t, Update(T("table_a")), `UPDATE table_a`, []interface{}{})
	testEqual(t, Update(T("table_a"), T("table_b")), `UPDATE table_a, table_b`, []interface{}{})
	testEqual(t, Update(T("table_a"), T("table_b"), T("table_c")), `UPDATE table_a, table_b, table_c`, []interface{}{})
}

func TestUpdateAs(t *testing.T) {
	testEqual(t, Update(As(T("table_a"), "ta")), `UPDATE table_a AS ta`, []interface{}{})
	testEqual(t, Update(As(T("table_a"), "ta"), As(T("table_b"), "tb")), `UPDATE table_a AS ta, table_b AS tb`, []interface{}{})
	testEqual(t, Update(As(T("table_a"), "ta"), As(T("table_b"), "tb"), As(T("table_c"), "tc")), `UPDATE table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestUpdateAlias(t *testing.T) {
	testEqual(t, Update(Alias(T("table_a"), "ta")), `UPDATE table_a AS ta`, []interface{}{})
	testEqual(t, Update(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb")), `UPDATE table_a AS ta, table_b AS tb`, []interface{}{})
	testEqual(t, Update(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb"), Alias(T("table_c"), "tc")), `UPDATE table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestSet(t *testing.T) {
	testEqual(t, Set(C("t.a = t.b")), `SET t.a = t.b`, []interface{}{})
	testEqual(t, Set(C("t.a = t.b"), SetEq("t.c", 100)), `SET t.a = t.b, t.c = ?`, []interface{}{100})
	testEqual(t, Set(C("t.a = t.b"), C("t.xx = t.yy"), SetEq("t.c", 100)), `SET t.a = t.b, t.xx = t.yy, t.c = ?`, []interface{}{100})
}

func TestUpdateBuilder(t *testing.T) {
	builder := UpdateBuilder().
		Update(As(T("table_person"), "t")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("t.col1 = ta.col1")))).
		Set(SetEq("col1", 100), SetEq("col2", 200)).
		Where(AndGroup(Eq("a", 100), Eq("b", 200)))
	testEqual(t, builder, `UPDATE table_person AS t LEFT JOIN table_a AS ta ON (t.col1 = ta.col1) SET col1 = ?, col2 = ? WHERE ((a = ?) AND (b = ?))`,
		[]interface{}{100, 200, 100, 200})
}

func TestUpdateBuilder_Clear(t *testing.T) {
	builder := UpdateBuilder().Update(T("table"))
	builder.Clear()
	if len(builder.update) > 0 {
		t.Errorf("builder's update required empty")
	}
	if len(builder.set) > 0 {
		t.Errorf("builder's set required empty")
	}
	if len(builder.join) > 0 {
		t.Errorf("builder's join required empty")
	}
	if len(builder.where) > 0 {
		t.Errorf("builder's where required empty")
	}
}
