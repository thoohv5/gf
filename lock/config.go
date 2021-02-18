package lock

// 配置
type Config struct {
	// 类型: redis
	Type string
	// 配置，根据类型对应不同的类型
	Config interface{}
}
