package simple

import (
	"image"
	"image/color"
	"math"
)

func Gamma(input image.Image, gamma float64) image.Image {
	gammaCorrection := 1 / gamma

	output := image.NewRGBA(input.Bounds())
	forEachPixel(input, func(x, y int, c color.Color) {
		r, g, b := rgb(c)
		r = truncate(255 * math.Pow(float64(r)/255, gammaCorrection))
		g = truncate(255 * math.Pow(float64(g)/255, gammaCorrection))
		b = truncate(255 * math.Pow(float64(b)/255, gammaCorrection))
		output.Set(x, y, newRGB(r, g, b))
	})
	return output
}
