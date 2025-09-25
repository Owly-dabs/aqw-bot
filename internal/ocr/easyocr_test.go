package ocr

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunEasyOCR_MockedGetwd(t *testing.T) {
	// Save original and defer restore
	original := getWorkingDir
	defer func() { getWorkingDir = original }()

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	getWorkingDir = func() (string, error) {
		return dir + "/../..", nil
	}

	// Path to the altered screenshot
	imagePath := filepath.Join("screenshot.png")

	texts, err := RunEasyOCR(imagePath)
	if err != nil {
		t.Fatalf("RunEasyOCR failed: %v", err)
	}

	if len(texts) == 0 {
		t.Fatal("Expected at least one line of OCR output, got none")
	}

	t.Logf("OCR Output: %v", texts)
}
