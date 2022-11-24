package qsllog

import (
	"encoding/json"
	"runtime"
	"time"
)

type Logger struct {
	log *MyLog
}

func NewLogger(level Level, path string) *Logger {
	return &Logger{
		log: NewMyLog(level, path),
	}
}

func NewUniqueNum() string {
	return time.Now().String()
}

func (l *Logger) AddLog(msg string, err error, data ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	runtime.ReadTrace()

	info := map[string]interface{}{
		"file": file,
		"line": line,
		"time": time.Now().Format("2006-01-02 15:04:06"),
		"msg":  msg,
		"err":  err.Error(),
	}

	for k, v := range data {
		info[string(k)] = v
	}

	if str, err := json.Marshal(info); err == nil {
		l.log.AddLog(string(str))
	}
}
