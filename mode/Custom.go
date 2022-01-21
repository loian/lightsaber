package mode

import (
	"fmt"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"time"
)

type Custom struct {
	config config.Custom
	lights *hardware.LightsArray
}

func (c *Custom) Stop(port *serial.Port) {

	for i := 0; i < c.lights.NumberOfLights(); i++ {
		c.lights.SetLed(i, 0, 0, 0)
	}

	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)

	port.Write(c.lights.Buffer())
}

func (c *Custom) Render(serialPort *serial.Port, signal chan bool) {
	for i, rgb := range c.config.Leds {
		c.lights.SetLed(i, rgb.R, rgb.G, rgb.B)
		fmt.Println(rgb)
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

func NewCustom(conf config.Custom, array *hardware.LightsArray) Custom {
	return Custom{
		conf,
		array,
	}
}
