package sqlbuilder_test

import (
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	sql, err := sqlbuilder.SelectQuery().
		Select().Fields("u.id", "u.name").ToSelectQuery().
		From().TableWithAlias("users", "u").ToSelectQuery().
		Where().
		Eq("u.status", "active").
		And().
		In("u.role", "admin", "editor").
		And().
		StartGroup().Like("u.name", "john").Or().Like("u.last_name", "doe").EndGroup().
		ToSelectQuery().
		OrderBy().Asc("u.name").Desc("u.last_name").ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select u.id, u.name from users u where u.status = ? "+
		"and u.role in (?, ?) and (u.name like ? or u.last_name like ?) "+
		"order by u.name, u.last_name desc", sql.SelectSQL)
	require.Equal(t, 5, len(sql.SelectArgs))
	require.Equal(t, []interface{}{"active", "admin", "editor", "john", "doe"}, sql.SelectArgs)
	require.Equal(t, "select count(*) from users u where u.status = ? and u.role in (?, ?) "+
		"and (u.name like ? or u.last_name like ?)", sql.CountSQL)
	require.Equal(t, 5, len(sql.CountArgs))
	require.Equal(t, []interface{}{"active", "admin", "editor", "john", "doe"}, sql.CountArgs)
}
