package sgen

import "testing"

func TestUpdate(t *testing.T) {
	TestEqual(t, Update(T("table_a")), `UPDATE table_a`, []interface{}{})
	TestEqual(t, Update(T("table_a"), T("table_b")), `UPDATE table_a, table_b`, []interface{}{})
	TestEqual(t, Update(T("table_a"), T("table_b"), T("table_c")), `UPDATE table_a, table_b, table_c`, []interface{}{})
}

func TestUpdateAs(t *testing.T) {
	TestEqual(t, Update(As(T("table_a"), "ta")), `UPDATE table_a AS ta`, []interface{}{})
	TestEqual(t, Update(As(T("table_a"), "ta"), As(T("table_b"), "tb")), `UPDATE table_a AS ta, table_b AS tb`, []interface{}{})
	TestEqual(t, Update(As(T("table_a"), "ta"), As(T("table_b"), "tb"), As(T("table_c"), "tc")), `UPDATE table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestUpdateAlias(t *testing.T) {
	TestEqual(t, Update(Alias(T("table_a"), "ta")), `UPDATE table_a AS ta`, []interface{}{})
	TestEqual(t, Update(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb")), `UPDATE table_a AS ta, table_b AS tb`, []interface{}{})
	TestEqual(t, Update(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb"), Alias(T("table_c"), "tc")), `UPDATE table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestSet(t *testing.T) {
	TestEqual(t, Set(C("t.a = t.b")), `SET t.a = t.b`, []interface{}{})
	TestEqual(t, Set(C("t.a = t.b"), SetEq("t.c", 100)), `SET t.a = t.b, t.c = ?`, []interface{}{100})
	TestEqual(t, Set(C("t.a = t.b"), C("t.xx = t.yy"), SetEq("t.c", 100)), `SET t.a = t.b, t.xx = t.yy, t.c = ?`, []interface{}{100})
}
func TestUpdateBuilder(t *testing.T) {
	builder := UpdateBuilder().
		Update(As(T("table_person"), "t")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("t.col1 = ta.col1")))).
		Set(SetEq("col1", 100), SetEq("col2", 200)).
		Where(AndGroup(Eq("a", 100), Eq("b", 200)))
	TestEqual(t, builder, `UPDATE table_person AS t LEFT JOIN table_a AS ta ON (t.col1 = ta.col1) SET col1 = ?, col2 = ? WHERE ((a = ?) AND (b = ?))`,
		[]interface{}{100, 200, 100, 200})
}
