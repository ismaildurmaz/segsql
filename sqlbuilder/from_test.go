package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSingleTable(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestSingleTableWithAlias(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().TableWithAlias("users", "u").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users u", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users u", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestJoinTables(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").
		InnerJoin("orders", "users.id = orders.user_id").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users inner join orders on users.id = orders.user_id",
		sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users inner join orders on users.id = orders.user_id",
		sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestJoinTableWithAlias(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").
		InnerJoinWithAlias("orders", "users.id = o.user_id", "o").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users inner join orders o on users.id = o.user_id", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users inner join orders o on users.id = o.user_id",
		sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}
