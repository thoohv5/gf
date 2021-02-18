package db

import (
	"context"

	"github.com/thoohv5/gf/db/standard"
)

// 数据dao标准
type IDao interface {
	// Tabler
	// Connecter
	IQuery
	// IFilter
	// 添加
	Add(ctx context.Context, value Tabler) error
	// 更新
	Update(ctx context.Context, param IQuery, update map[string]interface{}) error
	// 删除
	Delete(ctx context.Context, param IQuery) error
	// 查询单条
	Find(ctx context.Context, param IQuery, result Tabler) error
	// 查询多条
	Get(ctx context.Context, param IQuery, result interface{}) error
	// 查询
	Where(query interface{}, args ...interface{}) standard.IBuilder
	// get builder
	GetBuilder() standard.IBuilder
	// set builder
	SetBuilder(build standard.IBuilder) IDao
}

// 数据库连接配置常量
type Connecter interface {
	Connection() string
}

// 数据库表标准
type Tabler interface {
	TableName() string
}

// 公共查询
type IQuery interface {
	GetCommonReq() *CommonReq
}

type CommonReq struct {
	Start int32 `json:"start,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	// 排序：sort=otc_type,-created_at,*custom
	// 以符号开头，可选符号：(+或空 正序）（- 倒序）（* 自定义复杂排序标识关键词）
	Sorts []string `json:"sorts,omitempty"`
}

// 数据库过滤标准
type IFilter interface {
	BuildFilterQuery(filter IQuery) standard.IBuilder
	Filter(build standard.IBuilder, condition *CommonReq) standard.IBuilder
}
