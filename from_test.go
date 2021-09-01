package sgen

import "testing"

func TestFrom(t *testing.T) {
	TestEqual(t, From(T("table_a"), T("table_b")), `FROM table_a, table_b`, []interface{}{})
	TestEqual(t, From(Table("table_a"), Table("table_b")), `FROM table_a, table_b`, []interface{}{})
}

func TestFromAs(t *testing.T) {
	TestEqual(t, From(As(T("table_a"), "ta"), As(T("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
	TestEqual(t, From(As(Table("table_a"), "ta"), As(Table("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
}

func TestFromAlias(t *testing.T) {
	TestEqual(t, From(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
	TestEqual(t, From(Alias(Table("table_a"), "ta"), Alias(Table("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
}
