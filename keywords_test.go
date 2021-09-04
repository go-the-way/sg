package sgen

import "testing"

func TestKeywords(t *testing.T) {
	testEqualWithoutPs(t, Table, `TABLE`)
	testEqualWithoutPs(t, Drop, `DROP`)
	testEqualWithoutPs(t, View, `VIEW`)
	testEqualWithoutPs(t, Event, `EVENT`)
	testEqualWithoutPs(t, Procedure, `PROCEDURE`)
	testEqualWithoutPs(t, NotNull, `NOT NULL`)
	testEqualWithoutPs(t, Null, `NULL`)
}
