package cmds

import (
	"fmt"
	"os"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// ExecuteCommand return cli.Command
func ExecuteCommand() cli.Command {
	return cli.Command{
		Name:    "execute",
		Aliases: []string{"e"},
		Usage: `Import the original invoice data file from the E-Invoice platform of Ministry of Finance,
backup to invoices.db (sqlite3), and output specified file-type.`,
		Action: executeAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "case,c",
				Usage: "specify the case file",
			},
		},
	}
}

func executeAction(c *cli.Context) error {
	// level := strings.ToLower(c.GlobalString("verbose")) // check command line options: "verbose"
	// setglog(level)
	// util.Glog.Debugf("log level: %s\n", level)

	util.DebugPrintCaller()

	if err := vp.Cfg.ReadConfigs(); err != nil { // reading config
		return err
	}

	fln := os.ExpandEnv(c.String("case")) // check command line options: "case"
	if len(fln) > 0 {
		if !util.IsFileExist(fln) {
			// ut.Perr("The specified case-configuration-file %q does not exist!\n", fln)
			util.Glog.Errorf("The specified case-configuration-file %q does not exist!", fln)
			os.Exit(-1)
		}
		vp.Cfg.CaseFilename = fln
	}

	if err := execute(); err != nil { // run procedures of program
		// ut.Pwarn(err.Error())
		util.Glog.Error(err.Error())
		os.Exit(-1)
	}
	return nil
}

//
func execute() (err error) {
	util.DebugPrintCaller()

	util.Glog.Infof("\n%v", vp.Cfg) // print out config

	vp.Cases, err = vp.Cfg.ReadCaseConfigs() // reading settings of cases
	if err != nil {
		util.Glog.Error(err.Error())
		return err
	}

	vp.Connectdb() // connect to database

	// for _, c := range vp.Cases { // run every case
	// 	util.Glog.Infof("\n%v", c)

	// 	if err := c.UpdateFileBunker(); err != nil {
	// 		util.Glog.Error(err.Error())
	// 		return err
	// 	}

	// 	pvs, err := (&c.Input).ReadInvoices()
	// 	if err != nil {
	// 		util.Glog.Errorf("%v\n", err)
	// 		return err
	// 	}

	// 	for _, out := range c.Outputs {
	// 		if out.IsOutput {
	// 			err = out.WriteInvoices(pvs)
	// 			if err != nil {
	// 				util.Glog.Error(err.Error())
	// 				return err
	// 			}
	// 		}
	// 	}
	// }

	ch := make(chan string)
	for idx, c := range vp.Cases { // run every case

		go func(c *vp.Case, idx int) {

			util.Glog.Infof("\n%v", c)

			if err := c.UpdateFileBunker(); err != nil {
				util.Glog.Error(err.Error())
				// return err
			}

			pvs, err := (&c.Input).ReadInvoices()
			if err != nil {
				util.Glog.Errorf("%v\n", err)
				// return err
			}

			for _, out := range c.Outputs {
				if out.IsOutput {
					err = out.WriteInvoices(pvs)
					if err != nil {
						util.Glog.Error(err.Error())
						// return err
					}
				}
			}
			ch <- fmt.Sprintf("* ch: run case %d", idx+1)
		}(c, idx)
	}

	for range vp.Cases {
		fmt.Println(<-ch)
	}

	return nil
}

type chCase struct {
	invoices *vp.InvoiceCollection
	outs     []*vp.OutputFile
	idx      int
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
