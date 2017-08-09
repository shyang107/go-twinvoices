package ansi8

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"

	"github.com/shyang107/go-twinvoices/pencil"
	// "github.com/shyang107/go-twinvoices/util"
)

var (

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorCache   = make(map[pencil.Attribute]*Color)
	colorCacheMu sync.Mutex // protects colorsCache

)

// Color is a alias of "color.Color"
type Color struct {
	// color.Color
	params  []pencil.Attribute
	noColor *bool
}

//---------------------------------------------------------
// Foreground text colors
// Set text color (foreground) : 30+ n, where n is from the color table (30-37)
const (
	FgBlack pencil.Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack pencil.Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors (or bright colors )
// Set background color: 40 + n, where n is from the color table (40-47)
const (
	BgBlack pencil.Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack pencil.Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
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

// sequence returns a formated SGR sequence to be plugged into a
// ESC[38;2;<r>;<g>;<b>m... Select foreground color
// ESC[48;2;<r>;<g>;<b>m... Select background color
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, val := range c.params {
		// format[i] = fmt.Sprintf("%v", val)
		format[i] = strconv.Itoa(int(val))
	}

	return strings.Join(format, ";")
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", pencil.Escape, c.sequence())
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

func getCachedColor(p pencil.Attribute) *Color {
	colorCacheMu.Lock()
	defer colorCacheMu.Unlock()

	c, ok := colorCache[p]
	if !ok {
		c = New(p)
		colorCache[p] = c
	}

	return c
}

// colorString returns a formatted colorful string with specified "colorname"
func colorString(format string, p pencil.Attribute, a ...interface{}) string {
	c := getCachedColor(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

//---------------------------------------------------------
