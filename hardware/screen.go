package hardware

import (
	color_extractor "github.com/marekm4/color-extractor"
	"image"
	"image/color"
)

type Screen struct {
}

func (s *Screen) Sample(screenshot image.Image, samples []image.Rectangle) []color.Color {

	var colors []color.Color

	for _, rect := range samples {
		element := screenshot.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(rect)

		dominantColors := color_extractor.ExtractColors(element)
		dominantColor := dominantColors[0]
		r, g, b, _ := dominantColors[0].RGBA()
		if r < 50 && g < 50 && b < 50 {
			if len(dominantColors) > 1 {
				dominantColor = dominantColors[1]
			}
		}
		colors = append(colors, dominantColor)
	}

	return colors
}
