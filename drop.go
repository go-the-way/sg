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

var (
	// DropTable defines gen SQL like `DROP TABLE table_name`
	DropTable = func(table Ge) Ge { return NewJoiner([]Ge{Drop, Table, table}, " ", "", "", false) }
	// DropView defines gen SQL like `DROP VIEW view_name`
	DropView = func(view Ge) Ge { return NewJoiner([]Ge{Drop, View, view}, " ", "", "", false) }
	// DropEvent defines gen SQL like `DROP EVENT event_name`
	DropEvent = func(event Ge) Ge { return NewJoiner([]Ge{Drop, Event, event}, " ", "", "", false) }
	// DropProcedure defines gen SQL like `DROP PROCEDURE procedure_name`
	DropProcedure = func(procedure Ge) Ge { return NewJoiner([]Ge{Drop, Procedure, procedure}, " ", "", "", false) }
)
