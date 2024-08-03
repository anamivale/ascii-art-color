package colors

import (
	"fmt"
	"strings"
)

func PrintColored(color string) (int, int, int, error) {
	color = strings.ToLower(color)
	type ColorMap map[string][3]int

	colors := ColorMap{
		"black":            [3]int{0, 0, 0},
		"orange":			[3]int{255,165,0},
		"red":              [3]int{255, 0, 0},
		"green":            [3]int{0, 128, 0},
		"yellow":           [3]int{255, 255, 0},
		"blue":             [3]int{0, 0, 255},
		"magenta":          [3]int{255, 0, 255},
		"cyan":             [3]int{0, 255, 255},
		"white":            [3]int{255, 255, 255},
		"gray":             [3]int{128, 128, 128},
		"darkred":          [3]int{139, 0, 0},
		"darkgreen":        [3]int{0, 100, 0},
		"darkyellow":       [3]int{139, 139, 0},
		"darkblue":         [3]int{0, 0, 139},
		"darkmagenta":      [3]int{139, 0, 139},
		"darkcyan":         [3]int{0, 139, 139},
		"lightgray":        [3]int{211, 211, 211},
		"lightred":         [3]int{255, 102, 102},
		"lightgreen":       [3]int{144, 238, 144},
		"lightyellow":      [3]int{255, 255, 224},
		"lightblue":        [3]int{173, 216, 230},
		"lightmagenta":     [3]int{255, 182, 193},
		"lightcyan":        [3]int{224, 255, 255},
		"pink":             [3]int{255, 192, 203},
		"peach":            [3]int{255, 218, 185},
		"lavender":         [3]int{230, 230, 250},
		"turquoise":        [3]int{64, 224, 208},
		"violet":           [3]int{238, 130, 238},
		"gold":             [3]int{255, 215, 0},
		"silver":           [3]int{192, 192, 192},
		"maroon":           [3]int{128, 0, 0},
		"olive":            [3]int{128, 128, 0},
		"navy":             [3]int{0, 0, 128},
		"purple":           [3]int{128, 0, 128},
		"teal":             [3]int{0, 128, 128},
		"bisque":           [3]int{255, 228, 196},
		"skyblue":          [3]int{135, 206, 235},
		"salmon":           [3]int{250, 128, 114},
		"khaki":            [3]int{240, 230, 140},
		"orchid":           [3]int{218, 112, 214},
		"thistle":          [3]int{216, 191, 216},
		"tan":              [3]int{210, 180, 140},
		"sienna":           [3]int{160, 82, 45},
		"slategray":        [3]int{112, 128, 144},
		"wheat":            [3]int{245, 222, 179},
		"indigo":           [3]int{75, 0, 130},
		"peru":             [3]int{205, 133, 63},
		"burlywood":        [3]int{222, 184, 135},
		"firebrick":        [3]int{178, 34, 34},
		"cadetblue":        [3]int{95, 158, 160},
		"darkorchid":       [3]int{153, 50, 204},
		"darkslategray":    [3]int{47, 79, 79},
		"darkgoldenrod":    [3]int{184, 134, 11},
		"rosybrown":        [3]int{188, 143, 143},
		"mediumaquamarine": [3]int{102, 205, 170},
		"darkseagreen":     [3]int{143, 188, 143},
		"slateblue":        [3]int{106, 90, 205},
		"mediumslateblue":  [3]int{123, 104, 238},
		"sandybrown":       [3]int{244, 164, 96},
		"darkkhaki":        [3]int{189, 183, 107},
		"palevioletred":    [3]int{219, 112, 147},
		"mediumvioletred":  [3]int{199, 21, 133},
	}
	r,g,b := 0,0,0
	if strings.HasPrefix(color,"rgb"){
		r,g,b, err :=  RgbTo256ColorCode(color)
		if err != nil{
			return 0,0,0,err
		}
		return r,g,b ,nil
	} else if strings.HasPrefix(color,"#"){
		r,g,b, err :=  HexTo256(color)
		if err != nil{
			return 0,0,0,err
		}
		return r,g,b ,nil
	} else {
	if value, ok := colors[color]; ok {
		r = value[0]
		g = value[1]
		b = value[2]
	} else {
		return 0, 0, 0, fmt.Errorf(" Error: %v is not available",color)
	}
	} 
	return r, g, b, nil

}

// HexTo256 converts a hex color string (e.g., "FFA500") to a 256-color code.
// Returns an error if the input format is incorrect.
func HexTo256(hex string) (int,int,int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0,0,0, fmt.Errorf("invalid Hex color code lenth: %v", hex)
	}

	var r, g, b int
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b) // Parse hex into RGB values.
	if err != nil {
		return 0,0,0, fmt.Errorf("invalid Hex color code: %v", hex)
	}
	return r, g, b, nil // Convert RGB to 256-color code.
}

// RgbTo256ColorCode converts a string representation of an RGB color
// (e.g., "rgb(255,165,0)") to a 256-color code.
// Returns an error if the input format is incorrect.
func RgbTo256ColorCode(rgb string) (int,int,int, error) {
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")

	var r, g, b int
	_, err := fmt.Sscanf(rgb, "%d,%d,%d", &r, &g, &b) // Parse RGB values.
	if err != nil {
		return 0,0,0, fmt.Errorf("invalid RGB color code: format %v", rgb)
	}
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return 0,0,0, fmt.Errorf("RGB values must be in the range 0-255: %v", rgb)
	}
	return r, g, b, nil
}
