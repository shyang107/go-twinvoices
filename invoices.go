package invoices

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unsafe"

	"github.com/cpmech/gosl/chk"
)

var (
	// Verbose activates display of messages on console
	Verbose = true

	// ColorsOn activates use of colors on console
	ColorsOn = true
)

// math ---------------------------------------------------------

// imax returns the maximum value
func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// imax returns the minimum value
func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isum returns the summation
func isum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

// string ---------------------------------------------------------

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

// BytesSizeToString convert bytes to a human-readable size
func BytesSizeToString(byteCount int) string {
	suf := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"} //Longs run out around EB
	if byteCount == 0 {
		return "0" + suf[0]
	}
	bytes := math.Abs(float64(byteCount))
	place := int32(math.Floor(math.Log2(bytes) / 10))
	num := bytes / math.Pow(1024.0, float64(place))
	var strnum string
	if place == 0 {
		strnum = fmt.Sprintf("%.0f", num) + suf[place]
	} else {
		strnum = fmt.Sprintf("%.1f", num) + suf[place]
	}
	return strnum
}

// ConvertBytesToString convert []byte to string
func ConvertBytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// GetColStr return string use in field
func GetColStr(s string, size int, isleft bool) string {
	_, _, n := CountChars(s)
	spaces := strings.Repeat(" ", size-n)
	// size := nc*2 + ne // s 實際佔位數
	var tab string
	if isleft {
		tab = fmt.Sprintf("%[1]s%[2]s", s, spaces)
	} else {
		tab = fmt.Sprintf("%[2]s%[1]s", s, spaces)
	}
	return " " + tab
}

// CountChars returns the number of each other of chinses and english characters
func CountChars(str string) (nc, ne, n int) {
	for _, r := range str {
		lchar := len(string(r))
		// n += lchar
		if lchar > 1 {
			nc++
		} else {
			ne++
		}
	}
	n = 2*nc + ne
	return nc, ne, n
}

// IsChineseChar judges whether the chinese character exists ?
func IsChineseChar(str string) bool {
	// n := 0
	for _, r := range str {
		// io.Pf("%q ", r)
		if unicode.Is(unicode.Scripts["Han"], r) {
			// n++
			return true
		}
	}
	return false
}

// ArgsTable prints a nice table with input arguments
//  Input:
//   title -- title of table; e.g. INPUT ARGUMENTS
//   data  -- sets of THREE items in the following order:
//                 description, key, value, ...
//                 description, key, value, ...
//                      ...
//                 description, key, value, ...
func ArgsTable(title string, data ...interface{}) string {
	heads := []string{"description", "key", "value"}
	return ArgsTableN(title, 0, heads, data...)
}

// ArgsTableN prints a nice table with input arguments
//  Input:
//   title -- title of table; e.g. INPUT ARGUMENTS
//	 heads -- heads of table; e.g. []string{ col1,  col2, ... }
//	 nledsp -- length of leading spaces in every row
//   data  -- sets of THREE items in the following order:
//                 column1, column2, column3, ...
//                 column1, column2, column3, ...
//                      ...
//                 column1, column2, column3, ...
func ArgsTableN(title string, nledsp int, heads []string, data ...interface{}) string {
	Sf := fmt.Sprintf
	nf := len(heads)
	ndat := len(data)
	if ndat < nf {
		return ""
	}
	if nledsp < 0 {
		nledsp = 0
	}
	lspaces := StrSpaces(nledsp)
	nlines := ndat / nf
	sizes := make([]int, nf)
	for i := 0; i < nf; i++ {
		_, _, sizes[i] = CountChars(heads[i])
	}
	for i := 0; i < nlines; i++ {
		if i*nf+(nf-1) >= ndat {
			return Sf("ArgsTable: input arguments are not a multiple of %d\n", nf)
		}
		for j := 0; j < nf; j++ {
			str := Sf("%v", data[i*nf+j])
			_, _, nmix := CountChars(str)
			sizes[j] = imax(sizes[j], nmix)
		}
	}
	// strfmt := Sf("%%v  %%v  %%v\n")
	n := isum(sizes...) + nf + (nf-1)*2 + 1 // sizes[0] + sizes[1] + sizes[2] + 3 + 4
	_, _, l := CountChars(title)
	m := (n - l) / 2
	//
	var b bytes.Buffer
	bw := b.WriteString
	//
	bw(StrSpaces(m+nledsp) + title + "\n")
	bw(lspaces + StrThickLine(n))
	isleft := true
	sfields := make([]string, nf)
	for i := 0; i < nf; i++ {
		sfields[i] = GetColStr(heads[i], sizes[i], isleft)
		switch i {
		case 0:
			bw(Sf("%v", lspaces+sfields[i]))
		default:
			bw(Sf("  %v", sfields[i]))
		}
	}
	bw("\n")
	bw(lspaces + StrThinLine(n))
	for i := 0; i < nlines; i++ {
		for j := 0; j < nf; j++ {
			sfields[j] = GetColStr(Sf("%v", data[i*nf+j]), sizes[j], isleft)
			switch j {
			case 0:
				bw(Sf("%v", lspaces+sfields[j]))
			default:
				bw(Sf("  %v", sfields[j]))
			}
		}
		bw("\n")
	}
	bw(lspaces + StrThickLine(n))
	return b.String()
}

// StrThickLine returns a thick line (using '=')
func StrThickLine(n int) (l string) {
	l = strings.Repeat("=", n)
	return l + "\n"
}

// StrThinLine returns a thin line (using '-')
func StrThinLine(n int) (l string) {
	l = strings.Repeat("-", n)
	return l + "\n"
}

// StrSpaces returns a line with spaces
func StrSpaces(n int) (l string) {
	l = strings.Repeat(" ", n)
	return
}

// file ---------------------------------------------------------

// isFileExist checks whether a file exist
func isFileExist(filename string) bool {
	path := os.ExpandEnv(filename)
	isExist := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		isExist = false
	}
	return isExist
}

// convert ---------------------------------------------------------

// Atob converts string to bool
func Atob(val string) (bres bool) {
	if strings.ToLower(val) == "true" {
		return true
	}
	if strings.ToLower(val) == "false" {
		return false
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		chk.Panic("cannot parse string representing integer: %s", val)
	}
	if res != 0 {
		bres = true
	}
	return
}

// Atoi converts string to integer
func Atoi(val string) (res int) {
	res, err := strconv.Atoi(val)
	if err != nil {
		chk.Panic("cannot parse string representing integer number: %s", val)
	}
	return
}

// Atof converts string to float64
func Atof(val string) (res float64) {
	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		chk.Panic("cannot parse string representing float number: %s", val)
	}
	return
}

// Itob converts from integer to bool
//  Note: only zero returns false
//        anything else returns true
func Itob(val int) bool {
	if val == 0 {
		return false
	}
	return true
}

// Btoi converts flag to interger
//  Note: true  => 1
//        false => 0
func Btoi(flag bool) int {
	if flag {
		return 1
	}
	return 0
}

// Btoa converts flag to string
//  Note: true  => "true"
//        false => "false"
func Btoa(flag bool) string {
	if flag {
		return "true"
	}
	return "false"
}

// IntSf is the Sprintf for a slice of integers (without brackets)
func IntSf(msg string, slice []int) string {
	return strings.Trim(fmt.Sprintf(msg, slice), "[]")
}

// DblSf is the Sprintf for a slice of float64 (without brackets)
func DblSf(msg string, slice []float64) string {
	return strings.Trim(fmt.Sprintf(msg, slice), "[]")
}

// StrSf is the Sprintf for a slice of string (without brackets)
func StrSf(msg string, slice []string) string {
	return strings.Trim(fmt.Sprintf(msg, slice), "[]")
}

// Sf wraps Sprintf
func Sf(msg string, prm ...interface{}) string {
	return fmt.Sprintf(msg, prm...)
}

// Ff wraps Fprintf
func Ff(b *bytes.Buffer, msg string, prm ...interface{}) {
	fmt.Fprintf(b, msg, prm...)
}

// print ---------------------------------------------------------

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
