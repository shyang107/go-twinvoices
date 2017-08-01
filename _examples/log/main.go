package main

import (
	"github.com/kataras/golog"
	"github.com/shyang107/go-twinvoices/util"
)

func main() {
	testlog()
}

func testlog() {
	util.Glog.SetLevel("debug")
	var log golog.ExternalLogger
	log = util.Glog
	log.Info("info")
	log.Error("Error")
	log.Warn("Warn")
	log.Debug("Debug")
}
