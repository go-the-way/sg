package sgen

import "testing"

func TestInsert(t *testing.T) {
	testEqual(t, Insert(C("table"), C("aa")), `INSERT INTO table (aa)`, []interface{}{})
	testEqual(t, Insert(C("table"), C("aa"), C("bb")), `INSERT INTO table (aa, bb)`, []interface{}{})
	testEqual(t, Insert(C("table"), C("aa"), C("bb"), C("cc")), `INSERT INTO table (aa, bb, cc)`, []interface{}{})
}

func TestValues(t *testing.T) {
	testEqual(t, Values(Arg(100)), `VALUES (?)`, []interface{}{100})
	testEqual(t, Values(Arg(100), Arg(200)), `VALUES (?, ?)`, []interface{}{100, 200})
	testEqual(t, Values(Arg(100), Arg(200), Arg(300)), `VALUES (?, ?, ?)`, []interface{}{100, 200, 300})
}

func TestInsertBuilder(t *testing.T) {
	builder := InsertBuilder().
		Table(T("table_person")).
		Column(C("col1"), C("col2")).
		Value(Arg(100), Arg(200))
	testEqual(t, builder, `INSERT INTO table_person (col1, col2) VALUES (?, ?)`, []interface{}{100, 200})
}
