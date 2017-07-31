package util

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/kataras/golog"
	"github.com/kataras/pio"
)

var (
	glog = golog.New()

	// Glog is a logger of golog
	Glog MyLogger = glog
	// enableColor          = &glog.Printer.IsTerminal
)

// InitLogger initialize logger's environment
func InitLogger() {
	golog.NewLine("\n")
	// golog.InfoText("[INFO]", ColorString("info", "[INFO]"))
	// golog.WarnText("[WARN]", ColorString("warn", "[WARN]"))
	// golog.ErrorText("[ERRO]", ColorString("error", "[ERRO]"))
	// golog.DebugText("[DBUG]", ColorString("debug", "[DBUG]"))
	golog.InfoText("[INFO]", InfoColorString("[INFO]"))
	golog.WarnText("[WARN]", WarnColorString("[WARN]"))
	golog.ErrorText("[ERRO]", ErrorColorString("[ERRO]"))
	golog.DebugText("[DBUG]", DebugColorString("[DBUG]"))
	// Default Output is `os.Stderr`, but you can change it:
	// glSetOutput(os.Stdout)
	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	Glog.SetTimeFormat("")
	Glog.SetLevel("info")
}

//---------------------------------------------------------

// Level is a number which defines the log level.
type Level golog.Level

// The available log levels.
const (
	// DisableLevel will disable printer
	DisableLevel Level = iota
	// ErrorLevel will print only errors
	ErrorLevel
	// WarnLevel will print errors and warnings
	WarnLevel
	// InfoLevel will print errors, warnings and infos
	InfoLevel
	// DebugLevel will print on any level, errors, warnings, infos and debug messages
	DebugLevel
)

var (
	// ColorStringFuncs maps to a serious of [*]ColorString functions with key as golog.Level
	ColorStringFuncs = map[Level]func(string) string{
		DisableLevel: NonColorString,
		ErrorLevel:   ErrorColorString,
		WarnLevel:    WarnColorString,
		InfoLevel:    InfoColorString,
		DebugLevel:   DebugColorString,
	}
)

// fromLevelName return Level code w.r.t levelname
func fromLevelName(levelName string) Level {
	switch levelName {
	case "error":
		return ErrorLevel
	case "warning":
		fallthrough
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	default:
		return DisableLevel
	}
}

// MyLogger is a interface to inport logger
type MyLogger interface {
	// SetLevel accepts a string representation of
	// a `Level` and returns a `Level` value based on that "levelName".
	//
	// Available level names are:
	// "disable"
	// "error"
	// "warn"
	// "info"
	// "debug"
	//
	// Alternatively you can use the exported `Default.Level` field, i.e `Default.Level = golog.ErrorLevel`
	SetLevel(levelName string)
	// SetTimeFormat sets time format for logs,
	// if "s" is empty then time representation will be off.
	SetTimeFormat(s string)
	// AddOutput adds one or more `io.Writer` to the Default Logger's Printer.
	//
	// If one of the "writers" is not a terminal-based (i.e File)
	// then colors will be disabled for all outputs.
	AddOutput(writers ...io.Writer)
	// Print prints a log message without levels and colors.
	Print(v ...interface{})
	// Println prints a log message without levels and colors.
	// It adds a new line at the end.
	Println(v ...interface{})
	// Error will print only when logger's Level is error.
	Error(v ...interface{})
	// Errorf will print only when logger's Level is error.
	Errorf(format string, args ...interface{})
	// Warn will print when logger's Level is error, or warning.
	Warn(v ...interface{})
	// Warnf will print when logger's Level is error, or warning.
	Warnf(format string, args ...interface{})
	// Info will print when logger's Level is error, warning or info.
	Info(v ...interface{})
	// Infof will print when logger's Level is error, warning or info.
	Infof(format string, args ...interface{})
	// Debug will print when logger's Level is error, warning,info or debug.
	Debug(v ...interface{})
	// Debugf will print when logger's Level is error, warning,info or debug.
	Debugf(format string, args ...interface{})
}

// ColorString gives a colorful string w.r.t. levelname
func ColorString(levelname, s string) string {
	if !glog.Printer.IsTerminal {
		return s
	}
	return ColorStringFuncs[fromLevelName(levelname)](s)
	// return ColorStringFuncs[Level(glog.Level)](s)
}

// NonColorString return a colorful info-string
func NonColorString(s string) string {
	return s
}

// InfoColorString return a colorful info-string
func InfoColorString(s string) string {
	return color.HiCyanString(s)
}

// WarnColorString return a colorful warn-string
func WarnColorString(s string) string {
	return color.HiGreenString(s)
}

// ErrorColorString return a colorful error-string
func ErrorColorString(s string) string {
	return color.HiGreenString(s)
}

// DebugColorString return a colorful error-string
func DebugColorString(s string) string {
	return fmt.Sprintf("[38;5;202m"+"%s"+"[0m", s)
}

//---------------------------------------------------------
var (
	// Glog           = Golog //&Mlogger{}
	errorWithColor = pio.Red
	warnWithColor  = pio.Purple
	infoWithColor  = pio.LightGreen
	debugWithColor = pio.Yellow
)

// Level is a number which defines the log level.
type colorLevel uint32

type colorCode uint32

// const (
// 	colorCyan colorCode =
// )

const (
	nonColor colorLevel = iota
	errorColor
	warnColor
	infoColor
	debugColor
)

type Mlogger golog.Logger

func colormsg(clevel colorLevel, fmt string, prm ...interface{}) string {
	var msg string
	switch clevel {
	case errorColor:
		msg = errorWithColor(Sf(fmt, prm...))
	case warnColor:
		msg = warnWithColor(Sf(fmt, prm...))
	case infoColor:
		msg = infoWithColor(Sf(fmt, prm...))
	case debugColor:
		msg = debugWithColor(Sf(fmt, prm...))
	default: // nonColor
		msg = Sf(fmt, prm...)
	}
	return msg
}

func (*Mlogger) Debugf(fmt string, prm ...interface{}) {
	// msg := colormsg(debugColor, fmt, prm...)
	// Golog.Debugf(msg,prm...)
	Glog.Debugf(fmt, prm...)
}
func (*Mlogger) Debug(prm ...interface{}) {
	// msg := colormsg(debugColor, fmt, prm...)
	// Glog.Debugf(msg,prm...)
	Glog.Debug(prm...)
}

func (*Mlogger) Warnf(fmt string, prm ...interface{}) {
	// msg := colormsg(warnColor, fmt, prm...)
	// Glog.Warnf(msg)
	Glog.Warnf(fmt, prm...)
}

func (*Mlogger) Warn(prm ...interface{}) {
	// msg := colormsg(warnColor, fmt, prm...)
	// Glog.Warnf(msg)
	Glog.Warn(prm...)
}

func (*Mlogger) Infof(fmt string, prm ...interface{}) {
	// msg := colormsg(infoColor, fmt, prm...)
	// Glog.Infof(msg)
	Glog.Infof(fmt, prm...)
}
func (*Mlogger) Info(prm ...interface{}) {
	// msg := colormsg(infoColor, fmt, prm...)
	// Glog.Infof(msg)
	Glog.Info(prm...)
}

func (*Mlogger) Errorf(fmt string, prm ...interface{}) {
	// msg := colormsg(debugColor, fmt, prm...)
	// Glog.Errorf(msg)
	Glog.Errorf(fmt, prm...)
}
func (*Mlogger) Error(prm ...interface{}) {
	// msg := colormsg(debugColor, fmt, prm...)
	// Glog.Errorf(msg)
	Glog.Error(prm...)
}

func (*Mlogger) Println(prm ...interface{}) {
	Glog.Println(prm...)
}

func (*Mlogger) Print(prm ...interface{}) {
	Glog.Print(prm...)
}

//---------------------------------------------------------

// TodayFilename get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func TodayFilename() string {
	// today := time.Now().Format("Jan 02 2006")
	today := time.Now().Format("2006-01-02")
	return today + ".log"
}

// NewLogFile get a new logging file handler
func NewLogFile() (*os.File, error) {
	filename := TodayFilename()
	// open an output file, this will append to the today's file if server restarted.
	// f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
}
