package main

func AreKeysThere() (bool, error) {
	screenshot, err := screenShot()
	if err != nil {
		return false, err
	}

	// Calculate average color of the keys area
	r8, g8, b8 := avg(*screenshot)

	// above 125 there is not much black, so no keys
	var keysThere = r8 < 125 && g8 < 125 && b8 < 125
	return keysThere, nil
}
