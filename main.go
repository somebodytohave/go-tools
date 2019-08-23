package main

import (
	"github.com/sun-wenming/go-tools/mlog"
	"os"
)

func main() {
	mlog.GetLogger().Out = os.Stdout
	mlog.Errorln("--")
	mlog.FatallnlErr(nil)
	mlog.Errorln("--")

}
