package util

import (
	"fmt"
)

// code of string format
const (
	Fcstart = 101
	Fcstop  = 102
	Fostart = 111
	Fostop  = 112
	Ffstart = 21
	Ffstop  = 22
)

var (

	// Verbose activates display of messages on console
	Verbose = true

	// ColorsOn activates use of colors on console
	ColorsOn = true

	// Format is the format of message
	Format = map[int]string{
		// config
		Fcstart: "# Start to configure. -- %q\n",
		Fcstop:  "# Configuration has been concluded. -- %q\n",
		// option
		Fostart: "# Start to get case-options. -- %q\n",
		Fostop:  "# Case-options has been concluded. -- %q\n",
		// start/end function
		Ffstart: "* Function %q start.\n",
		Ffstop:  "* Function %q stop.\n",
	}
)

// specified function for output
var (
	Pfstart = PfCyan
	Pfstop  = PfBlue
	Pfsep   = Pfdyel2
	Prun    = PfYel
	Pchk    = Pfgreen2
	Pstat   = Pfyel
	Plog    = Pf
	Pinfo   = Pfcyan2
	Pwarn   = Pforan
	Perr    = Pfred
	Pdebug  = Pfgreen2
)

// print ---------------------------------------------------------

// PrintSepline print the separate-line
func PrintSepline(n int) {
	if n <= 0 {
		n = 60
	}
	Pfsep("%s", StrThinLine(n))
}

// PrintFormat commands ---------------------------------------------------------
// modified from "github.com/cpmech/gosl/io"

// Pl prints new line
func Pl() {
	if !Verbose {
		return
	}
	fmt.Println()
}

// low intensity

// Pf prints formatted string
func Pf(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	fmt.Printf(msg, prm...)
}

// Pfcyan prints formatted string in cyan
func Pfcyan(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;36m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfcyan2 prints formatted string in another shade of cyan
func Pfcyan2(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;50m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfyel prints formatted string in yellow
func Pfyel(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;33m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfdyel prints formatted string in dark yellow
func Pfdyel(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;58m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfdyel2 prints formatted string in another shade of dark yellow
func Pfdyel2(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;94m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfred prints formatted string in red
func Pfred(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;31m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfgreen prints formatted string in green
func Pfgreen(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;32m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfblue prints formatted string in blue
func Pfblue(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;34m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfmag prints formatted string in magenta
func Pfmag(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;35m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pflmag prints formatted string in light magenta
func Pflmag(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[0;95m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfpink prints formatted string in pink
func Pfpink(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;205m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfdgreen prints formatted string in dark green
func Pfdgreen(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;22m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfgreen2 prints formatted string in another shade of green
func Pfgreen2(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;2m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfpurple prints formatted string in purple
func Pfpurple(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;55m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfgrey prints formatted string in grey
func Pfgrey(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;59m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfblue2 prints formatted string in another shade of blue
func Pfblue2(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;69m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pfgrey2 prints formatted string in another shade of grey
func Pfgrey2(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;60m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// Pforan prints formatted string in orange
func Pforan(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[38;5;202m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// high intensity

// PfCyan prints formatted string in high intensity cyan
func PfCyan(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;36m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfYel prints formatted string in high intensity yello
func PfYel(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;33m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfRed prints formatted string in high intensity red
func PfRed(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;31m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfGreen prints formatted string in high intensity green
func PfGreen(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;32m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfBlue prints formatted string in high intensity blue
func PfBlue(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;34m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfMag prints formatted string in high intensity magenta
func PfMag(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;35m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}

// PfWhite prints formatted string in high intensity white
func PfWhite(msg string, prm ...interface{}) {
	if !Verbose {
		return
	}
	if ColorsOn {
		fmt.Printf("[1;37m"+msg+"[0m", prm...)
	} else {
		fmt.Printf(msg, prm...)
	}
}
