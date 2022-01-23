package mode

import (
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"time"
)

type Vader struct {
	config          config.Vader
	colorAdjustment ColorAdjustment
	lights          *hardware.LightsArray
}

func (v *Vader) Stop(port *serial.Port) {

	for i := 0; i < v.lights.NumberOfLights(); i++ {
		v.lights.SetLed(i,
			v.colorAdjustment.Adjust(hardware.Led{0, 0, 0}))
	}

	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)

	port.Write(v.lights.Buffer())
}
func (v *Vader) Render(serialPort *serial.Port, signal chan bool) {
	n := v.lights.NumberOfLights()
	i := 0
	for k := 0; k < n; k++ {
		v.lights.SetLed(k, v.colorAdjustment.Adjust(hardware.Led{0, 0, 0}))
	}

	var terminate = false
	for terminate == false {
		for k := 1; k < 10; k++ {

			v.lights.SetLed((i+k)%n, v.colorAdjustment.Adjust(hardware.Led{255, 0, 0}))
			v.lights.SetLed((i+k+1)%n, v.colorAdjustment.Adjust(hardware.Led{155, 0, 0}))
			v.lights.SetLed((i+k+2)%n, v.colorAdjustment.Adjust(hardware.Led{0, 0, 0}))
			v.lights.SetLed((i+k+n/2)%n, v.colorAdjustment.Adjust(hardware.Led{255, 0, 0}))
			v.lights.SetLed((i+k+n/+1)%n, v.colorAdjustment.Adjust(hardware.Led{155, 0, 0}))
			v.lights.SetLed((i+k+n/2+2)%n, v.colorAdjustment.Adjust(hardware.Led{0, 0, 0}))
			i++
		}
		i = i + n - 10

		time.Sleep(time.Duration(1-*v.config.Speed) * time.Millisecond)
		serialPort.Write(v.lights.Buffer())
		select {
		case terminate = <-signal:
			v.Stop(serialPort)
			return
		default:

		}
	}
}

func NewVader(vader config.Vader, colorAdjustment ColorAdjustment, array *hardware.LightsArray) Vader {
	return Vader{
		vader,
		colorAdjustment,
		array,
	}
}
