package rgb16b

import "github.com/shyang107/go-twinvoices/pencil"

// SetAttribute set the attributes of the object "Color"
func (c *Color) SetAttribute(attrs ...pencil.Attribute) {
	c = c.Add(attrs...)
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
