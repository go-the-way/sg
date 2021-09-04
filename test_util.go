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

func testEqualInterface(t *testing.T, expect, v interface{}) {
	if !reflect.DeepEqual(expect, v) {
		t.Fatalf("%v expect = [%v] \n    get = [%v]\n", t.Name(), expect, v)
	}
	t.Logf("%v expect = [%v] \nget    get = [%v]\n", t.Name(), expect, v)
}

func testEqual(t *testing.T, g Ge, expectSql string, expectPs []interface{}) {
	getSql, getPs := g.SQL()
	if getSql != expectSql {
		t.Fatalf("%v \nexpect sql = [%v] \nget    sql = [%v]\n", t.Name(), expectSql, getSql)
	}
	t.Logf("%v \nexpect sql = [%v] \nget    sql = [%v]\n", t.Name(), expectSql, getSql)
	if expectPs != nil {
		if !reflect.DeepEqual(getPs, expectPs) {
			t.Fatalf("%v \nexpect ps = [%v] \nget    ps = [%v]\n", t.Name(), expectPs, getPs)
		}
		t.Logf("%v \nexpect ps = [%v] \nget    ps = [%v]\n", t.Name(), expectPs, getPs)
	}
}

func testEqualWithoutPs(t *testing.T, g Ge, expectSql string) {
	testEqual(t, g, expectSql, nil)
}
