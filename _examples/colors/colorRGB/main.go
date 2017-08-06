package main

import (
	"fmt"
	"image/color"
	"image/color/palette"
	"math"
	"os"

	"golang.org/x/image/colornames"

	"github.com/shyang107/go-twinvoices/pencil"
	"github.com/shyang107/go-twinvoices/pencil/ansirgb"
	"github.com/shyang107/go-twinvoices/pencil/rgb16b"
	"github.com/shyang107/go-twinvoices/util"
	// "github.com/stroborobo/ansirgb"
	// sansirgb "github.com/stroborobo/ansirgb"
)

func main() {
	testcolortext()
	// printAnsiPalette()
}

// var paletteMap = make(map[string]color.RGBA)

func printAnsiPalette() {
	for idx, p := range ansirgb.Palette {
		// r, g, b, a := p.RGBA()
		fmt.Printf("%3d. %v -- %#02x\n", idx, p, uint8(ansirgb.Convert(p).Code))
	}
	os.Remove("ansiPalette.log")
	output, _ := os.OpenFile("ansiPalette.log", os.O_CREATE|os.O_WRONLY, 0666)
	// output := os.Stdout
	fmt.Fprintf(output, "var Palette color.Palette = []color.Color{\n")
	for _, p := range ansirgb.Palette {
		r, g, b, a := p.RGBA()
		// fmt.Fprintf(output, "- %4sansirgb.Color{Color:color.RGBA{%#02x,%#02x,%#02x,%#02x}, Code:%d},\n", "", uint8(r), uint8(b), uint8(g), uint8(a), ansirgb.Convert(p).Code)
		if ansirgb.Convert(p).Code == 16 {
			fmt.Fprintf(output, "%4s// 216 colors: 16-232; 6 * 6 * 6 = 216 colors\n", "")
		}
		if ansirgb.Convert(p).Code == 232 {
			fmt.Fprintf(output, "%4s// grayscale: 233-255\n", "")
		}
		if ansirgb.Convert(p).Code == -1 {
			fmt.Fprintf(output, "%4s// transparent: 256\n", "")
		}
		fmt.Fprintf(output, "%4s&Color{&color.RGBA{%#02x,%#02x,%#02x,%#02x}, %d},\n", "", r>>8, b>>8, g>>8, a>>8, ansirgb.Convert(p).Code)
	}
	fmt.Fprintf(output, "}\n")
	output.Close()

	maps := make(map[string]ansirgb.Color)
	cnames := colornames.Names
	num := len(cnames)
	size := 0
	for i := 0; i < num; i++ {
		rgb := colornames.Map[cnames[i]]
		cl := ansirgb.Convert(&rgb)
		maps[cnames[i]] = *cl
		_, _, n := util.CountChars(cnames[i])
		size = util.Imax(size, n+2)
	}
	// i := 0
	// for k, c := range maps {
	// 	i++
	// 	cl := ansirgb.Map[k]
	// 	// r, g, b, _ := cl.RGBA()
	// 	// s := fmt.Sprintf("%3d: \033[38;5;%dm%04X %04X %04X\033[0m", cl.Code, cl.Code, r, g, b)
	// 	fmt.Printf("%3d. %*q %v -- %*q %v\n", i, size, k, &c, size, k, cl.String())
	// }
	// fmt.Fprintf(os.Stdout, "\n%#v\n", maps)
	i := 0
	os.Remove("ansirgbMap.log")
	output, _ = os.OpenFile("ansirgbMap.log", os.O_CREATE|os.O_WRONLY, 0666)
	// output := os.Stdout
	fmt.Fprintf(output, "var Map = map[string]Color{\n")
	for k := range maps {
		_, _, n := util.CountChars(k)
		size = util.Imax(size, n)
	}
	size += 2
	for k, c := range maps {
		i++
		r, g, b, a := c.Color.RGBA()
		name := fmt.Sprintf("%q", k)
		name += ":"
		fmt.Fprintf(output, "%4s%-*sColor{color.RGBA{%#02x,%#02x,%#02x,%#02x}, %d},\n", "", size, name, r>>8, b>>8, g>>8, a>>8, c.Code)
	}
	fmt.Fprintf(output, "}\n")
	output.Close()
}

func testcolortext() {
	// i := 0
	// for k, c := range util.Colors {
	// 	i++
	// 	fmt.Printf("%3d.  %v : %+v\n", i, k, c)
	// }
	var plte color.Palette = palette.Plan9
	s := "這是彩色文字測試！This is a test of colorful text!"
	cnames := colornames.Names
	num := len(cnames)
	size := findMaxSize()
	for i := 0; i < num; i++ {
		name1, name2 := cnames[i], cnames[num-i-1]
		cl1, cl2 := colornames.Map[name1], colornames.Map[name2]
		cl3 := ansirgb.Convert(&cl1)
		idx1, idx2, idx3 := plte.Index(cl1), plte.Index(cl2), cl3.Code

		clfmt1, clfmt2, clfm8bit :=
			fmt.Sprintf("svg[fg:%d <%s>]", idx1, name1),
			fmt.Sprintf("svg[fg:%d<%s>|bg:%d<%s>]", idx1, name1, idx2, name2),
			fmt.Sprintf("256[fg:%v]", idx3)

		cl1Sf := rgb16b.New(
			rgb16b.RGBAttribute{GroundFlag: pencil.Foreground, Color: cl1},
			// rgb16b.RGBAttribute{GroundFlag: pencil.Background, Color:cl2},
		).SprintfFunc()
		cl2Sf := rgb16b.New(
			rgb16b.RGBAttribute{GroundFlag: pencil.Foreground, Color: cl1},
			rgb16b.RGBAttribute{GroundFlag: pencil.Background, Color: cl2},
		).SprintfFunc()
		cl3Sf := func(format string, colorIndex int, a ...interface{}) string {
			return fmt.Sprintf("\x1b[38;5;%dm", colorIndex) + fmt.Sprintf(format, a...) + "\x1b[0m"
		}

		fmt.Printf("%*s %s\n", size, clfmt1, cl1Sf("%s", s))
		fmt.Printf("%*s %s\n", size, clfm8bit, cl3Sf("%s", idx3, s))
		fmt.Printf("%*s %s\n\n", size, clfmt2, cl2Sf("%s", s))
	}

}

func findMaxSize() (size int) {
	var plte color.Palette = palette.Plan9
	cnames := rgb16b.Names
	num := len(cnames)
	for i := 0; i < num; i++ {
		name1, name2 := cnames[i], cnames[num-i-1]
		cl1, cl2 := rgb16b.Map[name1], rgb16b.Map[name2]
		idx1, idx2 := plte.Index(cl1), plte.Index(cl2)
		clfmt1, clfmt2 :=
			fmt.Sprintf("[fg:%d <%s>]", idx1, name1),
			fmt.Sprintf("[fg:%d<%s>|bg:%d<%s>]", idx1, name1, idx2, name2)
		_, _, n1 := util.CountChars(clfmt1)
		_, _, n2 := util.CountChars(clfmt2)
		n := util.Imax(n1, n2)
		size = util.Imax(size, n)
	}
	return size
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
