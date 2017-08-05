package ansisvg

import (
	"fmt"
	"image/color"
	"io"
	"strings"
	"sync"

	"github.com/shyang107/go-twinvoices/pencil"
)

var (

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorCache   = make(map[RGBAttribute]*Color)
	colorCacheMu sync.Mutex // protects colorsCache

)

// RGBAttribute define a key for a color
type RGBAttribute struct {
	color.Color
	GroundFlag pencil.GroundFlag
}

// Color is a alias of "color.Color"
type Color struct {
	// color.Color
	params  []RGBAttribute
	noColor *bool
}

//---------------------------------------------------------

const (
	fgleading = "\x1b[38;2;"
	bgleading = "\x1b[48;2;"
)

//---------------------------------------------------------

// New returns a newly created color object.
func New(value ...RGBAttribute) *Color {
	c := &Color{params: make([]RGBAttribute, 0)}
	c.Add(value...)
	return c
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...RGBAttribute) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value RGBAttribute) {
	c.params = append(c.params, RGBAttribute{})
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...RGBAttribute) *Color {
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
	var leadcfmt string
	format := make([]string, len(c.params))
	for i, val := range c.params {
		if pencil.IsForeground(val.GroundFlag) {
			leadcfmt = fgleading
		} else {
			leadcfmt = bgleading
		}
		r, g, b, _ := val.RGBA()
		format[i] = fmt.Sprintf("%s%v;%v;%vm", leadcfmt, r, g, b)
		// format[i] = fmt.Sprintf("%s%v;%v;%vm", leadcfmt, r, g, b)
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

func getCachedColor(k RGBAttribute) *Color {
	colorCacheMu.Lock()
	defer colorCacheMu.Unlock()

	c, ok := colorCache[k]
	if !ok {
		c = New(k)
		colorCache[k] = c
	}

	return c
}

// colorString returns a formatted colorful string with specified "colorname"
func colorString(format string, color RGBAttribute, a ...interface{}) string {
	c := getCachedColor(RGBAttribute{Color: color, GroundFlag: pencil.Foreground})

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

//---------------------------------------------------------
