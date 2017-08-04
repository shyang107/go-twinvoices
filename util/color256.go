// modify and simply from "github.com/fatih/color"
// refers: https://en.wikipedia.org/wiki/ANSI_escape_code

package util

import (
	"fmt"
	"strings"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	// NoColor = os.Getenv("TERM") == "dumb" ||
	// 	(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())) //&& Glog.Printer.IsTerminal

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colors256Cache = make(map[Attribute]*Color256)
	// colorsCacheMu sync.Mutex // protects colorsCache
)

// Color256 defines a custom color object which is defined by 256-color mode parameters.
type Color256 struct {
	params  []Attribute
	noColor *bool
}

// // Attribute defines a single SGR Code
// type Attribute int

// Base attributes
const (
	Foreground256 Attribute = 385 // ESC[38;5;<n>m
	Background256 Attribute = 485 // ESC[48;5;<n>m
)

// Foreground Standard colors: n, where n is from the color table (0-7)
// (as in ESC[30–37m) <- SGR code
const (
	FgBlack256 Attribute = iota << 8
	FgRed256
	FgGreen256
	FgYellow256
	FgBlue256
	FgMagenta256
	FgCyan256
	FgWhite256
)

// Foreground High-intensity colors: n, where n is from the color table (8-15)
// (as in ESC [ 90–97 m) <- SGR code
const (
	FgHiBlack256 Attribute = (iota + 8) << 8
	FgHiRed256
	FgHiGreen256
	FgHiYellow256
	FgHiBlue256
	FgHiMagenta256
	FgHiCyan256
	FgHiWhite256
)

// Foreground Grayscale colors: grayscale from black to white in 24 steps (232-255)
const (
	FgGrayscale01 Attribute = (iota + 232) << 8
	FgGrayscale02
	FgGrayscale03
	FgGrayscale04
	FgGrayscale05
	FgGrayscale06
	FgGrayscale07
	FgGrayscale08
	FgGrayscale09
	FgGrayscale10
	FgGrayscale11
	FgGrayscale12
	FgGrayscale13
	FgGrayscale14
	FgGrayscale15
	FgGrayscale16
	FgGrayscale17
	FgGrayscale18
	FgGrayscale19
	FgGrayscale20
	FgGrayscale21
	FgGrayscale22
	FgGrayscale23
	FgGrayscale24
)

const bgzone = 256

// Background Standard colors: n, where n is from the color table (0-7)
// (as in ESC[30–37m) <- SGR code
const (
	BgBlack256 Attribute = (iota + bgzone) << 8
	BgRed256
	BgGreen256
	BgYellow256
	BgBlue256
	BgMagenta256
	BgCyan256
	BgWhite256
)

// Background High-intensity colors: n, where n is from the color table (8-15)
// (as in ESC [ 90–97 m) <- SGR code
const (
	BgHiBlack256 Attribute = (iota + 8 + bgzone) << 8
	BgHiRed256
	BgHiGreen256
	BgHiYellow256
	BgHiBlue256
	BgHiMagenta256
	BgHiCyan256
	BgHiWhite256
)

// Background Grayscale colors: grayscale from black to white in 24 steps (232-255)
const (
	BgGrayscale01 Attribute = (iota + 232 + bgzone) << 8
	BgGrayscale02
	BgGrayscale03
	BgGrayscale04
	BgGrayscale05
	BgGrayscale06
	BgGrayscale07
	BgGrayscale08
	BgGrayscale09
	BgGrayscale10
	BgGrayscale11
	BgGrayscale12
	BgGrayscale13
	BgGrayscale14
	BgGrayscale15
	BgGrayscale16
	BgGrayscale17
	BgGrayscale18
	BgGrayscale19
	BgGrayscale20
	BgGrayscale21
	BgGrayscale22
	BgGrayscale23
	BgGrayscale24
)

// Color256RGB return index n of 6 × 6 × 6 cube (216 colors) (16-231)
// n = 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
func Color256RGB(r, g, b int, isForeground bool) (n Attribute) {
	r = checkRGBcode(r)
	g = checkRGBcode(g)
	b = checkRGBcode(b)
	if isForeground {
		n = Attribute(16+36*r+6*g+b) << 8
	} else {
		n = Attribute(16+36*r+6*g+b+bgzone) << 8
	}
	return n
}

func checkRGBcode(code int) int {
	switch {
	case code < 0:
		return 0
	case code > 5:
		return 5
	default:
		return code
	}
}

// DecodeColor256 decode a color attribute (fore- and back-ground) to true 256 colors code
func DecodeColor256(value Attribute) int {
	return int(value >> 8)
}

// EncodeColor256 encode a true 256 colors code to a color attribute
func EncodeColor256(value int, isForeground bool) (n Attribute) {
	if isForeground {
		n = Attribute(value) << 8
	} else {
		n = Attribute(value+bgzone) << 8
	}
	return n
}

// New256 returns a newly created color object.
// value is from EncodeColor256(value int, isForeground bool), constant color256-attributes
func New256(value ...Attribute) *Color256 {
	c := &Color256{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color256) Add(value ...Attribute) *Color256 {
	c.params = append(c.params, value...)
	return c
}

func (c *Color256) prepend(value Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color256) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color256) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color256) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color256) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color256) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color256) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

// sequence returns a formated SGR sequence to be plugged into a
// ESC[38;5;<n>m Select foreground color
// ESC[48;5;<n>m Select background color
// an example output might be: "38;15;12" -> foreground high-intensity blue
func (c *Color256) sequence() string {
	var leadcfmt string
	format := make([]string, len(c.params))
	for i, v := range c.params {
		// format[i] = strconv.Itoa(int(v))
		code := DecodeColor256(v)
		if code < bgzone {
			leadcfmt = fg256
		} else {
			leadcfmt = bg256
			code -= bgzone
		}
		format[i] = fmt.Sprintf("%s%dm", leadcfmt, code)
	}

	return strings.Join(format, "")
}

// wrap wraps the s string with the colors Attributes. The string is ready to
// be printed.
func (c *Color256) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}
func (c *Color256) format() string {
	// return fmt.Sprintf("%s[%sm", escape, c.sequence())
	return c.sequence()
}

func (c *Color256) unformat() string {
	// return fmt.Sprintf("%s[%dm", escape, Reset)
	// return clear
	return clear + reset
	// var unf string
	// for i := 0; i < len(c.params); i++ {
	// 	unf += reset
	// }
	// return unf
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color256) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color256) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *Color256) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

// func boolPtr(v bool) *bool {
// 	return &v
// }

func getCachedColor256(p Attribute) *Color256 {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colors256Cache[p]
	if !ok {
		c = New256(p)
		colors256Cache[p] = c
	}

	return c
}

// Color256String retrives the specified colorful string
func Color256String(format string, p Attribute, a ...interface{}) string {
	c := getCachedColor256(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

// ShadeCyanString retrive a formatted string in another shade of cyan
func ShadeCyanString(format string, a ...interface{}) string {
	return Color256String(format, Attribute(50<<8), a...)
}

// ShadeYellowString retrive a formatted string in another shade of Yellow (dark yellow)
func ShadeYellowString(format string, a ...interface{}) string {
	return Color256String(format, Attribute(58)<<8, a...)
}

// ShadeYellowString2 retrive a formatted string in another shade of Yellow (dark yellow2)
func ShadeYellowString2(format string, a ...interface{}) string {
	return Color256String(format, Attribute(94)<<8, a...)
}

// ShadePinkString retrive a formatted string in another shade of Pink
func ShadePinkString(format string, a ...interface{}) string {
	return Color256String(format, Attribute(205)<<8, a...)
}

// ShadeGreenString retrive a formatted string in another shade of Green (dark Green)
func ShadeGreenString(format string, a ...interface{}) string {
	return Color256String(format, Attribute(22)<<8, a...)
}

// ShadePurpleString retrive a formatted string in another shade of Purple
func ShadePurpleString(format string, a ...interface{}) string {
	return Color256String(format, Attribute(55)<<8, a...)
}

// ShadeBlueString2 retrive a formatted string in another shade of blue
func ShadeBlueString2(format string, a ...interface{}) string {
	return Color256String(format, Attribute(69)<<8, a...)
}

// ShadeGrayString1 retrive a formatted string in another shade of gray
func ShadeGrayString1(format string, a ...interface{}) string {
	return Color256String(format, Attribute(59)<<8, a...)
}

// ShadeGrayString2 retrive a formatted string in another shade of gray
func ShadeGrayString2(format string, a ...interface{}) string {
	return Color256String(format, Attribute(60)<<8, a...)
}

// FgOrange is the code of orange in 256-colors
const FgOrange = 202 << 8

// OrangeString retrive a formatted string in orange
func OrangeString(format string, a ...interface{}) string {
	return Color256String(format, FgOrange, a...)
}

// Four levels of gray
const (
	FgGray1 = 238 << 8
	FgGray2 = 243 << 8
	FgGray3 = 248 << 8
	FgGray4 = 258 << 8
)

// GrayString1 retrive a formatted string in Grayscale = 238
func GrayString1(format string, a ...interface{}) string {
	return Color256String(format, FgGray1, a...)
}

// GrayString2 retrive a formatted string in Grayscale = 243
func GrayString2(format string, a ...interface{}) string {
	return Color256String(format, FgGray2, a...)
}

// GrayString3 retrive a formatted string in Grayscale = 248
func GrayString3(format string, a ...interface{}) string {
	return Color256String(format, FgGray3, a...)
}

// GrayString4 retrive a formatted string in Grayscale = 253
func GrayString4(format string, a ...interface{}) string {
	return Color256String(format, FgGray4, a...)
}
