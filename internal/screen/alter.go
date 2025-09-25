package screen

import (
	"image"

	"github.com/disintegration/imaging"
)

func MakeReadable(img image.Image) (*image.NRGBA, error) {
	// Convert the image to grayscale
	grayImg := imaging.Grayscale(img)

	// Apply a threshold to create a binary image
	binaryImg := imaging.AdjustContrast(grayImg, 50)

	// Optionally, apply a Gaussian blur to reduce noise
	blurredImg := imaging.Blur(binaryImg, 1.0)

	return blurredImg, nil
}
