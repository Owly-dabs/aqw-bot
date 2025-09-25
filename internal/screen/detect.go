package screen

import (
	"fmt"
	"image/color"

	"github.com/kbinani/screenshot"
)

// absDiff returns the absolute difference between two uint8 values
func absDiff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}

// isColorMatch checks if two colors are close enough within a tolerance
func isColorMatch(c1, c2 color.RGBA, tolerance uint8) bool {
	return absDiff(c1.R, c2.R) <= tolerance &&
		absDiff(c1.G, c2.G) <= tolerance &&
		absDiff(c1.B, c2.B) <= tolerance
}

// IsColorAt returns true if the pixel at (x, y) matches the target color within tolerance
func IsColorAt(x, y int, target color.RGBA, tolerance uint8) (bool, error) {
	bounds := screenshot.GetDisplayBounds(0)

	if x < bounds.Min.X || y < bounds.Min.Y || x >= bounds.Max.X || y >= bounds.Max.Y {
		return false, fmt.Errorf("coordinates (%d,%d) out of screen bounds X:(%d, %d), Y:(%d, %d)", 
			x, y, bounds.Min.X, bounds.Max.X, bounds.Min.Y, bounds.Max.Y,
		)
	}

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return false, err
	}

	r, g, b, _ := img.At(x, y).RGBA()
	actual := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: 255,
	}

	return isColorMatch(actual, target, tolerance), nil
}

func GetColorAt(x, y int) (color.RGBA, error) {
	bounds := screenshot.GetDisplayBounds(0)

	// Check if coordinates are within screen bounds
	if x < bounds.Min.X || y < bounds.Min.Y || x >= bounds.Max.X || y >= bounds.Max.Y {
		return color.RGBA{}, fmt.Errorf("coordinates (%d,%d) are out of screen bounds", x, y)
	}

	// Capture the full screen (or just bounds)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return color.RGBA{}, err
	}

	r, g, b, _ := img.At(x, y).RGBA()

	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: 255,
	}, nil
}