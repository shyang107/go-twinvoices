package pencil

import (
	"io"
	"os"

	isatty "github.com/mattn/go-isatty"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())) //&& Glog.Printer.IsTerminal

	// Output defines the standard output of the print functions. By default
	// os.Stdout is used.
	Output = NewColorableStdout()
)

// GeneralColor is use to handle ANSI or RGB colors
type GeneralColor interface {
	SetAttribute(attrs ...interface{}) error

	DisableColor()
	EnableColor()

	Sprint(a ...interface{}) string
	Sprintln(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
	SprintFunc() func(a ...interface{}) string
	SprintfFunc() func(format string, a ...interface{}) string
	SprintlnFunc() func(a ...interface{}) string

	Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	FprintFunc() func(w io.Writer, a ...interface{})
	FprintfFunc() func(w io.Writer, format string, a ...interface{})
	FprintlnFunc() func(w io.Writer, a ...interface{})

	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	PrintFunc() func(a ...interface{})
	PrintfFunc() func(format string, a ...interface{})
	PrintlnFunc() func(a ...interface{})
}

// ColorMode defines the mode for ANSI basic 8-colors mode, ANSI 256 indexed colors
// or SVG colors (alpha-premultiplied 16-bits per channel RGBA)
type ColorMode uint

// // ColorCode defines the code of the indexed color (ANSI)
// type ColorCode int

// Color modes
const (
	ModeANSI8 ColorMode = iota
	ModeANSI256
	ModeRGB
)

// // GroundFlag define color in foreground or background
// type GroundFlag uint

// // Settings for the flag of foreground or background
// const (
// 	Foreground GroundFlag = 1 << iota
// 	Background
// )

// // IsForeground return true if flag = Foreground
// func IsForeground(flag GroundFlag) bool {
// 	switch flag {
// 	case Background:
// 		return false
// 	default: // Foreground
// 		return true
// 	}
// }

// // IsBackground return true if flag = Background
// func IsBackground(flag GroundFlag) bool {
// 	return !IsForeground(flag)
// }

// NewColorableStdout return new instance of Writer which handle escape sequence for stdout.
func NewColorableStdout() io.Writer {
	return os.Stdout
}

// NewColorableStderr return new instance of Writer which handle escape sequence for stderr.
func NewColorableStderr() io.Writer {
	return os.Stderr
}

// BoolPtr return &{bool}
func BoolPtr(v bool) *bool {
	return &v
}
