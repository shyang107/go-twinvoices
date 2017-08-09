package ansi256

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/shyang107/go-twinvoices/pencil"
)

// const (
// 	fgleading = "\x1b[38;5;"
// 	bgleading = "\x1b[48;5;"
// )

var (
	colorsCache   = make(map[pencil.ColorCode]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// Color defines a custom color object which is defined by 256-color mode parameters.
// "params" contains color index and it's attributes, such as foreground or
// background, ...; if not specify, default foreground color
type Color struct {
	Code    pencil.ColorCode // color index
	params  []pencil.Attribute
	noColor *bool // use DisableColor() or EnableColor() to setup
}

//---------------------------------------------------------

// Specilized colors: n, where n is from the color table (0-7, 8-15)
const (
	// Standard colors: 0-7 (as in ESC [ 30–37 m)
	Black pencil.ColorCode = iota // (avoid to confuse with non-color code)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	// High-intensity colors: 8-15 (as in ESC [ 90–97 m)
	HiBlack
	HiRed
	HiGreen
	HiYellow
	HiBlue
	HiMagenta
	HiCyan
	HiWhite
	// color index 16-231 = 6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
)

// Grayscale colors: grayscale from black to white in 24 steps (232-255)
const (
	Grayscale01 pencil.ColorCode = (iota + 232) // (avoid to confuse with non-color code)
	Grayscale02
	Grayscale03
	Grayscale04
	Grayscale05
	Grayscale06
	Grayscale07
	Grayscale08
	Grayscale09
	Grayscale10
	Grayscale11
	Grayscale12
	Grayscale13
	Grayscale14
	Grayscale15
	Grayscale16
	Grayscale17
	Grayscale18
	Grayscale19
	Grayscale20
	Grayscale21
	Grayscale22
	Grayscale23
	Grayscale24
)

//---------------------------------------------------------

// New returns a newly created color object.
func New(code pencil.ColorCode, params ...pencil.Attribute) *Color {
	c := &Color{Code: code, params: make([]pencil.Attribute, 0)}
	c.Add(params...)
	return c
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...pencil.Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value pencil.Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(code pencil.ColorCode, p ...pencil.Attribute) *Color {
	c := New(code, p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if pencil.NoColor {
		return
	}

	fmt.Fprintf(pencil.Output, "%s[%dm", pencil.Escape, pencil.Reset)
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(pencil.Output, c.format())
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

func (c *Color) setWriter(w io.Writer) *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(w, c.format())
	return c
}

func (c *Color) unsetWriter(w io.Writer) {
	if c.isNoColorSet() {
		return
	}

	if pencil.NoColor {
		return
	}

	fmt.Fprintf(w, "%s[%dm", pencil.Escape, pencil.Reset)
}

//---------------------------------------------------------

// wrap wraps the s string with the colors Attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

// // decode decode a color attribute (fore- and back-ground) to true 256 colors code
// func decode(value pencil.Attribute) int {
// 	return int(value >> 8)
// }

// // Encode encode a true 256 colors code to a color attribute
// func Encode(value int, isForeground bool) (n pencil.Attribute) {
// 	if isForeground {
// 		n = pencil.Attribute(value) << 8
// 	} else {
// 		n = pencil.Attribute(value+backgroundGate) << 8
// 	}
// 	return n
// }

// sequence returns a formated SGR sequence to be plugged into a
// ESC[38;5;<n>m Select foreground color
// ESC[48;5;<n>m Select background color
// an example output might be: "38;15;12" -> foreground high-intensity blue
func (c *Color) sequence() string {
	var colorfmt string
	format := make([]string, 0)
	for _, val := range c.params {
		sgr := pencil.GetSGR(val)
		if len(sgr) > 0 {
			format = append(format, sgr)
			continue
		}
		switch val {
		case pencil.Background:
			colorfmt, _ = pencil.GetBackground(pencil.SelectColorIndex, c.Code)
		case pencil.DefaultForeground:
			colorfmt = pencil.GetDefaultForeground()
		case pencil.DefaultBackground:
			colorfmt = pencil.GetDefaultBackground()
		default: // pencil.Foreground
			colorfmt, _ = pencil.GetForeground(pencil.SelectColorIndex, c.Code)
		}
		format = append(format, colorfmt)
	}

	return strings.Join(format, "")
}

// Fg retrive a leading sring in foreground color
func (c *Color) Fg() string {
	if c.isNoColorSet() {
		return ""
	}

	if pencil.NoColor {
		return ""
	}

	format, _ := pencil.GetForeground(pencil.SelectColorIndex, c.Code)
	return format
}

// Bg retrive a leading sring in background color
func (c *Color) Bg() string {
	if c.isNoColorSet() {
		return ""
	}

	if pencil.NoColor {
		return ""
	}

	format, _ := pencil.GetBackground(pencil.SelectColorIndex, c.Code)
	return format
}

func (c *Color) format() string {
	// return fmt.Sprintf("%s[%sm", escape, c.sequence())
	return c.sequence()
}

func (c *Color) unformat() string {
	return pencil.GetDefaultGround() + pencil.GetRest()
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return pencil.NoColor
}

func getCachedColor(k pencil.ColorCode) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[k]
	if !ok {
		c = New(k)
		colorsCache[k] = c
	}

	return c
}

// colorString returns a formatted colorful string with specified "colorname"
func colorString(format string, color pencil.ColorCode, a ...interface{}) string {
	c := getCachedColor(color)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

//---------------------------------------------------------
