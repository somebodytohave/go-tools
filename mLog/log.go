package mLog

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mFile"
	"log"
)

var (
	logger *logrus.Logger
)

func Setup() {
	// Create a new instance of the logger. You can have any number of instances.
	logger = logrus.New()
	var err error
	//You could set this to any `io.Writer` such as a mFile
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err := mFile.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatal(err)
	}

	//logger.SetOutput(os.Stdout)
	// 输出到文件中
	logger.SetOutput(f)

	// If you wish to add the calling method as a field
	logger.SetReportCaller(true)

	//logger.Formatter = new(logrus.JSONFormatter)

	//- Fatal：网站挂了，或者极度不正常
	//- Error：跟遇到的用户说对不起，可能有bug
	//- Warn：记录一下，某事又发生了
	//- Info：提示一切正常
	//- debug：没问题，就看看堆栈

	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// GetLogger Logger
func GetLogger() *logrus.Logger {
	return logger
}

func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args)

}

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
