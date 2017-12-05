package simple

import (
	"image"
	"image/color"
)

func forEachPixel(img image.Image, fn func(x, y int, c color.Color)) {
	bounds := img.Bounds()
	for x := bounds.Min.X; x != bounds.Max.X; x++ {
		for y := bounds.Min.Y; y != bounds.Max.Y; y++ {
			fn(x, y, img.At(x, y))
		}
	}
}
