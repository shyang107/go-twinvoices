package rgb16b

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/shyang107/go-twinvoices/pencil"
)

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(pencil.Output, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(pencil.Output, format, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintln(pencil.Output, a...)
}

// PrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Print(a...)
	}
}

// PrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		c.Printf(format, a...)
	}
}

// PrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Println(a...)
	}
}

// FBPrint formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func FBPrint(foregroundColor, backgroundColor color.Color, a ...interface{}) (n int, err error) {
	Set(foregroundColor, pencil.Foreground).Set()
	Set(backgroundColor, pencil.Background).Set()
	fc := Set(foregroundColor, pencil.Foreground).Set()
	bc := Set(backgroundColor, pencil.Background).Set()
	defer fc.unset()
	defer bc.unset()
	if fc.isNoColorSet() {
		return fmt.Fprint(pencil.Output, a...)
	}
	m := len(a)
	if a[m-1] == "\n" {
		a = append(a[:m-1], pencil.GetRest(), "\n")
		return fmt.Fprint(pencil.Output, a...)
	}
	a = append(a, pencil.GetRest())
	return fmt.Fprint(pencil.Output, a...)
}

// FBPrintf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func FBPrintf(foregroundColor, backgroundColor color.Color,
	format string, a ...interface{}) (n int, err error) {
	fc := Set(foregroundColor, pencil.Foreground).Set()
	bc := Set(backgroundColor, pencil.Background).Set()
	defer fc.unset()
	defer bc.unset()
	if fc.isNoColorSet() {
		return fmt.Fprintf(pencil.Output, format, a...)
	}
	fr := strings.TrimRight(format, " ")
	m := strings.LastIndex(fr, "\n")
	if m != -1 && m == len(fr)-1 {
		fr = fr[:m] + pencil.GetRest() + "\n"
	}
	fr += strings.Repeat(" ", len(format)-len(fr))
	return fmt.Fprintf(pencil.Output, fr, a...)
}

// FBPrintln formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func FBPrintln(foregroundColor, backgroundColor color.Color, a ...interface{}) (n int, err error) {
	fc := Set(foregroundColor, pencil.Foreground).Set()
	bc := Set(backgroundColor, pencil.Background).Set()
	defer fc.unset()
	defer bc.unset()
	if fc.isNoColorSet() {
		return fmt.Fprintln(pencil.Output, a...)
	}
	a = append(a, pencil.GetRest())
	return fmt.Fprintln(pencil.Output, a...)
}

// FBPrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func FBPrintFunc(foregroundColor, backgroundColor color.Color) func(a ...interface{}) {
	return func(a ...interface{}) {
		FBPrint(foregroundColor, backgroundColor, a...)
	}
}

// FBPrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func FBPrintfFunc(foregroundColor, backgroundColor color.Color) func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		FBPrintf(foregroundColor, backgroundColor, format, a...)
	}
}

// FBPrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func FBPrintlnFunc(foregroundColor, backgroundColor color.Color) func(a ...interface{}) {
	return func(a ...interface{}) {
		FBPrintln(foregroundColor, backgroundColor, a...)
	}
}
