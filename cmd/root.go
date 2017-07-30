package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/kataras/golog"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// RootApp represents the base command when called without any subcommands
var RootApp = cli.NewApp()

var (
	Glog     = util.Glog
	glInfo   = Glog.Info
	glInfof  = Glog.Infof
	glWarn   = Glog.Warn
	glWarnf  = Glog.Warnf
	glError  = Glog.Error
	glErrorf = Glog.Errorf
	glDebug  = Glog.Debug
	glDebugf = Glog.Debugf
)

// Rlog is a new instance of logger of logrus
// var Rlog = logrus.New()

func init() {
	// ut.Verbose = vp.Cfg.Verbose
	// ut.Pdebug("root.init called\n")
	util.DebugPrintCaller()
	RootApp.Version = vp.Version
	RootApp.Authors = []cli.Author{
		{Name: "S.H. Yang", Email: "shyang107@gmail.com"},
	}
	RootApp.Usage = "Handling the invoices from the E-Invoice platform"
	RootApp.Description = `The invoices mailed from the E-Invoice platform of Ministry of Finance, R.O.C. (Taiwan)
is encoded by Big-5 of Chinese character encoding method. Unfortunately, most OS and 
applocation use utf-8 encoding. This command can transfer a original Big-5 .csv file 
to other file types using utf-8 encoding; these file types include .json, .xml, .xlsx, 
or .yaml.`
	// RootApp.Before = initConfig
	RootApp.Action = rootAction
	RootApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "case,c",
			Usage: "specify the case file",
		},
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
	///
	// Default Output is `os.Stderr`, but you can change it:
	// glSetOutput(os.Stdout)
	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	golog.NewLine("\n")
	Glog.SetTimeFormat("")
	Glog.SetLevel("info")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootApp.
func Execute() {
	// ut.Pdebug("> root.Execute called\n")
	util.DebugPrintCaller()
	// initConfig()
	sort.Sort(cli.FlagsByName(RootApp.Flags))
	sort.Sort(cli.CommandsByName(RootApp.Commands))
	//
	// RootApp.RunAndExitOnError()
	if err := RootApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//
}

func rootAction(c *cli.Context) error {
	// ut.Pdebug("root.rootAction called\n")
	util.DebugPrintCaller()
	//
	if err := vp.Cfg.ReadConfigs(); err != nil {
		return err
	}
	// cli.ShowCommandHelpAndExit(c)
	if c.GlobalBool("verbose") {
		vp.Cfg.Verbose = c.GlobalBool("verbose")
		util.Verbose = vp.Cfg.Verbose
	}
	//
	fln := os.ExpandEnv(c.GlobalString("case"))
	if len(fln) > 0 {
		if !util.IsFileExist(fln) {
			// ut.Perr("The specified case-configuration-file %q does not exist!\n", fln)
			glErrorf("The specified case-configuration-file %q does not exist!", fln)
			os.Exit(-1)
		}
		vp.Cfg.CaseFilename = fln
	}
	//
	level := strings.ToLower(c.String("verbose"))
	if len(level) > 0 {
		vp.Cfg.Verbose, util.Verbose = true, true
		util.ColorsOn = vp.Cfg.ColorsOn
	}
	switch level {
	case "warn":
	case "error":
	case "debug":
	case "disable":
	default:
		level = "info"
	}
	Glog.SetLevel(level)
	//
	if err := execute(); err != nil {
		// ut.Pwarn(err.Error())
		glError(err.Error())
		os.Exit(-1)
	}
	return nil
}

func execute() (err error) {
	util.DebugPrintCaller()
	// ut.Pinfo("%v\n", vp.Cfg)
	glInfof("\n%v", vp.Cfg)
	vp.Cases, err = vp.Cfg.ReadCaseConfigs()
	if err != nil {
		return err
	}
	//
	vp.Connectdb()
	//
	for i := 0; i < len(vp.Cases); i++ {
		c := vp.Cases[i]
		// ut.Plog("%s", c)
		glInfof("\n%v", c)
		//
		if err := c.UpdateFileBunker(); err != nil {
			return err
		}
		//
		pvs, err := (&c.Input).ReadInvoices()
		if err != nil {
			// ut.Perr("%v\n", err)
			// ut.Glog.Errorf("%v\n", err)
			return err
		}
		for j := 0; j < len(c.Outputs); j++ {
			out := c.Outputs[j]
			if out.IsOutput {
				err = out.WriteInvoices(pvs)
				if err != nil {
					return err
				}
			}
		}
	}
	// pchk(GetFileBunkerTable(fbs, 0))
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig(c *cli.Context) error {
	// ut.Pdebug(">> root.initConfig called\n")
	util.DebugPrintCaller()
	if err := vp.Cfg.ReadConfigs(); err != nil {
		return err
	}
	return nil
}
