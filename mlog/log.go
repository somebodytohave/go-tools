package mlog

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mfile"
	"log"
)

var (
	Logger  *logrus.Logger
	InitLog bool
)

// 初始化日志配置
func Setup() {
	// Create a new instance of the Logger. You can have any number of instances.
	Logger = logrus.New()
	InitLog = true

	var err error
	//You could set this to any `io.Writer` such as a file
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err := mfile.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatal(err)
	}

	// 输出到控制台
	//Logger.SetOutput(os.Stdout)
	// 输出到文件中
	Logger.SetOutput(f)

	// TODO 增加 example_custom_caller_test https://github.com/sirupsen/logrus/blob/master/example_custom_caller_test.go

	// If you wish to add the calling method as a field
	Logger.SetReportCaller(true)

	//Logger.Formatter = new(logrus.JSONFormatter)
	//Logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// GetLogger Logger
func GetLogger() *logrus.Logger {
	return Logger
}

//- debug：没问题，就看看堆栈
func Debugln(args ...interface{}) {
	Logger.Debugln(args)
}

//- Info：提示一切正常
func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}

//- Warn：记录一下，某事又发生了
func Warnln(args ...interface{}) {
	Logger.Warnln(args)
}

//- Error：跟遇到的用户说对不起，可能有bug
func Errorln(args ...interface{}) {
	Logger.Errorln(args)

}

//- Fatal：网站挂了，或者极度不正常
func Fatalln(args ...interface{}) {
	Logger.Fatalln(args)
}

//func setPrefix(level Level) {
//	_, mfile, line, ok := runtime.Caller(DefaultCallerDepth)
//	if ok {
//		logPrefix = fmt.Sprintf("[%s:%d]", filepath.Base(mfile), line)
//	} else {
//		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
//	}
//	log.SetPrefix(logPrefix)
//}
