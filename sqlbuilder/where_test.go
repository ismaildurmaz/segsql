package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestWhere(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ?", sql.SelectSQL)
	require.Equal(t, 1, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ?", sql.CountSQL)
	require.Equal(t, 1, len(sql.CountArgs))
}

func TestWhereAnd(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).And().Like("name", "john").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ? and name like ?", sql.SelectSQL)
	require.Equal(t, 2, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ? and name like ?", sql.CountSQL)
	require.Equal(t, 2, len(sql.CountArgs))
}

func TestWhereOr(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().Eq("id", 1).Or().Like("name", "john").ToSelectQuery().ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select * from users where id = ? or name like ?", sql.SelectSQL)
	require.Equal(t, 2, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users where id = ? or name like ?", sql.CountSQL)
	require.Equal(t, 2, len(sql.CountArgs))
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
	sql, err := sqlbuilder.SelectQuery().From().Table("users").ToSelectQuery().
		Where().StartGroup().Eq("id", 1).EndGroup().ToSelectQuery().ToSQL()
	require.NoError(t, err)
	require.Equal(t, "select * from users where (id = ?)", sql.SelectSQL)
	require.Equal(t, 1, len(sql.SelectArgs))
	require.Equal(t, "select count(*) from users where (id = ?)", sql.CountSQL)
	require.Equal(t, 1, len(sql.CountArgs))
}
