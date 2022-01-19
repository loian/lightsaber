package hardware

import (
	"image"
	"image/color"
)

type Screen struct {
}

func (s *Screen) DominantColors(screenshot image.Image, samples []image.Rectangle) []color.Color {

	var colors []color.Color

	for _, rect := range samples {

		dominantColors := ExtractColorsWithConfig(screenshot, rect, Config{
			StepX:       25,
			StepY:       25,
			SmallBucket: .1,
		})

		if len(dominantColors) == 0 {
			dominantColor := color.RGBA{5, 5, 5, 255}
			colors = append(colors, dominantColor)
		} else {
			dominantColor := dominantColors[0]
			r, g, b, _ := dominantColors[0].RGBA()
			if r == 0 && g == 0 && b == 0 {
				dominantColor = color.RGBA{5, 5, 5, 255}
			}
			colors = append(colors, dominantColor)

		}
	}

	return colors
}
