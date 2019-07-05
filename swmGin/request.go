package swmGin

import (
	"github.com/sun-wenming/go-tools/mlog"
)

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	mlog.GetLogger().Error(v...)
	return
}
