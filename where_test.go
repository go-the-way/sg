package sgen

import (
	"fmt"
	"testing"
)

func testWhere(t *testing.T, c Column, g Ge, op string, v interface{}) {
	expectSql, expectPs := fmt.Sprintf("(%s %s ?)", c, op), []interface{}{v}
	testEqual(t, g, expectSql, expectPs)
}

func testWhereLike(t *testing.T, c Column, g Ge, l, r string, v interface{}) {
	expectSql, expectPs := fmt.Sprintf("(%s LIKE CONCAT('%s', ?, '%s'))", c, l, r), []interface{}{v}
	testEqual(t, g, expectSql, expectPs)
}

func testWhereInstr(t *testing.T, c Column, g Ge, v interface{}) {
	expectSql, expectPs := fmt.Sprintf("(INSTR(%s, ?) > 0)", c), []interface{}{v}
	testEqual(t, g, expectSql, expectPs)
}

func testWhereIn(t *testing.T, c Column, g Ge, inStr string, ps ...interface{}) {
	expectSql, expectPs := fmt.Sprintf("(%s IN (%s))", c, inStr), ps
	testEqual(t, g, expectSql, expectPs)
}

func testWhereBetweenAnd(t *testing.T, c Column, g Ge, ps ...interface{}) {
	expectSql, expectPs := fmt.Sprintf("(%s BETWEEN ? AND ?)", c), ps
	testEqual(t, g, expectSql, expectPs)
}

func TestEq(t *testing.T) {
	c := Column("col")
	op := "="
	testWhere(t, c, Eq(c, rune(100)), op, rune(100))
	testWhere(t, c, Eq(c, 'c'), op, 'c')
	testWhere(t, c, Eq(c, byte(100)), op, byte(100))
	testWhere(t, c, Eq(c, 100), op, 100)
	testWhere(t, c, Eq(c, int8(100)), op, int8(100))
	testWhere(t, c, Eq(c, int16(100)), op, int16(100))
	testWhere(t, c, Eq(c, int32(100)), op, int32(100))
	testWhere(t, c, Eq(c, int64(100)), op, int64(100))
	testWhere(t, c, Eq(c, float32(100)), op, float32(100))
	testWhere(t, c, Eq(c, float64(100)), op, float64(100))
	testWhere(t, c, Eq(c, complex64(100)), op, complex64(100))
	testWhere(t, c, Eq(c, complex128(100)), op, complex128(100))
	testWhere(t, c, Eq(c, "hole hole"), op, "hole hole")
}

func TestGt(t *testing.T) {
	c := Column("col")
	op := ">"
	testWhere(t, c, Gt(c, rune(100)), op, rune(100))
	testWhere(t, c, Gt(c, 'c'), op, 'c')
	testWhere(t, c, Gt(c, byte(100)), op, byte(100))
	testWhere(t, c, Gt(c, 100), op, 100)
	testWhere(t, c, Gt(c, int8(100)), op, int8(100))
	testWhere(t, c, Gt(c, int16(100)), op, int16(100))
	testWhere(t, c, Gt(c, int32(100)), op, int32(100))
	testWhere(t, c, Gt(c, int64(100)), op, int64(100))
	testWhere(t, c, Gt(c, float32(100)), op, float32(100))
	testWhere(t, c, Gt(c, float64(100)), op, float64(100))
	testWhere(t, c, Gt(c, complex64(100)), op, complex64(100))
	testWhere(t, c, Gt(c, complex128(100)), op, complex128(100))
	testWhere(t, c, Gt(c, "hole hole"), op, "hole hole")
}

func TestGtEq(t *testing.T) {
	c := Column("col")
	op := ">="
	testWhere(t, c, GtEq(c, rune(100)), op, rune(100))
	testWhere(t, c, GtEq(c, 'c'), op, 'c')
	testWhere(t, c, GtEq(c, byte(100)), op, byte(100))
	testWhere(t, c, GtEq(c, 100), op, 100)
	testWhere(t, c, GtEq(c, int8(100)), op, int8(100))
	testWhere(t, c, GtEq(c, int16(100)), op, int16(100))
	testWhere(t, c, GtEq(c, int32(100)), op, int32(100))
	testWhere(t, c, GtEq(c, int64(100)), op, int64(100))
	testWhere(t, c, GtEq(c, float32(100)), op, float32(100))
	testWhere(t, c, GtEq(c, float64(100)), op, float64(100))
	testWhere(t, c, GtEq(c, complex64(100)), op, complex64(100))
	testWhere(t, c, GtEq(c, complex128(100)), op, complex128(100))
	testWhere(t, c, GtEq(c, "hole hole"), op, "hole hole")
}

func TestLt(t *testing.T) {
	c := Column("col")
	op := "<"
	testWhere(t, c, Lt(c, rune(100)), op, rune(100))
	testWhere(t, c, Lt(c, 'c'), op, 'c')
	testWhere(t, c, Lt(c, byte(100)), op, byte(100))
	testWhere(t, c, Lt(c, 100), op, 100)
	testWhere(t, c, Lt(c, int8(100)), op, int8(100))
	testWhere(t, c, Lt(c, int16(100)), op, int16(100))
	testWhere(t, c, Lt(c, int32(100)), op, int32(100))
	testWhere(t, c, Lt(c, int64(100)), op, int64(100))
	testWhere(t, c, Lt(c, float32(100)), op, float32(100))
	testWhere(t, c, Lt(c, float64(100)), op, float64(100))
	testWhere(t, c, Lt(c, complex64(100)), op, complex64(100))
	testWhere(t, c, Lt(c, complex128(100)), op, complex128(100))
	testWhere(t, c, Lt(c, "hole hole"), op, "hole hole")
}

func TestLtEq(t *testing.T) {
	c := Column("col")
	op := "<="
	testWhere(t, c, LtEq(c, rune(100)), op, rune(100))
	testWhere(t, c, LtEq(c, 'c'), op, 'c')
	testWhere(t, c, LtEq(c, byte(100)), op, byte(100))
	testWhere(t, c, LtEq(c, 100), op, 100)
	testWhere(t, c, LtEq(c, int8(100)), op, int8(100))
	testWhere(t, c, LtEq(c, int16(100)), op, int16(100))
	testWhere(t, c, LtEq(c, int32(100)), op, int32(100))
	testWhere(t, c, LtEq(c, int64(100)), op, int64(100))
	testWhere(t, c, LtEq(c, float32(100)), op, float32(100))
	testWhere(t, c, LtEq(c, float64(100)), op, float64(100))
	testWhere(t, c, LtEq(c, complex64(100)), op, complex64(100))
	testWhere(t, c, LtEq(c, complex128(100)), op, complex128(100))
	testWhere(t, c, LtEq(c, "hole hole"), op, "hole hole")
}

func TestLike(t *testing.T) {
	c := Column("col")
	l := "%"
	r := "%"
	testWhereLike(t, c, Like(c, rune(100)), l, r, rune(100))
	testWhereLike(t, c, Like(c, 'c'), l, r, 'c')
	testWhereLike(t, c, Like(c, byte(100)), l, r, byte(100))
	testWhereLike(t, c, Like(c, 100), l, r, 100)
	testWhereLike(t, c, Like(c, int8(100)), l, r, int8(100))
	testWhereLike(t, c, Like(c, int16(100)), l, r, int16(100))
	testWhereLike(t, c, Like(c, int32(100)), l, r, int32(100))
	testWhereLike(t, c, Like(c, int64(100)), l, r, int64(100))
	testWhereLike(t, c, Like(c, float32(100)), l, r, float32(100))
	testWhereLike(t, c, Like(c, float64(100)), l, r, float64(100))
	testWhereLike(t, c, Like(c, complex64(100)), l, r, complex64(100))
	testWhereLike(t, c, Like(c, complex128(100)), l, r, complex128(100))
	testWhereLike(t, c, Like(c, "hole hole"), l, r, "hole hole")
}

func TestLeftLike(t *testing.T) {
	c := Column("col")
	l := "%"
	r := ""
	testWhereLike(t, c, LeftLike(c, rune(100)), l, r, rune(100))
	testWhereLike(t, c, LeftLike(c, 'c'), l, r, 'c')
	testWhereLike(t, c, LeftLike(c, byte(100)), l, r, byte(100))
	testWhereLike(t, c, LeftLike(c, 100), l, r, 100)
	testWhereLike(t, c, LeftLike(c, int8(100)), l, r, int8(100))
	testWhereLike(t, c, LeftLike(c, int16(100)), l, r, int16(100))
	testWhereLike(t, c, LeftLike(c, int32(100)), l, r, int32(100))
	testWhereLike(t, c, LeftLike(c, int64(100)), l, r, int64(100))
	testWhereLike(t, c, LeftLike(c, float32(100)), l, r, float32(100))
	testWhereLike(t, c, LeftLike(c, float64(100)), l, r, float64(100))
	testWhereLike(t, c, LeftLike(c, complex64(100)), l, r, complex64(100))
	testWhereLike(t, c, LeftLike(c, complex128(100)), l, r, complex128(100))
	testWhereLike(t, c, LeftLike(c, "hole hole"), l, r, "hole hole")
}

func TestLeftRight(t *testing.T) {
	c := Column("col")
	l := ""
	r := "%"
	testWhereLike(t, c, RightLike(c, rune(100)), l, r, rune(100))
	testWhereLike(t, c, RightLike(c, 'c'), l, r, 'c')
	testWhereLike(t, c, RightLike(c, byte(100)), l, r, byte(100))
	testWhereLike(t, c, RightLike(c, 100), l, r, 100)
	testWhereLike(t, c, RightLike(c, int8(100)), l, r, int8(100))
	testWhereLike(t, c, RightLike(c, int16(100)), l, r, int16(100))
	testWhereLike(t, c, RightLike(c, int32(100)), l, r, int32(100))
	testWhereLike(t, c, RightLike(c, int64(100)), l, r, int64(100))
	testWhereLike(t, c, RightLike(c, float32(100)), l, r, float32(100))
	testWhereLike(t, c, RightLike(c, float64(100)), l, r, float64(100))
	testWhereLike(t, c, RightLike(c, complex64(100)), l, r, complex64(100))
	testWhereLike(t, c, RightLike(c, complex128(100)), l, r, complex128(100))
	testWhereLike(t, c, RightLike(c, "hole hole"), l, r, "hole hole")
}

func TestInstr(t *testing.T) {
	c := Column("col")
	testWhereInstr(t, c, Instr(c, rune(100)), rune(100))
	testWhereInstr(t, c, Instr(c, 'c'), 'c')
	testWhereInstr(t, c, Instr(c, byte(100)), byte(100))
	testWhereInstr(t, c, Instr(c, 100), 100)
	testWhereInstr(t, c, Instr(c, int8(100)), int8(100))
	testWhereInstr(t, c, Instr(c, int16(100)), int16(100))
	testWhereInstr(t, c, Instr(c, int32(100)), int32(100))
	testWhereInstr(t, c, Instr(c, int64(100)), int64(100))
	testWhereInstr(t, c, Instr(c, float32(100)), float32(100))
	testWhereInstr(t, c, Instr(c, float64(100)), float64(100))
	testWhereInstr(t, c, Instr(c, complex64(100)), complex64(100))
	testWhereInstr(t, c, Instr(c, complex128(100)), complex128(100))
	testWhereInstr(t, c, Instr(c, "hole hole"), "hole hole")
}

func TestIn(t *testing.T) {
	c := Column("col")
	inStr := "?"
	testWhereIn(t, c, In(c, rune(100)), inStr, rune(100))
	testWhereIn(t, c, In(c, 'c'), inStr, 'c')
	testWhereIn(t, c, In(c, byte(100)), inStr, byte(100))
	testWhereIn(t, c, In(c, 100), inStr, 100)
	testWhereIn(t, c, In(c, int8(100)), inStr, int8(100))
	testWhereIn(t, c, In(c, int16(100)), inStr, int16(100))
	testWhereIn(t, c, In(c, int32(100)), inStr, int32(100))
	testWhereIn(t, c, In(c, int64(100)), inStr, int64(100))
	testWhereIn(t, c, In(c, float32(100)), inStr, float32(100))
	testWhereIn(t, c, In(c, float64(100)), inStr, float64(100))
	testWhereIn(t, c, In(c, complex64(100)), inStr, complex64(100))
	testWhereIn(t, c, In(c, complex128(100)), inStr, complex128(100))
	testWhereIn(t, c, In(c, "hole hole"), inStr, "hole hole")
}

func TestBetweenAnd(t *testing.T) {
	c := Column("col")
	testWhereBetweenAnd(t, c, BetweenAnd(c, rune(100), rune(100)), rune(100), rune(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, 'c', 'c'), 'c', 'c')
	testWhereBetweenAnd(t, c, BetweenAnd(c, byte(100), byte(100)), byte(100), byte(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, 100, 100), 100, 100)
	testWhereBetweenAnd(t, c, BetweenAnd(c, int8(100), int8(100)), int8(100), int8(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, int16(100), int16(100)), int16(100), int16(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, int32(100), int32(100)), int32(100), int32(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, int64(100), int64(100)), int64(100), int64(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, float32(100), float32(100)), float32(100), float32(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, float64(100), float64(100)), float64(100), float64(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, complex64(100), complex64(100)), complex64(100), complex64(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, complex128(100), complex128(100)), complex128(100), complex128(100))
	testWhereBetweenAnd(t, c, BetweenAnd(c, "hole hole", "hole hole"), "hole hole", "hole hole")
}

func TestAnd(t *testing.T) {
	testEqual(t, And(Eq("c", 100)), `AND (c = ?)`, []interface{}{100})
	testEqual(t, And(Gt("c", 100)), `AND (c > ?)`, []interface{}{100})
	testEqual(t, And(GtEq("c", 100)), `AND (c >= ?)`, []interface{}{100})
	testEqual(t, And(Lt("c", 100)), `AND (c < ?)`, []interface{}{100})
	testEqual(t, And(LtEq("c", 100)), `AND (c <= ?)`, []interface{}{100})
	testEqual(t, And(Like("c", 100)), `AND (c LIKE CONCAT('%', ?, '%'))`, []interface{}{100})
	testEqual(t, And(LeftLike("c", 100)), `AND (c LIKE CONCAT('%', ?, ''))`, []interface{}{100})
	testEqual(t, And(RightLike("c", 100)), `AND (c LIKE CONCAT('', ?, '%'))`, []interface{}{100})
	testEqual(t, And(In("c", 100)), `AND (c IN (?))`, []interface{}{100})
	testEqual(t, And(In("c", 100, 200)), `AND (c IN (?, ?))`, []interface{}{100, 200})
	testEqual(t, And(In("c", 100, 200, 300)), `AND (c IN (?, ?, ?))`, []interface{}{100, 200, 300})
	testEqual(t, And(Instr("c", 100)), `AND (INSTR(c, ?) > 0)`, []interface{}{100})
}

func TestOr(t *testing.T) {
	testEqual(t, Or(Eq("c", 100)), `OR (c = ?)`, []interface{}{100})
	testEqual(t, Or(Gt("c", 100)), `OR (c > ?)`, []interface{}{100})
	testEqual(t, Or(GtEq("c", 100)), `OR (c >= ?)`, []interface{}{100})
	testEqual(t, Or(Lt("c", 100)), `OR (c < ?)`, []interface{}{100})
	testEqual(t, Or(LtEq("c", 100)), `OR (c <= ?)`, []interface{}{100})
	testEqual(t, Or(Like("c", 100)), `OR (c LIKE CONCAT('%', ?, '%'))`, []interface{}{100})
	testEqual(t, Or(LeftLike("c", 100)), `OR (c LIKE CONCAT('%', ?, ''))`, []interface{}{100})
	testEqual(t, Or(RightLike("c", 100)), `OR (c LIKE CONCAT('', ?, '%'))`, []interface{}{100})
	testEqual(t, Or(In("c", 100)), `OR (c IN (?))`, []interface{}{100})
	testEqual(t, Or(In("c", 100, 200)), `OR (c IN (?, ?))`, []interface{}{100, 200})
	testEqual(t, Or(In("c", 100, 200, 300)), `OR (c IN (?, ?, ?))`, []interface{}{100, 200, 300})
	testEqual(t, Or(Instr("c", 100)), `OR (INSTR(c, ?) > 0)`, []interface{}{100})
}

func TestWhere(t *testing.T) {
	testEqual(t, Where(AndGroup(Eq("a", 1))), "WHERE ((a = ?))", []interface{}{1})
	testEqual(t, Where(AndGroup(Eq("a", 1), Eq("b", 100))), "WHERE ((a = ?) AND (b = ?))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(Gt("a", 1))), "WHERE ((a > ?))", []interface{}{1})
	testEqual(t, Where(AndGroup(Gt("a", 1), Gt("b", 100))), "WHERE ((a > ?) AND (b > ?))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(GtEq("a", 1))), "WHERE ((a >= ?))", []interface{}{1})
	testEqual(t, Where(AndGroup(GtEq("a", 1), GtEq("b", 100))), "WHERE ((a >= ?) AND (b >= ?))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(Lt("a", 1))), "WHERE ((a < ?))", []interface{}{1})
	testEqual(t, Where(AndGroup(Lt("a", 1), Lt("b", 100))), "WHERE ((a < ?) AND (b < ?))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(LtEq("a", 1))), "WHERE ((a <= ?))", []interface{}{1})
	testEqual(t, Where(AndGroup(LtEq("a", 1), LtEq("b", 100))), "WHERE ((a <= ?) AND (b <= ?))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(Like("a", 1))), "WHERE ((a LIKE CONCAT('%', ?, '%')))", []interface{}{1})
	testEqual(t, Where(AndGroup(Like("a", 1), Like("b", 100))), "WHERE ((a LIKE CONCAT('%', ?, '%')) AND (b LIKE CONCAT('%', ?, '%')))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(LeftLike("a", 1))), "WHERE ((a LIKE CONCAT('%', ?, '')))", []interface{}{1})
	testEqual(t, Where(AndGroup(LeftLike("a", 1), LeftLike("b", 100))), "WHERE ((a LIKE CONCAT('%', ?, '')) AND (b LIKE CONCAT('%', ?, '')))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(RightLike("a", 1))), "WHERE ((a LIKE CONCAT('', ?, '%')))", []interface{}{1})
	testEqual(t, Where(AndGroup(RightLike("a", 1), RightLike("b", 100))), "WHERE ((a LIKE CONCAT('', ?, '%')) AND (b LIKE CONCAT('', ?, '%')))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(In("a", 1, 100))), "WHERE ((a IN (?, ?)))", []interface{}{1, 100})
	testEqual(t, Where(AndGroup(In("a", 1, 100), In("b", 100, 200))), "WHERE ((a IN (?, ?)) AND (b IN (?, ?)))", []interface{}{1, 100, 100, 200})
	testEqual(t, Where(AndGroup(Instr("a", 1))), "WHERE ((INSTR(a, ?) > 0))", []interface{}{1})
	testEqual(t, Where(AndGroup(Instr("a", 1), Instr("b", 100))), "WHERE ((INSTR(a, ?) > 0) AND (INSTR(b, ?) > 0))", []interface{}{1, 100})

	testEqual(t, Where(OrGroup(Eq("a", 1))), "WHERE ((a = ?))", []interface{}{1})
	testEqual(t, Where(OrGroup(Eq("a", 1), Eq("b", 100))), "WHERE ((a = ?) OR (b = ?))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(Gt("a", 1))), "WHERE ((a > ?))", []interface{}{1})
	testEqual(t, Where(OrGroup(Gt("a", 1), Gt("b", 100))), "WHERE ((a > ?) OR (b > ?))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(GtEq("a", 1))), "WHERE ((a >= ?))", []interface{}{1})
	testEqual(t, Where(OrGroup(GtEq("a", 1), GtEq("b", 100))), "WHERE ((a >= ?) OR (b >= ?))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(Lt("a", 1))), "WHERE ((a < ?))", []interface{}{1})
	testEqual(t, Where(OrGroup(Lt("a", 1), Lt("b", 100))), "WHERE ((a < ?) OR (b < ?))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(LtEq("a", 1))), "WHERE ((a <= ?))", []interface{}{1})
	testEqual(t, Where(OrGroup(LtEq("a", 1), LtEq("b", 100))), "WHERE ((a <= ?) OR (b <= ?))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(Like("a", 1))), "WHERE ((a LIKE CONCAT('%', ?, '%')))", []interface{}{1})
	testEqual(t, Where(OrGroup(Like("a", 1), Like("b", 100))), "WHERE ((a LIKE CONCAT('%', ?, '%')) OR (b LIKE CONCAT('%', ?, '%')))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(LeftLike("a", 1))), "WHERE ((a LIKE CONCAT('%', ?, '')))", []interface{}{1})
	testEqual(t, Where(OrGroup(LeftLike("a", 1), LeftLike("b", 100))), "WHERE ((a LIKE CONCAT('%', ?, '')) OR (b LIKE CONCAT('%', ?, '')))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(RightLike("a", 1))), "WHERE ((a LIKE CONCAT('', ?, '%')))", []interface{}{1})
	testEqual(t, Where(OrGroup(RightLike("a", 1), RightLike("b", 100))), "WHERE ((a LIKE CONCAT('', ?, '%')) OR (b LIKE CONCAT('', ?, '%')))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(In("a", 1, 100))), "WHERE ((a IN (?, ?)))", []interface{}{1, 100})
	testEqual(t, Where(OrGroup(In("a", 1, 100), In("b", 100, 200))), "WHERE ((a IN (?, ?)) OR (b IN (?, ?)))", []interface{}{1, 100, 100, 200})
	testEqual(t, Where(OrGroup(Instr("a", 1))), "WHERE ((INSTR(a, ?) > 0))", []interface{}{1})
	testEqual(t, Where(OrGroup(Instr("a", 1), Instr("b", 100))), "WHERE ((INSTR(a, ?) > 0) OR (INSTR(b, ?) > 0))", []interface{}{1, 100})
}
