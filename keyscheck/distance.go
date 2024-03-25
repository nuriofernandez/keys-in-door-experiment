package keyscheck

import "math"

func distance(aR, aG, aB int, bR, bG, bB int) int {
	dR := math.Abs(float64(aR) - float64(bR))
	dG := math.Abs(float64(aG) - float64(bG))
	dB := math.Abs(float64(aB) - float64(bB))

	return int((dR + dG + dB) / 3)
}
