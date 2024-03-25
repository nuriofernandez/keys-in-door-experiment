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

From the keys area, I defined an area where the keys should be visible:

![](https://i.imgur.com/Nz5YnPk.png)

I get all pixels in that area and calculate the average color.

```go
for y := startY; y >= endY; y-- {
    for x := startX; x <= endX; x++ {
        // Get the color of the pixel at the current position
        pixelColor := img.At(x, y)
        // Extract the RGB components
        r, g, b, _ := pixelColor.RGBA()
    }
}
```

Since the keys are darker than the door,
if the average color is below 125, that means the keys are there.

```go
// above 125 there is not much black, so no keys
var keysThere = r < 125 && g < 125 && b < 125
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