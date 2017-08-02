package main

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/shyang107/go-twinvoices/util"
)

func main() {
	testcolortext()
}
func testcolortext() {
	// i := 0
	// for k, c := range util.Colors {
	// 	i++
	// 	fmt.Printf("%3d.  %v : %+v\n", i, k, c)
	// }
	s := "這是彩色文字測試！This is a test of colorful text!"
	cnames := colornames.Names
	num := len(cnames)
	for i, n := range cnames {
		fmt.Printf("%3d <%s> %s\n", i+1, n, util.ColorRGBString("%s", util.ColorsSVG[n], s))
		c := util.NewRGB(
			util.ColorAttribute{IsForeground: true, RGB: util.ColorsSVG[n]},
			util.ColorAttribute{IsForeground: false, RGB: util.ColorsSVG[cnames[num-i-1]]},
		)
		cstr := c.Sprintf("%s", s)
		fmt.Println(cstr)
	}
}
