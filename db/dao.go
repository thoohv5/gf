package db

import (
	"context"

	"github.com/thoohv5/gf/db/standard"
)

var (
	// DBs map[DB_NAME]*gorm.DB
	DbWrapper map[string]*standard.Wrapper
)

type Dao struct {
	*CommonReq
	tabler  Tabler
	filter  IFilter
	wrapper *standard.Wrapper
}

// NewDao new
func NewDao(model IDao) *Dao {
	connect, ok := model.(Connecter)
	if !ok {
		return nil
	}
	dbWrapper := DbWrapper[connect.Connection()]
	table, ok := model.(Tabler)
	if !ok {
		return nil
	}
	filter, ok := model.(IFilter)
	if !ok {
		return nil
	}
	return &Dao{
		tabler:  table,
		filter:  filter,
		wrapper: dbWrapper,
	}
}

// register dao
func RegisterDao(model IDao) *Dao {
	return NewDao(model)
}

func (d *Dao) Add(ctx context.Context, value Tabler) error {
	return d.wrapper.Write().Add(ctx, value)
}

func (d *Dao) Update(ctx context.Context, param IQuery, update map[string]interface{}) error {
	return d.filter.BuildFilterQuery(d.wrapper.Write(), param).Update(ctx, update)
}

func (d *Dao) Delete(ctx context.Context, param IQuery) error {
	return d.filter.BuildFilterQuery(d.wrapper.Write(), param).Delete(ctx, d.tabler)
}

func (d *Dao) Find(ctx context.Context, param IQuery, result Tabler) error {
	return d.filter.Filter(d.filter.BuildFilterQuery(d.wrapper.Read(), param), param.GetCommonReq()).Find(ctx, result)
}

func (d *Dao) Get(ctx context.Context, param IQuery, result interface{}) error {
	return d.filter.Filter(d.filter.BuildFilterQuery(d.wrapper.Write(), param), param.GetCommonReq()).Get(ctx, result)
}

// func (d *Dao) Where(query interface{}, args ...interface{}) standard.IBuilder {
// 	return d.builder.Where(query, args...)
// }

func (d *Dao) GetCommonReq() *CommonReq {
	if nil == d {
		return &CommonReq{}
	}
	return d.CommonReq
}

// func (d *Dao) GetBuilder() standard.IBuilder {
// 	return d.builder
// }

// filter
func (d *Dao) Filter(build standard.IBuilder, condition *CommonReq) standard.IBuilder {
	if nil != condition {
		// Start
		if start := condition.Start; start > 0 {
			build = build.Offset(start)
		}
		// Limit
		if limit := condition.Limit; limit > 0 {
			build = build.Limit(limit)
		}
		// 排序
		for _, sort := range condition.Sorts {
			if string(sort[0]) == "+" {
				sort = string(sort[1:]) + " ASC"
			} else if string(sort[0]) == "-" {
				sort = string(sort[1:]) + " DESC"
			} else {
				sort = sort + " ASC"
			}
			build = build.Order(sort)
		}
	}

	return build
}
