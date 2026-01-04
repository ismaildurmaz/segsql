package sqlbuilder

import (
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type SelectList = utils.GenericArray[selecting]

type selecting interface {
	generateSQL() (string, utils.AnyList)
}

type SelectBuilder struct {
	owner *SelectQueryBuilder
	items SelectList
}

func (s *SelectBuilder) All() *SelectBuilder {
	s.addSelects(newSelectField("*", "", s))

	return s
}

func (s *SelectBuilder) Count() *SelectBuilder {
	s.addSelects(newSelectField("count(*)", "", s))

	return s
}

func (s *SelectBuilder) Avg(field string) *SelectBuilder {
	s.addSelects(newProjectionField("avg", field, "", s))

	return s
}

func (s *SelectBuilder) Sum(field string) *SelectBuilder {
	s.addSelects(newProjectionField("sum", field, "", s))

	return s
}

func (s *SelectBuilder) Max(field string) *SelectBuilder {
	s.addSelects(newProjectionField("max", field, "", s))

	return s
}

func (s *SelectBuilder) Min(field string) *SelectBuilder {
	s.addSelects(newProjectionField("min", field, "", s))

	return s
}

func (s *SelectBuilder) Field(field string) *SelectBuilder {
	s.addSelects(newSelectField(field, "", s))

	return s
}

func (s *SelectBuilder) Fields(fields ...string) *SelectBuilder {
	items := make([]selecting, len(fields))
	for i, field := range fields {
		items[i] = newSelectField(field, "", s)
	}
	s.addSelects(items...)

	return s
}

func (s *SelectBuilder) ToSelectQuery() *SelectQueryBuilder {
	return s.owner
}

func newSelectList() SelectList {
	return SelectList{}
}

func (s *SelectBuilder) addSelects(child ...selecting) {
	s.items.AddList(child)
}

func newSelectBuilder(owner *SelectQueryBuilder) *SelectBuilder {
	return &SelectBuilder{owner: owner, items: newSelectList()}
}

func (s *SelectBuilder) generateSQL() (string, utils.AnyList) {
	fields := make([]string, len(s.items))
	args := utils.NewAnyList()
	for i, f := range s.items {
		sql, sargs := f.generateSQL()
		fields[i] = sql

		if sargs != nil && sargs.Len() > 0 {
			args.AddList(sargs)
		}
	}
	sql := strings.Join(fields, ", ")

	return sql, args
}

type selectField struct {
	selecting

	field         string
	alias         string
	selectBuilder *SelectBuilder
}

func newSelectField(field string, alias string, selectBuilder *SelectBuilder) *selectField {
	return &selectField{field: field, alias: alias, selectBuilder: selectBuilder}
}

func (s *selectField) generateSQL() (string, utils.AnyList) {
	sb := strings.Builder{}
	sb.WriteString(s.field)
	if s.alias != "" {
		sb.WriteString(" as ")
		sb.WriteString(s.alias)
	}

	return sb.String(), nil
}

type projectionField struct {
	selecting

	funcName      string
	field         string
	alias         string
	selectBuilder *SelectBuilder
}

func newProjectionField(funcName, field string, alias string, selectBuilder *SelectBuilder) *projectionField {
	return &projectionField{funcName: funcName, field: field, alias: alias, selectBuilder: selectBuilder}
}

func (s *projectionField) generateSQL() (string, utils.AnyList) {
	sb := strings.Builder{}
	sb.WriteString(s.funcName)
	sb.WriteString("(")
	sb.WriteString(s.field)
	sb.WriteString(")")
	if s.alias != "" {
		sb.WriteString(" as ")
		sb.WriteString(s.alias)
	}

	return sb.String(), nil
}

// TODO implement subselect
