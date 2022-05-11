# sg

sg: A simple standard SQL generator written in Go.

[![CircleCI](https://circleci.com/gh/go-the-way/sg/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/sg/tree/main)
[![codecov](https://codecov.io/gh/go-the-way/sg/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/sg)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/sg)](https://goreportcard.com/report/github.com/go-the-way/sg)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/sg?status.svg)](https://pkg.go.dev/github.com/go-the-way/sg?tab=doc)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#sql-query-builders)

## Overview
- [Builders](#builders)
  - [Insert](#insert-builder)
  - [Delete](#delete-builder)
  - [Update](#update-builder)
  - [Select](#select-builder)
- [Generators](#generators)
  - [Create view](#create-view)
  - [Create index](#create-index)
  - [Create unique index](#create-unique-index)
  - [Index definition](#index-definition)
  - [Column definition](#column-definition)
  - [Primary key](#primary-key)
  - [Default](#default)
  - [Delete](#delete)
  - [Delete from](#delete-from)
  - [Drop table](#drop-table)
  - [Drop view](#drop-view)
  - [Drop event](#drop-event)
  - [Drop procedure](#drop-procedure)
  - [Insert](#insert)
  - [Values](#values)
  - [Alias](#alias)
  - [Arg](#arg)
  - [From](#from)
  - [Left join](#left-join)
  - [Right join](#right-join)
  - [Inner join](#inner-join)
  - [On](#on)
  - [Select](#select)
  - [Order by](#order-by)
  - [Asc](#asc)
  - [Desc](#desc)
  - [Asc group](#asc-group)
  - [Desc group](#desc-group)
  - [Group by](#group-by)
  - [Having](#having)
  - [Update](#update)
  - [Set](#set)
  - [Set eq](#set-eq)
  - [Where](#where)
  - [And](#and)
  - [Or](#or)
  - [Not](#not)
  - [And group](#and-group)
  - [Or group](#or-group)
  - [Eq](#eq)
  - [Not eq](#not-eq)
  - [Gt](#gt)
  - [Gt eq](#gt-eq)
  - [Lt](#lt)
  - [Lt eq](#lt-eq)
  - [Like](#like)
  - [Left like](#left-like)
  - [Right like](#right-like)
  - [Instr](#instr)
  - [In](#in)
  - [Between and](#between-and)

###  Builders

#### Insert Builder

```go
package main

import (
  "fmt"
  . "github.com/go-the-way/sg"
)

func main() {
  builder := InsertBuilder().
    Table(T("table_person")).
    Column(C("col1"), C("col2")).
    Value(Arg(100), Arg(200))
  fmt.Println(builder.Build())
  // Output:
  // INSERT INTO table_person (col1, col2) VALUES (?, ?) [100 200]
}
```

#### Delete Builder

```go
package main

import (
	"fmt"
	. "github.com/go-the-way/sg"
)

func main() {
	builder := DeleteBuilder().
		Delete(T("t1.*")).
		From(As(C("table1"), "t1"), As(C("table2"), "t2")).
		Where(AndGroup(Gt("t1.col1", 100), Gt("t2.col2", 200)))
	fmt.Println(builder.Build())
	// Output:
	// DELETE t1.* FROM table1 AS t1, table2 AS t2 WHERE ((t1.col1 > ?) AND (t2.col2 > ?)) [100 200]
}
```

#### Update Builder

```go
package main

import (
	"fmt"
	. "github.com/go-the-way/sg"
)

func main() {
	builder := UpdateBuilder().
		Update(As(T("table_person"), "t")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("t.col1 = ta.col1")))).
		Set(SetEq("col1", 100), SetEq("col2", 200)).
		Where(AndGroup(Eq("a", 100), Eq("b", 200)))
	fmt.Println(builder.Build())
	// Output:
	// UPDATE table_person AS t LEFT JOIN table_a AS ta ON (t.col1 = ta.col1) SET col1 = ?, col2 = ? WHERE ((a = ?) AND (b = ?)) [100 200 100 200]
}
```

#### Select Builder

```go
package main

import (
	"fmt"
	. "github.com/go-the-way/sg"
)

func main() {
	builder := SelectBuilder().
		Select(C("a"), C("b")).
		From(T("table_person")).
		Join(LeftJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1")))).
		Where(AndGroup(Eq("a", 100), Eq("b", 200))).
		OrderBy(DescGroup(C("a"), C("b")))
	fmt.Println(builder.Build())
	// Output:
	// SELECT a, b FROM table_person LEFT JOIN table_a AS ta ON (ta.col1 = tb.col1) WHERE ((a = ?) AND (b = ?)) ORDER BY a DESC, b DESC [100 200]
}
```

###  Generators

#### Create view

```
CreateView(P("vm_nowTime"), P(`select NOW() AS t`)
```

#### Create index

```
CreateIndex(false, P("idx_name"), T("table"), C("name"))
```

#### Create unique index

```
CreateUniqueIndex(P("idx_name"), T("table"), C("name"))
```

#### Index definition

```
IndexDefinition(false, P("idx_name"), C("name"))
```

#### Column definition

```
ColumnDefinition(P("id"), P("int"), false, true, false, "", "ID")
```

#### Primary key

```
PrimaryKey(C("id"))
```

#### Default

```
Default(C("1"))
```

#### Delete

```
Delete([]Ge{}, T("table_a"))
```

#### Delete from

```
DeleteFrom(T("table_a"))
```

#### Drop table

```
DropTable(T("table"))
```

#### Drop view

```
DropView(T("view"))
```

#### Drop event

```
DropEvent(T("event"))
```

#### Drop procedure

```
DropProcedure(T("procedure"))
```

#### Insert

```
Insert(C("table"), C("aa"))
```

#### Values

```
Values(Arg(100))
```

#### Alias

```
Alias(C("hello"), "hello_lo")
```

#### Arg

```
Arg(100)
```

#### From

```
From(T("table_a"), T("table_b"))
```

#### Left join

```
LeftJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1"))
```

#### Right join

```
RightJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1"))
```

#### Inner join

```
InnerJoin(As(T("table_a"), "ta"), On(C("ta.col1 = tb.col1"))
```

#### On

```
On(C("ta.col1 = tb.col1"))
```

#### Select

```
Select(C("t.col_a"))
```

#### Order by

```
OrderBy(AscGroup(C("t.abc"), C("t.xxx")))
```

#### Asc

```
Asc(C("t.abc"))
```

#### Desc

```
Desc(C("t.abc"))
```

#### Asc group

```
AscGroup(C("t.abc"), C("t.xxx"))
```

#### Desc group

```
DescGroup(C("t.abc"), C("t.xxx"))
```

#### Group by

```
GroupBy(C("t.abc"), C("t.xyz"))
```

#### Having

```
Having(AndGroup(Eq("a", 1)))
```

#### Update

```
Update(T("table_a"))
```

#### Set

```
Set(C("t.a = t.b"))
```

#### Set eq

```
SetEq("col1", 100)
```

#### Where

```
Where(AndGroup(Eq("a", 1)))
```

#### And

```
And(Eq("c", 100))
```

#### Or

```
Or(Eq("c", 100))
```

#### Not

```
Not(Eq("c", 100))
```

#### And group

```
AndGroup(Gt("t1.col1", 100), Gt("t2.col2", 200))
```

#### Or group

```
OrGroup(Eq("a", 1), Eq("b", 100))
```

#### Eq

```
Eq("a", 1))
```

#### Not eq

```
NotEq("a", 1))
```

#### Gt

```
Gt("a", 1))
```

#### Gt eq

```
Lt("a", 1))
```

#### Lt

```
Lt("a", 1))
```

#### Lt eq

```
LtEq("a", 1))
```

#### Like

```
Like("a", 1)
```

#### Left like

```
LeftLike("a", 1)
```

#### Right like

```
RightLike("a", 1)
```

#### Instr

```
Instr("a", 1)
```

#### In

```
In("a", 1)
```

#### Between and

```
BetweenAnd(c, rune(100), rune(100))
```

