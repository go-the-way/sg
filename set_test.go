package sgen

import "testing"

func TestSet(t *testing.T) {
	TestEqual(t, Set(C("t.a = t.b")), `SET t.a = t.b`, []interface{}{})
	TestEqual(t, Set(C("t.a = t.b"), SetEq("t.c", 100)), `SET t.a = t.b, t.c = ?`, []interface{}{100})
	TestEqual(t, Set(C("t.a = t.b"), C("t.xx = t.yy"), SetEq("t.c", 100)), `SET t.a = t.b, t.xx = t.yy, t.c = ?`, []interface{}{100})
}
