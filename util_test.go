/*
 * Copyright 2021 sg(go-the-way) Author. All Rights Reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sg

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
