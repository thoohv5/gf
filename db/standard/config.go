package standard

// 数据库日志配置
type Log struct {
	Mode int    // 数据库日志：0-无日志, 1-写日志, 2-读写日志
	Cat  string // 日志类别
}

// 数据库配置
type Config struct {
	Driver          string // 数据库驱动
	Source          string // 数据库源
	ConnMaxLifeTime int    // 数据库最大连接时长
	MaxIdleConns    int
	MaxOpenConns    int
	Slave           []struct{ Source string } // 数据库备用源
	Log             Log                       // 数据库日志
}
