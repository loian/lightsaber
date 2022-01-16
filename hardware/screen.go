package hardware

import "image"

type Screen struct {
	Elements []image.Image
}

func (s *Screen) Sample(screenshot image.Image, samples []image.Rectangle) {
	s.Elements = []image.Image{}

	for _, rect := range samples {
		element := screenshot.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(rect)

		s.Elements = append(s.Elements, element)
	}

}
