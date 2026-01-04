package sqlbuilder

import (
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type groupBy interface {
	generateSQL() string
}

type GroupByList = utils.GenericArray[groupBy]

func newGroupByList() GroupByList {
	return GroupByList{}
}

type GroupByBuilder struct {
	groupBy

	owner *SelectQueryBuilder
	items GroupByList
}

func newGroupByBuilder(owner *SelectQueryBuilder) *GroupByBuilder {
	return &GroupByBuilder{owner: owner, items: newGroupByList()}
}

func (g *GroupByBuilder) Fields(fields ...string) *GroupByBuilder {
	g.items.Add(newGroupByField(fields))

	return g
}

func (g *GroupByBuilder) ToSelectQuery() *SelectQueryBuilder {
	return g.owner
}

func (g *GroupByBuilder) generateSQL() string {
	sb := strings.Builder{}
	for i, child := range g.items {
		sql := child.generateSQL()
		sb.WriteString(sql)
		if i < g.items.Len()-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}

type groupByField struct {
	groupBy

	fields []string
}

func newGroupByField(fields []string) *groupByField {
	return &groupByField{fields: fields}
}

func (g *groupByField) generateSQL() string {
	sb := strings.Builder{}
	for i, s := range g.fields {
		sb.WriteString(s)
		if i < len(g.fields)-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}
