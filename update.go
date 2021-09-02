package sgen

var (
	Update = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "UPDATE ", "", false) }
	Set    = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "SET ", "", false) }
	SetEq  = func(c C, v interface{}) Ge { return &wC{c, "=", v, "", ""} }
)

// updateBuilder
/*
 A standard SELECT :
 UPDATE table AS t
 SET t.a = ?, t.b = ?, t.c = ?
 WHERE t.a > 0 AND t.b > 0
**/
type updateBuilder struct {
	updates []Ge
	joins   []Ge
	sets    []Ge
	wheres  []Ge
}

func UpdateBuilder() *updateBuilder {
	return &updateBuilder{
		updates: []Ge{},
		joins:   []Ge{},
		sets:    []Ge{},
		wheres:  []Ge{},
	}
}

func (ub *updateBuilder) Update(updates ...Ge) *updateBuilder {
	ub.updates = append(ub.updates, updates...)
	return ub
}

func (ub *updateBuilder) Join(joins ...Ge) *updateBuilder {
	ub.joins = append(ub.joins, joins...)
	return ub
}

func (ub *updateBuilder) Set(sets ...Ge) *updateBuilder {
	ub.sets = append(ub.sets, sets...)
	return ub
}

func (ub *updateBuilder) Where(wheres ...Ge) *updateBuilder {
	ub.wheres = append(ub.wheres, wheres...)
	return ub
}

func (ub *updateBuilder) Clear() *updateBuilder {
	ub.updates = []Ge{}
	ub.joins = []Ge{}
	ub.sets = []Ge{}
	ub.wheres = []Ge{}
	return ub
}

func (ub *updateBuilder) SQL() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Update(ub.updates...),
		NewJoiner(ub.joins, "", "", "", false),
		Set(ub.sets...),
		Where(ub.wheres...)}, " ", "", "", false)
	return joiner.SQL()
}

func (ub *updateBuilder) Build() (string, []interface{}) {
	return ub.SQL()
}
