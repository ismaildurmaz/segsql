package sqlbuilder

import (
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type orderBy interface {
	generateSQL() string
}

type OrderByList = utils.GenericArray[orderBy]

func newOrderByList() OrderByList {
	return OrderByList{}
}

type OrderByBuilder struct {
	orderBy

	owner *SelectQueryBuilder
	items OrderByList
}

func (o *OrderByBuilder) Asc(fields ...string) *OrderByBuilder {
	o.items.Add(newOrderByField(fields, false))

	return o
}

func (o *OrderByBuilder) Desc(fields ...string) *OrderByBuilder {
	o.items.Add(newOrderByField(fields, true))

	return o
}

func (o *OrderByBuilder) ToSelectQuery() *SelectQueryBuilder {
	return o.owner
}
func (o *OrderByBuilder) generateSQL() string {
	sb := strings.Builder{}

	for i, f := range o.items {
		sql := f.generateSQL()
		sb.WriteString(sql)

		if i < o.items.Len()-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}

func newOrderByBuilder(owner *SelectQueryBuilder) *OrderByBuilder {
	return &OrderByBuilder{owner: owner, items: newOrderByList()}
}

type orderByField struct {
	orderBy

	fields []string
	desc   bool
}

func newOrderByField(fields []string, desc bool) *orderByField {
	return &orderByField{fields: fields, desc: desc}
}

func (o *orderByField) generateSQL() string {
	sb := strings.Builder{}
	for i, f := range o.fields {
		sb.WriteString(f)

		if o.desc {
			sb.WriteString(" desc")
		}

		if i < len(o.fields)-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}
