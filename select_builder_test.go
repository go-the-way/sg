package sgen

import "testing"

func TestSelectBuilder(t *testing.T) {
	sb := SelectBuilder().
		Select(C("a"), C("b")).
		From(T("table_person")).
		Where(AndGroup(Eq("a", 100), Eq("b", 200))).
		OrderBy(DescGroup(C("a"), C("b")))
	TestEqual(t, sb, `SELECT a, b FROM table_person WHERE ((a = ?) AND (b = ?)) ORDER BY a DESC, b DESC`,
		[]interface{}{100, 200})
}
