package ansi8

import "github.com/shyang107/go-twinvoices/pencil"

// SetAttribute set the color of the object "Color"
func (c *Color) SetAttribute(attrs ...interface{}) error {
	for _, intf := range attrs {
		attr, ok := intf.(pencil.Attribute)
		if !ok {
			continue
		}
		c = c.Add(attr)
	}
	return nil
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	c.noColor = pencil.BoolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color) EnableColor() {
	c.noColor = pencil.BoolPtr(false)
}
