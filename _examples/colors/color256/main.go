package main

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/shyang107/go-twinvoices/util"
)

func main() {
	// testcolor256code()
	// testcolortext()
	testcolortext2()
}
func testcolortext2() {
	// i := 0
	// for k, c := range util.Colors {
	// 	i++
	// 	fmt.Printf("%3d.  %v : %+v\n", i, k, c)
	// }
	s := "這是彩色文字測試！This is a test of colorful text!"
	cnames := colornames.Names
	for i, n := range cnames {
		fmt.Printf("%3d <%s> %s\n", i+1, n, util.ColorRGBString("%s", util.Colors[n], s))
	}
}

func testcolortext() {
	s := "這是彩色文字測試！This is a test of colorful text!"
	fmt.Printf("\n%v\n", "*** 256 foreground colors ***")
	for i := 0; i < 256; i++ {
		fgcode := util.Attribute(i) << 8
		code := util.DecodeColor256(fgcode)
		fmt.Printf("[fg:%3d] %s\n", code, util.Color256String("%s", fgcode, s))
	}
	fmt.Printf("\n%v\n", "*** 256 background colors ***")
	for i := 0; i < 256; i++ {
		bgcode := util.Attribute(i+256) << 8
		code := util.DecodeColor256(bgcode)
		fmt.Printf("[bg:%3d] %s\n", code, util.Color256String("%s", bgcode, s))
	}
	fmt.Printf("\n%v\n", "*** 256 foreground and background colors ***")
	for i := 0; i < 256; i++ {
		fgcode := util.Attribute(255-i) << 8
		fcode := util.DecodeColor256(fgcode)
		bgcode := util.Attribute(i+256) << 8
		bcode := util.DecodeColor256(bgcode) - 256
		c := util.New256(fgcode, bgcode)
		fmt.Printf("[fg:%3d][bg:%3d] %s\n", fcode, bcode, c.Sprintf("%s", s))
	}
}

func testcolor256code() {
	type Attribute int
	// Foreground Standard colors: n, where n is from the color table (0-7)
	// (as in ESC[30–37m) <- SGR code
	const (
		FgBlack256 Attribute = iota << 8
		FgRed256
		FgGreen256
		FgYellow256
		FgBlue256
		FgMagenta256
		FgCyan256
		FgWhite256
	)

	// Foreground High-intensity colors: n, where n is from the color table (8-15)
	// (as in ESC [ 90–97 m) <- SGR code
	const (
		FgHiBlack256 Attribute = (iota + 8) << 8
		FgHiRed256
		FgHiGreen256
		FgHiYellow256
		FgHiBlue256
		FgHiMagenta256
		FgHiCyan256
		FgHiWhite256
	)

	// Foreground Grayscale colors: grayscale from black to white in 24 steps (232-255)
	const (
		FgGrayscale01 Attribute = (iota + 232) << 8
		FgGrayscale02
		FgGrayscale03
		FgGrayscale04
		FgGrayscale05
		FgGrayscale06
		FgGrayscale07
		FgGrayscale08
		FgGrayscale09
		FgGrayscale10
		FgGrayscale11
		FgGrayscale12
		FgGrayscale13
		FgGrayscale14
		FgGrayscale15
		FgGrayscale16
		FgGrayscale17
		FgGrayscale18
		FgGrayscale19
		FgGrayscale20
		FgGrayscale21
		FgGrayscale22
		FgGrayscale23
		FgGrayscale24
	)
	const bgzone = 256
	// Background Standard colors: n, where n is from the color table (0-7)
	// (as in ESC[30–37m) <- SGR code
	const (
		BgBlack256 Attribute = (iota + bgzone) << 8
		BgRed256
		BgGreen256
		BgYellow256
		BgBlue256
		BgMagenta256
		BgCyan256
		BgWhite256
	)

	// Background High-intensity colors: n, where n is from the color table (8-15)
	// (as in ESC [ 90–97 m) <- SGR code
	const (
		BgHiBlack256 Attribute = (iota + 8 + bgzone) << 8
		BgHiRed256
		BgHiGreen256
		BgHiYellow256
		BgHiBlue256
		BgHiMagenta256
		BgHiCyan256
		BgHiWhite256
	)

	// Background Grayscale colors: grayscale from black to white in 24 steps (232-255)
	const (
		BgGrayscale01 Attribute = (iota + 232 + bgzone) << 8
		BgGrayscale02
		BgGrayscale03
		BgGrayscale04
		BgGrayscale05
		BgGrayscale06
		BgGrayscale07
		BgGrayscale08
		BgGrayscale09
		BgGrayscale10
		BgGrayscale11
		BgGrayscale12
		BgGrayscale13
		BgGrayscale14
		BgGrayscale15
		BgGrayscale16
		BgGrayscale17
		BgGrayscale18
		BgGrayscale19
		BgGrayscale20
		BgGrayscale21
		BgGrayscale22
		BgGrayscale23
		BgGrayscale24
	)
	const one = 1
	fmt.Println()
	fmt.Printf("1 << 8 = %d\n", one<<8)
	fmt.Printf("(1 << 8)>>8 = %d\n", (one<<8)>>8)
	fmt.Println("*** 256 foreground Colors: Standard colors ***")
	fmt.Printf("%s = %d >> 8 = %v\n", "FgBlack256  ", FgBlack256, (FgBlack256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgRed256    ", FgRed256, (FgRed256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGreen256  ", FgGreen256, (FgGreen256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgYellow256 ", FgYellow256, (FgYellow256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgBlue256   ", FgBlue256, (FgBlue256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgMagenta256", FgMagenta256, (FgMagenta256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgCyan256   ", FgCyan256, (FgCyan256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgWhite256  ", FgWhite256, (FgWhite256)>>8)
	fmt.Println("*** 256 foreground Colors: High-intensity colors ***")
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiBlack256  ", FgHiBlack256, (FgHiBlack256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiRed256    ", FgHiRed256, (FgHiRed256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiGreen256  ", FgHiGreen256, (FgHiGreen256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiYellow256 ", FgHiYellow256, (FgHiYellow256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiBlue256   ", FgHiBlue256, (FgHiBlue256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiMagenta256", FgHiMagenta256, (FgHiMagenta256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiCyan256   ", FgHiCyan256, (FgHiCyan256)>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgHiWhite256  ", FgHiWhite256, (FgHiWhite256)>>8)
	fmt.Println("*** 256 foreground Colors: Grayscale ***")
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale01", FgGrayscale01, FgGrayscale01>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale02", FgGrayscale02, FgGrayscale02>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale03", FgGrayscale03, FgGrayscale03>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale04", FgGrayscale04, FgGrayscale04>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale05", FgGrayscale05, FgGrayscale05>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale06", FgGrayscale06, FgGrayscale06>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale07", FgGrayscale07, FgGrayscale07>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale08", FgGrayscale08, FgGrayscale08>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale09", FgGrayscale09, FgGrayscale09>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale10", FgGrayscale10, FgGrayscale10>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale11", FgGrayscale11, FgGrayscale11>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale12", FgGrayscale12, FgGrayscale12>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale13", FgGrayscale13, FgGrayscale13>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale14", FgGrayscale14, FgGrayscale14>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale15", FgGrayscale15, FgGrayscale15>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale16", FgGrayscale16, FgGrayscale16>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale17", FgGrayscale17, FgGrayscale17>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale18", FgGrayscale18, FgGrayscale18>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale19", FgGrayscale19, FgGrayscale19>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale20", FgGrayscale20, FgGrayscale20>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale21", FgGrayscale21, FgGrayscale21>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale22", FgGrayscale22, FgGrayscale22>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale23", FgGrayscale23, FgGrayscale23>>8)
	fmt.Printf("%s = %d >> 8 = %v\n", "FgGrayscale24", FgGrayscale24, FgGrayscale24>>8)

	fmt.Println("*** 256 background Colors: Standard colors ***")
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgBlack256  ", BgBlack256, BgBlack256>>8, BgBlack256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgRed256    ", BgRed256, BgRed256>>8, BgRed256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGreen256  ", BgGreen256, BgGreen256>>8, BgGreen256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgYellow256 ", BgYellow256, BgYellow256>>8, BgYellow256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgBlue256   ", BgBlue256, BgBlue256>>8, BgBlue256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgMagenta256", BgMagenta256, BgMagenta256>>8, BgMagenta256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgCyan256   ", BgCyan256, BgCyan256>>8, BgCyan256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgWhite256  ", BgWhite256, BgWhite256>>8, BgWhite256>>8-bgzone)
	fmt.Println("*** 256 background Colors: High-intensity colors ***")
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiBlack256  ", BgHiBlack256, BgHiBlack256>>8, BgHiBlack256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiRed256    ", BgHiRed256, BgHiRed256>>8, BgHiRed256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiGreen256  ", BgHiGreen256, BgHiGreen256>>8, BgHiGreen256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiYellow256 ", BgHiYellow256, BgHiYellow256>>8, BgHiYellow256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiBlue256   ", BgHiBlue256, BgHiBlue256>>8, BgHiBlue256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiMagenta256", BgHiMagenta256, BgHiMagenta256>>8, BgHiMagenta256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiCyan256   ", BgHiCyan256, BgHiCyan256>>8, BgHiCyan256>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgHiWhite256  ", BgHiWhite256, BgHiWhite256>>8, BgHiWhite256>>8-bgzone)
	fmt.Println("*** 256 background Colors: Grayscale ***")
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale01", BgGrayscale01, BgGrayscale01>>8, BgGrayscale01>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale02", BgGrayscale02, BgGrayscale02>>8, BgGrayscale02>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale03", BgGrayscale03, BgGrayscale03>>8, BgGrayscale03>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale04", BgGrayscale04, BgGrayscale04>>8, BgGrayscale04>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale05", BgGrayscale05, BgGrayscale05>>8, BgGrayscale05>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale06", BgGrayscale06, BgGrayscale06>>8, BgGrayscale06>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale07", BgGrayscale07, BgGrayscale07>>8, BgGrayscale07>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale08", BgGrayscale08, BgGrayscale08>>8, BgGrayscale08>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale09", BgGrayscale09, BgGrayscale09>>8, BgGrayscale09>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale10", BgGrayscale10, BgGrayscale10>>8, BgGrayscale10>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale11", BgGrayscale11, BgGrayscale11>>8, BgGrayscale11>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale12", BgGrayscale12, BgGrayscale12>>8, BgGrayscale12>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale13", BgGrayscale13, BgGrayscale13>>8, BgGrayscale13>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale14", BgGrayscale14, BgGrayscale14>>8, BgGrayscale14>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale15", BgGrayscale15, BgGrayscale15>>8, BgGrayscale15>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale16", BgGrayscale16, BgGrayscale16>>8, BgGrayscale16>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale17", BgGrayscale17, BgGrayscale17>>8, BgGrayscale17>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale18", BgGrayscale18, BgGrayscale18>>8, BgGrayscale18>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale19", BgGrayscale19, BgGrayscale19>>8, BgGrayscale19>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale20", BgGrayscale20, BgGrayscale20>>8, BgGrayscale20>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale21", BgGrayscale21, BgGrayscale21>>8, BgGrayscale21>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale22", BgGrayscale22, BgGrayscale22>>8, BgGrayscale22>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale23", BgGrayscale23, BgGrayscale23>>8, BgGrayscale23>>8-bgzone)
	fmt.Printf("%s = %d >> 8 = %d - 256 = %v\n", "BgGrayscale24", BgGrayscale24, BgGrayscale24>>8, BgGrayscale24>>8-bgzone)
}
