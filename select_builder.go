package sgen

// selectBuilder
/*
 A standard SELECT :
 SELECT t.a, t.b, t.c
 FROM table AS t
 WHERE t.a > 0 AND t.b > 0
 ORDER BY t.a DESC, t.b ASC
**/
type selectBuilder struct {
	selects  []Ge
	from     []Ge
	joins    []Ge
	wheres   []Ge
	orderBys []Ge
}

func SelectBuilder() *selectBuilder {
	return &selectBuilder{
		selects:  []Ge{},
		from:     []Ge{},
		wheres:   []Ge{},
		orderBys: []Ge{},
	}
}

func (sb *selectBuilder) Select(selects ...Ge) *selectBuilder {
	sb.selects = append(sb.selects, selects...)
	return sb
}

func (sb *selectBuilder) From(from ...Ge) *selectBuilder {
	sb.from = append(sb.from, from...)
	return sb
}

func (sb *selectBuilder) Join(joins ...Ge) *selectBuilder {
	sb.joins = append(sb.joins, joins...)
	return sb
}

func (sb *selectBuilder) Where(wheres ...Ge) *selectBuilder {
	sb.wheres = append(sb.wheres, wheres...)
	return sb
}

func (sb *selectBuilder) OrderBy(orderBys ...Ge) *selectBuilder {
	sb.orderBys = append(sb.orderBys, orderBys...)
	return sb
}

func (sb *selectBuilder) Clear() *selectBuilder {
	sb.selects = []Ge{}
	sb.from = []Ge{}
	sb.wheres = []Ge{}
	sb.orderBys = []Ge{}
	return sb
}

func (sb *selectBuilder) SQL() (string, []interface{}) {
	joiner := NewJoiner([]Ge{
		Select(sb.selects...),
		From(sb.from...),
		NewJoiner(sb.joins, "", "", "", false),
		Where(sb.wheres...),
		OrderBy(sb.orderBys...)}, " ", "", "", false)
	return joiner.SQL()
}

func (sb *selectBuilder) Build() (string, []interface{}) {
	return sb.SQL()
}
