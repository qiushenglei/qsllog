package main

import (
	"errors"
	"github.com/qiushenglei/qsllog"
)

func main() {
	err := errors.New("is test error")
	logger := qsllog.NewLogger(qsllog.WARNING, "./")
	logger.AddLog("测试", err, []int{1, 2, 3}, []int{1, 2, 3})
}
