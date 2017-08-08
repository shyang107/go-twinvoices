package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shyang107/go-twinvoices/pencil"
	"github.com/shyang107/go-twinvoices/pencil/rgb16b"
	"github.com/shyang107/go-twinvoices/util"
	"golang.org/x/image/colornames"
)

func main() {
	testColors()
}

func testColors() {
	color.Black("color.Black:%v ", "This is testing the colorful text!")
	color.HiBlack("color.HiBlack:%v ", "This is testing the colorful text!")
	color.Blue("color.Blue:%v ", "This is testing the colorful text!")
	color.HiBlue("color.HiBlue:%v ", "This is testing the colorful text!")
	color.Cyan("color.Cyan:%v ", "This is testing the colorful text!")
	color.HiCyan("color.HiCyan:%v ", "This is testing the colorful text!")
	color.Green("color.Green:%v ", "This is testing the colorful text!")
	color.HiGreen("color.HiGreen:%v ", "This is testing the colorful text!")
	color.Magenta("color.Magenta:%v ", "This is testing the colorful text!")
	color.HiMagenta("color.HiMagenta:%v ", "This is testing the colorful text!")
	color.Red("color.Red:%v ", "This is testing the colorful text!")
	color.HiRed("color.HiRed:%v ", "This is testing the colorful text!")
	color.White("color.White:%v ", "This is testing the colorful text!")
	color.HiWhite("color.HiWhite:%v ", "This is testing the colorful text!")
	color.Yellow("color.Yellow:%v ", "This is testing the colorful text!")
	color.HiYellow("color.HiYellow:%v ", "This is testing the colorful text!")
	cl := color.New(color.BgGreen, color.FgBlack)
	cl.Print("color.New(color.BgRed,color.FgYellow) ", "This is testing the colorful text!", "\x1b[0m\n")
	util.Pforan("%v %v\n", "Pforan", "This is testing the colorful text!")
	fgcolor := rgb16b.New(colornames.Blue, pencil.Foreground)
	bgcolor := rgb16b.New(colornames.Darkorange, pencil.Background)
	// fgcolor.DisableColor()
	// bgcolor.DisableColor()
	fmt.Println("rgb16b:  This is testing the colorful text!",
		fgcolor.Fg(), bgcolor.Bg(), "This is testing the colorful text!",
		pencil.GetRest(),
	)
	fmt.Println("rgb16b.FBSprint:  This is testing the colorful text!",
		rgb16b.FBSprint(colornames.Blue, colornames.Orangered, "This is testing the colorful text!"),
	)
	fmt.Println("rgb16b.FBSprintln:  This is testing the colorful text!",
		rgb16b.FBSprintln(colornames.Blue, colornames.Orangered, "This is testing the colorful text!"),
	)
	fmt.Println("rgb16b.FBSprintf:  This is testing the colorful text!",
		rgb16b.FBSprintf(colornames.Blue, colornames.Orangered, "%s", "This is testing the colorful text!"),
	)
	fmt.Println()
	fbprint := rgb16b.FBPrintFunc(colornames.Gold, colornames.Darkred)
	fmt.Print("rgb16b.FBPrintFunc(colornames.Darkred, colornames.Gold): ")
	fbprint("This is testing the colorful text!", "\n")
	fmt.Println("=====================")
	fbprintf := rgb16b.FBPrintfFunc(colornames.Gold, colornames.Darkred)
	fmt.Print("rgb16b.FBPrintfFunc(colornames.Darkred, colornames.Gold): ")
	fbprintf("%s\n", "This is testing the colorful text!")
	fmt.Println("=====================")
	fbprintln := rgb16b.FBPrintlnFunc(colornames.Gold, colornames.Darkred)
	fmt.Print("rgb16b.FBPrintlnFunc(colornames.Darkred, colornames.Gold): ")
	fbprintln("This is testing the colorful text!")
	fmt.Println("=====================")

	// format := " %s%v%s\n %s%v%s\n "
	// formatR := strings.TrimRight(format, " ")
	// n := strings.LastIndex(formatR, "\n")
	// formatR = formatR[:n] + "^"
	// fmt.Println(string([]rune(format)), "len = ", len(format),
	// 	`; LastIndex \n = `, n, formatR)
}
