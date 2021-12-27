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

import "testing"

func TestKeywords(t *testing.T) {
	testEqualWithoutPs(t, Table, "TABLE")
	testEqualWithoutPs(t, Create, "CREATE")
	testEqualWithoutPs(t, Drop, "DROP")
	testEqualWithoutPs(t, View, "VIEW")
	testEqualWithoutPs(t, Unique, "UNIQUE")
	testEqualWithoutPs(t, Index, "INDEX")
	testEqualWithoutPs(t, Event, "EVENT")
	testEqualWithoutPs(t, Procedure, "PROCEDURE")
	testEqualWithoutPs(t, AutoIncrement, "AUTO_INCREMENT")
	testEqualWithoutPs(t, NotNull, "NOT NULL")
	testEqualWithoutPs(t, Null, "NULL")
}
