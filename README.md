# Are keys in the door?

I keep forgetting my keys on the outside of the door, and realize next morning I slept
with the keys in the door. 😱

So, I decided to create this service to be able to send me notifications.

## Requirements

- A security camera that supports `RTSP` (I use a TPlink/TAPO security camera)
- A device that can serve as a server
- FFmpeg installed

## My setup

I have my security camera looking to the door:

![](https://i.imgur.com/izcFbCu.jpeg)

## Technical explanation

From the keys area, I defined an area where the keys should be visible, and
for color comparaison, there is a control point to compare the colors with the
keys area.

![](https://i.imgur.com/4inD0In.png)

I get all pixels in that area and calculate the average color,
ignoring those too close to the control point.

```go
for y := startY; y >= endY; y-- {
    for x := startX; x <= endX; x++ {
        // Get the color of the pixel at the current position
        pixelColor := img.At(x, y)
        // Extract the RGB components
        r, g, b, _ := pixelColor.RGBA()
		
		if distance(r,g,b, controlPoint) < 10 {
			// Ignore colors too close to the control point
			// so later the comparator will be more precise.
		    continue;	
        }
    }
}
```

Since the keys are darker than the door,
the difference between the control point should be noticeable.

```go
// Calculate the difference between the avg and the control point colors
difference := distance(
    avgR, avgB, avgG,
    controlPoint
)

// if the color from the control point differs more than 5 points,
// then the keys are there.
keysThere := difference > 5
return keysThere, nil
```

## What I do with the bool?

For now, I just established a simple http server that prints a json with it 🤣

```bash
GET http://localhost:8090/status
```
```json
{
  "keysInDoor": true
}
```