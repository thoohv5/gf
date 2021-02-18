package standard

// 连接标准
type Connecter interface {
	// 连接
	Connect(config *Config) (*Wrapper, error)
}
