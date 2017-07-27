package cmd

import (
	"fmt"
	"os"

	"github.com/koding/multiconfig"
	inv "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/spf13/cobra"
)

// var cfgFile string
var name string
var age int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "invoices",
	Short: "Handling the invoices from the E-Invoice platform",
	Long: `The invoices mailed from the E-Invoice platform of Ministry of Finance, R.O.C. (Taiwan)
is encoded by Big-5 of Chinese character encoding method. Unfortunately, most OS and 
applocation use utf-8 encoding. This command can transfer a original Big-5 .csv file 
to other file types using utf-8 encoding; these file types include .json, .xml, .xlsx, 
or .yaml.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) < 2 {
		// 	cmd.Help()
		// 	return
		// }
		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	util.PfBlue("> Execute\n")
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	util.PfBlue("init\n")
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra_demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	// RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	util.PfBlue(">> initConfig\n")
	m := multiconfig.NewWithPath(inv.CfgFile)
	// Populated the serverConf struct
	cfg := new(inv.Config)
	err := m.Load(cfg) // Check for error
	if err != nil {
		util.Pfred(">>> initConfig: err\n")
		inv.Cfg = inv.GetDefualtConfig()
	} else {
		util.PfBlue(">>> initConfig: set cfg\n")
		inv.Cfg = cfg
	}
	//
	util.Verbose = inv.Cfg.Verbose
	// m.MustLoad(inv.Cfg) // Panic's if there is any error
	util.Pfgreen2("%v\n", inv.Cfg)
}
