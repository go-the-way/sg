package sgen

import "testing"

func TestDelete(t *testing.T) {
	TestEqual(t, Delete([]Ge{T("ta.*")}, As(T("table_a"), "ta")), `DELETE ta.* FROM table_a AS ta`, []interface{}{})
	TestEqual(t, Delete([]Ge{T("ta.*"), T("tb.*")}, As(T("table_a"), "ta"), As(T("table_b"), "tb")),
		`DELETE ta.*, tb.* FROM table_a AS ta, table_b AS tb`, []interface{}{})
	TestEqual(t, Delete([]Ge{T("ta.*"), T("tb.*"), T("tc.*")}, As(T("table_a"), "ta"), As(T("table_b"), "tb"), As(T("table_c"), "tc")),
		`DELETE ta.*, tb.*, tc.* FROM table_a AS ta, table_b AS tb, table_c AS tc`, []interface{}{})
}

func TestDeleteFrom(t *testing.T) {
	TestEqual(t, DeleteFrom(T("table_a")), `DELETE FROM table_a`, []interface{}{})
	TestEqual(t, DeleteFrom(T("table_a"), T("table_b")), `DELETE FROM table_a, table_b`, []interface{}{})
	TestEqual(t, DeleteFrom(T("table_a"), T("table_b"), T("table_c")), `DELETE FROM table_a, table_b, table_c`, []interface{}{})
}
