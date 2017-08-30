package cmds

import (
	"fmt"
	"os"
	"strings"

	vp "github.com/shyang107/go-twinvoices"
	util "github.com/shyang107/gout"
	"github.com/urfave/cli"
)

// RecoveryCommand represents the init command
func RecoveryCommand() cli.Command {
	return cli.Command{
		Name:        "recovery",
		Aliases:     []string{"r"},
		Usage:       "Empty and recovery the table of database",
		Description: `Empty the tables of inveoices and details, and recover them from filebunker`,
		Action:      recoveryAction,
	}
}

func droptables() {
	vp.Connectdb()

	if vp.DB.HasTable(&vp.Invoice{}) && vp.DB.HasTable(&vp.Detail{}) {
		util.Glog.Debugf("♲  Drop the tables %q and %q ...",
			vp.Invoice{}.TableName(), vp.Detail{}.TableName())
		vp.DB.DropTable(&vp.Invoice{}, &vp.Detail{})
	}

	util.Glog.Debugf("♲  Recreate the tables %q and %q ...",
		vp.Invoice{}.TableName(), vp.Detail{}.TableName())
	vp.DB.AutoMigrate(&vp.Invoice{}, &vp.Detail{})
	vp.DB.Model(&vp.Invoice{}).Related(&vp.Detail{}, "uin")
}

func recoveryAction(c *cli.Context) error {
	checkVerbose(c)

	util.DebugPrintCaller()

	droptables()

	fbc := vp.GetAllOriginalDataFromDB()

	ch := make(chan string)
	for _, f := range *fbc {

		go func(f *vp.FileBunker) {
			util.Glog.Infof("\n%s", f.GetArgsTable("", 0))
			// util.Glog.Debugf("%v\n", string(f.Contents))

			var vslice vp.InvoiceCollection
			var dslice vp.DetailCollection
			var cb util.ReadLinesCallback = func(idx int, line string) (stop bool) {
				// if f.Encoding == "big-5" {
				// line = vp.Big5ToUtf8(line)
				// }
				line = vp.Big5ToUtf8(line)
				recs := strings.Split(line, "|")
				head := recs[0]
				switch head {
				case "M": // invoice
					pinv := vp.UnmarshalCSVInvoice(recs)
					vslice.Add(pinv)
				case "D": // deltail of invoice
					pdet := vp.UnmarshalCSVDetail(recs)
					dslice.Add(pdet)
				}
				return
			}
			util.ReadLines(f.Contents, cb)

			vslice.Combine(&dslice)
			// util.Glog.Infof("♲  Invoices list:\n%s", vslice.GetInvoicesTable())
			vslice.AddToDB()

			ch <- fmt.Sprintf("* ch: ending %q", f.Name)
		}(f)
	}

	for range *fbc {
		fmt.Println(<-ch)
		// util.Glog.Info(<-ch)
	}

	os.Exit(0)
	return nil
}
