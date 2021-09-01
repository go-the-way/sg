package sgen

import "testing"

func TestAs(t *testing.T) {
	TestEqual(t, As(C("hello"), "hello_lo"), "hello AS hello_lo", nil)
}

func TestAlias(t *testing.T) {
	TestEqual(t, Alias(C("hello"), "hello_lo"), "hello AS hello_lo", nil)
}
