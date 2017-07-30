package util

import (
	"os"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/pio"
)

var (
	// Golog loggin information
	Glog = golog.New()
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

//-------------------

// TodayFilename get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func TodayFilename() string {
	// today := time.Now().Format("Jan 02 2006")
	today := time.Now().Format("2006-01-02")
	return today + ".log"
}

// NewLogFile get a new logging file handler
func NewLogFile() *os.File {
	filename := TodayFilename()
	// open an output file, this will append to the today's file if server restarted.
	// f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
