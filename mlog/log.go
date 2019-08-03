package mlog

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mfile"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	logger *logrus.Logger
)

// 初始化日志配置
func Setup() {
	// Create a new instance of the logger. You can have any number of instances.
	logger = logrus.New()
	var err error
	//You could set this to any `io.Writer` such as a file
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err := mfile.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 输出到文件中
	logger.SetOutput(f)
	// 输出到控制台
	logger.SetOutput(os.Stdout)

	// 获取调用 日志的具体位置
	logger.SetReportCaller(true)
	// 被调用者的 方法名称 与 行数
	formatter := &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			filename := f.File + ", func: " + funcname + "~" + strconv.Itoa(f.Line)
			return "", filename
		},
	}
	logger.Formatter = formatter

	// TODO 增加 example_custom_caller_test https://github.com/sirupsen/logrus/blob/master/example_custom_caller_test.go

	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// GetLogger logger
func GetLogger() *logrus.Logger {
	if logger == nil {
		Setup()
	}
	return logger
}

//- debug：没问题，就看看堆栈
func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

//- Info：提示一切正常
func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

//- Warn：记录一下，某事又发生了
func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

//- Error：跟遇到的用户说对不起，可能有bug
func Errorln(args ...interface{}) {
	logger.Errorln(args)

}

//- Fatal：网站挂了，或者极度不正常
func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

// If you wish to add the calling method as a field
// 获取调用 日志的具体位置
func SetReportCaller(reportCaller bool) {
	logger.SetReportCaller(reportCaller)
}

