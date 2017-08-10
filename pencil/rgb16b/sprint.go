package rgb16b

import (
	"fmt"
	"image/color"

	"github.com/shyang107/go-twinvoices/pencil"
)

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(color.Color, ...Attribute).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

func fbcolor(foregroundColor, backgroundColor color.Color) string {
	fc := New(foregroundColor, pencil.Foreground)
	bc := New(backgroundColor, pencil.Background)
	if fc.isNoColorSet() {
		return ""
	}
	return fc.Fg() + bc.Bg()
}

// FBSprint is just like Print, but returns a string instead of printing it.
func FBSprint(foregroundColor, backgroundColor color.Color, a ...interface{}) string {
	fb := fbcolor(foregroundColor, backgroundColor)
	if len(fb) > 0 && !pencil.NoColor {
		a = append(a, pencil.GetRest())
	}
	return fb + fmt.Sprint(a...)
}

// FBSprintln is just like Println, but returns a string instead of printing it.
func FBSprintln(foregroundColor, backgroundColor color.Color, a ...interface{}) string {
	fb := fbcolor(foregroundColor, backgroundColor)
	if len(fb) > 0 && !pencil.NoColor {
		a = append(a, pencil.GetRest())
	}
	return fb + fmt.Sprintln(a...)
}

// FBSprintf is just like Printf, but returns a string instead of printing it.
func FBSprintf(foregroundColor, backgroundColor color.Color,
	format string, a ...interface{}) string {
	fb := fbcolor(foregroundColor, backgroundColor)
	if len(fb) > 0 && !pencil.NoColor {
		format = fb + format + pencil.GetRest()
	}
	return fmt.Sprintf(format, a...)
}

// FBSprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(color.Color, ...Attribute).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func FBSprintFunc(foregroundColor, backgroundColor color.Color) func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return FBSprint(foregroundColor, backgroundColor, a...)
	}
}

// FBSprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func FBSprintfFunc(foregroundColor, backgroundColor color.Color) func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return FBSprintf(foregroundColor, backgroundColor, format, a...)
	}
}

// FBSprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func FBSprintlnFunc(foregroundColor, backgroundColor color.Color) func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return FBSprintln(foregroundColor, backgroundColor, a...)
	}
}
