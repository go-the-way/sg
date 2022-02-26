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
	"fmt"
	"strings"
)

type Joiner struct {
	gs     []Generator
	sep    string
	prefix string
	suffix string
	append string
	group  bool
}

func NewJoiner(gs []Generator, sep, prefix, suffix string, group bool) *Joiner {
	return &Joiner{gs, sep, prefix, suffix, "", group}
}

func NewJoinerWithAppend(gs []Generator, sep, prefix, suffix, append string, group bool) *Joiner {
	return &Joiner{gs, sep, prefix, suffix, append, group}
}

func (j *Joiner) SQL() (string, []interface{}) {
	ss := make([]string, 0)
	params := make([]interface{}, 0)
	for _, g := range j.gs {
		sql, ps := g.SQL()
		if sql != "" {
			ss = append(ss, sql+j.append)
		}
		if ps != nil && len(ps) > 0 {
			params = append(params, ps...)
		}
	}
	if len(ss) == 0 {
		return "", params
	}
	gl, gr := "", ""
	if j.group {
		gl, gr = "(", ")"
	}
	return fmt.Sprintf("%s%s%s%s%s", j.prefix, gl, strings.Join(ss, j.sep), gr, j.suffix), params
}
