package cmds

import (
	"os"
	"strings"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// DumpCommand represents the dump command
func DumpCommand() cli.Command {
	return cli.Command{
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
}

func dumpAction(c *cli.Context) error {
	level := strings.ToLower(c.GlobalString("verbose")) // check command line options: "verbose"
	// util.Glog.Debugf("log level: %s\n", level)
	setglog(level)

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
