package sqlbuilder

import (
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type SelectQueryBuilder struct {
	selects  SelectList
	froms    FromList
	wheres   WhereList
	orderBys OrderByList
	groupBys GroupByList
	offset   *int
	limit    *int
}

func SelectQuery() *SelectQueryBuilder {
	return &SelectQueryBuilder{
		selects:  newSelectList(),
		froms:    newFromList(),
		wheres:   newWhereList(),
		groupBys: newGroupByList(),
	}
}

const (
	minimumPage     = 1
	defaultPageSize = 10
)

func (q *SelectQueryBuilder) Select() *SelectBuilder {
	s := newSelectBuilder(q)
	q.selects.Add(s)

	return s
}

func (q *SelectQueryBuilder) From() *FromBuilder {
	f := newFromBuilder(q)
	q.froms.Add(f)

	return f
}

func (q *SelectQueryBuilder) Where() *WhereBuilder {
	r := newWhereBuilder(q)
	q.wheres.Add(r)

	return r
}

func (q *SelectQueryBuilder) OrderBy() *OrderByBuilder {
	o := newOrderByBuilder(q)
	q.orderBys.Add(o)

	return o
}

func (q *SelectQueryBuilder) GroupBy() *GroupByBuilder {
	g := newGroupByBuilder(q)
	q.groupBys.Add(g)

	return g
}

func (q *SelectQueryBuilder) Page(page int, pageSize int) *SelectQueryBuilder {
	if page <= 0 {
		page = minimumPage
	}

	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	return q.Offset((page - 1) * pageSize).Limit(pageSize)
}

func (q *SelectQueryBuilder) Offset(offset int) *SelectQueryBuilder {
	if offset < 0 {
		offset = 0
	}

	q.offset = &offset

	return q
}

func (q *SelectQueryBuilder) Limit(limit int) *SelectQueryBuilder {
	if limit <= 0 {
		limit = defaultPageSize
	}

	q.limit = &limit

	return q
}

func (q *SelectQueryBuilder) ToSQL() (SelectQuerySQL, error) {
	selectPart, selectArgs := q.prepareSelectPart()
	fromPart, fromArgs := q.prepareFromPart()
	wherePart, whereArgs, err := q.prepareWherePart()
	if err != nil {
		return SelectQuerySQL{}, err
	}

	groupByPart := q.prepareGroupByPart()
	orderByPart := q.prepareOrderByPart()

	countSql := q.prepareSql("count(*)", fromPart, wherePart, groupByPart, orderByPart)
	selectSql := q.prepareSql(selectPart, fromPart, wherePart, groupByPart, orderByPart)

	return SelectQuerySQL{
		SelectSQL:  selectSql,
		SelectArgs: utils.MergeLists(selectArgs, fromArgs, whereArgs),
		CountSQL:   countSql,
		CountArgs:  utils.MergeLists(fromArgs, whereArgs),
	}, nil
}

func (q *SelectQueryBuilder) prepareSql(
	selectPart string, fromPart string, wherePart string,
	groupByPart string, orderByPart string,
) string {
	selectSql := strings.Builder{}
	selectSql.WriteString("select ")
	selectSql.WriteString(selectPart)

	if len(fromPart) > 0 {
		selectSql.WriteString(" from ")
		selectSql.WriteString(fromPart)
	}

	if len(wherePart) > 0 {
		selectSql.WriteString(" where ")
		selectSql.WriteString(wherePart)
	}

	if len(groupByPart) > 0 {
		selectSql.WriteString(" group by ")
		selectSql.WriteString(groupByPart)
	}

	if len(orderByPart) > 0 {
		selectSql.WriteString(" order by ")
		selectSql.WriteString(orderByPart)
	}

	return selectSql.String()
}

func (q *SelectQueryBuilder) prepareOrderByPart() string {
	orderByPart := strings.Builder{}
	for i, o := range q.orderBys {
		sql := o.generateSQL()
		orderByPart.WriteString(sql)
		if i < len(q.orderBys)-1 {
			orderByPart.WriteString(", ")
		}
	}

	return orderByPart.String()
}

func (q *SelectQueryBuilder) prepareGroupByPart() string {
	groupByPart := strings.Builder{}
	for i, g := range q.groupBys {
		sql := g.generateSQL()
		groupByPart.WriteString(sql)
		if i < len(q.groupBys)-1 {
			groupByPart.WriteString(", ")
		}
	}

	return groupByPart.String()
}

func (q *SelectQueryBuilder) prepareWherePart() (string, utils.AnyList, error) {
	wherePart := strings.Builder{}
	whereArgs := utils.NewAnyList()
	for i, w := range q.wheres {
		sql, args, err := w.generateSQL()
		if err != nil {
			return "", nil, err
		}

		wherePart.WriteString(sql)
		if i < len(q.wheres)-1 {
			wherePart.WriteString(" ")
		}

		if args != nil && args.Len() > 0 {
			whereArgs.AddList(args)
		}
	}

	return wherePart.String(), whereArgs, nil
}

func (q *SelectQueryBuilder) prepareFromPart() (string, utils.AnyList) {
	fromPart := strings.Builder{}
	fromArgs := utils.NewAnyList()
	for i, f := range q.froms {
		sql, args := f.generateSQL()
		fromPart.WriteString(sql)
		if i < len(q.froms)-1 {
			fromPart.WriteString(", ")
		}

		if args != nil && args.Len() > 0 {
			fromArgs.AddList(args)
		}
	}

	return fromPart.String(), fromArgs
}

func (q *SelectQueryBuilder) prepareSelectPart() (string, utils.AnyList) {
	selectPart := strings.Builder{}
	selectArgs := utils.NewAnyList()

	if q.selects.IsEmpty() {
		selectPart.WriteString("*")
	} else {
		for i, s := range q.selects {
			sql, args := s.generateSQL()
			selectPart.WriteString(sql)
			if i < len(q.selects)-1 {
				selectPart.WriteString(", ")
			}

			if args != nil && args.Len() > 0 {
				selectArgs.AddList(args)
			}
		}
	}

	return selectPart.String(), selectArgs
}

type SelectQuerySQL struct {
	SelectSQL  string
	CountSQL   string
	SelectArgs []any
	CountArgs  []any
}
