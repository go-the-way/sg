package sgen

import "testing"

func TestLeftJoin(t *testing.T) {
	TestEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `LEFT JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	TestEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `LEFT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	TestEqual(t, LeftJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `LEFT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}

func TestRightJoin(t *testing.T) {
	TestEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `RIGHT JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	TestEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `RIGHT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	TestEqual(t, RightJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `RIGHT JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}

func TestInnerJoin(t *testing.T) {
	TestEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(C("ta.col1 = tb.col1")),
	), `INNER JOIN table_a AS ta ON (ta.col1 = tb.col1)`, []interface{}{})

	TestEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"))),
	), `INNER JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2))`, []interface{}{})

	TestEqual(t, InnerJoin(As(T("table_a"), "ta"),
		On(AndGroup(C("ta.col1 = tb.col1"), C("ta.col2 = tb.col2"), C("ta.col3 = tc.col3"))),
	), `INNER JOIN table_a AS ta ON ((ta.col1 = tb.col1 AND ta.col2 = tb.col2 AND ta.col3 = tc.col3))`, []interface{}{})
}
