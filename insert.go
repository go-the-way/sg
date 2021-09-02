package sgen

var (
	Insert = func(table Ge, gs ...Ge) Ge {
		return NewJoiner([]Ge{P("INSERT INTO"), table,
			NewJoiner(gs, ", ", "", "", true)},
			" ", "", "", false)
	}

	Values = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "VALUES ", "", true) }
)

// insertBuilder
/*
 A standard INSERT :
 INSERT INTO table (col1, col2) VALUES (?, ?)
**/
type insertBuilder struct {
	table  Ge
	column []Ge
	value  []Ge
}

func InsertBuilder() *insertBuilder {
	return &insertBuilder{
		column: []Ge{},
		value:  []Ge{},
	}
}

func (b *insertBuilder) Table(table Ge) *insertBuilder {
	b.table = table
	return b
}

func (b *insertBuilder) Column(columns ...Ge) *insertBuilder {
	b.column = append(b.column, columns...)
	return b
}

func (b *insertBuilder) Value(values ...Ge) *insertBuilder {
	b.value = append(b.value, values...)
	return b
}

func (b *insertBuilder) Clear() *insertBuilder {
	b.table = nil
	b.column = []Ge{}
	b.value = []Ge{}
	return b
}

func (b *insertBuilder) SQL() (string, []interface{}) {
	joiner := NewJoiner([]Ge{Insert(b.table),
		NewJoiner(b.column, ", ", "", "", true),
		Values(b.value...)}, " ", "", "", false)
	return joiner.SQL()
}

func (b *insertBuilder) Build() (string, []interface{}) {
	return b.SQL()
}
