package util

import (
	"image/color"

	"github.com/stroborobo/ansirgb"
)

const (
	escape  = "\x1b" // "\e" (033 or escape control code)
	reset   = "\x1b[0m"
	fgClose = "\x1b[39m"
	bgClose = "\x1b[49m"
	clear   = "\x1b[39;49m"

	fgsvg = "\x1b[38;2;"
	bgsvg = "\x1b[48;2;"

	fg256 = "\x1b[38;5;"
	bg256 = "\x1b[48;5;"

	colorsep = "|"
)

// Base attributes
const (
	Reset        Attribute = iota // Reset / Normal: all attributes off
	Bold                          // Bold or increased intensity
	Faint                         // Faint (decreased intensity): Not widely supported.
	Italic                        // Italic: on: Not widely supported. Sometimes treated as inverse.
	Underline                     // Underline: Single
	BlinkSlow                     // less than 150 per minute
	BlinkRapid                    // MS-DOS ANSI.SYS; 150+ per minute; not widely supported
	ReverseVideo                  // Image: Negative: inverse or reverse; swap foreground and background
	Concealed                     // Not widely supported.
	CrossedOut                    // Characters legible, but marked for deletion. Not widely supported.
)

// PaletteANSI is a palette of ANSI-colors.
//	0-16: ignore 8 Bit colors, they are very dependent on the user's terminal
//	while other colors usually aren't.
//	216 colors: 16-232; 6 * 6 * 6 = 216 colors
//	grayscale: 233-255
//	transparent: 256
var PaletteANSI = []ansirgb.Color{
	// 6 * 6 * 6 = 216 colors
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0x00, 0xff}, Code: 16},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0x00, 0xff}, Code: 17},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0x00, 0xff}, Code: 18},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0x00, 0xff}, Code: 19},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0x00, 0xff}, Code: 20},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0x00, 0xff}, Code: 21},
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0x5f, 0xff}, Code: 22},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0x5f, 0xff}, Code: 23},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0x5f, 0xff}, Code: 24},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0x5f, 0xff}, Code: 25},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0x5f, 0xff}, Code: 26},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0x5f, 0xff}, Code: 27},
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0x87, 0xff}, Code: 28},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0x87, 0xff}, Code: 29},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0x87, 0xff}, Code: 30},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0x87, 0xff}, Code: 31},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0x87, 0xff}, Code: 32},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0x87, 0xff}, Code: 33},
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0xaf, 0xff}, Code: 34},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0xaf, 0xff}, Code: 35},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0xaf, 0xff}, Code: 36},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0xaf, 0xff}, Code: 37},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0xaf, 0xff}, Code: 38},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0xaf, 0xff}, Code: 39},
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0xd7, 0xff}, Code: 40},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0xd7, 0xff}, Code: 41},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0xd7, 0xff}, Code: 42},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0xd7, 0xff}, Code: 43},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0xd7, 0xff}, Code: 44},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0xd7, 0xff}, Code: 45},
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0xff, 0xff}, Code: 46},
	ansirgb.Color{Color: color.RGBA{0x00, 0x5f, 0xff, 0xff}, Code: 47},
	ansirgb.Color{Color: color.RGBA{0x00, 0x87, 0xff, 0xff}, Code: 48},
	ansirgb.Color{Color: color.RGBA{0x00, 0xaf, 0xff, 0xff}, Code: 49},
	ansirgb.Color{Color: color.RGBA{0x00, 0xd7, 0xff, 0xff}, Code: 50},
	ansirgb.Color{Color: color.RGBA{0x00, 0xff, 0xff, 0xff}, Code: 51},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0x00, 0xff}, Code: 52},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0x00, 0xff}, Code: 53},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x00, 0xff}, Code: 54},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0x00, 0xff}, Code: 55},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0x00, 0xff}, Code: 56},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0x00, 0xff}, Code: 57},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0x5f, 0xff}, Code: 58},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0x5f, 0xff}, Code: 59},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x5f, 0xff}, Code: 60},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0x5f, 0xff}, Code: 61},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0x5f, 0xff}, Code: 62},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0x5f, 0xff}, Code: 63},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0x87, 0xff}, Code: 64},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0x87, 0xff}, Code: 65},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x87, 0xff}, Code: 66},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0x87, 0xff}, Code: 67},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0x87, 0xff}, Code: 68},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0x87, 0xff}, Code: 69},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0xaf, 0xff}, Code: 70},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0xaf, 0xff}, Code: 71},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0xaf, 0xff}, Code: 72},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0xaf, 0xff}, Code: 73},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0xaf, 0xff}, Code: 74},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0xaf, 0xff}, Code: 75},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0xd7, 0xff}, Code: 76},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0xd7, 0xff}, Code: 77},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0xd7, 0xff}, Code: 78},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0xd7, 0xff}, Code: 79},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0xd7, 0xff}, Code: 80},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0xd7, 0xff}, Code: 81},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x00, 0xff, 0xff}, Code: 82},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0xff, 0xff}, Code: 83},
	ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0xff, 0xff}, Code: 84},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0xff, 0xff}, Code: 85},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0xff, 0xff}, Code: 86},
	ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0xff, 0xff}, Code: 87},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0x00, 0xff}, Code: 88},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0x00, 0xff}, Code: 89},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x00, 0xff}, Code: 90},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0x00, 0xff}, Code: 91},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x00, 0xff}, Code: 92},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0x00, 0xff}, Code: 93},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0x5f, 0xff}, Code: 94},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0x5f, 0xff}, Code: 95},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x5f, 0xff}, Code: 96},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0x5f, 0xff}, Code: 97},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x5f, 0xff}, Code: 98},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0x5f, 0xff}, Code: 99},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0x87, 0xff}, Code: 100},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0x87, 0xff}, Code: 101},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x87, 0xff}, Code: 102},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0x87, 0xff}, Code: 103},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x87, 0xff}, Code: 104},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0x87, 0xff}, Code: 105},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0xaf, 0xff}, Code: 106},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0xaf, 0xff}, Code: 107},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xaf, 0xff}, Code: 108},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0xaf, 0xff}, Code: 109},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0xaf, 0xff}, Code: 110},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0xaf, 0xff}, Code: 111},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0xd7, 0xff}, Code: 112},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0xd7, 0xff}, Code: 113},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xd7, 0xff}, Code: 114},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0xd7, 0xff}, Code: 115},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0xd7, 0xff}, Code: 116},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0xd7, 0xff}, Code: 117},
	ansirgb.Color{Color: color.RGBA{0x87, 0x00, 0xff, 0xff}, Code: 118},
	ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0xff, 0xff}, Code: 119},
	ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xff, 0xff}, Code: 120},
	ansirgb.Color{Color: color.RGBA{0x87, 0xaf, 0xff, 0xff}, Code: 121},
	ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0xff, 0xff}, Code: 122},
	ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0xff, 0xff}, Code: 123},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0x00, 0xff}, Code: 124},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0x00, 0xff}, Code: 125},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0x00, 0xff}, Code: 126},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0x00, 0xff}, Code: 127},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0x00, 0xff}, Code: 128},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0x00, 0xff}, Code: 129},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0x5f, 0xff}, Code: 130},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0x5f, 0xff}, Code: 131},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0x5f, 0xff}, Code: 132},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0x5f, 0xff}, Code: 133},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0x5f, 0xff}, Code: 134},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0x5f, 0xff}, Code: 135},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0x87, 0xff}, Code: 136},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0x87, 0xff}, Code: 137},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0x87, 0xff}, Code: 138},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0x87, 0xff}, Code: 139},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0x87, 0xff}, Code: 140},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0x87, 0xff}, Code: 141},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0xaf, 0xff}, Code: 142},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0xaf, 0xff}, Code: 143},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0xaf, 0xff}, Code: 144},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0xaf, 0xff}, Code: 145},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xaf, 0xff}, Code: 146},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0xaf, 0xff}, Code: 147},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0xd7, 0xff}, Code: 148},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0xd7, 0xff}, Code: 149},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0xd7, 0xff}, Code: 150},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0xd7, 0xff}, Code: 151},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xd7, 0xff}, Code: 152},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0xd7, 0xff}, Code: 153},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x00, 0xff, 0xff}, Code: 154},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0xff, 0xff}, Code: 155},
	ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0xff, 0xff}, Code: 156},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xaf, 0xff, 0xff}, Code: 157},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xff, 0xff}, Code: 158},
	ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0xff, 0xff}, Code: 159},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0x00, 0xff}, Code: 160},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x00, 0xff}, Code: 161},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x00, 0xff}, Code: 162},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0x00, 0xff}, Code: 163},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0x00, 0xff}, Code: 164},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0x00, 0xff}, Code: 165},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0x5f, 0xff}, Code: 166},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x5f, 0xff}, Code: 167},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x5f, 0xff}, Code: 168},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0x5f, 0xff}, Code: 169},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0x5f, 0xff}, Code: 170},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0x5f, 0xff}, Code: 171},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0x87, 0xff}, Code: 172},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x87, 0xff}, Code: 173},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x87, 0xff}, Code: 174},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0x87, 0xff}, Code: 175},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0x87, 0xff}, Code: 176},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0x87, 0xff}, Code: 177},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0xaf, 0xff}, Code: 178},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0xaf, 0xff}, Code: 179},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0xaf, 0xff}, Code: 180},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0xaf, 0xff}, Code: 181},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0xaf, 0xff}, Code: 182},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0xaf, 0xff}, Code: 183},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0xd7, 0xff}, Code: 184},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0xd7, 0xff}, Code: 185},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0xd7, 0xff}, Code: 186},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0xd7, 0xff}, Code: 187},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0xd7, 0xff}, Code: 188},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0xd7, 0xff}, Code: 189},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x00, 0xff, 0xff}, Code: 190},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0xff, 0xff}, Code: 191},
	ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0xff, 0xff}, Code: 192},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xaf, 0xff, 0xff}, Code: 193},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0xff, 0xff}, Code: 194},
	ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0xff, 0xff}, Code: 195},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0x00, 0xff}, Code: 196},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x00, 0xff}, Code: 197},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0x00, 0xff}, Code: 198},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0x00, 0xff}, Code: 199},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0x00, 0xff}, Code: 200},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x00, 0xff}, Code: 201},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0x5f, 0xff}, Code: 202},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x5f, 0xff}, Code: 203},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0x5f, 0xff}, Code: 204},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0x5f, 0xff}, Code: 205},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0x5f, 0xff}, Code: 206},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x5f, 0xff}, Code: 207},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0x87, 0xff}, Code: 208},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x87, 0xff}, Code: 209},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0x87, 0xff}, Code: 210},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0x87, 0xff}, Code: 211},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0x87, 0xff}, Code: 212},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x87, 0xff}, Code: 213},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0xaf, 0xff}, Code: 214},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0xaf, 0xff}, Code: 215},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0xaf, 0xff}, Code: 216},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xaf, 0xff}, Code: 217},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xaf, 0xff}, Code: 218},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xaf, 0xff}, Code: 219},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0xd7, 0xff}, Code: 220},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0xd7, 0xff}, Code: 221},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0xd7, 0xff}, Code: 222},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xd7, 0xff}, Code: 224},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xd7, 0xff}, Code: 225},
	ansirgb.Color{Color: color.RGBA{0xff, 0x00, 0xff, 0xff}, Code: 226},
	ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0xff, 0xff}, Code: 227},
	ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0xff, 0xff}, Code: 228},
	ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xff, 0xff}, Code: 229},
	ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	ansirgb.Color{Color: color.RGBA{0x08, 0x08, 0x08, 0xff}, Code: 232},
	// grayscale: 233-255
	ansirgb.Color{Color: color.RGBA{0x12, 0x12, 0x12, 0xff}, Code: 233},
	ansirgb.Color{Color: color.RGBA{0x1c, 0x1c, 0x1c, 0xff}, Code: 234},
	ansirgb.Color{Color: color.RGBA{0x26, 0x26, 0x26, 0xff}, Code: 235},
	ansirgb.Color{Color: color.RGBA{0x30, 0x30, 0x30, 0xff}, Code: 236},
	ansirgb.Color{Color: color.RGBA{0x3a, 0x3a, 0x3a, 0xff}, Code: 237},
	ansirgb.Color{Color: color.RGBA{0x44, 0x44, 0x44, 0xff}, Code: 238},
	ansirgb.Color{Color: color.RGBA{0x4e, 0x4e, 0x4e, 0xff}, Code: 239},
	ansirgb.Color{Color: color.RGBA{0x58, 0x58, 0x58, 0xff}, Code: 240},
	ansirgb.Color{Color: color.RGBA{0x62, 0x62, 0x62, 0xff}, Code: 241},
	ansirgb.Color{Color: color.RGBA{0x6c, 0x6c, 0x6c, 0xff}, Code: 242},
	ansirgb.Color{Color: color.RGBA{0x76, 0x76, 0x76, 0xff}, Code: 243},
	ansirgb.Color{Color: color.RGBA{0x80, 0x80, 0x80, 0xff}, Code: 244},
	ansirgb.Color{Color: color.RGBA{0x8a, 0x8a, 0x8a, 0xff}, Code: 245},
	ansirgb.Color{Color: color.RGBA{0x94, 0x94, 0x94, 0xff}, Code: 246},
	ansirgb.Color{Color: color.RGBA{0x9e, 0x9e, 0x9e, 0xff}, Code: 247},
	ansirgb.Color{Color: color.RGBA{0xa8, 0xa8, 0xa8, 0xff}, Code: 248},
	ansirgb.Color{Color: color.RGBA{0xb2, 0xb2, 0xb2, 0xff}, Code: 249},
	ansirgb.Color{Color: color.RGBA{0xbc, 0xbc, 0xbc, 0xff}, Code: 250},
	ansirgb.Color{Color: color.RGBA{0xc6, 0xc6, 0xc6, 0xff}, Code: 251},
	ansirgb.Color{Color: color.RGBA{0xd0, 0xd0, 0xd0, 0xff}, Code: 252},
	ansirgb.Color{Color: color.RGBA{0xda, 0xda, 0xda, 0xff}, Code: 253},
	ansirgb.Color{Color: color.RGBA{0xe4, 0xe4, 0xe4, 0xff}, Code: 254},
	ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	// transparent
	ansirgb.Color{Color: color.RGBA{0x00, 0x00, 0x00, 0x00}, Code: -1},
}

// MapSVG2ANSI contains named colors defined in the SVG 1.1 spec. convert to ANSI colors
var MapSVG2ANSI = map[string]ansirgb.Color{
	"indianred":            ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x5f, 0xff}, Code: 167},
	"red":                  ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0x0, 0xff}, Code: 196},
	"slategray":            ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x87, 0xff}, Code: 66},
	"cadetblue":            ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0xaf, 0xff}, Code: 73},
	"chartreuse":           ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0xff, 0xff}, Code: 118},
	"darkgoldenrod":        ansirgb.Color{Color: color.RGBA{0xaf, 0x0, 0x87, 0xff}, Code: 136},
	"dodgerblue":           ansirgb.Color{Color: color.RGBA{0x0, 0xff, 0x87, 0xff}, Code: 33},
	"darkslateblue":        ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x5f, 0xff}, Code: 60},
	"grey":                 ansirgb.Color{Color: color.RGBA{0x80, 0x80, 0x80, 0xff}, Code: 244},
	"lemonchiffon":         ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"antiquewhite":         ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xd7, 0xff}, Code: 224},
	"blue":                 ansirgb.Color{Color: color.RGBA{0x0, 0xff, 0x0, 0xff}, Code: 21},
	"blueviolet":           ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x0, 0xff}, Code: 92},
	"coral":                ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x87, 0xff}, Code: 209},
	"navy":                 ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0x0, 0xff}, Code: 18},
	"snow":                 ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"black":                ansirgb.Color{Color: color.RGBA{0x0, 0x0, 0x0, 0xff}, Code: 16},
	"darkcyan":             ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0x87, 0xff}, Code: 30},
	"lightcyan":            ansirgb.Color{Color: color.RGBA{0xd7, 0xff, 0xff, 0xff}, Code: 195},
	"lightgrey":            ansirgb.Color{Color: color.RGBA{0xd0, 0xd0, 0xd0, 0xff}, Code: 252},
	"chocolate":            ansirgb.Color{Color: color.RGBA{0xd7, 0x0, 0x5f, 0xff}, Code: 166},
	"darkorange":           ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0x87, 0xff}, Code: 208},
	"tan":                  ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0xaf, 0xff}, Code: 180},
	"thistle":              ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0xaf, 0xff}, Code: 182},
	"gray":                 ansirgb.Color{Color: color.RGBA{0x80, 0x80, 0x80, 0xff}, Code: 244},
	"olive":                ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0x87, 0xff}, Code: 100},
	"plum":                 ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0xaf, 0xff}, Code: 182},
	"tomato":               ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x5f, 0xff}, Code: 203},
	"aquamarine":           ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0xff, 0xff}, Code: 122},
	"azure":                ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"rosybrown":            ansirgb.Color{Color: color.RGBA{0xaf, 0x87, 0x87, 0xff}, Code: 138},
	"cornflowerblue":       ansirgb.Color{Color: color.RGBA{0x5f, 0xff, 0x87, 0xff}, Code: 69},
	"darkkhaki":            ansirgb.Color{Color: color.RGBA{0xaf, 0x5f, 0xaf, 0xff}, Code: 143},
	"green":                ansirgb.Color{Color: color.RGBA{0x0, 0x0, 0x87, 0xff}, Code: 28},
	"mediumspringgreen":    ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0xff, 0xff}, Code: 48},
	"darkgray":             ansirgb.Color{Color: color.RGBA{0xa8, 0xa8, 0xa8, 0xff}, Code: 248},
	"lightblue":            ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xd7, 0xff}, Code: 152},
	"lightslategrey":       ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x87, 0xff}, Code: 102},
	"navajowhite":          ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	"palegreen":            ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xff, 0xff}, Code: 120},
	"mistyrose":            ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xd7, 0xff}, Code: 224},
	"blanchedalmond":       ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xd7, 0xff}, Code: 224},
	"darkviolet":           ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x0, 0xff}, Code: 92},
	"greenyellow":          ansirgb.Color{Color: color.RGBA{0xaf, 0x0, 0xff, 0xff}, Code: 154},
	"mintcream":            ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"sandybrown":           ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0xaf, 0xff}, Code: 215},
	"darkgrey":             ansirgb.Color{Color: color.RGBA{0xa8, 0xa8, 0xa8, 0xff}, Code: 248},
	"darkturquoise":        ansirgb.Color{Color: color.RGBA{0x0, 0xd7, 0xd7, 0xff}, Code: 44},
	"lime":                 ansirgb.Color{Color: color.RGBA{0x0, 0x0, 0xff, 0xff}, Code: 46},
	"limegreen":            ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0xd7, 0xff}, Code: 77},
	"mediumslateblue":      ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0x5f, 0xff}, Code: 99},
	"mediumvioletred":      ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x0, 0xff}, Code: 162},
	"peru":                 ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x87, 0xff}, Code: 173},
	"saddlebrown":          ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0x5f, 0xff}, Code: 94},
	"bisque":               ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xd7, 0xff}, Code: 224},
	"lawngreen":            ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0xff, 0xff}, Code: 118},
	"lightsteelblue":       ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xd7, 0xff}, Code: 152},
	"lightyellow":          ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"seagreen":             ansirgb.Color{Color: color.RGBA{0x0, 0x5f, 0x87, 0xff}, Code: 29},
	"lightcoral":           ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0x87, 0xff}, Code: 210},
	"paleturquoise":        ansirgb.Color{Color: color.RGBA{0xaf, 0xff, 0xff, 0xff}, Code: 159},
	"purple":               ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x0, 0xff}, Code: 90},
	"dimgray":              ansirgb.Color{Color: color.RGBA{0x6c, 0x6c, 0x6c, 0xff}, Code: 242},
	"dimgrey":              ansirgb.Color{Color: color.RGBA{0x6c, 0x6c, 0x6c, 0xff}, Code: 242},
	"fuchsia":              ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x0, 0xff}, Code: 201},
	"ivory":                ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"floralwhite":          ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"palegoldenrod":        ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	"pink":                 ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xaf, 0xff}, Code: 218},
	"maroon":               ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0x0, 0xff}, Code: 88},
	"turquoise":            ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0xd7, 0xff}, Code: 80},
	"gainsboro":            ansirgb.Color{Color: color.RGBA{0xda, 0xda, 0xda, 0xff}, Code: 253},
	"honeydew":             ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	"indigo":               ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x0, 0xff}, Code: 54},
	"lightgreen":           ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xff, 0xff}, Code: 120},
	"steelblue":            ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0x87, 0xff}, Code: 67},
	"whitesmoke":           ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	"aqua":                 ansirgb.Color{Color: color.RGBA{0x0, 0xff, 0xff, 0xff}, Code: 51},
	"ghostwhite":           ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"lightgoldenrodyellow": ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"powderblue":           ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0xd7, 0xff}, Code: 152},
	"silver":               ansirgb.Color{Color: color.RGBA{0xbc, 0xbc, 0xbc, 0xff}, Code: 250},
	"burlywood":            ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0xaf, 0xff}, Code: 180},
	"darksalmon":           ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x87, 0xff}, Code: 174},
	"hotpink":              ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0x5f, 0xff}, Code: 205},
	"palevioletred":        ansirgb.Color{Color: color.RGBA{0xd7, 0x87, 0x5f, 0xff}, Code: 168},
	"darkorchid":           ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x5f, 0xff}, Code: 98},
	"linen":                ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	"yellowgreen":          ansirgb.Color{Color: color.RGBA{0x87, 0x5f, 0xd7, 0xff}, Code: 113},
	"salmon":               ansirgb.Color{Color: color.RGBA{0xff, 0x5f, 0x87, 0xff}, Code: 209},
	"beige":                ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"darkgreen":            ansirgb.Color{Color: color.RGBA{0x0, 0x0, 0x5f, 0xff}, Code: 22},
	"deepskyblue":          ansirgb.Color{Color: color.RGBA{0x0, 0xff, 0xaf, 0xff}, Code: 39},
	"lightskyblue":         ansirgb.Color{Color: color.RGBA{0x87, 0xff, 0xd7, 0xff}, Code: 117},
	"slategrey":            ansirgb.Color{Color: color.RGBA{0x5f, 0x87, 0x87, 0xff}, Code: 66},
	"violet":               ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x87, 0xff}, Code: 213},
	"white":                ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"aliceblue":            ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"brown":                ansirgb.Color{Color: color.RGBA{0xaf, 0x0, 0x0, 0xff}, Code: 124},
	"skyblue":              ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0xd7, 0xff}, Code: 116},
	"slateblue":            ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0x5f, 0xff}, Code: 62},
	"springgreen":          ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0xff, 0xff}, Code: 48},
	"teal":                 ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0x87, 0xff}, Code: 30},
	"cyan":                 ansirgb.Color{Color: color.RGBA{0x0, 0xff, 0xff, 0xff}, Code: 51},
	"deeppink":             ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0x0, 0xff}, Code: 198},
	"goldenrod":            ansirgb.Color{Color: color.RGBA{0xd7, 0x0, 0xaf, 0xff}, Code: 178},
	"midnightblue":         ansirgb.Color{Color: color.RGBA{0x0, 0x5f, 0x0, 0xff}, Code: 17},
	"firebrick":            ansirgb.Color{Color: color.RGBA{0xaf, 0x0, 0x0, 0xff}, Code: 124},
	"gold":                 ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0xd7, 0xff}, Code: 220},
	"lavenderblush":        ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Code: 231},
	"lightgray":            ansirgb.Color{Color: color.RGBA{0xd0, 0xd0, 0xd0, 0xff}, Code: 252},
	"sienna":               ansirgb.Color{Color: color.RGBA{0xaf, 0x0, 0x5f, 0xff}, Code: 130},
	"magenta":              ansirgb.Color{Color: color.RGBA{0xff, 0xff, 0x0, 0xff}, Code: 201},
	"mediumseagreen":       ansirgb.Color{Color: color.RGBA{0x5f, 0x5f, 0xaf, 0xff}, Code: 71},
	"orchid":               ansirgb.Color{Color: color.RGBA{0xd7, 0xd7, 0x5f, 0xff}, Code: 170},
	"seashell":             ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	"darkmagenta":          ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x0, 0xff}, Code: 90},
	"forestgreen":          ansirgb.Color{Color: color.RGBA{0x0, 0x0, 0x87, 0xff}, Code: 28},
	"lavender":             ansirgb.Color{Color: color.RGBA{0xee, 0xee, 0xee, 0xff}, Code: 255},
	"lightpink":            ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xaf, 0xff}, Code: 217},
	"darkseagreen":         ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0xaf, 0xff}, Code: 108},
	"khaki":                ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0xd7, 0xff}, Code: 222},
	"lightslategray":       ansirgb.Color{Color: color.RGBA{0x87, 0x87, 0x87, 0xff}, Code: 102},
	"mediumblue":           ansirgb.Color{Color: color.RGBA{0x0, 0xd7, 0x0, 0xff}, Code: 20},
	"cornsilk":             ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"moccasin":             ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	"wheat":                ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	"mediumorchid":         ansirgb.Color{Color: color.RGBA{0xaf, 0xd7, 0x5f, 0xff}, Code: 134},
	"oldlace":              ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"orange":               ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0xaf, 0xff}, Code: 214},
	"peachpuff":            ansirgb.Color{Color: color.RGBA{0xff, 0xaf, 0xd7, 0xff}, Code: 223},
	"darkslategray":        ansirgb.Color{Color: color.RGBA{0x44, 0x44, 0x44, 0xff}, Code: 238},
	"lightseagreen":        ansirgb.Color{Color: color.RGBA{0x0, 0xaf, 0xaf, 0xff}, Code: 37},
	"darkblue":             ansirgb.Color{Color: color.RGBA{0x0, 0x87, 0x0, 0xff}, Code: 18},
	"darkred":              ansirgb.Color{Color: color.RGBA{0x87, 0x0, 0x0, 0xff}, Code: 88},
	"mediumaquamarine":     ansirgb.Color{Color: color.RGBA{0x5f, 0xaf, 0xd7, 0xff}, Code: 79},
	"olivedrab":            ansirgb.Color{Color: color.RGBA{0x5f, 0x0, 0x87, 0xff}, Code: 64},
	"crimson":              ansirgb.Color{Color: color.RGBA{0xd7, 0x5f, 0x0, 0xff}, Code: 161},
	"darkslategrey":        ansirgb.Color{Color: color.RGBA{0x44, 0x44, 0x44, 0xff}, Code: 238},
	"mediumturquoise":      ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0xd7, 0xff}, Code: 80},
	"royalblue":            ansirgb.Color{Color: color.RGBA{0x5f, 0xd7, 0x5f, 0xff}, Code: 62},
	"yellow":               ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0xff, 0xff}, Code: 226},
	"papayawhip":           ansirgb.Color{Color: color.RGBA{0xff, 0xd7, 0xff, 0xff}, Code: 230},
	"darkolivegreen":       ansirgb.Color{Color: color.RGBA{0x4e, 0x4e, 0x4e, 0xff}, Code: 239},
	"lightsalmon":          ansirgb.Color{Color: color.RGBA{0xff, 0x87, 0xaf, 0xff}, Code: 216},
	"mediumpurple":         ansirgb.Color{Color: color.RGBA{0x87, 0xd7, 0x5f, 0xff}, Code: 98},
	"orangered":            ansirgb.Color{Color: color.RGBA{0xff, 0x0, 0x5f, 0xff}, Code: 202},
}
