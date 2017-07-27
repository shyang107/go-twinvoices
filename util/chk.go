package util

import (
	"fmt"
	"runtime"
)

// Panic panicks
func Panic(msg string, prm ...interface{}) {
	CallerInfo(4)
	CallerInfo(3)
	CallerInfo(2)
	panic(fmt.Sprintf(msg, prm...))
}

// CallerInfo returns the file and line positions where an error occurred
//  idx -- use idx=2 to get the caller of Panic
func CallerInfo(idx int) {
	pc, file, line, ok := runtime.Caller(idx)
	if !ok {
		file, line = "?", 0
	}
	var fname string
	f := runtime.FuncForPC(pc)
	if f != nil {
		fname = f.Name()
	}
	if Verbose {
		fmt.Printf("file = %s:%d\n", file, line)
		fmt.Printf("func = %s\n", fname)
	}
}
