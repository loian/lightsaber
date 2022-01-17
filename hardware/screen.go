package hardware

import (
	color_extractor "github.com/marekm4/color-extractor"
	"image"
	"image/color"
)

type Screen struct {
}

func (s *Screen) DominantColors(screenshot image.Image, samples []image.Rectangle) []color.Color {

	var colors []color.Color

	for _, rect := range samples {
		element := screenshot.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(rect)

		dominantColors := color_extractor.ExtractColors(element)
		dominantColor := dominantColors[0]
		r, g, b, _ := dominantColors[0].RGBA()
		if r == 0 && g == 0 && b == 0 {
			dominantColor = color.RGBA{5, 5, 5, 155}
		}

		colors = append(colors, dominantColor)
	}

	return colors
}
