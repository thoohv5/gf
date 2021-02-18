package impl

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/thoohv5/gf/log/standard"
)

type zLogger struct {
	*zap.Logger
}

func NewZLogger() standard.ILogger {

	hook := lumberjack.Logger{
		Filename:   "./logs/log.log", // 日志文件路径
		MaxSize:    128,              // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,               // 日志文件最多保存多少个备份
		MaxAge:     7,                // 文件最多保存多少天
		Compress:   true,             // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	caller := zap.AddCaller()

	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	// logger := zap.New(core, caller, zap.Development(), filed)
	logger := zap.New(core, caller, zap.Development())

	return &zLogger{
		Logger: logger,
	}
}

func (l *zLogger) Info(msg string, fields ...standard.Field) {
	l.Logger.Info(msg, l.deal(fields...)...)
}

func (l *zLogger) Error(msg string, fields ...standard.Field) {
	l.Logger.Error(msg, l.deal(fields...)...)
}

func (l *zLogger) deal(fields ...standard.Field) []zap.Field {
	list := make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		list = append(list, zap.Any(field.GetKey(), field.GetVal()))
	}
	return list
}
