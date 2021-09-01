package orderby

import (
	"github.com/billcoding/sgen"
	"testing"
)

func TestAsc(t *testing.T) {
	sgen.TestEqual(t, Asc(sgen.C("t.abc")), "t.abc ASC", []interface{}{})
}

func TestDesc(t *testing.T) {
	sgen.TestEqual(t, Desc(sgen.C("t.abc")), "t.abc DESC", []interface{}{})
}

func TestAscGroup(t *testing.T) {
	sgen.TestEqual(t, AscGroup(sgen.C("t.abc")), "t.abc ASC", []interface{}{})
	sgen.TestEqual(t, AscGroup(sgen.C("t.abc"), sgen.C("t.xxx")), "t.abc ASC, t.xxx ASC", []interface{}{})
}

func TestDescGroup(t *testing.T) {
	sgen.TestEqual(t, DescGroup(sgen.C("t.abc")), "t.abc DESC", []interface{}{})
	sgen.TestEqual(t, DescGroup(sgen.C("t.abc"), sgen.C("t.xxx")), "t.abc DESC, t.xxx DESC", []interface{}{})
}

func TestGen(t *testing.T) {
	sgen.TestEqual(t, Gen(AscGroup(sgen.C("t.abc"), sgen.C("t.xxx"))), " ORDER BY t.abc ASC, t.xxx ASC", []interface{}{})
	sgen.TestEqual(t, Gen(AscGroup(sgen.C("t.abc"), sgen.C("t.xxx")), DescGroup(sgen.C("t.abc"), sgen.C("t.xxx"))), " ORDER BY t.abc ASC, t.xxx ASC, t.abc DESC, t.xxx DESC", []interface{}{})
}
