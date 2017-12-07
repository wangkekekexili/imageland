package simple

import (
	"image"
	"image/color"
	"math"
)

type Distance func(left, right color.Color) float64

func EuclideanDistance(left, right color.Color) float64 {
	lr, lg, lb := rgb(left)
	rr, rg, rb := rgb(right)
	return float64((lr-rr)*(lr-rr) + (lg-rg)*(lg-rg) + (lb-rb)*(lb-rb))
}

func ReduceToPalette(input image.Image, palette []color.Color, distanceFunc Distance) image.Image {
	output := image.NewRGBA(input.Bounds())
	forEachPixel(input, func(x, y int, c color.Color) {
		output.Set(x, y, findNearestColor(c, palette, distanceFunc))
	})
	return output
}

func findNearestColor(c color.Color, palette []color.Color, distanceFunc Distance) color.Color {
	var nearest color.Color
	minDistance := math.MaxFloat64
	for i := 0; i != len(palette); i++ {
		dist := distanceFunc(c, palette[i])
		if dist < minDistance {
			minDistance = dist
			nearest = palette[i]
		}
	}
	return nearest
}
