package simple

import (
	"image"
	"image/color"
)

func BrightnessRGB(input image.Image, delta int) image.Image {
	output := image.NewRGBA(input.Bounds())
	forEachPixel(input, func(x, y int, c color.Color) {
		r, g, b := rgb(c)
		r = truncate(float64(r + delta))
		g = truncate(float64(g + delta))
		b = truncate(float64(b + delta))
		output.Set(x, y, newRGB(r, g, b))
	})
	return output
}
