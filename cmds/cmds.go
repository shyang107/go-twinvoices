package cmds

import (
	"fmt"
	"os"
	"strings"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// Init uses to initial all program
func Init() {
	util.DebugPrintCaller()
	util.InitLogger()
	// ut.Verbose = vp.Cfg.Verbose
	// ut.Pdebug("root.init called\n")

	app := cli.NewApp()

	app.Name = "invoice"
	app.Version = vp.Version
	app.Authors = []cli.Author{
		{Name: "S.H. Yang", Email: "shyang107@gmail.com"},
	}

	app.Usage = "Handling the invoices from the E-Invoice platform"
	app.Description = `The invoices mailed from the E-Invoice platform of Ministry of Finance, R.O.C. (Taiwan)
	is encoded by Big-5 of Chinese character encoding method. Unfortunately, most OS and 
	applocation use utf-8 encoding. This command can transfer a original Big-5 .csv file 
	to other file types using utf-8 encoding; these file types include .json, .xml, .xlsx, 
	or .yaml.`

	app.Commands = []cli.Command{
		ExecuteCommand(),
		InitialCommand(),
		DumpCommand(),
		RecoveryCommand(),
		ListCommand(),
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "verbose,V",
			Usage: `verbose output of logging information (default log-level is "info") 
		logging-levle are "disable", "info", "warn", "error", and "debug"
		"disable" will disable printer
		"error" will print only errors
		"warn" will print errors and warnings
		"info" will print errors, warnings and infos
		"debug" will print on any level, errors, warnings, infos and debug messages`,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func checkVerbose(c *cli.Context) {
	level := strings.ToLower(c.GlobalString("verbose")) // check command line options: "verbose"
	// util.Glog.Debugf("log level: %s\n", level)
	setglog(level)
}

func setglog(level string) {
	if len(level) > 0 {
		vp.Cfg.Verbose, util.Verbose = true, true
		util.ColorsOn = vp.Cfg.ColorsOn
	}
	switch level {
	case "disable":
	case "fatal":
	case "error":
	case "warn":
	case "debug":
	// case "goro":
	default:
		level = "info"
	}
	util.Glog.SetLevel(level)
}
