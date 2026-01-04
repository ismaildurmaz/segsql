package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestGroupBy_SingleField(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().Fields("status").Count().ToSelectQuery().
		From().Table("orders").ToSelectQuery().
		GroupBy().Fields("status").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select status, count(*) from orders group by status", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from orders group by status", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestGroupBy_MultipleFields(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().Fields("status", "country").Count().Sum("total_amount").ToSelectQuery().
		From().Table("orders").ToSelectQuery().
		GroupBy().Fields("status", "country").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select status, country, count(*), sum(total_amount) from orders "+
		"group by status, country", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from orders group by status, country", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}
