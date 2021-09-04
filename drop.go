package sgen

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
