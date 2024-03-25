package keyscheck

import (
	"image"
)

func avg(img image.Image) (int, int, int) {
	// Define the region of interest
	startX := 1280
	endX := 1315
	startY := 805
	endY := 760

	// Variables to hold total color components
	totalR := 0
	totalG := 0
	totalB := 0
	totalPixels := 0

	// Iterate over the pixels in the specified region
	for y := startY; y >= endY; y-- {
		for x := startX; x <= endX; x++ {
			// Get the color of the pixel at the current position
			pixelColor := img.At(x, y)
			// Extract the RGB components
			r, g, b, _ := pixelColor.RGBA()
			// Accumulate the color components
			totalR += int(r >> 8)
			totalG += int(g >> 8)
			totalB += int(b >> 8)
			totalPixels++
		}
	}

	// Calculate the average color
	avgR := totalR / totalPixels
	avgG := totalG / totalPixels
	avgB := totalB / totalPixels

	return avgR, avgG, avgB
}
