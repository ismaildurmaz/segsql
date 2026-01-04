package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSomeFields(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().Fields("id", "name").ToSelectQuery().
		From().Table("users").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select id, name from users", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestAll(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().All().ToSelectQuery().
		From().Table("users").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestSomeFunctions(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().Field("city").Count().Avg("age").Sum("salary").Max("height").Min("weight").ToSelectQuery().
		From().Table("users").ToSelectQuery().
		GroupBy().Fields("city").ToSelectQuery().
		ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select city, count(*), avg(age), sum(salary), max(height), min(weight) from users group by city",
		sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users group by city", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}
