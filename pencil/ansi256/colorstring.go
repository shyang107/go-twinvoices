package ansi256

import "github.com/shyang107/go-twinvoices/pencil"

// Standard colors 0-7

// BlackString retrive a formatted string in color black
func BlackString(format string, a ...interface{}) string {
	return colorString(format, Black, a...)
}

// RedString retrive a formatted string in color Red
func RedString(format string, a ...interface{}) string {
	return colorString(format, Red, a...)
}

// GreenString retrive a formatted string in color Green
func GreenString(format string, a ...interface{}) string {
	return colorString(format, Green, a...)
}

// YellowString retrive a formatted string in color Yellow
func YellowString(format string, a ...interface{}) string {
	return colorString(format, Yellow, a...)
}

// BlueString retrive a formatted string in color Blue
func BlueString(format string, a ...interface{}) string {
	return colorString(format, Blue, a...)
}

// MagentaString retrive a formatted string in color Magenta
func MagentaString(format string, a ...interface{}) string {
	return colorString(format, Magenta, a...)
}

// CyanString retrive a formatted string in color Cyan
func CyanString(format string, a ...interface{}) string {
	return colorString(format, Cyan, a...)
}

// WhiteString retrive a formatted string in color White
func WhiteString(format string, a ...interface{}) string {
	return colorString(format, White, a...)
}

// High-intensity colors 8-15

// HiBlackString retrive a formatted string in color HiBlack
func HiBlackString(format string, a ...interface{}) string {
	return colorString(format, HiBlack, a...)
}

// HiRedString retrive a formatted string in color HiRed
func HiRedString(format string, a ...interface{}) string {
	return colorString(format, HiRed, a...)
}

// HiGreenString retrive a formatted string in color HiGreen
func HiGreenString(format string, a ...interface{}) string {
	return colorString(format, HiGreen, a...)
}

// HiYellowString retrive a formatted string in color HiYellow
func HiYellowString(format string, a ...interface{}) string {
	return colorString(format, HiYellow, a...)
}

// HiBlueString retrive a formatted string in color HiBlue
func HiBlueString(format string, a ...interface{}) string {
	return colorString(format, HiBlue, a...)
}

// HiMagentaString retrive a formatted string in color HiMagenta
func HiMagentaString(format string, a ...interface{}) string {
	return colorString(format, HiMagenta, a...)
}

// HiCyanString retrive a formatted string in color HiCyan
func HiCyanString(format string, a ...interface{}) string {
	return colorString(format, HiCyan, a...)
}

// HiWhiteString retrive a formatted string in color HiWhite
func HiWhiteString(format string, a ...interface{}) string {
	return colorString(format, HiWhite, a...)
}

// Specified colors in 216-colors: 16-231

// ShadeCyanString retrive a formatted string in another shade of cyan
func ShadeCyanString(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(50), a...)
}

// ShadeYellowString retrive a formatted string in another shade of Yellow (dark yellow)
func ShadeYellowString(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(58), a...)
}

// ShadeYellowString2 retrive a formatted string in another shade of Yellow (dark yellow2)
func ShadeYellowString2(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(94), a...)
}

// PinkString retrive a formatted string in another shade of Pink
func PinkString(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(205), a...)
}

// ShadeGreenString retrive a formatted string in another shade of Green (dark Green)
func ShadeGreenString(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(22), a...)
}

// ShadePurpleString retrive a formatted string in another shade of Purple
func ShadePurpleString(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(55), a...)
}

// ShadeBlueString2 retrive a formatted string in another shade of blue
func ShadeBlueString2(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(69), a...)
}

// ShadeGrayString1 retrive a formatted string in another shade of gray
func ShadeGrayString1(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(59), a...)
}

// ShadeGrayString2 retrive a formatted string in another shade of gray
func ShadeGrayString2(format string, a ...interface{}) string {
	return colorString(format, pencil.ColorCode(60), a...)
}

// Orange is the code of orange in 256-colors
const Orange = 202

// OrangeString retrive a formatted string in orange
func OrangeString(format string, a ...interface{}) string {
	return colorString(format, Orange, a...)
}

// Four levels of gray
const (
	FgGray1 pencil.ColorCode = 238
	FgGray2 pencil.ColorCode = 243
	FgGray3 pencil.ColorCode = 248
	FgGray4 pencil.ColorCode = 258
)

// GrayString1 retrive a formatted string in Grayscale = 238
func GrayString1(format string, a ...interface{}) string {
	return colorString(format, FgGray1, a...)
}

// GrayString2 retrive a formatted string in Grayscale = 243
func GrayString2(format string, a ...interface{}) string {
	return colorString(format, FgGray2, a...)
}

// GrayString3 retrive a formatted string in Grayscale = 248
func GrayString3(format string, a ...interface{}) string {
	return colorString(format, FgGray3, a...)
}

// GrayString4 retrive a formatted string in Grayscale = 253
func GrayString4(format string, a ...interface{}) string {
	return colorString(format, FgGray4, a...)
}
