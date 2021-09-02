package sgen

import "testing"

func TestInsert(t *testing.T) {
	TestEqual(t, Insert(C("table"), C("aa")), `INSERT INTO table (aa)`, []interface{}{})
	TestEqual(t, Insert(C("table"), C("aa"), C("bb")), `INSERT INTO table (aa, bb)`, []interface{}{})
	TestEqual(t, Insert(C("table"), C("aa"), C("bb"), C("cc")), `INSERT INTO table (aa, bb, cc)`, []interface{}{})
}

func TestValues(t *testing.T) {
	TestEqual(t, Values(Arg(100)), `VALUES(?)`, []interface{}{100})
	TestEqual(t, Values(Arg(100), Arg(200)), `VALUES(?, ?)`, []interface{}{100, 200})
	TestEqual(t, Values(Arg(100), Arg(200), Arg(300)), `VALUES(?, ?, ?)`, []interface{}{100, 200, 300})
}
