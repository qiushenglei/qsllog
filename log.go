package qsllog

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Level string

const DEBUG Level = "DEBUG"
const INFO Level = "INFO"
const WARNING Level = "WARNING"
const ERROR Level = "ERROR"

type log interface {
	AddLog(level string, data ...interface{})
}

type MyLog struct {
	fileName string
	file     *os.File
	level    Level
	path     string
	date     string
}

// NewMyLog 创建实例
func NewMyLog(level Level, path string) *MyLog {
	mylog := &MyLog{
		level: level,
		path:  path,
	}
	mylog.touchFile()
	return mylog
}

// AddLog 追加日志
func (l *MyLog) AddLog(data string) error {
	// 判断日志是否是当天的
	l.judgeTodayInstance()

	// 追加日志
	l.append(data)
	return nil
}

func (l *MyLog) append(data string) {
	w := bufio.NewWriter(l.file)
	_, err := w.WriteString(data + " \n")

	//l.file.WriteString(data + "\n")
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 打开文件
func (l *MyLog) touchFile() error {
	// 获取文件名
	l.fileName = l.getFileName()

	// 文件路径
	filePath := l.path + l.fileName

	// 判断文件是否存在
	_, err := os.Stat(filePath)
	if err != nil && os.IsExist(err) {
		panic(err)
	}

	var file *os.File
	if os.IsNotExist(err) {
		// 创建文件
		file, err = os.Create(l.fileName)
		if err != nil {
			panic(err)
		}
	} else {
		// 打开文件，追加可读写模式
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
	l.file = file

	return err
}

// 获取文件名
func (l *MyLog) getFileName() string {
	l.date = l.getTodayTimeFormat()
	return fmt.Sprintf("%s_%s", l.level, l.date)
}

// 获取当天时间
func (l *MyLog) getTodayTimeFormat() string {
	return time.Now().Format("2006-01-02")
}

// 判断是否是今天的实例
func (l *MyLog) judgeTodayInstance() {
	if l.date != l.getTodayTimeFormat() {
		beforFile := l.file
		// 新建日志日志
		if err := l.touchFile(); err != nil {
			panic(err)
		}
		// 关闭前一天的日志文件
		beforFile.Close()
	}
}
