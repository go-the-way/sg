package sgen

import "testing"

func TestSelectBuilder(t *testing.T) {
	sb := SelectBuilder().
		Select(C("a"), C("b")).
		From(T("table_person")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1")))).
		Where(AndGroup(Eq("a", 100), Eq("b", 200))).
		OrderBy(DescGroup(C("a"), C("b")))
	TestEqual(t, sb, `SELECT a, b FROM table_person LEFT JOIN table_a AS ta ON (ta.col1 = tb.col1) WHERE ((a = ?) AND (b = ?)) ORDER BY a DESC, b DESC`, []interface{}{100, 200})
}
