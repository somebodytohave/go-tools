package mGin

import (
	"github.com/sun-wenming/go-tools/mLog"
)

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	if mLog.InitLog {
		mLog.Errorln(v...)
	}
	return
}
