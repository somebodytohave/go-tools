package mGin

import (
	"github.com/sun-wenming/go-tools/mLog"
)

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	mLog.GetLogger().Error(v...)
	return
}
