package screen

import (
	"aqw/internal/window"
	"image/png"
	"os"
	"testing"
)

func TestCaptureScreen(t *testing.T) {
	// Define the region to capture
	win := &window.Window{ //REPLACE with actual window capture logic
		ID:     "60817424",
		X:      70,
		Y:      0, // Adjusted Y to 0 for full screen capture
		Width:  2490,
		Height: 1536,
	}

	// Capture the region
	img, err := CaptureScreen(win)
	if err != nil {
		t.Fatalf("Failed to capture screen region: %v", err)
	}

	// Check if the image is not nil and has the expected size
	if img == nil {
		t.Fatal("Captured image is nil")
	}
	if img.Bounds().Dx() != win.Width || img.Bounds().Dy() != win.Height {
		t.Errorf("Captured image size mismatch: got %dx%d, want %dx%d",
			img.Bounds().Dx(), img.Bounds().Dy(), win.Width, win.Height,
		)
	}

	// Save the image to a file for manual inspection
	file, err := os.Create("screenshot.png")
	if err != nil {
		t.Fatalf("❌ Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Fatalf("❌ Failed to encode PNG: %v\n", err)
		return
	}
}

func TestCaptureRegion(t *testing.T) {
	// Define the window and region to capture
	win := &window.Window{ //REPLACE with actual window capture logic
		ID:     "65011728",
		X:      70,
		Y:      0, // Adjusted Y to 0 for full screen capture
		Width:  2490,
		Height: 1536,
	}

	// Capture the region
	img, err := CaptureRegion(win, PlayerStatsRegion)
	if err != nil {
		t.Fatalf("Failed to capture screen region: %v", err)
	}

	// Check if the image is not nil and has the expected size
	if img == nil {
		t.Fatal("Captured image is nil")
	}
	if img.Bounds().Dx() != PlayerStatsRegion.Width || img.Bounds().Dy() != PlayerStatsRegion.Height {
		t.Errorf("Captured image size mismatch: got %dx%d, want %dx%d",
			img.Bounds().Dx(), img.Bounds().Dy(), PlayerStatsRegion.Width, PlayerStatsRegion.Height,
		)
	}

	// Save the image to a file for manual inspection
	file, err := os.Create("screenshot.png")
	if err != nil {
		t.Fatalf("❌ Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Fatalf("❌ Failed to encode PNG: %v\n", err)
		return
	}
}
