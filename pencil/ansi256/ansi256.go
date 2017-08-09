package ansi256

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/shyang107/go-twinvoices/pencil"
)

const (
	fgleading = "\x1b[38;5;"
	bgleading = "\x1b[48;5;"
)

var (
	colorsCache   = make(map[pencil.Attribute]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// Color defines a custom color object which is defined by 256-color mode parameters.
type Color struct {
	params  []pencil.Attribute
	noColor *bool
}

//---------------------------------------------------------
// // Attribute defines a single SGR Code
// type Attribute int

// // Base attributes
// const (
// 	Foreground pencil.Attribute = 385 // ESC[38;5;<n>m
// 	Background pencil.Attribute = 485 // ESC[48;5;<n>m
// )

// Foreground Standard colors: n, where n is from the color table (0-7)
// (as in ESC[30–37m) <- SGR code
const (
	FgBlack pencil.Attribute = iota << 8
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground High-intensity colors: n, where n is from the color table (8-15)
// (as in ESC [ 90–97 m) <- SGR code
const (
	FgHiBlack pencil.Attribute = (iota + 8) << 8
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Foreground Grayscale colors: grayscale from black to white in 24 steps (232-255)
const (
	FgGrayscale01 pencil.Attribute = (iota + 232) << 8
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

const backgroundGate = 1 << 8

// Background Standard colors: n, where n is from the color table (0-7)
// (as in ESC[30–37m) <- SGR code
const (
	BgBlack pencil.Attribute = (iota + backgroundGate) << 8
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background High-intensity colors: n, where n is from the color table (8-15)
// (as in ESC [ 90–97 m) <- SGR code
const (
	BgHiBlack pencil.Attribute = (iota + 8 + backgroundGate) << 8
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// Background Grayscale colors: grayscale from black to white in 24 steps (232-255)
const (
	BgGrayscale01 pencil.Attribute = (iota + 232 + backgroundGate) << 8
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

//---------------------------------------------------------

// New returns a newly created color object.
func New(value ...pencil.Attribute) *Color {
	c := &Color{params: make([]pencil.Attribute, 0)}
	c.Add(value...)
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
func Set(p ...pencil.Attribute) *Color {
	c := New(p...)
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

// decode decode a color attribute (fore- and back-ground) to true 256 colors code
func decode(value pencil.Attribute) int {
	return int(value >> 8)
}

// Encode encode a true 256 colors code to a color attribute
func Encode(value int, isForeground bool) (n pencil.Attribute) {
	if isForeground {
		n = pencil.Attribute(value) << 8
	} else {
		n = pencil.Attribute(value+backgroundGate) << 8
	}
	return n
}

// sequence returns a formated SGR sequence to be plugged into a
// ESC[38;5;<n>m Select foreground color
// ESC[48;5;<n>m Select background color
// an example output might be: "38;15;12" -> foreground high-intensity blue
func (c *Color) sequence() string {
	var leadcfmt string
	format := make([]string, len(c.params))
	for i, v := range c.params {
		// format[i] = strconv.Itoa(int(v))
		code := decode(v)
		if code < backgroundGate {
			leadcfmt = fgleading
		} else {
			leadcfmt = bgleading
			code -= backgroundGate
		}
		format[i] = fmt.Sprintf("%s%dm", leadcfmt, code)
	}

	return strings.Join(format, "")
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

func getCachedColor(k pencil.Attribute) *Color {
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
func colorString(format string, color pencil.Attribute, a ...interface{}) string {
	c := getCachedColor(color)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

//---------------------------------------------------------
