package dal

import (
	"context"
	"errors"
	"strconv"
)

type Query struct {
	orm *Orm
	ctx *context.Context
	from string
	where []string
	fields []string
	orderBy []string
	limit int
	isReady bool
	Result any
	Error error
}

func InitQuery (o *Orm, ctx *context.Context) *Query {
	return &Query{
		orm: o,
		ctx: ctx,
		isReady: false,
		Result: nil,
		Error: nil,
	}
}

func (q *Query) From(tableName string) *Query {
	q.from = tableName
	q.isReady = true
	return q
}

func (q *Query) Where(condition string) *Query {
	q.where = append(q.where, condition)
	return q
}

func (q *Query) Columns(fields []string) *Query {
	q.fields = append(q.fields, fields...)
	return q
}

func (q *Query) OrderBy(orderBy string) *Query {
	q.orderBy = append(q.orderBy, orderBy)
	return q
}

func (q *Query) Limit(limit int) *Query {
	q.limit = limit
	return q
}

func (q *Query) Select() *Query {
	if !q.isReady {
		q.Error = errors.New("запрос не готов к выполнению")
		return q
	}

	query := "SELECT "
	if len(q.fields) <= 0 {
		query += "* "
	} else {
		for i, f := range q.fields {
			query += f
			if i < len(q.fields)-1 {
				query += ", "
			}
		}
	}

	if len(q.from) <= 0 {
		q.Error = errors.New("отсутствует указание имени таблицы для запроса")
		return q
	}

	query += " FROM " + q.from

	if len(q.where) > 0 {
		query += " WHERE "
		for i, w := range q.where {
			if i > 0 {
				query += " AND "
			}
			query += w
		}
	}

	if len(q.orderBy) > 0 {
		query += " ORDER BY  "
		for i, w := range q.orderBy {
			if i > 0 {
				query += ", "
			}
			query += w
		}
	}

	if q.limit > 0 {
		query += " LIMIT " + strconv.Itoa(q.limit)
	}

	var result any
	var err error

	if len(q.fields) == 1 {
		result, err = q.orm.getOneColumn(query, q.ctx)
	} else {
		result, err = q.orm.getManyColumns(query, q.ctx)
	}

	if err != nil {
		q.Result = nil
		q.Error = err
	} else {
		q.Error = nil
		q.Result = result
	}

	return q
}

