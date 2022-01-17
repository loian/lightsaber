package hardware

import (
	"image"
)

type Margins struct {
	Right  int
	Top    int
	Left   int
	Bottom int
}
type SamplesGeometry struct {
	viewPort    image.Rectangle
	ledGeometry LedGeometry
	margins     Margins
	width       int
	heigh       int
}

func (s SamplesGeometry) Calculate() []image.Rectangle {
	yResolution := s.viewPort.Max.Y
	xResolution := s.viewPort.Max.X

	samples := []image.Rectangle{}

	rightHeight := int(yResolution / s.ledGeometry.Right)
	for y := yResolution; y >= rightHeight; y = y - rightHeight {
		rectangle := image.Rectangle{
			image.Point{
				xResolution - s.width - s.margins.Right,
				y - rightHeight,
			},
			image.Point{
				xResolution - s.margins.Right,
				y,
			},
		}

		samples = append(samples, rectangle)

	}

	topWidth := int(xResolution / s.ledGeometry.Top)
	for x := xResolution; x >= topWidth; x = x - topWidth {
		rectangle := image.Rectangle{
			image.Point{
				x - topWidth,
				0 + s.margins.Top,
			},
			image.Point{
				x,
				s.heigh + s.margins.Top,
			},
		}
		samples = append(samples, rectangle)
	}

	leftHeight := int(yResolution / s.ledGeometry.Left)
	for y := 0; y <= yResolution-leftHeight; y = y + leftHeight {
		rectangle := image.Rectangle{
			image.Point{
				0 + s.margins.Left,
				y,
			},
			image.Point{
				s.width + s.margins.Left,
				y + leftHeight,
			},
		}
		samples = append(samples, rectangle)

	}

	bottomWidth := int(xResolution / s.ledGeometry.Bottom)
	count := 0
	for x := 0; x <= xResolution-bottomWidth; x = x + bottomWidth {
		rectangle := image.Rectangle{
			image.Point{
				x,
				yResolution - s.heigh - s.margins.Bottom,
			},
			image.Point{
				x + bottomWidth,
				yResolution - s.margins.Bottom,
			},
		}
		count = count + 1
		samples = append(samples, rectangle)
	}

	return samples
}

func NewSamplesGeometry(viewPort image.Rectangle, ledGeometry LedGeometry, margins Margins, width int, heigh int) SamplesGeometry {
	return SamplesGeometry{
		viewPort,
		ledGeometry,
		margins,
		width,
		heigh,
	}
}
