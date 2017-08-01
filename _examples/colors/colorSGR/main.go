package main

import (
	"fmt"

	"github.com/shyang107/go-twinvoices/util"
)

func main() {
	testcolortext()
}

func testcolortext() {
	s := "這是彩色文字測試！This is a test of colorful text!"
	fmt.Printf("\n%v\n", "*** SGR standard colors ***")
	fmt.Printf("%v\n", util.RedString(s))
	fmt.Printf("%v\n", util.GreenString(s))
	fmt.Printf("%v\n", util.YellowString(s))
	fmt.Printf("%v\n", util.BlueString(s))
	fmt.Printf("%v\n", util.MagentaString(s))
	fmt.Printf("%v\n", util.CyanString(s))
	fmt.Printf("%v\n", util.WhiteString(s))
	fmt.Printf("\n%v\n", "*** SGR hi-intensity colors ***")
	fmt.Printf("%v\n", util.HiRedString(s))
	fmt.Printf("%v\n", util.HiGreenString(s))
	fmt.Printf("%v\n", util.HiYellowString(s))
	fmt.Printf("%v\n", util.HiBlueString(s))
	fmt.Printf("%v\n", util.HiMagentaString(s))
	fmt.Printf("%v\n", util.HiCyanString(s))
	fmt.Printf("%v\n", util.HiWhiteString(s))
}
