package sgen

var (
	Delete = func(dgs []Ge, fgs ...Ge) Ge {
		if dgs == nil || len(dgs) <= 0 {
			return DeleteFrom(fgs...)
		}
		return NewJoiner(
			[]Ge{NewJoiner(dgs, ", ", "DELETE ", "", false),
				NewJoiner(fgs, ", ", "FROM ", "", false)}, " ", "", "", false,
		)
	}
	DeleteFrom = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "DELETE FROM ", "", false) }
)

// deleteBuilder
/*
 A standard DELETE :
 DELETE FROM table WHERE col1 > 100
 DELETE t.* FROM table AS t WHERE t.col1 > 100
**/
type deleteBuilder struct {
	delete []Ge
	from   []Ge
	where  []Ge
}

func DeleteBuilder() *deleteBuilder {
	return &deleteBuilder{
		delete: []Ge{},
		from:   []Ge{},
		where:  []Ge{},
	}
}

func (b *deleteBuilder) Delete(deletes ...Ge) *deleteBuilder {
	b.delete = append(b.delete, deletes...)
	return b
}

func (b *deleteBuilder) From(from ...Ge) *deleteBuilder {
	b.from = append(b.from, from...)
	return b
}

func (b *deleteBuilder) Where(wheres ...Ge) *deleteBuilder {
	b.where = append(b.where, wheres...)
	return b
}

func (b *deleteBuilder) Clear() *deleteBuilder {
	b.delete = []Ge{}
	b.from = []Ge{}
	b.where = []Ge{}
	return b
}

func (b *deleteBuilder) SQL() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Delete(b.delete, b.from...),
		Where(b.where...)}, " ", "", "", false)
	return joiner.SQL()
}

func (b *deleteBuilder) Build() (string, []interface{}) {
	return b.SQL()
}
