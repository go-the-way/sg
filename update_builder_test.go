package sgen

import "testing"

func TestUpdateBuilder(t *testing.T) {
	sb := UpdateBuilder().
		Update(As(T("table_person"), "t")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("t.col1 = ta.col1")))).
		Set(SetEq("col1", 100), SetEq("col2", 200)).
		Where(AndGroup(Eq("a", 100), Eq("b", 200)))
	TestEqual(t, sb, `UPDATE table_person AS t LEFT JOIN table_a AS ta ON (t.col1 = ta.col1) SET col1 = ?, col2 = ? WHERE ((a = ?) AND (b = ?))`,
		[]interface{}{100, 200, 100, 200})
}
