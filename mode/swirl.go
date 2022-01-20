package mode

//ported from https://github.com/adafruit/Adalight

import (
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"math"
	"time"
)

type Swirl struct {
	ledGeometry config.LedGeometry
	lights      *hardware.LightsArray
}

func (s *Swirl) Render(serialPort *serial.Port) {
	sine1 := 0.
	hue1 := 0
	var r, g, b byte
	for {
		sine2 := sine1
		hue2 := hue1
		total := s.ledGeometry.Right + s.ledGeometry.Top + s.ledGeometry.Left + s.ledGeometry.Bottom

		for i := 0; i < total; i++ {
			lo := byte(hue2 & 255)
			switch (hue2 >> 8) % 6 {
			case 0:
				r = 255
				g = lo
				b = 0
			case 1:
				r = 255 - lo
				g = 255
				b = 0
			case 2:
				r = 0
				g = 255
				b = lo
			case 3:
				r = 0
				g = 255 - lo
				b = 255
			case 4:
				r = lo
				g = 0
				b = 255
			case 5:
				r = 255
				g = 0
				b = 255 - lo
			}

			brightness := math.Pow(0.5+math.Sin(sine2)*0.5, 3.0) * 255.0
			s.lights.SetLed(
				i,
				byte(float64(r)*brightness/255),
				byte(float64(g)*brightness/255),
				byte(float64(b)*brightness/255),
			)
		}
		hue1 = (hue1 + 5) % (1536)
		sine1 -= .03
		time.Sleep(60 * time.Millisecond)
		serialPort.Write(s.lights.Buffer())
	}
}

func NewSwirl(geometry config.LedGeometry, array *hardware.LightsArray) *Swirl {
	return &Swirl{
		geometry,
		array,
	}
}
