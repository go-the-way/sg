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

func testSingleEqual(t *testing.T, expect, v interface{}) {
	if !reflect.DeepEqual(expect, v) {
		t.Fatalf("%v expect = [%v] but get = [%v]\n", t.Name(), expect, v)
	}
	t.Logf("%v expect = [%v] get get = [%v]\n", t.Name(), expect, v)
}

func testEqual(t *testing.T, g Generator, expectSql string, expectPs []interface{}) {
	getSql, getPs := g.SQL()
	if getSql != expectSql {
		t.Fatalf("%v expect sql = [%v] but get sql = [%v]\n", t.Name(), expectSql, getSql)
	}
	if !reflect.DeepEqual(getPs, expectPs) {
		t.Fatalf("%v expect ps = [%v] but get ps = [%v]\n", t.Name(), expectPs, getPs)
	}
	t.Logf("%v expect sql = [%v] get sql = [%v]\n", t.Name(), expectSql, getSql)
	t.Logf("%v expect ps = [%v] get ps = [%v]\n", t.Name(), expectPs, getPs)
}
