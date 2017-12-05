package simple

import (
	"image"
	"image/color"
)

func GrayScale(input image.Image) image.Image {
	output := image.NewGray(input.Bounds())
	forEachPixel(input, func(x, y int, c color.Color) {
		output.Set(x, y, c)
	})
	return output
}
