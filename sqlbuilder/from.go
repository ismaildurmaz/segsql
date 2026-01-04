package sqlbuilder

import (
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type FromList = utils.GenericArray[from]

type joinType string

const (
	InnerJoin joinType = "inner"
	LeftJoin  joinType = "left"
	RightJoin joinType = "right"
	FullJoin  joinType = "full"
	CrossJoin joinType = "cross"
)

type from interface {
	generateSQL() (string, utils.AnyList)
}

type FromBuilder struct {
	owner *SelectQueryBuilder
	items FromList
}

func newFromList() FromList {
	return FromList{}
}

func (s *FromBuilder) AppendBuilder(other *FromBuilder) *FromBuilder {
	s.items.AddList(other.items)

	return s
}

func (s *FromBuilder) Table(table string) *FromBuilder {
	s.addFrom(newFromTable(table, ""))

	return s
}

func (s *FromBuilder) TableWithAlias(table string, alias string) *FromBuilder {
	s.addFrom(newFromTable(table, alias))

	return s
}

func (s *FromBuilder) InnerJoin(table string, on string) *FromBuilder {
	s.addFrom(newJoinTable(InnerJoin, table, on, ""))

	return s
}

func (s *FromBuilder) InnerJoinWithAlias(table string, on string, alias string) *FromBuilder {
	s.addFrom(newJoinTable(InnerJoin, table, on, alias))

	return s
}

func (s *FromBuilder) LeftJoin(table string, on string) *FromBuilder {
	s.addFrom(newJoinTable(LeftJoin, table, on, ""))

	return s
}

func (s *FromBuilder) LeftJoinWithAlias(table string, on string, alias string) *FromBuilder {
	s.addFrom(newJoinTable(LeftJoin, table, on, alias))

	return s
}

func (s *FromBuilder) RightJoin(table string, on string) *FromBuilder {
	s.addFrom(newJoinTable(RightJoin, table, on, ""))

	return s
}

func (s *FromBuilder) RightJoinWithAlias(table string, on string, alias string) *FromBuilder {
	s.addFrom(newJoinTable(RightJoin, table, on, alias))

	return s
}

func (s *FromBuilder) FullJoin(table string, on string) *FromBuilder {
	s.addFrom(newJoinTable(FullJoin, table, on, ""))

	return s
}

func (s *FromBuilder) FullJoinWithAlias(table string, on string, alias string) *FromBuilder {
	s.addFrom(newJoinTable(FullJoin, table, on, alias))

	return s
}

func (s *FromBuilder) CrossJoin(table string) *FromBuilder {
	s.addFrom(newJoinTable(CrossJoin, table, "", ""))

	return s
}

func (s *FromBuilder) CrossJoinWithAlias(table string, alias string) *FromBuilder {
	s.addFrom(newJoinTable(CrossJoin, table, "", alias))

	return s
}

func (s *FromBuilder) ToSelectQuery() *SelectQueryBuilder {
	return s.owner
}

func (s *FromBuilder) addFrom(child from) {
	s.items.Add(child)
}

func (s *FromBuilder) generateSQL() (string, utils.AnyList) {
	sb := strings.Builder{}
	args := utils.NewAnyList()

	for i, f := range s.items {
		sql, a := f.generateSQL()

		if i == 0 {
			sb.WriteString(sql)
		} else {
			sb.WriteString(" ")
			sb.WriteString(sql)
		}

		if a != nil && a.Len() > 0 {
			args.AddList(a)
		}
	}

	return sb.String(), args
}

func newFromBuilder(owner *SelectQueryBuilder) *FromBuilder {
	return &FromBuilder{owner: owner, items: newFromList()}
}

type fromTable struct {
	from

	table string
	alias string
}

func newFromTable(table string, alias string) *fromTable {
	return &fromTable{table: table, alias: alias}
}

func (f *fromTable) generateSQL() (string, utils.AnyList) {
	sql := f.table
	if f.alias != "" {
		sql = sql + " " + f.alias
	}

	return sql, nil
}

type joinTable struct {
	from

	joinType joinType
	table    string
	on       string
	alias    string
}

func newJoinTable(joinType joinType, table, on string, alias string) *joinTable {
	return &joinTable{joinType: joinType, table: table, on: on, alias: alias}
}

func (j *joinTable) generateSQL() (string, utils.AnyList) {
	sql := string(j.joinType) + " join " + j.table
	if j.alias != "" {
		sql += " " + j.alias
	}
	if j.on != "" {
		sql += " on " + j.on
	}

	return sql, nil
}
