package hardware

import (
	"fmt"
	"image"
)

type SamplesGeometry struct {
	ViewPort    image.Rectangle
	LedGeometry LedGeometry
	Width       int
	Heigh       int
}

func (s SamplesGeometry) Calculate() []image.Rectangle {
	yResolution := s.ViewPort.Max.Y
	xResolution := s.ViewPort.Max.X

	samples := []image.Rectangle{}

	rightHeight := int(yResolution / s.LedGeometry.Right)
	for y := yResolution; y > 0; y = y - rightHeight {
		rectangle := image.Rectangle{
			image.Point{
				xResolution - s.Width,
				y - rightHeight,
			},
			image.Point{
				xResolution,
				y,
			},
		}

		samples = append(samples, rectangle)
		fmt.Println(rectangle)
	}
	return samples
}
