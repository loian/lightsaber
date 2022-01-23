package mode

import (
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"time"
)

type CustomScene struct {
	config          config.Custom
	colorAdjustment ColorAdjustment
	lights          *hardware.LightsArray
}

func (c *CustomScene) Stop(port *serial.Port) {

	for i := 0; i < c.lights.NumberOfLights(); i++ {
		c.lights.SetLed(i, hardware.Led{0, 0, 0})
	}

	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)

	port.Write(c.lights.Buffer())
}

func (c *CustomScene) Render(serialPort *serial.Port, signal chan bool) {
	for i, rgb := range c.config.Leds {
		c.lights.SetLed(i,
			c.colorAdjustment.Adjust(hardware.Led{R: rgb.R, G: rgb.G, B: rgb.B}),
		)
	}

	var terminate = false
	for terminate == false {
		time.Sleep(100 * time.Millisecond)
		serialPort.Write(c.lights.Buffer())

		select {
		case terminate = <-signal:
			c.Stop(serialPort)
			return
		default:
		}
	}
}

func NewCustom(conf config.Custom, colorAdjustment ColorAdjustment, array *hardware.LightsArray) CustomScene {
	return CustomScene{
		conf,
		colorAdjustment,
		array,
	}
}
