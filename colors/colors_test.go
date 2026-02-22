package colors

import "testing"

import "math"

func hslEqual(a, b HSL) bool {
	tolerance := 0.001
	return math.Abs(a.H-b.H) < tolerance &&
		math.Abs(a.S-b.S) < tolerance &&
		math.Abs(a.L-b.L) < tolerance
}

func runTests[In, Out comparable](t *testing.T, fn func(In) Out, tests []struct {
	name     string
	input    In
	expected Out
}, equals func(Out, Out) bool) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fn(test.input)
			if !equals(result, test.expected) {
				t.Errorf("input %v: got %v: want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestRgbToHsl(t *testing.T) {

	tests := []struct {
		name     string
		input    RGB
		expected HSL
	}{
		{
			name:     "white",
			input:    RGB{255, 255, 255},
			expected: HSL{0.0, 0.0, 1.0},
		},
		{
			name:     "black",
			input:    RGB{0, 0, 0},
			expected: HSL{0.0, 0.0, 0.0},
		},
		{
			name:     "#335c99",
			input:    RGB{51, 92, 153},
			expected: HSL{0.60, 0.50, 0.40},
		},
		{
			name:     "pure red",
			input:    RGB{255, 0, 0},
			expected: HSL{0.0, 1.0, 0.50},
		},
		{
			name:     "pure green",
			input:    RGB{0, 255, 0},
			expected: HSL{0.333, 1.0, 0.50},
		},
		{
			name:     "pure blue",
			input:    RGB{0, 0, 255},
			expected: HSL{0.667, 1.0, 0.50},
		},
	}

	runTests(t, RgbToHsl, tests, hslEqual)

}

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected RGB
	}{
		{
			name:     "valid hex color",
			input:    "4e8f3d",
			expected: RGB{R: 78, G: 143, B: 61},
		},
		{
			name:     "black",
			input:    "000000",
			expected: RGB{R: 0, G: 0, B: 0},
		},
		{
			name:     "white",
			input:    "ffffff",
			expected: RGB{R: 255, G: 255, B: 255},
		},
		{
			name:     "uppercase",
			input:    "B44B4B",
			expected: RGB{R: 180, G: 75, B: 75},
		},
	}

	runTests(t, HexToRgb, tests, func(a, b RGB) bool { return a == b })
}
