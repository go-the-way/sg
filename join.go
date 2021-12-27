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

var (
	LeftJoin  = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "LEFT JOIN ", "", false) }
	RightJoin = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "RIGHT JOIN ", "", false) }
	InnerJoin = func(gs ...Ge) Ge { return NewJoiner(gs, " ", "INNER JOIN ", "", false) }

	On = func(g Ge) Ge { return NewJoiner([]Ge{g}, "", "ON ", "", true) }
)
