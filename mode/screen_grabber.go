package mode

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"lightsaber/util"
)

type ScreenGrabber struct {
	displayIndex    int
	colorAdjustment config.ColorAdjustment
	samplesGeometry hardware.SamplesGeometry
	lights          *hardware.LightsArray
}

func (sg *ScreenGrabber) Render(port *serial.Port, signal chan bool) {

	sampleAreas := sg.samplesGeometry.Calculate()
	screen := hardware.Screen{}
	terminate := false
	for terminate != true {

		select {
		case <-signal:
			terminate = true
			fmt.Println("sig")
		default:
		}

		img, _ := screenshot.CaptureDisplay(sg.displayIndex)
		colors := screen.DominantColors(img, sampleAreas)
		for pos, c := range colors {
			r, g, b := util.ToRGB256(c)
			r, g, b = util.Darken(r, g, b, sg.colorAdjustment.DarkenPercentage)
			sg.lights.SetLed(pos, r, g, b)
		}
		_, err := port.Write(sg.lights.Buffer())
		if err != nil {
			logrus.Error("unable to send data to the serial port: ", err)
		}
	}
}

func NewScreenGrabber(
	displayIndex int,
	colorAdjustment config.ColorAdjustment,
	samplesGeometry hardware.SamplesGeometry,
	lights *hardware.LightsArray) *ScreenGrabber {
	return &ScreenGrabber{
		displayIndex:    displayIndex,
		colorAdjustment: colorAdjustment,
		samplesGeometry: samplesGeometry,
		lights:          lights,
	}
}
