package screen

import (
	"aqw/internal/window"
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
)

func CaptureScreen(win *window.Window) (*image.RGBA, error) {
	// Activate window
	if err := win.Activate(); err != nil {
		return nil, err
	}
	// capture the screen region based on the provided coordinates and dimensions
	return screenshot.Capture(win.X, win.Y, win.Width, win.Height)
}

type Region struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Constant for player stats region
var PlayerStatsRegion = Region{
	X:      150,
	Y:      100, // Adjusted Y to 0 for full screen capture
	Width:  300,
	Height: 300,
}

func CaptureRegion(win *window.Window, region Region) (*image.RGBA, error) {
	x, y, w, h := region.X, region.Y, region.Width, region.Height

	// Activate window
	if err := win.Activate(); err != nil {
		return nil, err
	}

	// Ensure the coordinates are within the window bounds
	if x < 0 || y < 0 || x+w > win.Width || y+h > win.Height {
		return nil, fmt.Errorf("capture region out of bounds: (%d, %d, %d, %d) exceeds window size (%d, %d)",
			x, y, w, h, win.Width, win.Height)
	}

	// capture the specified region of the screen
	return screenshot.Capture(x+win.X, y+win.Y, w, h)
}
