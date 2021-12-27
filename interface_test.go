/*
 * Copyright 2021 sg(go-the-way) Author. All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sg

import (
	"reflect"
	"testing"
)

func TestColumn(t *testing.T) {
	test := func(col string, expect string) {
		testEqual(t, C(col), expect, nil)
	}
	test("hello", "hello")
	test("hello1", "hello1")
	test("hello2", "hello2")
	test("hello3", "hello3")
	test("hello4", "hello4")
}

func TestC(t *testing.T) {
	test := func(col string, expect string) {
		testEqual(t, T(col), expect, nil)
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

func TestTEqC(t *testing.T) {
	cn := "table"
	c := T(cn)
	col := C(cn)
	if !reflect.DeepEqual(c, col) {
		t.Fatalf("T[%v] is an alias for C[%v], but not equals", c, col)
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

func TestAs(t *testing.T) {
	testEqual(t, As(C("hello"), "hello_lo"), "hello AS hello_lo", nil)
}

func TestAlias(t *testing.T) {
	testEqual(t, Alias(C("hello"), "hello_lo"), "hello AS hello_lo", nil)
}

func TestFrom(t *testing.T) {
	testEqual(t, From(T("table_a"), T("table_b")), `FROM table_a, table_b`, []interface{}{})
	testEqual(t, From(C("table_a"), C("table_b")), `FROM table_a, table_b`, []interface{}{})
}

func TestFromAs(t *testing.T) {
	testEqual(t, From(As(T("table_a"), "ta"), As(T("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
	testEqual(t, From(As(C("table_a"), "ta"), As(C("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
}

func TestFromAlias(t *testing.T) {
	testEqual(t, From(Alias(T("table_a"), "ta"), Alias(T("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
	testEqual(t, From(Alias(C("table_a"), "ta"), Alias(C("table_b"), "tb")), `FROM table_a AS ta, table_b AS tb`, []interface{}{})
}
