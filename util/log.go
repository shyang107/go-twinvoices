package util

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/kataras/golog"
)

var (
	// Glog is a logger of golog
	Glog = golog.New()
	// Glog MyLogger = glog
)

// InitLogger initialize logger's environment
func InitLogger() {
	golog.NewLine("\n")
	golog.InfoText("[INFO]", InfoColorString("[INFO]"))
	golog.WarnText("[WARN]", WarnColorString("[WARN]"))
	golog.ErrorText("[ERRO]", ErrorColorString("[ERRO]"))
	golog.DebugText("[DBUG]", DebugColorString("[DBUG]"))
	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	// Glog.SetTimeFormat("2006/01/02 15:04:05")
	Glog.SetTimeFormat("")
	Glog.SetLevel("info")
	NoColor = !Glog.Printer.IsTerminal
}

// EnableLoggerOuttoFile enables the "Glog" out to os.Stderr and os.File
func EnableLoggerOuttoFile() {
	// file, err := os.OpenFile(util.TodayFilename(), os.O_CREATE|os.O_WRONLY, 0666)
	file, err := NewLogFile()
	// defer file.Close()
	if err == nil {
		Glog.AddOutput(file)
	} else {
		Glog.Errorf("Failed to log to file, using default stderr")
	}
	NoColor = true
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
	levelName = strings.ToLower(levelName)
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

// LogColorString gives a colorful string w.r.t. levelname
func LogColorString(levelname, s string) string {
	if !Glog.Printer.IsTerminal {
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
	return HiCyanString(s)
}

// WarnColorString return a colorful warn-string
func WarnColorString(s string) string {
	// return HiGreenString(s)
	return HiMagentaString(s)
}

// ErrorColorString return a colorful error-string
func ErrorColorString(s string) string {
	return HiGreenString(s)
}

// DebugColorString return a colorful error-string
func DebugColorString(s string) string {
	// return fmt.Sprintf("[38;5;202m"+"%s"+"[0m", s)
	// return Color256String("%s", EncodeColor256(202, true), s)
	return OrangeString("%s", s)
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
