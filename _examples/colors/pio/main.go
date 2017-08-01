package main

import (
	"os"

	"github.com/kataras/pio"
)

func main() {
	testColors()
}

func testColors() {
	p := pio.NewTextPrinter("color", os.Stdout)
	p.Println("*** pio ***")
	p.Println(pio.Blue("pio.Blue: this is a blue text"))
	p.Println(pio.Gray("pio.Gray: this is a gray text"))
	p.Println(pio.Red("pio.Red: this is a red text"))
	p.Println(pio.Purple("pio.Purple: this is a purple text"))
	p.Println(pio.Yellow("pio.Yellow: this is a yellow text"))
	p.Println(pio.Green("pio.Green: this is a green text"))
}
