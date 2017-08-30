package main

import (
	"time"

	"github.com/shyang107/go-twinvoices/cmds"
	util "github.com/shyang107/gout"
)

func init() {
	// util.EnableLoggerOutToFile("debug")
}

func main() {

	// os.Args = append(os.Args, "-V", "debug", "e")
	start := time.Now()

	cmds.Init()

	util.Glog.Println()
	util.Glog.Child("main()").Infof("run-time elapsed: %s\n", time.Since(start).String())
	// fmt.Println("run-time elapsed: ", time.Since(start).String())
}
