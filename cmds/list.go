package cmds

import (
	"os"
	"strings"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// ListCommand represents the list command to list invoices of database
func ListCommand() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases:     []string{"l"},
		Usage:       "List all invoices.",
		Description: "Gets all records of database and lists them.",
		Action:      listAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "format,f",
				Usage: `list format; 
	formats: (defualt is "simple")
			"simple" (or "s") "brief" (or "b"), "pretty(or "p")"`,
			},
		},
	}
}

func listAction(c *cli.Context) error {
	checkVerbose(c)

	util.DebugPrintCaller()

	vp.Connectdb()

	a, err := vp.DBGetAllInvoices()
	if err != nil {
		return err
	}

	format := strings.ToLower(c.String("format"))
	switch format {
	case "pretty", "p":
		util.Glog.Print(a.GetInvoicesTable())
	case "brief", "b":
		util.Glog.Print(a.GetList())
	default: // "simple","s"
		util.Glog.Print(a)
	}

	os.Exit(0)
	return nil
}
