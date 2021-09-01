package sgen

import "fmt"

type (
	C         = Column
	T         = Table
	Ge        = Generator
	Column    string
	Table     string
	Generator interface {
		SQL() (string, []interface{})
	}
)

func (c Column) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s", c), nil
}

func (t Table) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s", t), nil
}
