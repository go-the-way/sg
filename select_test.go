package sgen

import "testing"

func TestSelect(t *testing.T) {
	TestEqual(t, Select(C("t.col_a")), `SELECT t.col_a`, []interface{}{})
	TestEqual(t, Select(C("t.col_a"), C("t.col_b")), `SELECT t.col_a, t.col_b`, []interface{}{})
	TestEqual(t, Select(C("t.col_a"), C("t.col_b"), C("t.col_c")), `SELECT t.col_a, t.col_b, t.col_c`, []interface{}{})
}

func TestSelectAs(t *testing.T) {
	TestEqual(t, Select(As(C("t.col_a"), "colA")), `SELECT t.col_a AS colA`, []interface{}{})
	TestEqual(t, Select(As(C("t.col_a"), "colA"), As(C("t.col_b"), "colB")), `SELECT t.col_a AS colA, t.col_b AS colB`, []interface{}{})
	TestEqual(t, Select(As(C("t.col_a"), "colA"), As(C("t.col_b"), "colB"), As(C("t.col_c"), "colC")), `SELECT t.col_a AS colA, t.col_b AS colB, t.col_c AS colC`, []interface{}{})
}

func TestSelectAlias(t *testing.T) {
	TestEqual(t, Select(Alias(C("t.col_a"), "colA")), `SELECT t.col_a AS colA`, []interface{}{})
	TestEqual(t, Select(Alias(C("t.col_a"), "colA"), Alias(C("t.col_b"), "colB")), `SELECT t.col_a AS colA, t.col_b AS colB`, []interface{}{})
	TestEqual(t, Select(Alias(C("t.col_a"), "colA"), Alias(C("t.col_b"), "colB"), Alias(C("t.col_c"), "colC")), `SELECT t.col_a AS colA, t.col_b AS colB, t.col_c AS colC`, []interface{}{})
}
