// Copyright 2022 sg Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sg

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
