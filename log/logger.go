package log

import (
	"github.com/thoohv5/gf/log/impl"
	"github.com/thoohv5/gf/log/standard"
)

var _defaultLogger standard.ILogger

func init() {
	_defaultLogger = impl.NewZLogger()
}

func New(logType Type) standard.ILogger {
	_defaultLogger = get(logType)
	return _defaultLogger
}

type Type int

const (
	Log Type = 1
	Zap Type = 2
)

func get(logType Type) standard.ILogger {
	var logger standard.ILogger
	switch logType {
	case Log:
		logger = impl.NewLogger()
	case Zap:
		logger = impl.NewZLogger()
	default:
		logger = impl.NewZLogger()
	}
	return logger
}

func WithField(key string, val interface{}) standard.Field {
	return standard.NewField(key, val)
}

func Info(msg string, fields ...standard.Field) {
	_defaultLogger.Info(msg, fields...)
}
func Error(msg string, fields ...standard.Field) {
	_defaultLogger.Error(msg, fields...)
}
