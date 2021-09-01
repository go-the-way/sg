package sgen

import (
	"testing"
)

func TestJoiner(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true, append: "append"}
	TestSingleEqual(t, "sep", j.sep)
	TestSingleEqual(t, "prefix", j.prefix)
	TestSingleEqual(t, "suffix", j.suffix)
	TestSingleEqual(t, true, j.group)
	TestSingleEqual(t, "append", j.append)
}

func TestNewJoiner(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true}
	except := NewJoiner(nil, "sep", "prefix", "suffix", true)
	TestSingleEqual(t, except, j)
}

func TestNewJoinerWithAppend(t *testing.T) {
	j := &Joiner{sep: "sep", prefix: "prefix", suffix: "suffix", group: true, append: "append"}
	except := NewJoinerWithAppend(nil, "sep", "prefix", "suffix", "append", true)
	TestSingleEqual(t, except, j)
}

func TestJoiner_SQL(t *testing.T) {
	{
		joiner := NewJoiner([]Generator{&geImpl{}, &geImpl{}}, ",", "prefix ", " suffix", true)
		getSql, getPs := joiner.SQL()
		TestEqual(t, getSql, "prefix (geImpl,geImpl) suffix", getPs, []interface{}{})
	}

	{
		joiner := NewJoinerWithAppend([]Generator{&geImpl{}, &geImpl{}}, ",", "prefix ", " suffix", " append", true)
		getSql, getPs := joiner.SQL()
		TestEqual(t, getSql, "prefix (geImpl append,geImpl append) suffix", getPs, []interface{}{})
	}
}
