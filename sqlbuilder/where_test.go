package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestWhere(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ?", sqlResult.SelectSQL)
	require.Equal(t, 1, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ?", sqlResult.CountSQL)
	require.Equal(t, 1, len(sqlResult.CountArgs))
}

func TestWhereAnd(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).And().Like("name", "john").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ? and name like ?", sqlResult.SelectSQL)
	require.Equal(t, 2, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ? and name like ?", sqlResult.CountSQL)
	require.Equal(t, 2, len(sqlResult.CountArgs))
}

func TestWhereOr(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).Or().Like("name", "john").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ? or name like ?", sqlResult.SelectSQL)
	require.Equal(t, 2, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ? or name like ?", sqlResult.CountSQL)
	require.Equal(t, 2, len(sqlResult.CountArgs))
}

func TestWhereInvalidParenthesis(t *testing.T) {
	_, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().Where().EndGroup().
		ToSelectQuery().ToSQL()
	require.Error(t, err)

	_, err = sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().Where().StartGroup().
		ToSelectQuery().ToSQL()
	require.Error(t, err)

	_, err = sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().Where().StartGroup().
		Eq("id", 1).ToSelectQuery().ToSQL()
	require.Error(t, err)

	_, err = sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().Where().
		Eq("id", 1).Or().StartGroup().ToSelectQuery().ToSQL()
	require.Error(t, err)

	_, err = sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().Where().StartGroup().
		Eq("id", 1).EndGroup().EndGroup().ToSelectQuery().ToSQL()
	require.Error(t, err)
}

func TestWhereValidParenthesis(t *testing.T) {
	sqlResult, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().StartGroup().Eq("id", 1).EndGroup().ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users where (id = ?)", sqlResult.SelectSQL)
	require.Equal(t, 1, len(sqlResult.SelectArgs))
	require.Equal(t, "select count(*) from users where (id = ?)", sqlResult.CountSQL)
	require.Equal(t, 1, len(sqlResult.CountArgs))
}
