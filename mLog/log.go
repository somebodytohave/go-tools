package mLog

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mFile"
	"log"
)

var (
	logger  *logrus.Logger
	InitLog bool
)

// 初始化日志配置
func Setup() {
	// Create a new instance of the logger. You can have any number of instances.
	logger = logrus.New()
	InitLog = true
	
	var err error
	//You could set this to any `io.Writer` such as a file
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err := mFile.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatal(err)
	}

	// 输出到控制台
	//logger.SetOutput(os.Stdout)
	// 输出到文件中
	logger.SetOutput(f)

	// If you wish to add the calling method as a field
	logger.SetReportCaller(true)

	//logger.Formatter = new(logrus.JSONFormatter)
	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// GetLogger Logger
func GetLogger() *logrus.Logger {
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

//func setPrefix(level Level) {
//	_, mFile, line, ok := runtime.Caller(DefaultCallerDepth)
//	if ok {
//		logPrefix = fmt.Sprintf("[%s:%d]", filepath.Base(mFile), line)
//	} else {
//		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
//	}
//	log.SetPrefix(logPrefix)
//}
