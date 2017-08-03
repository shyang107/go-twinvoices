package main

import (
	"fmt"
	"image/color"
	"image/color/palette"
	"math"

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
	var plte color.Palette = palette.Plan9
	var cl1, cl2 color.Color
	s := "這是彩色文字測試！This is a test of colorful text!"
	cnames := colornames.Names
	num := len(cnames)
	for i := 0; i < num; i++ {
		name1, name2 := cnames[i], cnames[num-i-1]
		cl1, cl2 = util.ColorsSVG[name1], util.ColorsSVG[name2]
		idx1 := plte.Index(cl1)
		idx2 := plte.Index(cl2)
		// cl1rgb :=
		// r, g, b, _ := util.ColorsSVG[name1].RGBA()
		// _, idx8bit2 := convert24To8bitColor(r, g, b)
		cl3 := color.RGBAModel.Convert(cl1)
		idx3 := plte.Index(cl3)
		//
		clfmt1, clfmt2, clfm8bit := fmt.Sprintf("[fg:%d <%s>]", idx1, name1),
			fmt.Sprintf("[fg:%d<%s>|bg:%d<%s>]", idx1, name1, idx2, name2),
			fmt.Sprintf("[fg:%v,%v]", idx1, idx3)
		_, ne1, _ := util.CountChars(clfmt1)
		_, ne2, _ := util.CountChars(clfmt2)
		m := util.Imax(ne1, ne2)

		fmt.Printf("%*s %s\n", m, clfmt1, util.ColorRGBString("%s", util.ColorsSVG[name1], s))
		c := util.NewRGB(
			util.ColorAttribute{IsForeground: true, RGB: util.ColorsSVG[name1]},
			util.ColorAttribute{IsForeground: false, RGB: util.ColorsSVG[name2]},
		)

		fmt.Printf("%*s \x1b[38;5;%dm%s\x1b[0m \n", m, clfm8bit, idx3, s)
		cstr := c.Sprintf("%s", s)

		// fmt.Println(cstr)
		fmt.Printf("%*s %s\n\n", m, clfmt2, cstr)
	}

}

func convert24To8bitColor(r, g, b uint32) (idx1, idx2 int) {
	// 8bit Color = (Red * 7 / 255) << 5 + (Green * 7 / 255) << 2 + (Blue * 3 / 255)
	//  (Math.floor((red / 32)) << 5) + (Math.floor((green / 32)) << 2) + Math.floor((blue / 64));
	// (r*6/256)*36 + (g*6/256)*6 + (b*6/256)
	// idx1 = int((r*7/255)<<5 + (g*7/255)<<2 + (b * 3 / 255))
	// var rr, gg, bb byte
	rr, gg, bb := (r * 8 / 256), (g * 8 / 256), (b * 4 / 256)
	idx1 = int((rr << 5) | (gg << 2) | bb)
	red, green, blue := float64(r), float64(g), float64(b)
	idx2 = int(math.Floor(red/32))<<5 + int(math.Floor(green/32))<<2 + int(math.Floor(blue/64))
	return
}
