package pencil

import (
	"image/color/palette"
)

// PalettePlan9 is Plan9 in "image/color/palette"
//
// Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4 subdivision,
// with 4 shades in each subcube. Compared to the WebSafe, the idea is to reduce the color
// resolution by dicing the color cube into fewer cells, and to use the extra space to increase
// the intensity resolution. This results in 16 gray shades (4 gray subcubes with 4 samples in each),
// 13 shades of each primary and secondary color (3 subcubes with 4 samples plus black) and a
// reasonable selection of colors covering the rest of the color cube. The advantage is better
// representation of continuous tones.
var PalettePlan9 = palette.Plan9
