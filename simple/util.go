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

func rgb(c color.Color) (r, g, b int) {
	cr, cg, cb, _ := c.RGBA()
	return int(cr & 255), int(cg & 255), int(cb & 255)
}

func newRGB(r, g, b int) color.RGBA {
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}

func truncate(v float64) int {
	if v < 0 {
		v = 0
	}
	if v > 255 {
		v = 255
	}
	return int(v)
}
