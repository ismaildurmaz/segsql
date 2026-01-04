package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestOrderAsc(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Asc("id").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users order by id", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestOrderDesc(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Desc("id").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users order by id desc", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}

func TestOrderMultiple(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		OrderBy().Asc("id").Desc("name").ToSelectQuery().
		ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users order by id, name desc", sqlResult.SelectSQL)
	require.Equal(t, 0, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users", sqlResult.CountSQL)
	require.Equal(t, 0, len(sqlResult.CountArgs))
}
