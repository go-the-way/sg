package sgen

import (
	"testing"
)

func TestJoiner(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true, append: "append"}
	testEqualInterface(t, "sep", j.sep)
	testEqualInterface(t, "prefix", j.prefix)
	testEqualInterface(t, "suffix", j.suffix)
	testEqualInterface(t, true, j.group)
	testEqualInterface(t, "append", j.append)
}

func TestNewJoiner(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true}
	expect := NewJoiner(nil, "sep", "prefix", "suffix", true)
	testEqualInterface(t, expect, j)
}

func TestNewJoinerWithAppend(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true, append: "append"}
	expect := NewJoinerWithAppend(nil, "sep", "prefix", "suffix", "append", true)
	testEqualInterface(t, expect, j)
}

func TestJoiner_SQL(t *testing.T) {
	{
		joiner := NewJoiner([]Generator{&geImpl{}, &geImpl{}}, ",", "prefix ", " suffix", true)
		testEqual(t, joiner, "prefix (geImpl,geImpl) suffix", []interface{}{})
	}

	{
		joiner := NewJoinerWithAppend([]Generator{&geImpl{}, &geImpl{}}, ",", "prefix ", " suffix", " append", true)
		testEqual(t, joiner, "prefix (geImpl append,geImpl append) suffix", []interface{}{})
	}
}
