package util

import (
	"fmt"
	"image/color"
	"image/color/palette"
	"io"
	"strings"
	"sync"

	colorable "github.com/mattn/go-colorable"

	"golang.org/x/image/colornames"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	// NoColor = os.Getenv("TERM") == "dumb" ||
	// 	(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())) //&& Glog.Printer.IsTerminal

	// Output defines the standard output of the print functions. By default
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	crgbCache   = make(map[ColorAttribute]*ColorRGB)
	crgbCacheMu sync.Mutex // protects colorsCache

	// ColorsSVG is Colors map. Colors["SVGnames"]=color.RGBA
	ColorsSVG = colornames.Map

	plte color.Palette = palette.Plan9
)

const (
	escape = "\x1b" // "\e" (033 or escape control code)
	reset  = "\x1b[0m"

	fgRGB    = "38;2;"
	bgRGB    = "48;2;"
	colorsep = "|"
)

// ColorAttribute define a key for a color
type ColorAttribute struct {
	RGB          color.RGBA
	IsForeground bool
}

// ColorRGB is a alias of "color.Color"
type ColorRGB struct {
	// color.Color
	params  []ColorAttribute
	noColor *bool
}

//---------------------------------------------------------

// NewRGB returns a newly created color object.
func NewRGB(value ...ColorAttribute) *ColorRGB {
	c := &ColorRGB{params: make([]ColorAttribute, 0)}
	c.Add(value...)
	return c
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *ColorRGB) Add(value ...ColorAttribute) *ColorRGB {
	c.params = append(c.params, value...)
	return c
}

func (c *ColorRGB) prepend(value ColorAttribute) {
	c.params = append(c.params, ColorAttribute{})
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...ColorAttribute) *ColorRGB {
	c := NewRGB(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if NoColor {
		return
	}

	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.
func (c *ColorRGB) Set() *ColorRGB {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(Output, c.format())
	return c
}

func (c *ColorRGB) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

func (c *ColorRGB) setWriter(w io.Writer) *ColorRGB {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(w, c.format())
	return c
}

func (c *ColorRGB) unsetWriter(w io.Writer) {
	if c.isNoColorSet() {
		return
	}

	if NoColor {
		return
	}

	fmt.Fprintf(w, "%s[%dm", escape, Reset)
}

//---------------------------------------------------------

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *ColorRGB) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprint(w, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *ColorRGB) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *ColorRGB) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintf(w, format, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *ColorRGB) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
func (c *ColorRGB) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintln(w, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *ColorRGB) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintln(Output, a...)
}

// // Sprint is just like Print, but returns a string instead of printing it.
// func (c *ColorRGB) Sprint(a ...interface{}) string {
// 	return c.wrap(fmt.Sprint(a...))
// }

// // Sprintln is just like Println, but returns a string instead of printing it.
// func (c *ColorRGB) Sprintln(a ...interface{}) string {
// 	return c.wrap(fmt.Sprintln(a...))
// }

// // Sprintf is just like Printf, but returns a string instead of printing it.
// func (c *ColorRGB) Sprintf(format string, a ...interface{}) string {
// 	return c.wrap(fmt.Sprintf(format, a...))
// }

// FprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func (c *ColorRGB) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprint(w, a...)
	}
}

// PrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func (c *ColorRGB) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Print(a...)
	}
}

// FprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func (c *ColorRGB) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		c.Fprintf(w, format, a...)
	}
}

// PrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func (c *ColorRGB) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		c.Printf(format, a...)
	}
}

// FprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func (c *ColorRGB) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprintln(w, a...)
	}
}

// PrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func (c *ColorRGB) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Println(a...)
	}
}

//---------------------------------------------------------

// Sprint is just like Print, but returns a string instead of printing it.
func (c *ColorRGB) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *ColorRGB) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *ColorRGB) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *ColorRGB) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *ColorRGB) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *ColorRGB) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

// wrap wraps the s string with the colors Attributes. The string is ready to
// be printed.
func (c *ColorRGB) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func getRGBCodeString(c color.RGBA) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%v;%v;%vm", r, g, b)
}

// sequence returns a formated SGR sequence to be plugged into a
// ESC[38;2;<r>;<g>;<b>m... Select foreground color
// ESC[48;2;<r>;<g>;<b>m... Select background color
func (c *ColorRGB) sequence() string {
	var lcfmt string
	format := make([]string, len(c.params))
	for i, val := range c.params {
		if val.IsForeground {
			lcfmt = fgRGB
		} else {
			lcfmt = bgRGB
		}
		r, g, b, _ := val.RGB.RGBA()
		format[i] = fmt.Sprintf("%s[%s%v;%v;%vm", escape, lcfmt, r, g, b)
		// format[i] = fmt.Sprintf("%s%v;%v;%vm", lcfmt, r, g, b)
	}

	return strings.Join(format, "")
}

func (c *ColorRGB) format() string {
	// return fmt.Sprintf("%s[%sm", escape, c.sequence())
	return c.sequence()
}

func (c *ColorRGB) unformat() string {
	// return fmt.Sprintf("%s[%dm", escape, Reset)
	// return fmt.Sprintf("%s[%dm", escape, 0)
	var unf string
	for i := 0; i < len(c.params); i++ {
		unf += reset // "\x1b[0m"
	}
	return unf
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *ColorRGB) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.
func (c *ColorRGB) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *ColorRGB) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

func getCachedColorRGB(k ColorAttribute) *ColorRGB {
	crgbCacheMu.Lock()
	defer crgbCacheMu.Unlock()

	c, ok := crgbCache[k]
	if !ok {
		c = NewRGB(k)
		crgbCache[k] = c
	}

	return c
}

// colorRGBString returns a formatted colorful string with specified "colorname"
func colorRGBString(format string, rgb color.RGBA, a ...interface{}) string {
	c := getCachedColorRGB(ColorAttribute{RGB: rgb, IsForeground: true})

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

//---------------------------------------------------------

//---------------------------------------------------------
