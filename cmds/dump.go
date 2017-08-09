package cmds

import (
	"os"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

var dfile string

// dumpCmd represents the dump command
var dumpCmd = cli.Command{
	Name:        "dump",
	Aliases:     []string{"d"},
	Usage:       "Dump all records from database",
	Description: "Dump all recirds from database into .json file.",
	Action:      dumpAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file,f",
			Usage: "specify the dump path",
		},
	},
}

func init() {
	util.DebugPrintCaller()

	util.Verbose = vp.Cfg.Verbose
	util.ColorsOn = vp.Cfg.ColorsOn

	RootApp.Commands = append(RootApp.Commands, dumpCmd)
}

func dumpAction(c *cli.Context) error {
	util.DebugPrintCaller()

	vp.Cfg.IsDump = true
	dfn := c.String("file")
	if len(dfn) > 0 {
		vp.Cfg.DumpFilename = os.ExpandEnv(dfn)
	}

	vp.Connectdb()

	vp.DBDumpData(vp.Cfg.DumpFilename)
	os.Exit(0)
	return nil
}
