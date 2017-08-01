package main

import (
	"github.com/fatih/color"
	"github.com/shyang107/go-twinvoices/util"
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
	cl.Print("color.New(color.BgRed,color.FgYellow) ", "This is testing the colorful text!", "\n")
	util.Pforan("%v %v\n", "Pforan", "This is testing the colorful text!")
}
