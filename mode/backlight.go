package mode

import (
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"time"
)

type Backlight struct {
	config          config.Backlight
	colorAdjustment ColorAdjustment
	lights          *hardware.LightsArray
}

func (v *Backlight) Stop(port *serial.Port) {

	for i := 0; i < v.lights.NumberOfLights(); i++ {
		v.lights.SetLed(i, hardware.Led{0, 0, 0})
	}

	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)

	port.Write(v.lights.Buffer())
}

func (b *Backlight) Render(serialPort *serial.Port, signal chan bool) {
	n := b.lights.NumberOfLights()
	for k := 0; k < n; k++ {
		b.lights.SetLed(k,
			b.colorAdjustment.Adjust(hardware.Led{R: b.config.R, G: b.config.G, B: b.config.B}),
		)
	}
	var terminate = false
	for terminate == false {
		time.Sleep(100 * time.Millisecond)
		serialPort.Write(b.lights.Buffer())

		select {
		case terminate = <-signal:
			b.Stop(serialPort)
			return
		default:
		}
	}
}
func NewBacklight(backlight config.Backlight, colorAdjustment ColorAdjustment, array *hardware.LightsArray) Backlight {
	return Backlight{
		backlight,
		colorAdjustment,
		array,
	}
}
