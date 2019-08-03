package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/go-tools/mlog"
)


func main() {
	mlog.Warnln(logrus.GetLevel())
}
