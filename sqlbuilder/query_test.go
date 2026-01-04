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
		ToSelectQuery().
		ToSQL()

	require.NoError(t, err)
	require.Equal(t, "select u.id, u.name from users u where u.status = ? and u.role in (?, ?)", sql.SelectSQL)
	require.Equal(t, 3, len(sql.SelectArgs))
	require.Equal(t, []interface{}{"active", "admin", "editor"}, sql.SelectArgs)
	require.Equal(t, "select count(*) from users u where u.status = ? and u.role in (?, ?)", sql.CountSQL)
	require.Equal(t, 3, len(sql.CountArgs))
	require.Equal(t, []interface{}{"active", "admin", "editor"}, sql.CountArgs)
}
