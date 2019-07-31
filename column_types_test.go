package jet

import (
	"testing"
)

var subQuery = table1.SELECT(table1ColFloat, table1ColInt).AsTable("sub_query")

func TestNewBoolColumn(t *testing.T) {
	boolColumn := BoolColumn("colBool").From(subQuery)
	AssertPostgreClauseSerialize(t, boolColumn, `sub_query."colBool"`)
	AssertPostgreClauseSerialize(t, boolColumn.EQ(Bool(true)), `(sub_query."colBool" = $1)`, true)
	assertProjectionSerialize(t, boolColumn, `sub_query."colBool" AS "colBool"`)

	boolColumn2 := table1ColBool.From(subQuery)
	AssertPostgreClauseSerialize(t, boolColumn2, `sub_query."table1.col_bool"`)
	AssertPostgreClauseSerialize(t, boolColumn2.EQ(Bool(true)), `(sub_query."table1.col_bool" = $1)`, true)
	assertProjectionSerialize(t, boolColumn2, `sub_query."table1.col_bool" AS "table1.col_bool"`)
}

func TestNewIntColumn(t *testing.T) {
	intColumn := IntegerColumn("col_int").From(subQuery)
	AssertPostgreClauseSerialize(t, intColumn, `sub_query."col_int"`)
	AssertPostgreClauseSerialize(t, intColumn.EQ(Int(12)), `(sub_query."col_int" = $1)`, int64(12))
	assertProjectionSerialize(t, intColumn, `sub_query."col_int" AS "col_int"`)

	intColumn2 := table1ColInt.From(subQuery)
	AssertPostgreClauseSerialize(t, intColumn2, `sub_query."table1.col_int"`)
	AssertPostgreClauseSerialize(t, intColumn2.EQ(Int(14)), `(sub_query."table1.col_int" = $1)`, int64(14))
	assertProjectionSerialize(t, intColumn2, `sub_query."table1.col_int" AS "table1.col_int"`)

}

func TestNewFloatColumnColumn(t *testing.T) {
	floatColumn := FloatColumn("col_float").From(subQuery)
	AssertPostgreClauseSerialize(t, floatColumn, `sub_query."col_float"`)
	AssertPostgreClauseSerialize(t, floatColumn.EQ(Float(1.11)), `(sub_query."col_float" = $1)`, float64(1.11))
	assertProjectionSerialize(t, floatColumn, `sub_query."col_float" AS "col_float"`)

	floatColumn2 := table1ColFloat.From(subQuery)
	AssertPostgreClauseSerialize(t, floatColumn2, `sub_query."table1.col_float"`)
	AssertPostgreClauseSerialize(t, floatColumn2.EQ(Float(2.22)), `(sub_query."table1.col_float" = $1)`, float64(2.22))
	assertProjectionSerialize(t, floatColumn2, `sub_query."table1.col_float" AS "table1.col_float"`)

}
