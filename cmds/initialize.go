package cmds

import (
	"os"
	"strings"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// InitialCommand represents the init command
func InitialCommand() cli.Command {
	return cli.Command{
		Name:        "initial",
		Aliases:     []string{"i"},
		Usage:       "Empty and initialize the using database",
		Description: `Empty the database, and initializes database`,
		Action:      initialAction,
	}
}

func initialAction(c *cli.Context) error {
	level := strings.ToLower(c.GlobalString("verbose")) // check command line options: "verbose"
	// util.Glog.Debugf("log level: %s\n", level)
	setglog(level)

	util.DebugPrintCaller()

	if err := vp.Initialdb(); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
