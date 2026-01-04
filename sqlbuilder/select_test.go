package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSomeFields(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().
		Select().Fields("id", "name").ToSelectQuery().
		From().Table("users").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select id, name from users", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestAll(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().
		Select().Fields("*").ToSelectQuery().
		From().Table("users").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestSomeFunctions(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().
		Select().Count().Avg("age").Sum("salary").Max("height").Min("weight").ToSelectQuery().
		From().Table("users").ToSelectQuery().
		ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select count(*), avg(age), sum(salary), max(height), min(weight) from users",
		sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}
