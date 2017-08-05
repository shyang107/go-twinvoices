package ansi256

import "github.com/shyang107/go-twinvoices/pencil"

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

// ShadePinkString retrive a formatted string in another shade of Pink
func ShadePinkString(format string, a ...interface{}) string {
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
