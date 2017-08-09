package pencil

import (
	"fmt"
	"image/color"
)

// ansi control code
const (
	Escape = "\x1b" // "\e" (033 or escape control code)
)

// ColorCode is color code in colors
type ColorCode int

// Attribute defines a single SGR Code
type Attribute int

// SGR (Select Graphic Rendition) parameters
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

// SGR (Select Graphic Rendition) parameters
const (
	// Reserved for extended set foreground color
	// typical supported next arguments are 5; n where n is color index (0..255)
	// 	ESC[ … 38;5;<n> … m Select foreground color
	// or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
	// 	ESC[ … 38;2;<r>;<g>;<b> … m Select RGB foreground color
	Foreground Attribute = 38
	// Reserved for extended set background color
	// typical supported next arguments are 5; n where n is color index (0..255)
	// 	ESC[ … 48;5;<n> … m Select background color
	// or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
	// 	ESC[ … 48;2;<r>;<g>;<b> … m Select RGB background color
	Background Attribute = 48

	// Default foreground color (not supported on some terminals),
	DefaultForeground Attribute = 39
	// Default background color (not supported on some terminals),
	DefaultBackground Attribute = 49
)

// private SGR (Select Graphic Rendition) parameters
const (
	SelectColorIndex Attribute = 201
	SelectColorRGB   Attribute = 202
)

// GetSGR returns the ANSI SGR (Select Graphic Rendition) parameters
func GetSGR(a Attribute) string {
	switch a {
	case Reset, Bold, Faint, Italic, Underline, BlinkSlow,
		BlinkRapid, ReverseVideo, Concealed, CrossedOut:
		return fmt.Sprintf("%s[%vm", Escape, a)
	default:
		return ""
	}
}

// GetRest returns the ANSI rest color code
func GetRest() string {
	return GetSGR(Reset)
}

// GetForeground returns the ANSI foreground color code
// typical supported next arguments are 5; n where n is color index (0..255)
// 	ESC[ … 38;5;<n> … m Select foreground color
// or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
// 	ESC[ … 38;2;<r>;<g>;<b> … m Select RGB foreground color
func GetForeground(selectColor Attribute, cl interface{}) (string, error) {
	var colorfmt string
	switch selectColor {
	case SelectColorIndex:
		switch cl.(type) {
		case Attribute:
			code, ok := cl.(Attribute)
			if !ok {
				return "", fmt.Errorf("Selct color index; but <cl> = %v", cl)
			}
			colorfmt = GetForegroundIndex(int(code))
		case ColorCode:
			code, ok := cl.(ColorCode)
			if !ok {
				return "", fmt.Errorf("Selct color index; but <cl> = %v", cl)
			}
			colorfmt = GetForegroundIndex(int(code))
		}
		return colorfmt, nil
	default: // SelectColorRGB
		rgb, ok := cl.(color.Color)
		if !ok {
			return "", fmt.Errorf("Selct RGB color; but <cl> = %v", cl)
		}
		r, g, b, _ := rgb.RGBA()
		return GetForegroundRGB(int(r), int(g), int(b)), nil
	}
}

// GetBackground returns the ANSI  background color code
// typical supported next arguments are 5; n where n is color index (0..255)
// 	ESC[ … 48;5;<n> … m Select background color
// or 2;r;g;b where r,g,b are red, green and blue color channels (out of 255)
// 	ESC[ … 48;2;<r>;<g>;<b> … m Select RGB background color
func GetBackground(selectColor Attribute, cl interface{}) (string, error) {
	var colorfmt string
	switch selectColor {
	case SelectColorIndex:
		switch cl.(type) {
		case Attribute:
			code, ok := cl.(Attribute)
			if !ok {
				return "", fmt.Errorf("Selct color index; but <cl> = %v", cl)
			}
			colorfmt = GetBackgroundIndex(int(code))
		case ColorCode:
			code, ok := cl.(ColorCode)
			if !ok {
				return "", fmt.Errorf("Selct color index; but <cl> = %v", cl)
			}
			colorfmt = GetBackgroundIndex(int(code))
		}
		return colorfmt, nil
	default: // SelectColorRGB
		rgb, ok := cl.(color.Color)
		if !ok {
			return "", fmt.Errorf("Selct RGB color; but <cl> = %v", cl)
		}
		r, g, b, _ := rgb.RGBA()
		return GetBackgroundRGB(int(r), int(g), int(b)), nil
	}
}

// GetForegroundIndex returns the ANSI foreground color code
// typical supported next arguments are 5; n where n is color index (0..255)
// 	ESC[ … 38;5;<n> … m Select foreground color
func GetForegroundIndex(colorIndex int) string {
	return fmt.Sprintf("%s[%v;5;%vm", Escape, Foreground, colorIndex)
}

// GetBackgroundIndex returns the ANSI  background color code
// typical supported next arguments are 5; n where n is color index (0..255)
// 	ESC[ … 48;5;<n> … m Select background color
func GetBackgroundIndex(colorIndex int) string {
	return fmt.Sprintf("%s[%v;5;%vm", Escape, Background, colorIndex)
}

// GetForegroundRGB returns the ANSI  foreground color code
// typical supported next arguments are 2;
// r;g;b where r,g,b are red, green and blue color channels (out of 255)
// 	ESC[ … 38;2;<r>;<g>;<b> … m Select RGB foreground color
func GetForegroundRGB(r, g, b int) string {
	return fmt.Sprintf("%s[%v;2;%v;%v;%vm", Escape, Foreground, r, g, b)
}

// GetBackgroundRGB returns the ANSI background color code
// typical supported next arguments are 2;
// r;g;b where r,g,b are red, green and blue color channels (out of 255)
// 	ESC[ … 48;2;<r>;<g>;<b> … m Select RGB background color
func GetBackgroundRGB(r, g, b int) string {
	return fmt.Sprintf("%s[%v;2;%v;%v;%vm", Escape, Background, r, g, b)
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
