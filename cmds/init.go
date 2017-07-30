package cmds

import (
	"os"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// initCmd represents the init command
var initCmd = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Empty and initialize the using database",
	Description: `Empty the database, and initializes database`,
	Action:      initAction,
}

func init() {
	// ut.Pdebug("init.init called\n")
	// ut.Pdebug("init.init called\n")
	util.DebugPrintCaller()
	RootApp.Commands = append(RootApp.Commands, initCmd)
}

func initAction(c *cli.Context) error {
	// ut.Pdebug("init.initAction called\n")
	util.DebugPrintCaller()
	if err := vp.Initialdb(); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
