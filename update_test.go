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
