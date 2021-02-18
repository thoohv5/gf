package impl

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/thoohv5/gf/log/standard"
)

type logger struct {
	*log.Logger
}

func NewLogger() standard.ILogger {
	file, err := os.OpenFile("logs/log.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	return &logger{
		Logger: log.New(io.MultiWriter(file, os.Stdout), "LOG: ", log.Ldate|log.Lmicroseconds|log.Llongfile),
	}
}

func (l *logger) Info(msg string, fields ...standard.Field) {
	l.Logger.Println(l.deal(msg, fields...)...)
}
func (l *logger) Error(msg string, fields ...standard.Field) {
	l.Logger.Fatalln(l.deal(msg, fields...)...)
}

func (l *logger) deal(msg string, fields ...standard.Field) []interface{} {
	var v []interface{}
	format := []string{"msg: %s"}
	formatValues := []interface{}{msg}
	for _, field := range fields {
		format = append(format, fmt.Sprintf("%s: %s", field.GetKey(), "%s"))
		formatValues = append(formatValues, field.GetVal())
	}
	v = append(v, fmt.Sprintf(strings.Join(format, ","), formatValues...))
	return v
}
