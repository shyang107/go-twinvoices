package cmd

import (
	"fmt"
	"os"
	"sort"

	homedir "github.com/mitchellh/go-homedir"
	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

var (
	// utilverbose = &util.Verbose
	// print
	pfstart = util.PfCyan
	pfstop  = util.PfBlue
	pfsep   = util.Pfdyel2
	plog    = util.Pf
	pwarn   = util.Pfred
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
	util.PfBlue("> root.Execute called\n")
	// initConfig()
	sort.Sort(cli.FlagsByName(RootApp.Flags))
	sort.Sort(cli.CommandsByName(RootApp.Commands))
	if err := RootApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//
	if err := execute(); err != nil {
		pwarn(err.Error())
		os.Exit(-1)
	}
}

func rootAction(c *cli.Context) error {
	util.PfBlue("root.rootAction called\n")
	//
	if c.GlobalBool("verbose") {
		vp.Cfg.Verbose = c.GlobalBool("verbose")
		util.Verbose = vp.Cfg.Verbose
	}
	//
	fln := os.ExpandEnv(c.GlobalString("case"))
	if len(fln) > 0 {
		if !util.IsFileExist(fln) {
			pwarn("The specified case-configuration-file %q does not exist!\n", fln)
			os.Exit(-1)
		}
		vp.Cfg.CaseFilename = fln
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig(c *cli.Context) error {
	util.PfBlue(">> root.initConfig called\n")
	if err := vp.Cfg.ReadConfigs(); err != nil {
		return err
	}
	return nil
}

func execute() (err error) {
	plog("%v\n", vp.Cfg)
	home, err := homedir.Dir()
	pchk("home:%v\n", home)
	fln := vp.Cfg.CaseFilename
	vp.Cases, err = vp.Cfg.ReadCaseConfigs(fln)
	if err != nil {
		return err
	}
	//
	vp.Connectdb()
	//
	// var fbs = make([]*FileBunker, 0)
	// for _, o := range ol.List {
	for i := 0; i < len(vp.Cases); i++ {
		c := vp.Cases[i]
		plog("%s", c)
		//
		if err := c.UpdateFileBunker(); err != nil {
			return err
		}
		//
		pvs, err := (&c.Input).ReadInvoices()
		if err != nil {
			perr("%v\n", err)
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
