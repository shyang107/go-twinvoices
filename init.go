package invoices

import (
	"fmt"
	"runtime"

	"github.com/shyang107/go-twinvoices/util"
)

const (
	//
	fcstart = 101
	fcstop  = 102
	fostart = 111
	fostop  = 112
	ffstart = 21
	ffstop  = 22
	csvSep  = "|"
)

var (
	format = map[int]string{
		// config
		fcstart: "# Start to configure. -- %q\n",
		fcstop:  "# Configuration has been concluded. -- %q\n",
		// option
		fostart: "# Start to get case-options. -- %q\n",
		fostop:  "# Case-options has been concluded. -- %q\n",
		// start/end function
		ffstart: "* Function %q start.\n",
		ffstop:  "* Function %q stop.\n",
	}

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

// misc ---------------------------------------------------------

func callerName(idx int) string {
	pc, _, _, _ := runtime.Caller(idx) //idx = 0 self, 1 for caller, 2 for upper caller
	return runtime.FuncForPC(pc).Name()
}

func startfunc(fid int) {
	pfstart(format[fid], callerName(2))
}

func stopfunc(fid int) {
	pfstop(format[fid], callerName(2))
	printSepline(60)
}

func printSepline(n int) {
	if n <= 0 {
		n = 60
	}
	pfsep("%s", util.StrThinLine(n))
}

// getErrMessage get error message if error
func getErrMessage(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

func checkErr(err error) {
	if err != nil {
		// perr(getErrMessage(err))
		panic(err)
	}
}
