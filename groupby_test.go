package sgen

import "testing"

func TestGroupBy(t *testing.T) {
	TestEqual(t, GroupBy(C("t.abc")), "GROUP BY t.abc", []interface{}{})
	TestEqual(t, GroupBy(C("t.abc"), C("t.xyz")), "GROUP BY t.abc, t.xyz", []interface{}{})
	TestEqual(t, GroupBy(C("t.abc"), C("t.xyz"), C("t.uvw")), "GROUP BY t.abc, t.xyz, t.uvw", []interface{}{})
}
