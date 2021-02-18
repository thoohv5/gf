package standard

type ILogger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

type Field struct {
	key string
	val interface{}
}

func NewField(key string, val interface{}) Field {
	return Field{
		key: key,
		val: val,
	}
}

func (f *Field) GetKey() string {
	return f.key
}

func (f *Field) GetVal() interface{} {
	return f.val
}
