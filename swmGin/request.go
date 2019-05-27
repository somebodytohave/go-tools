package swmGin

import (
	"github.com/sun-wenming/go-tools/logging"
)

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	logging.GetLogger().Error(v...)
	return
}
