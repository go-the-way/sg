package sgen

import (
	"reflect"
	"testing"
)

func TestColumn(t *testing.T) {
	test := func(col string, cc string) {
		getSql, getPs := C(col).SQL()
		TestEqual(t, getSql, cc, getPs, nil)
	}
	test("hello", "hello")
	test("hello1", "hello1")
	test("hello2", "hello2")
	test("hello3", "hello3")
	test("hello4", "hello4")
}

func TestCEqColumn(t *testing.T) {
	cn := "hello"
	c := C(cn)
	col := Column(cn)
	if !reflect.DeepEqual(c, col) {
		t.Fatalf("C[%v] is an alias for Column[%v], but not equals", c, col)
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
