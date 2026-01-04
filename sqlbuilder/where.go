package sqlbuilder

import (
	"fmt"
	"strings"

	"github.com/ismaildurmaz/segsql/utils"
)

type where interface {
	generateSQL() (string, utils.AnyList, error)
}

type WhereList = utils.GenericArray[where]

type WhereBuilder struct {
	where

	owner *SelectQueryBuilder
	items WhereList
}

func (w *WhereBuilder) ToSelectQuery() *SelectQueryBuilder {
	return w.owner
}

func (w *WhereBuilder) Eq(expr string, param interface{}) *WhereJoinBuilder {
	w.addWheres(newSimpleWhere(expr, "=", param))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) NotEq(expr string, param interface{}) *WhereJoinBuilder {
	w.addWheres(newSimpleWhere(expr, "!=", param))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) Like(expr string, param interface{}) *WhereJoinBuilder {
	w.addWheres(newSimpleWhere(expr, "like", param))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) Between(expr string, param1 interface{}, param2 interface{}) *WhereJoinBuilder {
	w.addWheres(newBetweenWhere(expr, param1, param2, true))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) NotBetween(expr string, param1 interface{}, param2 interface{}) *WhereJoinBuilder {
	w.addWheres(newBetweenWhere(expr, param1, param2, false))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) In(expr string, params ...interface{}) *WhereJoinBuilder {
	w.addWheres(newInWhere(expr, params, false))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) NotIn(expr string, params ...interface{}) *WhereJoinBuilder {
	w.addWheres(newInWhere(expr, params, true))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) IsNull(expr string) *WhereJoinBuilder {
	w.addWheres(newSimpleWhere(expr, "is null", nil))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) IsNotNull(expr string) *WhereJoinBuilder {
	w.addWheres(newSimpleWhere(expr, "is not null", nil))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) Expression(expr string, params ...interface{}) *WhereJoinBuilder {
	w.addWheres(newExpressionWhere(expr, params))

	return w.newWhereJoinBuilder()
}

func (w *WhereBuilder) StartGroup() *WhereBuilder {
	w.addWheres(startGroup{})

	return w
}

func (w *WhereBuilder) EndGroup() *WhereBuilder {
	w.addWheres(endGroup{})

	return w
}

func (w *WhereBuilder) generateSQL() (string, utils.AnyList, error) {
	args, sqls, s, list, err := w.prepareItems()
	if err != nil {
		return s, list, err
	}

	sb := strings.Builder{}
	for i, sql := range sqls {
		sb.WriteString(sql)
		if i < len(sqls)-1 {
			if sql != "(" && sqls[i+1] != ")" {
				sb.WriteString(" ")
			}
		}
	}

	return sb.String(), args, nil
}

func (w *WhereBuilder) prepareItems() (utils.AnyList, []string, string, utils.AnyList, error) {
	args := utils.NewAnyList()

	parenthesisCount := 0

	sqls := make([]string, w.items.Len())
	for i, child := range w.items {
		sql, sargs, err := child.generateSQL()
		if err != nil {
			return nil, nil, "", nil, err
		}

		sqls[i] = sql

		if sql == "(" {
			parenthesisCount++
		} else if sql == ")" {
			parenthesisCount--
		}

		if parenthesisCount < 0 {
			return nil, nil, "", nil, fmt.Errorf("invalid where sql, too many closing parentheses")
		}

		if sargs != nil && sargs.Len() > 0 {
			args.AddList(sargs)
		}
	}

	if parenthesisCount > 0 {
		return nil, nil, "", nil, fmt.Errorf("invalid where sql, too many opening parentheses")
	}

	return args, sqls, "", nil, nil
}

func newWhereList() WhereList {
	return WhereList{}
}

func newWhereBuilder(owner *SelectQueryBuilder) *WhereBuilder {
	return &WhereBuilder{owner: owner, items: newWhereList()}
}

func (w *WhereBuilder) addWheres(child ...where) {
	w.items.AddList(child)
}

func (w *WhereBuilder) newWhereJoinBuilder() *WhereJoinBuilder {
	return &WhereJoinBuilder{owner: w}
}

type WhereJoinBuilder struct {
	owner *WhereBuilder
	op    string
}

func (w *WhereJoinBuilder) And() *WhereBuilder {
	w.owner.items.Add(joinWhere{op: "and"})

	return w.owner
}

func (w *WhereJoinBuilder) Or() *WhereBuilder {
	w.owner.items.Add(joinWhere{op: "or"})

	return w.owner
}

func (w *WhereJoinBuilder) StartGroup() *WhereBuilder {
	w.owner.items.Add(startGroup{})

	return w.owner
}

func (w *WhereJoinBuilder) EndGroup() *WhereBuilder {
	w.owner.items.Add(endGroup{})

	return w.owner
}

func (w *WhereJoinBuilder) ToSelectQuery() *SelectQueryBuilder {
	return w.owner.owner
}

type simpleWhere struct {
	where

	expr  string
	op    string
	param interface{}
}

func newSimpleWhere(expr string, op string, param interface{}) *simpleWhere {
	return &simpleWhere{expr: expr, op: op, param: param}
}

func (w *simpleWhere) generateSQL() (string, utils.AnyList, error) {
	mutList := utils.NewAnyList()

	if w.op != "is null" && w.op != "is not null" {
		mutList.Add(w.param)
	}

	return w.expr + " " + w.op + " ?", mutList, nil
}

type betweenWhere struct {
	where

	expr   string
	param1 interface{}
	param2 interface{}
	not    bool
}

func newBetweenWhere(expr string, param1 interface{}, param2 interface{}, not bool) *betweenWhere {
	return &betweenWhere{expr: expr, param1: param1, param2: param2, not: not}
}

func (w *betweenWhere) generateSQL() (string, utils.AnyList, error) {
	sql := w.expr + " "
	if w.not {
		sql += "not "
	}
	sql += "between ? and ?"
	arr := utils.NewAnyList()
	arr.AddItems(w.param1, w.param2)

	return sql, arr, nil
}

type inWhere struct {
	where

	expr   string
	params []interface{}
	not    bool
}

func newInWhere(expr string, params []interface{}, not bool) *inWhere {
	return &inWhere{expr: expr, params: params, not: not}
}

func (i *inWhere) generateSQL() (string, utils.AnyList, error) {
	p := make([]string, len(i.params))

	if len(i.params) == 0 {
		return "", nil, fmt.Errorf("in where params cannot be empty")
	}

	for range i.params {
		p = append(p, "?")
	}

	sql := i.expr + " "
	if i.not {
		sql += "not "
	}
	sql += "in (" + strings.Join(p, ", ") + ")"

	arr := utils.NewAnyList()
	arr.AddItems(i.params...)

	return sql, arr, nil
}

type joinWhere struct {
	where

	op string
}

func (j joinWhere) generateSQL() (string, utils.AnyList, error) {
	return j.op, nil, nil
}

type startGroup struct {
	where
}

func (s startGroup) generateSQL() (string, utils.AnyList, error) {
	return "(", nil, nil
}

type endGroup struct {
	where
}

func (e endGroup) generateSQL() (string, utils.AnyList, error) {
	return ")", nil, nil
}

type expressionWhere struct {
	where

	expr   string
	params []interface{}
}

func (e expressionWhere) generateSQL() (string, utils.AnyList, error) {
	return e.expr, e.params, nil
}

func newExpressionWhere(expr string, params []interface{}) *expressionWhere {
	return &expressionWhere{expr: expr, params: params}
}
