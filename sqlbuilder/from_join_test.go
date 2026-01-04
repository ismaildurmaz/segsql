package sqlbuilder_test

import (
	"fmt"
	"testing"

	"github.com/ismaildurmaz/segsql/sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestFromBuilder_Joins(t *testing.T) {
	tests := []struct {
		name     string
		build    func(*sqlbuilder.FromBuilder)
		expected string
	}{
		{
			name: "inner join",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("users").
					InnerJoin("roles", "roles.id = users.role_id")
			},
			expected: "users inner join roles on roles.id = users.role_id",
		},
		{
			name: "inner join with alias",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("users u").
					InnerJoinWithAlias("roles", "r.id = u.role_id", "r")
			},
			expected: "users u inner join roles r on r.id = u.role_id",
		},
		{
			name: "left join",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("orders").
					LeftJoin("payments", "payments.order_id = orders.id")
			},
			expected: "orders left join payments on payments.order_id = orders.id",
		},
		{
			name: "right join with alias",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("a").
					RightJoinWithAlias("b", "b.a_id = a.id", "b")
			},
			expected: "a right join b b on b.a_id = a.id",
		},
		{
			name: "full join",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("table1").
					FullJoin("table2", "table2.id = table1.id")
			},
			expected: "table1 full join table2 on table2.id = table1.id",
		},
		{
			name: "cross join",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("users").
					CrossJoin("roles")
			},
			expected: "users cross join roles",
		},
		{
			name: "cross join with alias",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("users u").
					CrossJoinWithAlias("roles", "r")
			},
			expected: "users u cross join roles r",
		},
		{
			name: "multiple joins chained",
			build: func(f *sqlbuilder.FromBuilder) {
				f.
					Table("users u").
					InnerJoinWithAlias("roles", "r.id = u.role_id", "r").
					LeftJoin("permissions", "permissions.role_id = r.id")
			},
			expected: "users u inner join roles r on r.id = u.role_id left join permissions on permissions.role_id = r.id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := sqlbuilder.SelectQuery().From()

			tt.build(f)

			sql, err := f.ToSelectQuery().ToSQL()
			expected := fmt.Sprintf("select * from %s", tt.expected)

			require.NoError(t, err)
			require.Equal(t, expected, sql.SelectSQL)
			require.Equal(t, len(sql.SelectArgs), 0)
		})
	}
}
