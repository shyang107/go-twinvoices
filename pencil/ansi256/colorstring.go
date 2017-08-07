package ansi256

import "github.com/shyang107/go-twinvoices/pencil"

// Standard colors 0-7

// BlackString retrive a formatted string in color black
func BlackString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(0), a...)
}

// RedString retrive a formatted string in color Red
func RedString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(1)<<8, a...)
}

// GreenString retrive a formatted string in color Green
func GreenString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(2)<<8, a...)
}

// YellowString retrive a formatted string in color Yellow
func YellowString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(3)<<8, a...)
}

// BlueString retrive a formatted string in color Blue
func BlueString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(4)<<8, a...)
}

// MagentaString retrive a formatted string in color Magenta
func MagentaString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(5)<<8, a...)
}

// CyanString retrive a formatted string in color Cyan
func CyanString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(6)<<8, a...)
}

// GreyString retrive a formatted string in color Grey
func GreyString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(7)<<8, a...)
}

// High-intensity colors 8-15

// HiBlackString retrive a formatted string in color HiBlack
func HiBlackString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(8), a...)
}

// HiRedString retrive a formatted string in color HiRed
func HiRedString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(9)<<8, a...)
}

// HiGreenString retrive a formatted string in color HiGreen
func HiGreenString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(10)<<8, a...)
}

// HiYellowString retrive a formatted string in color HiYellow
func HiYellowString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(11)<<8, a...)
}

// HiBlueString retrive a formatted string in color HiBlue
func HiBlueString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(12)<<8, a...)
}

// HiMagentaString retrive a formatted string in color HiMagenta
func HiMagentaString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(13)<<8, a...)
}

// HiCyanString retrive a formatted string in color HiCyan
func HiCyanString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(14)<<8, a...)
}

// HiGreyString retrive a formatted string in color HiGrey
func HiGreyString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(15)<<8, a...)
}

// Specified colors in 216-colors: 16-231

// ShadeCyanString retrive a formatted string in another shade of cyan
func ShadeCyanString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(50<<8), a...)
}

// ShadeYellowString retrive a formatted string in another shade of Yellow (dark yellow)
func ShadeYellowString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(58)<<8, a...)
}

// ShadeYellowString2 retrive a formatted string in another shade of Yellow (dark yellow2)
func ShadeYellowString2(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(94)<<8, a...)
}

// PinkString retrive a formatted string in another shade of Pink
func PinkString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(205)<<8, a...)
}

// ShadeGreenString retrive a formatted string in another shade of Green (dark Green)
func ShadeGreenString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(22)<<8, a...)
}

// ShadePurpleString retrive a formatted string in another shade of Purple
func ShadePurpleString(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(55)<<8, a...)
}

// ShadeBlueString2 retrive a formatted string in another shade of blue
func ShadeBlueString2(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(69)<<8, a...)
}

// ShadeGrayString1 retrive a formatted string in another shade of gray
func ShadeGrayString1(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(59)<<8, a...)
}

// ShadeGrayString2 retrive a formatted string in another shade of gray
func ShadeGrayString2(format string, a ...interface{}) string {
	return colorString(format, pencil.Attribute(60)<<8, a...)
}

// FgOrange is the code of orange in 256-colors
const FgOrange = 202 << 8

// OrangeString retrive a formatted string in orange
func OrangeString(format string, a ...interface{}) string {
	return colorString(format, FgOrange, a...)
}

// Four levels of gray
const (
	FgGray1 = 238 << 8
	FgGray2 = 243 << 8
	FgGray3 = 248 << 8
	FgGray4 = 258 << 8
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
