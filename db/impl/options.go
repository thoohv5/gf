package impl

// 可选参数列表
type options struct {
	isWrite bool
	txCount uint32
}

// 为可选参数赋值的函数
type ServerOption func(*options)

// 是否为写库
func WithIsWrite(isWrite bool) ServerOption {
	return func(o *options) {
		o.isWrite = isWrite
	}
}

// 事物统计
func WithTxCount(count uint32) ServerOption {
	return func(o *options) {
		o.txCount = count
	}
}
