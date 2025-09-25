package screen

import (
	"image/png"
	"os"
	"testing"
)

func TestMakeReadable(t *testing.T) {
	file, err := os.Open("./screenshot.png")
	if err != nil {
		t.Fatalf("Failed to open screenshot: %v", err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf("Failed to decode PNG: %v", err)
	}

	// Assuming MakeReadable is a function that processes a screen image
	alteredImg, err := MakeReadable(img)
	if err != nil {
		t.Fatalf("Failed to make screen readable: %v", err)
	}

	// Save the image to a file for manual inspection
	newFile, err := os.Create("altered_screenshot.png")
	if err != nil {
		t.Fatalf("❌ Failed to create file: %v\n", err)
		return
	}
	defer newFile.Close()

	err = png.Encode(newFile, alteredImg)
	if err != nil {
		t.Fatalf("❌ Failed to encode PNG: %v\n", err)
		return
	}

	t.Logf("Image processed and saved successfully.")
}
