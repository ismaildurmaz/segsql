package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestOrderAsc(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Asc("id").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users order by id", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestOrderDesc(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Desc("id").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users order by id desc", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}

func TestOrderMultiple(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Asc("id").Desc("name").ToSelectQuery().
		ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users order by id, name desc", sql.SelectSQL)
	require.Equal(t, 0, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users", sql.CountSQL)
	require.Equal(t, 0, len(sql.CountArgs))
}
