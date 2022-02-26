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

func TestLeftJoin(t *testing.T) {
	testEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `LEFT JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	testEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `LEFT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	testEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `LEFT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}

func TestRightJoin(t *testing.T) {
	testEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `RIGHT JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	testEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `RIGHT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	testEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `RIGHT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}

func TestInnerJoin(t *testing.T) {
	testEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `INNER JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	testEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `INNER JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	testEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `INNER JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}
