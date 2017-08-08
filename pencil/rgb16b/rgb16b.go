package rgb16b

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
	colorCache   = make(map[pencil.Attribute]*Color)
	colorCacheMu sync.Mutex // protects colorsCache

)

// Color is a alias of "color.Color"
type Color struct {
	color.Color
	params  []pencil.Attribute
	noColor *bool
}

//---------------------------------------------------------

const (
	fgleading = "\x1b[38;2;"
	bgleading = "\x1b[48;2;"
)

//---------------------------------------------------------

// New returns a newly created color object.
func New(cl color.Color, params ...pencil.Attribute) *Color {
	c := &Color{Color: cl, params: make([]pencil.Attribute, 0)}
	c.Add(params...)
	return c
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(params ...pencil.Attribute) *Color {
	c.params = append(c.params, params...)
	return c
}

func (c *Color) prepend(param pencil.Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = param
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(cl color.Color, p ...pencil.Attribute) *Color {
	c := New(cl, p...)
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
			colorfmt, _ = pencil.GetBackground(pencil.SelectColorRGB, c.Color)
		case pencil.DefaultForeground:
			colorfmt = pencil.GetDefaultForeground()
		case pencil.DefaultBackground:
			colorfmt = pencil.GetDefaultBackground()
		default: // pencil.Foreground
			colorfmt, _ = pencil.GetForeground(pencil.SelectColorRGB, c.Color)
		}
		format = append(format, colorfmt)
	}

	// format[i] = fmt.Sprintf("%s%v;%v;%vm", leadcfmt, r, g, b)
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

	format, _ := pencil.GetForeground(pencil.SelectColorRGB, c.Color)
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

	format, _ := pencil.GetBackground(pencil.SelectColorRGB, c.Color)
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

// func getCachedColor(k pencil.Attribute) *Color {
// 	colorCacheMu.Lock()
// 	defer colorCacheMu.Unlock()

// 	c, ok := colorCache[k]
// 	if !ok {
// 		c = New(k)
// 		colorCache[k] = c
// 	}

// 	return c
// }

// // colorString returns a formatted colorful string with specified "colorname"
// func colorString(format string, color pencil.Attribute, a ...interface{}) string {
// 	c := getCachedColor(pencil.Attribute{Color: color, GroundFlag: pencil.Foreground})

// 	if len(a) == 0 {
// 		return c.SprintFunc()(format)
// 	}

// 	return c.SprintfFunc()(format, a...)
// }

//---------------------------------------------------------
