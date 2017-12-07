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

func rgb(c color.Color) (r, g, b uint8) {
	cr, cg, cb, _ := c.RGBA()
	return uint8(cr & 255), uint8(cg & 255), uint8(cb & 255)
}
