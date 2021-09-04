package sgen

import "testing"

func TestDropTable(t *testing.T) {
	testEqualWithoutPs(t, DropTable(T("table")), `DROP TABLE table`)
}

func TestDropView(t *testing.T) {
	testEqualWithoutPs(t, DropView(T("view")), `DROP VIEW view`)
}

func TestDropEvent(t *testing.T) {
	testEqualWithoutPs(t, DropEvent(T("event")), `DROP EVENT event`)
}

func TestDropProcedure(t *testing.T) {
	testEqualWithoutPs(t, DropProcedure(T("procedure")), `DROP PROCEDURE procedure`)
}
