package pencil

import "fmt"

// ansi control code
const (
	Escape = "\x1b" // "\e" (033 or escape control code)
)

// Attribute defines a single SGR Code
type Attribute int

// Base attributes
const (
	Reset        Attribute = iota // Reset / Normal: all attributes off
	Bold                          // Bold or increased intensity
	Faint                         // Faint (decreased intensity): Not widely supported.
	Italic                        // Italic: on: Not widely supported. Sometimes treated as inverse.
	Underline                     // Underline: Single
	BlinkSlow                     // less than 150 per minute
	BlinkRapid                    // MS-DOS ANSI.SYS; 150+ per minute; not widely supported
	ReverseVideo                  // Image: Negative: inverse or reverse; swap foreground and background
	Concealed                     // Not widely supported.
	CrossedOut                    // Characters legible, but marked for deletion. Not widely supported.
)

// Base attributes
const (
	DefaultForeground Attribute = 39 // Default foreground color
	DefaultBackground Attribute = 49 // Default background color
)

// GetRest returns the  ANSI rest color code
func GetRest() string {
	return fmt.Sprintf("%s[%vm", Escape, Reset)
}

// GetDefaultForeground returns the ANSI default foreground color code
func GetDefaultForeground() string {
	return fmt.Sprintf("%s[%vm", Escape, DefaultForeground)
}

// GetDefaultBackground returns the ANSI default background color code
func GetDefaultBackground() string {
	return fmt.Sprintf("%s[%vm", Escape, DefaultBackground)
}

// GetDefaultGround returns the ANSI default color code
func GetDefaultGround() string {
	return fmt.Sprintf("%s[%v;%vm", Escape, DefaultForeground, DefaultBackground)
}
