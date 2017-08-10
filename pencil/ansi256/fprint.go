package ansi256

import (
	"fmt"
	"io"
	"strings"

	"github.com/shyang107/go-twinvoices/pencil"
)

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprint(w, a...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintf(w, format, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintln(w, a...)
}

// FprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprint(w, a...)
	}
}

// FprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		c.Fprintf(w, format, a...)
	}
}

// FprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprintln(w, a...)
	}
}

// FBFprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func FBFprint(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode,
	a ...interface{}) (n int, err error) {
	fc := Set(foregroundColor, pencil.Foreground).setWriter(w)
	bc := Set(backgroundColor, pencil.Background).setWriter(w)
	defer fc.unsetWriter(w)
	defer bc.unsetWriter(w)

	// return fmt.Fprint(w, a...)

	if fc.isNoColorSet() {
		return fmt.Fprint(w, a...)
	}
	m := len(a)
	if a[m-1] == "\n" {
		a = append(a[:m-1], pencil.GetRest(), "\n")
		return fmt.Fprint(w, a...)
	}
	a = append(a, pencil.GetRest())
	return fmt.Fprint(w, a...)
}

// FBFprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func FBFprintf(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode,
	format string, a ...interface{}) (n int, err error) {
	fc := Set(foregroundColor, pencil.Foreground).setWriter(w)
	bc := Set(backgroundColor, pencil.Background).setWriter(w)
	defer fc.unsetWriter(w)
	defer bc.unsetWriter(w)

	// return fmt.Fprintf(w, format, a...)

	if fc.isNoColorSet() {
		return fmt.Fprintf(w, format, a...)
	}
	fr := strings.TrimRight(format, " ")
	m := strings.LastIndex(fr, "\n")
	if m != -1 && m == len(fr)-1 {
		fr = fr[:m] + pencil.GetRest() + "\n"
	}
	fr += strings.Repeat(" ", len(format)-len(fr))
	return fmt.Fprintf(w, fr, a...)

}

// FBFprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func FBFprintln(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode,
	a ...interface{}) (n int, err error) {
	fc := Set(foregroundColor, pencil.Foreground).setWriter(w)
	bc := Set(backgroundColor, pencil.Background).setWriter(w)
	defer fc.unsetWriter(w)
	defer bc.unsetWriter(w)

	// return fmt.Fprintln(w, a...)

	if fc.isNoColorSet() {
		return fmt.Fprintln(w, a...)
	}
	a = append(a, pencil.GetRest())
	return fmt.Fprintln(w, a...)
}

// FBFprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func FBFprintFunc(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode) func(a ...interface{}) {
	return func(a ...interface{}) {
		FBFprint(w, foregroundColor, backgroundColor, a...)
	}
}

// FBFprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func FBFprintfFunc(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode) func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		FBFprintf(w, foregroundColor, backgroundColor, format, a...)
	}
}

// FBFprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func FBFprintlnFunc(w io.Writer, foregroundColor, backgroundColor pencil.ColorCode) func(a ...interface{}) {
	return func(a ...interface{}) {
		FBFprintln(w, foregroundColor, backgroundColor, a...)
	}
}
