package sgen

import "fmt"

type (
	C         = Column
	Ge        = Generator
	Column    string
	Generator interface {
		SQL() (string, []interface{})
	}
)

func (c Column) SQL() (string, []interface{}) {
	return fmt.Sprintf("%s", c), nil
}
