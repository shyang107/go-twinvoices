package invoices

import (
	"fmt"
	"os"
	"strings"
)

var (
	trim = strings.TrimSpace
	sf   = fmt.Sprintf
)

// math ---------------------------------------------------------

// imax returns the maximum value
func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// imax returns the minimum value
func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isum returns the summation
func isum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

// string ---------------------------------------------------------

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

// file ---------------------------------------------------------

// isFileExist checks whether a file exist
func isFileExist(filename string) bool {
	path := os.ExpandEnv(filename)
	isExist := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		isExist = false
	}
	return isExist
}

// print ---------------------------------------------------------
