package colors

import (
	"fmt"
	"strings"
)
// RgbTo256 converts RGB values to a 256-color code value
// based on the 6x6x6 color cube used in trminal color schemes.
func RgbTo256(r, g, b int) int {
	var rI, gI, bI int = r / 51, g / 51, b / 51 // Reduce each color component to fit in 6 levels (0-5).
	return  16 + 36*rI + 6 * gI + bI  // Claculate 256-color code.
}

// HexTo256 converts a hex color string (e.g., "FFA500") to a 256-color code.
// Returns an error if the input format is incorrect.
func HexTo256(hex string) (int, error) {
	hex = strings.TrimPrefix(hex, "H")
	if len(hex) != 6 {
		return 0, fmt.Errorf("invalid Hex color code lenth: %v", hex)
	}

	var r, g, b int
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b) // Parse hex into RGB values.
	if err != nil {
		return 0, fmt.Errorf("invalid Hex color code: %v", hex)
	}
	return RgbTo256(r, g, b), nil // Convert RGB to 256-color code.
}

// RgbTo256ColorCode converts a string representation of an RGB color
// (e.g., "rgb(255,165,0)") to a 256-color code.
// Returns an error if the input format is incorrect.
func RgbTo256ColorCode(rgb string) (int, error) {
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")

	var r, g, b int
	_, err := fmt.Sscanf(rgb, "%d,%d,%d", &r, &g, &b) // Parse RGB values.
	if err != nil {
		return 0, fmt.Errorf("invalid RGB color code: format %v", rgb)
	}
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return 0, fmt.Errorf("RGB values must be in the range 0-255: %v", rgb)
	}
	return RgbTo256(r, g, b), nil
}
