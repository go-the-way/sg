package sgen

import (
	"reflect"
	"testing"
)

type geImpl struct {
}

func (g *geImpl) SQL() (string, []interface{}) {
	return "geImpl", nil
}

func TestSingleEqual(t *testing.T, except, v interface{}) {
	if !reflect.DeepEqual(except, v) {
		t.Fatalf("%v except = [%v] but get = [%v]\n", t.Name(), except, v)
	}
	t.Logf("%v except = [%v] get get = [%v]\n", t.Name(), except, v)
}

func TestEqual(t *testing.T, g Generator, exceptSql string, exceptPs []interface{}) {
	getSql, getPs := g.SQL()
	if getSql != exceptSql {
		t.Fatalf("%v except sql = [%v] but get sql = [%v]\n", t.Name(), exceptSql, getSql)
	}
	if !reflect.DeepEqual(getPs, exceptPs) {
		t.Fatalf("%v except ps = [%v] but get ps = [%v]\n", t.Name(), exceptPs, getPs)
	}
	t.Logf("%v except sql = [%v] get sql = [%v]\n", t.Name(), exceptSql, getSql)
	t.Logf("%v except ps = [%v] get ps = [%v]\n", t.Name(), exceptPs, getPs)
}
