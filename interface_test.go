package sgen

import (
	"reflect"
	"testing"
)

func TestColumn(t *testing.T) {
	test := func(col string, except string) {
		TestEqual(t, C(col), except, nil)
	}
	test("hello", "hello")
	test("hello1", "hello1")
	test("hello2", "hello2")
	test("hello3", "hello3")
	test("hello4", "hello4")
}

func TestTable(t *testing.T) {
	test := func(col string, except string) {
		TestEqual(t, T(col), except, nil)
	}
	test("table", "table")
	test("table2", "table2")
	test("table3", "table3")
	test("table4", "table4")
	test("table5", "table5")
}

func TestCEqColumn(t *testing.T) {
	cn := "hello"
	c := C(cn)
	col := Column(cn)
	if !reflect.DeepEqual(c, col) {
		t.Fatalf("C[%v] is an alias for Column[%v], but not equals", c, col)
	}
}

func TestTEqTable(t *testing.T) {
	cn := "table"
	c := T(cn)
	col := Table(cn)
	if !reflect.DeepEqual(c, col) {
		t.Fatalf("T[%v] is an alias for Table[%v], but not equals", c, col)
	}
}

func TestGeEqGenerator(t *testing.T) {
	var (
		ge  Ge
		ger Generator
	)
	ge = &geImpl{}
	ger = &geImpl{}
	if !reflect.DeepEqual(ge, ger) {
		t.Fatalf("Ge[%v] is an alias for Generator[%v], but not equals", ge, ger)
	}
}
