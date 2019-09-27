package mlog

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mfile"
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
		logger.Errorln(err)
	}
	// 输出到文件中
	logger.SetOutput(f)
	// 输出到控制台
	//logger.SetOutput(os.Stdout)

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
	GetLogger().Debugln(args)
}

//- Info：提示一切正常
func Infoln(args ...interface{}) {
	GetLogger().Infoln(args)
}

//- Warn：记录一下，某事又发生了
func Warnln(args ...interface{}) {
	GetLogger().Warnln(args)
}

//- Error：跟遇到的用户说对不起，可能有bug
func Errorln(args ...interface{}) {
	GetLogger().Errorln(args)

}

//- Fatal：网站挂了，或者极度不正常
func Fatalln(args ...interface{}) {
	GetLogger().Fatalln(args)
}

//- debug：没问题，就看看堆栈
func DebuglnErr(err error) {
	if err != nil {
		GetLogger().Debugln(err)
	}
}

//- Info：提示一切正常
func InfolnErr(err error) {
	if err != nil {
		GetLogger().Infoln(err)
	}
}

//- Warn：记录一下，某事又发生了
func WarnlnErr(err error) {
	if err != nil {
		GetLogger().Warningln(err)
	}
}

//- Error：跟遇到的用户说对不起，可能有bug
func ErrorlnErr(err error) {
	if err != nil {
		GetLogger().Errorln(err)
	}
}

//- Fatal：网站挂了，或者极度不正常
func FatallnlErr(err error) {
	if err != nil {
		GetLogger().Fatalln(err)
	}
}

// If you wish to add the calling method as a field
// 获取调用 日志的具体位置
func SetReportCaller(reportCaller bool) {
	if reportCaller {
		// 被调用者的 方法名称 与 行数
		formatter := &logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				filename := f.File + ", func: " + funcName + "~" + strconv.Itoa(f.Line)
				return "", filename
			},
		}
		GetLogger().Formatter = formatter
	}
	GetLogger().SetReportCaller(reportCaller)
}

// getCaller retrieves the name of the first non-logrus calling function
// 获取调用者的文件、方法名、行数。
func GetCaller(skip int) string {
	pcs := make([]uintptr, 1)
	depth := runtime.Callers(1+skip, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	frame, _ := frames.Next()
	s := strings.Split(frame.Function, ".")
	funcName := s[len(s)-1]
	filename := frame.File + ", func: " + funcName + "~" + strconv.Itoa(frame.Line)
	return filename
}
