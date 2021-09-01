package sgen

import "testing"

func TestValues(t *testing.T) {
	TestEqual(t, Values(Arg(100)), `VALUES(?)`, []interface{}{100})
	TestEqual(t, Values(Arg(100), Arg(200)), `VALUES(?, ?)`, []interface{}{100, 200})
	TestEqual(t, Values(Arg(100), Arg(200), Arg(300)), `VALUES(?, ?, ?)`, []interface{}{100, 200, 300})
}
