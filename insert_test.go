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

func TestInsert(t *testing.T) {
	testEqual(t, Insert(C("table"), C("aa")), `INSERT INTO table (aa)`, []interface{}{})
	testEqual(t, Insert(C("table"), C("aa"), C("bb")), `INSERT INTO table (aa, bb)`, []interface{}{})
	testEqual(t, Insert(C("table"), C("aa"), C("bb"), C("cc")), `INSERT INTO table (aa, bb, cc)`, []interface{}{})
}

func TestValues(t *testing.T) {
	testEqual(t, Values(Arg(100)), `VALUES (?)`, []interface{}{100})
	testEqual(t, Values(Arg(100), Arg(200)), `VALUES (?, ?)`, []interface{}{100, 200})
	testEqual(t, Values(Arg(100), Arg(200), Arg(300)), `VALUES (?, ?, ?)`, []interface{}{100, 200, 300})
}

func TestInsertBuilder(t *testing.T) {
	builder := InsertBuilder().
		Table(T("table_person")).
		Column(C("col1"), C("col2")).
		Value(Arg(100), Arg(200))
	testEqual(t, builder, `INSERT INTO table_person (col1, col2) VALUES (?, ?)`, []interface{}{100, 200})
}

func TestInsertBuilder_Clear(t *testing.T) {
	builder := InsertBuilder()
	builder.Clear()
	if builder.table != nil {
		t.Errorf("builder's table required nil")
	}
	if len(builder.column) > 0 {
		t.Errorf("builder's column required empty")
	}
	if len(builder.value) > 0 {
		t.Errorf("builder's value required empty")
	}
}
