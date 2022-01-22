package hardware

import (
	"image"
	"lightsaber/config"
)

type SamplesGeometry struct {
	viewPort    image.Rectangle
	ledGeometry config.LedGeometry
	margins     config.Margins
	size        config.Size
}

func (s SamplesGeometry) Calculate() []image.Rectangle {
	yResolution := s.viewPort.Max.Y
	xResolution := s.viewPort.Max.X

	samples := []image.Rectangle{}

	rightHeight := int(yResolution / *s.ledGeometry.Right)
	for y := yResolution; y >= rightHeight; y = y - rightHeight {
		rectangle := image.Rectangle{
			image.Point{
				xResolution - *s.size.Width - *s.margins.Right,
				y - rightHeight,
			},
			image.Point{
				xResolution - *s.margins.Right,
				y,
			},
		}

		samples = append(samples, rectangle)

	}

	topWidth := int(xResolution / *s.ledGeometry.Top)
	for x := xResolution; x >= topWidth; x = x - topWidth {
		rectangle := image.Rectangle{
			image.Point{
				x - topWidth,
				0 + *s.margins.Top,
			},
			image.Point{
				x,
				*s.size.Height + *s.margins.Top,
			},
		}
		samples = append(samples, rectangle)
	}

	leftHeight := int(yResolution / *s.ledGeometry.Left)
	for y := 0; y <= yResolution-leftHeight; y = y + leftHeight {
		rectangle := image.Rectangle{
			image.Point{
				0 + *s.margins.Left,
				y,
			},
			image.Point{
				*s.size.Width + *s.margins.Left,
				y + leftHeight,
			},
		}
		samples = append(samples, rectangle)

	}

	bottomWidth := int(xResolution / *s.ledGeometry.Bottom)
	count := 0
	for x := 0; x <= xResolution-bottomWidth; x = x + bottomWidth {
		rectangle := image.Rectangle{
			image.Point{
				x,
				yResolution - *s.size.Height - *s.margins.Bottom,
			},
			image.Point{
				x + bottomWidth,
				yResolution - *s.margins.Bottom,
			},
		}
		count = count + 1
		samples = append(samples, rectangle)
	}

	return samples
}

func NewSamplesGeometry(viewPort image.Rectangle, ledGeometry config.LedGeometry, grabber config.ScreenGrabber) SamplesGeometry {
	return SamplesGeometry{
		viewPort,
		ledGeometry,
		grabber.Margin,
		grabber.Size,
	}
}
