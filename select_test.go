package sgen

import "testing"

func TestSelect(t *testing.T) {
	TestEqual(t, Select(C("t.col_a")), `SELECT t.col_a`, []interface{}{})
	TestEqual(t, Select(C("t.col_a"), C("t.col_b")), `SELECT t.col_a, t.col_b`, []interface{}{})
	TestEqual(t, Select(C("t.col_a"), C("t.col_b"), C("t.col_c")), `SELECT t.col_a, t.col_b, t.col_c`, []interface{}{})
}

func TestSelectAs(t *testing.T) {
	TestEqual(t, Select(As(C("t.col_a"), "colA")), `SELECT t.col_a AS colA`, []interface{}{})
	TestEqual(t, Select(As(C("t.col_a"), "colA"), As(C("t.col_b"), "colB")), `SELECT t.col_a AS colA, t.col_b AS colB`, []interface{}{})
	TestEqual(t, Select(As(C("t.col_a"), "colA"), As(C("t.col_b"), "colB"), As(C("t.col_c"), "colC")), `SELECT t.col_a AS colA, t.col_b AS colB, t.col_c AS colC`, []interface{}{})
}

func TestSelectAlias(t *testing.T) {
	TestEqual(t, Select(Alias(C("t.col_a"), "colA")), `SELECT t.col_a AS colA`, []interface{}{})
	TestEqual(t, Select(Alias(C("t.col_a"), "colA"), Alias(C("t.col_b"), "colB")), `SELECT t.col_a AS colA, t.col_b AS colB`, []interface{}{})
	TestEqual(t, Select(Alias(C("t.col_a"), "colA"), Alias(C("t.col_b"), "colB"), Alias(C("t.col_c"), "colC")), `SELECT t.col_a AS colA, t.col_b AS colB, t.col_c AS colC`, []interface{}{})
}

func TestAsc(t *testing.T) {
	TestEqual(t, Asc(C("t.abc")), "t.abc ASC", []interface{}{})
}

func TestDesc(t *testing.T) {
	TestEqual(t, Desc(C("t.abc")), "t.abc DESC", []interface{}{})
}

func TestAscGroup(t *testing.T) {
	TestEqual(t, AscGroup(C("t.abc")), "t.abc ASC", []interface{}{})
	TestEqual(t, AscGroup(C("t.abc"), C("t.xxx")), "t.abc ASC, t.xxx ASC", []interface{}{})
}

func TestDescGroup(t *testing.T) {
	TestEqual(t, DescGroup(C("t.abc")), "t.abc DESC", []interface{}{})
	TestEqual(t, DescGroup(C("t.abc"), C("t.xxx")), "t.abc DESC, t.xxx DESC", []interface{}{})
}

func TestOrderBy(t *testing.T) {
	TestEqual(t, OrderBy(AscGroup(C("t.abc"), C("t.xxx"))), "ORDER BY t.abc ASC, t.xxx ASC", []interface{}{})
	TestEqual(t, OrderBy(AscGroup(C("t.abc"), C("t.xxx")), DescGroup(C("t.abc"), C("t.xxx"))), "ORDER BY t.abc ASC, t.xxx ASC, t.abc DESC, t.xxx DESC", []interface{}{})
}

func TestGroupBy(t *testing.T) {
	TestEqual(t, GroupBy(C("t.abc")), "GROUP BY t.abc", []interface{}{})
	TestEqual(t, GroupBy(C("t.abc"), C("t.xyz")), "GROUP BY t.abc, t.xyz", []interface{}{})
	TestEqual(t, GroupBy(C("t.abc"), C("t.xyz"), C("t.uvw")), "GROUP BY t.abc, t.xyz, t.uvw", []interface{}{})
}

func TestHaving(t *testing.T) {
	TestEqual(t, Having(AndGroup(Eq("a", 1))), "HAVING ((a = ?))", []interface{}{1})
	TestEqual(t, Having(AndGroup(Eq("a", 1), Eq("b", 100))), "HAVING ((a = ?) AND (b = ?))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(Gt("a", 1))), "HAVING ((a > ?))", []interface{}{1})
	TestEqual(t, Having(AndGroup(Gt("a", 1), Gt("b", 100))), "HAVING ((a > ?) AND (b > ?))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(GtEq("a", 1))), "HAVING ((a >= ?))", []interface{}{1})
	TestEqual(t, Having(AndGroup(GtEq("a", 1), GtEq("b", 100))), "HAVING ((a >= ?) AND (b >= ?))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(Lt("a", 1))), "HAVING ((a < ?))", []interface{}{1})
	TestEqual(t, Having(AndGroup(Lt("a", 1), Lt("b", 100))), "HAVING ((a < ?) AND (b < ?))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(LtEq("a", 1))), "HAVING ((a <= ?))", []interface{}{1})
	TestEqual(t, Having(AndGroup(LtEq("a", 1), LtEq("b", 100))), "HAVING ((a <= ?) AND (b <= ?))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(Like("a", 1))), "HAVING ((a LIKE CONCAT('%', ?, '%')))", []interface{}{1})
	TestEqual(t, Having(AndGroup(Like("a", 1), Like("b", 100))), "HAVING ((a LIKE CONCAT('%', ?, '%')) AND (b LIKE CONCAT('%', ?, '%')))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(LeftLike("a", 1))), "HAVING ((a LIKE CONCAT('%', ?, '')))", []interface{}{1})
	TestEqual(t, Having(AndGroup(LeftLike("a", 1), LeftLike("b", 100))), "HAVING ((a LIKE CONCAT('%', ?, '')) AND (b LIKE CONCAT('%', ?, '')))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(RightLike("a", 1))), "HAVING ((a LIKE CONCAT('', ?, '%')))", []interface{}{1})
	TestEqual(t, Having(AndGroup(RightLike("a", 1), RightLike("b", 100))), "HAVING ((a LIKE CONCAT('', ?, '%')) AND (b LIKE CONCAT('', ?, '%')))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(In("a", 1, 100))), "HAVING ((a IN (?, ?)))", []interface{}{1, 100})
	TestEqual(t, Having(AndGroup(In("a", 1, 100), In("b", 100, 200))), "HAVING ((a IN (?, ?)) AND (b IN (?, ?)))", []interface{}{1, 100, 100, 200})
	TestEqual(t, Having(AndGroup(Instr("a", 1))), "HAVING ((INSTR(a, ?) > 0))", []interface{}{1})
	TestEqual(t, Having(AndGroup(Instr("a", 1), Instr("b", 100))), "HAVING ((INSTR(a, ?) > 0) AND (INSTR(b, ?) > 0))", []interface{}{1, 100})

	TestEqual(t, Having(OrGroup(Eq("a", 1))), "HAVING ((a = ?))", []interface{}{1})
	TestEqual(t, Having(OrGroup(Eq("a", 1), Eq("b", 100))), "HAVING ((a = ?) OR (b = ?))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(Gt("a", 1))), "HAVING ((a > ?))", []interface{}{1})
	TestEqual(t, Having(OrGroup(Gt("a", 1), Gt("b", 100))), "HAVING ((a > ?) OR (b > ?))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(GtEq("a", 1))), "HAVING ((a >= ?))", []interface{}{1})
	TestEqual(t, Having(OrGroup(GtEq("a", 1), GtEq("b", 100))), "HAVING ((a >= ?) OR (b >= ?))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(Lt("a", 1))), "HAVING ((a < ?))", []interface{}{1})
	TestEqual(t, Having(OrGroup(Lt("a", 1), Lt("b", 100))), "HAVING ((a < ?) OR (b < ?))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(LtEq("a", 1))), "HAVING ((a <= ?))", []interface{}{1})
	TestEqual(t, Having(OrGroup(LtEq("a", 1), LtEq("b", 100))), "HAVING ((a <= ?) OR (b <= ?))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(Like("a", 1))), "HAVING ((a LIKE CONCAT('%', ?, '%')))", []interface{}{1})
	TestEqual(t, Having(OrGroup(Like("a", 1), Like("b", 100))), "HAVING ((a LIKE CONCAT('%', ?, '%')) OR (b LIKE CONCAT('%', ?, '%')))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(LeftLike("a", 1))), "HAVING ((a LIKE CONCAT('%', ?, '')))", []interface{}{1})
	TestEqual(t, Having(OrGroup(LeftLike("a", 1), LeftLike("b", 100))), "HAVING ((a LIKE CONCAT('%', ?, '')) OR (b LIKE CONCAT('%', ?, '')))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(RightLike("a", 1))), "HAVING ((a LIKE CONCAT('', ?, '%')))", []interface{}{1})
	TestEqual(t, Having(OrGroup(RightLike("a", 1), RightLike("b", 100))), "HAVING ((a LIKE CONCAT('', ?, '%')) OR (b LIKE CONCAT('', ?, '%')))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(In("a", 1, 100))), "HAVING ((a IN (?, ?)))", []interface{}{1, 100})
	TestEqual(t, Having(OrGroup(In("a", 1, 100), In("b", 100, 200))), "HAVING ((a IN (?, ?)) OR (b IN (?, ?)))", []interface{}{1, 100, 100, 200})
	TestEqual(t, Having(OrGroup(Instr("a", 1))), "HAVING ((INSTR(a, ?) > 0))", []interface{}{1})
	TestEqual(t, Having(OrGroup(Instr("a", 1), Instr("b", 100))), "HAVING ((INSTR(a, ?) > 0) OR (INSTR(b, ?) > 0))", []interface{}{1, 100})
}

func TestSelectBuilder(t *testing.T) {
	sb := SelectBuilder().
		Select(C("a"), C("b")).
		From(T("table_person")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1")))).
		Where(AndGroup(Eq("a", 100), Eq("b", 200))).
		OrderBy(DescGroup(C("a"), C("b")))
	TestEqual(t, sb, `SELECT a, b FROM table_person LEFT JOIN table_a AS ta ON (ta.col1 = tb.col1) WHERE ((a = ?) AND (b = ?)) ORDER BY a DESC, b DESC`, []interface{}{100, 200})
}
