package grpc

type Error struct {
	code    Code
	message string
	detail  string
}

func NewError(code Code, message string) *Error {
	return &Error{
		code:    code,
		message: message,
		detail:  "",
	}
}

type Code = string

const (
	UnknownErrCode  = Code(1000)
	InvalidArgument = Code(1001)
)

var mapErrMsg = map[Code]string{
	UnknownErrCode:  "未知错误",
	InvalidArgument: "无效参数",
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Error() string {
	msg := e.message
	if len(msg) > 0 {
		return msg
	}
	mErrMsg, ok := mapErrMsg[e.code]
	if ok {
		return mErrMsg
	}
	return "未知错误"
}

func (e *Error) Detail() string {
	return e.detail
}
