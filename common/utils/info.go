package utils

import (
	"fmt"
	"runtime"
)

// 获取嵌套的上层函数信息，通过调用栈, 0表示当前函数
func GetUpFuncInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "Reward/utils/info.go:66"
	}

	fc := runtime.FuncForPC(pc).Name()

	return fmt.Sprintf("%s:%d:%s", file, line, fc)
}
