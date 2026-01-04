package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSingleTable(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestSingleTableWithAlias(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().TableWithAlias("users", "u").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users u", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users u", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestJoinTables(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").
		InnerJoin("orders", "users.id = orders.user_id").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users inner join orders on users.id = orders.user_id",
		sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users inner join orders on users.id = orders.user_id",
		sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestJoinTableWithAlias(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").
		InnerJoinWithAlias("orders", "users.id = o.user_id", "o").ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users inner join orders o on users.id = o.user_id", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users inner join orders o on users.id = o.user_id",
		sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}
