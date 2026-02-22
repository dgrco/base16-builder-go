package colors

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Values in the range 0-255
type RGB struct {
	R, G, B uint64
}

// Return normalized values for the RGB struct.
func (rgb RGB) Normalize() (float64, float64, float64) {
	return float64(rgb.R) / 255.0,
		float64(rgb.G) / 255.0,
		float64(rgb.B) / 255.0
}

// Values in the range 0.0-1.0
type HSL struct {
	H, S, L float64
}

func calculateHue(r, g, b, chroma, xMax float64) float64 {
	switch {
	case chroma == 0:
		return 0
	case xMax == r:
		return 60 * math.Mod((g-b)/chroma, 6)
	case xMax == g:
		return 60 * ((b-r)/chroma + 2)
	case xMax == b:
		return 60 * ((r-g)/chroma + 4)
	}
	return 0
}

func calculateSaturation(lightness float64, xMax float64) float64 {
	if lightness == 0 || lightness == 1 {
		return 0
	}
	return (xMax - lightness) / min(lightness, 1-lightness)
}

// Convert an RGB struct to an HSL struct.
func RgbToHsl(rgb RGB) HSL {

	r, g, b := rgb.Normalize()

	xMax := max(r, g, b)
	xMin := min(r, g, b)

	chroma := xMax - xMin

	lightness := xMax - chroma/2

	hue := calculateHue(r, g, b, chroma, xMax) / 360

	saturation := calculateSaturation(lightness, xMax)

	return HSL{hue, saturation, lightness}
}

// Converts a hexadecimal color string to an RGB struct.
//
// The hex string must be 6-characters long (0-9, a-f).
//
// example: "40bf73"
func HexToRgb(hex string) RGB {

	hex_r := hex[:2]
	hex_g := hex[2:4]
	hex_b := hex[4:6]

	return RGB{
		R: hexToDec(hex_r),
		G: hexToDec(hex_g),
		B: hexToDec(hex_b),
	}
}

// Convert a hexadecimal string to an integer.
//
// example hexToDec("FF") returns 255.
func hexToDec(hex string) uint64 {
	number, err := strconv.ParseUint(hex, 16, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing hex '%s': %v\n", hex, err)
	}
	return number
}

// Reverse a hex string.
func ReverseRgb(hex string) string {
	first := hex[:2]
	second := hex[2:4]
	third := hex[4:6]
	return third + second + first
}
