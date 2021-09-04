package sgen

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
