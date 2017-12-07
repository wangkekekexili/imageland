package simple

import (
	"image"
	"image/color"
)

const (
	minContrast = -255
	maxContrast = 255
)

func Contrast(input image.Image, contrast int) image.Image {
	if contrast < minContrast {
		contrast = minContrast
	}
	if contrast > maxContrast {
		contrast = maxContrast
	}
	factor := 259 * float64(contrast+255) / (255 * float64(259-contrast))
	output := image.NewRGBA(input.Bounds())
	forEachPixel(input, func(x, y int, c color.Color) {
		r, g, b := rgb(c)
		nr := truncate(factor*float64(r-128) + 128)
		ng := truncate(factor*float64(g-128) + 128)
		nb := truncate(factor*float64(b-128) + 128)
		output.Set(x, y, newRGB(nr, ng, nb))
	})
	return output
}
