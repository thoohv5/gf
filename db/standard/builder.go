package standard

import "context"

// builder标准
type IBuilder interface {
	// 添加
	Add(ctx context.Context, value interface{}) error
	// 更新
	Update(ctx context.Context, attrs ...interface{}) error
	// 删除
	Delete(ctx context.Context, value interface{}, where ...interface{}) error
	// 查询单条
	Find(ctx context.Context, out interface{}, where ...interface{}) error
	// 查询多条
	Get(ctx context.Context, out interface{}, where ...interface{}) error
	// 统计
	Count(ctx context.Context, value interface{}) error
	// 查询且条件
	Where(query interface{}, args ...interface{}) IBuilder
	// 查询或条件
	Or(query interface{}, args ...interface{}) IBuilder
	// 查询定位
	Offset(offset interface{}) IBuilder
	// 查询区间
	Limit(limit interface{}) IBuilder
	// 排序
	Order(value interface{}, reorder ...bool) IBuilder
	// 事务开始
	Begin(ctx context.Context) (IBuilder, error)
	// 事务回滚
	Rollback() error
	// 事务提交
	Commit() error
	// 模型
	Model(value interface{}) IBuilder
	// 检查数据是否为空
	IsEmpty(e error) bool
	// 检查是否为写库
	IsWrite() bool
	// 检查是否在事物中
	IsStartTx() bool
	// 日志
	SetLogger(log Logger) IBuilder
	// Exec
	Exec(ctx context.Context, sql string, values ...interface{}) error
	// Query
	Query(ctx context.Context, dest interface{}, sql string, values ...interface{}) error
}

// 日志标准
type Logger interface {
	Print(v ...interface{})
}
