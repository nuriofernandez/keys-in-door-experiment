package keyscheck

import "github.com/nuriofernandez/keys-in-door-experiment/camera"

func AreKeysThere() (bool, error) {
	screenshot, err := camera.ScreenShot()
	if err != nil {
		return false, err
	}

	// Calculate average color of the keys area
	aR, aB, aG := avg(*screenshot)

	// Get control point color to compare
	cR, cG, cB := control(*screenshot)

	// Calculate the difference between the avg and the control point colors
	var difference = distance(
		aR, aB, aG,
		cR, cG, cB,
	)

	// if the color from the control point differs more than 5 points,
	// then the keys are there.
	keysThere := difference > 5
	return keysThere, nil
}
