package ansi8

import (
	"fmt"
	"io"
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
