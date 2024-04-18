package keyscheck

import (
	"fmt"
	"image"
	"math"
)

func avg(img image.Image) (int, int, int, float64) {
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

	// Get control point color to compare
	cR, cG, cB := control(img)

	// Iterate over the pixels in the specified region
	for y := startY; y >= endY; y-- {
		for x := startX; x <= endX; x++ {
			// Get the color of the pixel at the current position
			pixelColor := img.At(x, y)
			// Extract the RGB components
			r, g, b, _ := pixelColor.RGBA()

			// Accumulate the color components
			rR := int(r >> 8)
			rG := int(g >> 8)
			rB := int(b >> 8)

			// Calculate the difference between the avg and the control point colors
			var difference = distance(
				rR, rB, rG,
				cR, cG, cB,
			)

			if difference < 15 {
				// ignore those colors close to the control point
				continue
			}

			// Accumulate the color components
			totalR += rR
			totalG += rG
			totalB += rB
			totalPixels++
		}
	}

	fmt.Printf("Total number of pixels: %d\n", totalPixels)
	if totalPixels <= 0 {
		totalPixels = 1
	}

	// Precision calculation
	diffX := math.Abs(float64(startX-endX)) + 1
	diffY := math.Abs(float64(startY-endY)) + 1
	area := diffY * diffX
	precision := float64(totalPixels) / area
	fmt.Printf("Precision: %f\n", precision)

	// Calculate the average color
	avgR := totalR / totalPixels
	avgG := totalG / totalPixels
	avgB := totalB / totalPixels

	return avgR, avgG, avgB, precision
}
