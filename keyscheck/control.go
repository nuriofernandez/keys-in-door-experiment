package keyscheck

import "image"

func control(img image.Image) (int, int, int) {
	pixelColor := img.At(1270, 780)
	// Extract the RGB components
	r, g, b, _ := pixelColor.RGBA()

	// cast to int
	return int(r >> 8), int(g >> 8), int(b >> 8)
}
