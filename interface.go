package sgen

type (
	Column string

	Generator interface {
		SQL() (string, []interface{})
	}
)
