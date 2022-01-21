package mode

import (
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"math/rand"
	"time"
)

type RGB config.RGB

type Ocean struct {
	lights *hardware.LightsArray
}

func (o *Ocean) Stop(port *serial.Port) {

	for i := 0; i < o.lights.NumberOfLights(); i++ {
		o.lights.SetLed(i, 0, 0, 0)

	}

	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)

	port.Write(o.lights.Buffer())
}

func (o *Ocean) Render(serialPort *serial.Port, signal chan bool) {
	terminate := false

	baseColor := RGB{4, 133, 162}
	color2 := RGB{4, 173, 211}
	color3 := RGB{4, 173, 226}
	color4 := RGB{4, 173, 235}
	color5 := RGB{4, 173, 245}
	color6 := RGB{0, 0, 245}

	for terminate == false {

		for i := 0; i < o.lights.NumberOfLights(); i++ {
			o.lights.SetLed(i, baseColor.R, baseColor.G, baseColor.B)
		}

		for i := 0; i < o.lights.NumberOfLights()/9; i++ {
			pos := rand.Intn(o.lights.NumberOfLights())
			o.lights.SetLed(pos, color2.R, color2.G, color2.B)
		}

		for i := 0; i < o.lights.NumberOfLights()/9; i++ {
			pos := rand.Intn(o.lights.NumberOfLights())
			o.lights.SetLed(pos, color3.R, color3.G, color3.B)
		}

		for i := 0; i < o.lights.NumberOfLights()/9; i++ {
			pos := rand.Intn(o.lights.NumberOfLights())
			o.lights.SetLed(pos, color4.R, color4.G, color4.B)
		}

		for i := 0; i < o.lights.NumberOfLights()/6; i++ {
			pos := rand.Intn(o.lights.NumberOfLights())
			o.lights.SetLed(pos, color5.R, color5.G, color5.B)
		}

		for i := 0; i < o.lights.NumberOfLights()/24; i++ {
			pos := rand.Intn(o.lights.NumberOfLights())
			o.lights.SetLed(pos, color6.R, color6.G, color6.B)
		}

		time.Sleep(90 * time.Millisecond)
		serialPort.Write(o.lights.Buffer())

		select {
		case terminate = <-signal:
			o.Stop(serialPort)
			return
		default:
		}
	}
}
func NewOcean(array *hardware.LightsArray) Ocean {
	return Ocean{
		array,
	}
}
