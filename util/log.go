package util

import (
	"os"
	"strings"
	"time"

	"github.com/kataras/golog"
)

var (
	// Glog is a logger of golog
	Glog = golog.New()
)

// InitLogger initialize logger's environment
func InitLogger() {
	golog.NewLine("\n")
	golog.InfoText("[INFO]", LogColorString("info", "[INFO]"))
	golog.WarnText("[WARN]", LogColorString("warn", "[WARN]"))
	golog.ErrorText("[ERRO]", LogColorString("error", "[ERRO]"))
	golog.DebugText("[DBUG]", LogColorString("debug", "[DBUG]"))
	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	// Glog.SetTimeFormat("2006/01/02 15:04:05")
	Glog.SetTimeFormat("")
	Glog.SetLevel("info")
	NoColor = !Glog.Printer.IsTerminal
}

//---------------------------------------------------------

var (
	// LogColorStringFuncs maps to a serious of [*]ColorString functions with key as golog.Level
	LogColorStringFuncs = map[golog.Level]func(string) string{
		golog.DisableLevel: func(s string) string { return s },
		golog.ErrorLevel:   func(s string) string { return HiRedString(s) },
		golog.WarnLevel:    func(s string) string { return HiMagentaString(s) },
		golog.InfoLevel:    func(s string) string { return HiCyanString(s) },
		golog.DebugLevel:   func(s string) string { return OrangeString(s) },
	}
)

// fromLevelName return Level code w.r.t levelname
func fromLevelName(levelName string) golog.Level {
	switch strings.ToLower(levelName) {
	case "error":
		return golog.ErrorLevel
	case "warning":
		fallthrough
	case "warn":
		return golog.WarnLevel
	case "info":
		return golog.InfoLevel
	case "debug":
		return golog.DebugLevel
	default:
		return golog.DisableLevel
	}
}

// LogColorString gives a colorful string w.r.t. levelname
func LogColorString(levelname, s string) string {
	if !Glog.Printer.IsTerminal {
		return s
	}
	return LogColorStringFuncs[fromLevelName(levelname)](s)
	// return ColorStringFuncs[Level(glog.Level)](s)
}

//---------------------------------------------------------

// EnableLoggerOutToFile enables the "Glog" out to os.Stderr and os.File
func EnableLoggerOutToFile(levelname string) {
	// file, err := os.OpenFile(util.TodayFilename(), os.O_CREATE|os.O_WRONLY, 0666)
	Glog.SetLevel(levelname)
	file, err := NewLogFile()
	// defer file.Close()
	if err == nil {
		Glog.AddOutput(file)
	} else {
		Glog.Errorf("Failed to log to file, using default stderr")
	}
	NoColor = true
}

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
