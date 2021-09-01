package sgen

import "testing"

func TestInsert(t *testing.T) {
	TestEqual(t, Insert(C("table"), C("aa")), `INSERT INTO table (aa)`, []interface{}{})
	TestEqual(t, Insert(C("table"), C("aa"), C("bb")), `INSERT INTO table (aa, bb)`, []interface{}{})
	TestEqual(t, Insert(C("table"), C("aa"), C("bb"), C("cc")), `INSERT INTO table (aa, bb, cc)`, []interface{}{})
}
