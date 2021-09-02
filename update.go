package sgen

var (
	Update = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "UPDATE ", "", false) }
	Set    = func(gs ...Ge) Ge { return NewJoiner(gs, ", ", "SET ", "", false) }
	SetEq  = func(c C, v interface{}) Ge { return &wC{c, "=", v, "", ""} }
)

// updateBuilder
/*
 A standard UPDATE :
 UPDATE table AS t
 SET t.a = ?, t.b = ?, t.c = ?
 WHERE t.a > 0 AND t.b > 0
**/
type updateBuilder struct {
	update []Ge
	join   []Ge
	set    []Ge
	where  []Ge
}

func UpdateBuilder() *updateBuilder {
	return &updateBuilder{
		update: []Ge{},
		join:   []Ge{},
		set:    []Ge{},
		where:  []Ge{},
	}
}

func (ub *updateBuilder) Update(updates ...Ge) *updateBuilder {
	ub.update = append(ub.update, updates...)
	return ub
}

func (ub *updateBuilder) Join(joins ...Ge) *updateBuilder {
	ub.join = append(ub.join, joins...)
	return ub
}

func (ub *updateBuilder) Set(sets ...Ge) *updateBuilder {
	ub.set = append(ub.set, sets...)
	return ub
}

func (ub *updateBuilder) Where(wheres ...Ge) *updateBuilder {
	ub.where = append(ub.where, wheres...)
	return ub
}

func (ub *updateBuilder) Clear() *updateBuilder {
	ub.update = []Ge{}
	ub.join = []Ge{}
	ub.set = []Ge{}
	ub.where = []Ge{}
	return ub
}

func (ub *updateBuilder) SQL() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Update(ub.update...),
		NewJoiner(ub.join, "", "", "", false),
		Set(ub.set...),
		Where(ub.where...)}, " ", "", "", false)
	return joiner.SQL()
}

func (ub *updateBuilder) Build() (string, []interface{}) {
	return ub.SQL()
}
