package cmd

import (
	"fmt"
	"os"
	"sort"

	vp "github.com/shyang107/go-twinvoices"
	ut "github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// RootApp represents the base command when called without any subcommands
var RootApp = cli.NewApp()

func init() {
	// ut.Verbose = vp.Cfg.Verbose
	// ut.Pdebug("root.init called\n")
	ut.Glog.Debugf("* (cmd.root.init) %q called by %q", ut.CallerName(1), ut.CallerName(2))
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
	RootApp.Before = initConfig
	RootApp.Action = rootAction
	RootApp.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose,b",
			Usage: "verbose output",
		},
		cli.StringFlag{
			Name:  "case,c",
			Usage: "specify the case file",
		},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootApp.
func Execute() {
	// ut.Pdebug("> root.Execute called\n")
	ut.Glog.Debugf("* %q called by %q", ut.CallerName(1), ut.CallerName(2))
	// initConfig()
	sort.Sort(cli.FlagsByName(RootApp.Flags))
	sort.Sort(cli.CommandsByName(RootApp.Commands))
	if err := RootApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//
	if err := execute(); err != nil {
		// ut.Pwarn(err.Error())
		ut.Glog.Error(err.Error())
		os.Exit(-1)
	}
}

func rootAction(c *cli.Context) error {
	// ut.Pdebug("root.rootAction called\n")
	ut.Glog.Debugf("* %q called by %q", ut.CallerName(1), ut.CallerName(2))
	//
	if c.GlobalBool("verbose") {
		vp.Cfg.Verbose = c.GlobalBool("verbose")
		ut.Verbose = vp.Cfg.Verbose
	}
	//
	fln := os.ExpandEnv(c.GlobalString("case"))
	if len(fln) > 0 {
		if !ut.IsFileExist(fln) {
			// ut.Perr("The specified case-configuration-file %q does not exist!\n", fln)
			ut.Glog.Errorf("The specified case-configuration-file %q does not exist!", fln)
			os.Exit(-1)
		}
		vp.Cfg.CaseFilename = fln
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig(c *cli.Context) error {
	// ut.Pdebug(">> root.initConfig called\n")
	ut.Glog.Debugf("* %q called by %q", ut.CallerName(1), ut.CallerName(2))
	if err := vp.Cfg.ReadConfigs(); err != nil {
		return err
	}
	return nil
}

func execute() (err error) {
	ut.Glog.Debugf("* %q called by %q", ut.CallerName(1), ut.CallerName(2))
	// ut.Pinfo("%v\n", vp.Cfg)
	ut.Glog.Infof("\n%v", vp.Cfg)
	vp.Cases, err = vp.Cfg.ReadCaseConfigs(vp.Cfg.CaseFilename)
	if err != nil {
		return err
	}
	//
	vp.Connectdb()
	//
	for i := 0; i < len(vp.Cases); i++ {
		c := vp.Cases[i]
		// ut.Plog("%s", c)
		ut.Glog.Infof("\n%v", c)
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
