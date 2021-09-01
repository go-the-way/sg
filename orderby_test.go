package sgen

import (
	"testing"
)

func TestAsc(t *testing.T) {
	TestEqual(t, Asc(C("t.abc")), "t.abc ASC", []interface{}{})
}

func TestDesc(t *testing.T) {
	TestEqual(t, Desc(C("t.abc")), "t.abc DESC", []interface{}{})
}

func TestAscGroup(t *testing.T) {
	TestEqual(t, AscGroup(C("t.abc")), "t.abc ASC", []interface{}{})
	TestEqual(t, AscGroup(C("t.abc"), C("t.xxx")), "t.abc ASC, t.xxx ASC", []interface{}{})
}

func TestDescGroup(t *testing.T) {
	TestEqual(t, DescGroup(C("t.abc")), "t.abc DESC", []interface{}{})
	TestEqual(t, DescGroup(C("t.abc"), C("t.xxx")), "t.abc DESC, t.xxx DESC", []interface{}{})
}

func TestOrderBy(t *testing.T) {
	TestEqual(t, OrderBy(AscGroup(C("t.abc"), C("t.xxx"))), "ORDER BY t.abc ASC, t.xxx ASC", []interface{}{})
	TestEqual(t, OrderBy(AscGroup(C("t.abc"), C("t.xxx")), DescGroup(C("t.abc"), C("t.xxx"))), "ORDER BY t.abc ASC, t.xxx ASC, t.abc DESC, t.xxx DESC", []interface{}{})
}
