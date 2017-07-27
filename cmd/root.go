package cmd

import (
	"fmt"
	"os"
	"sort"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

var (
	utilverbose = &util.Verbose
	// print
	pfstart = util.PfCyan
	pfstop  = util.PfBlue
	pfsep   = util.Pfdyel2
	plog    = util.Pf
	pwarn   = util.Pforan
	perr    = util.Pflmag
	prun    = util.PfYel
	pchk    = util.Pfgreen2
	pstat   = util.Pfyel
)

// RootApp represents the base command when called without any subcommands
var RootApp = cli.NewApp()

func init() {
	// util.Verbose = vp.Cfg.Verbose
	util.PfBlue("root.init called\n")
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
	util.PfBlue("> root.Execute called\n")
	initConfig()
	sort.Sort(cli.FlagsByName(RootApp.Flags))
	sort.Sort(cli.CommandsByName(RootApp.Commands))
	if err := RootApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	plog("%v\n", vp.Cfg)
}

func rootAction(c *cli.Context) error {
	util.PfBlue("root.rootAction called\n")
	//
	if c.GlobalBool("verbose") {
		vp.Cfg.Verbose = c.GlobalBool("verbose")
		util.Verbose = vp.Cfg.Verbose
	}
	//
	fln := c.GlobalString("case")
	if len(fln) > 0 {
		if !util.IsFileExist(fln) {
			util.Panic("The specified config-file %q of case does not exist!", fln)
		}
		vp.Cfg.CaseFilename = fln
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	util.PfBlue(">> root.initConfig called\n")
	if err := vp.Cfg.ReadConfigs(); err != nil {
		util.Panic("%v\n", err.Error())
	}
}
