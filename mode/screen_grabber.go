package mode

import (
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

func (sg *ScreenGrabber) Render(port *serial.Port) {

	sampleAreas := sg.samplesGeometry.Calculate()
	screen := hardware.Screen{}
	img, _ := screenshot.CaptureDisplay(sg.displayIndex)
	colors := screen.DominantColors(img, sampleAreas)
	for pos, c := range colors {
		//TODO: move the color adjustment in a separate class or decorate the lights struct
		r, g, b := util.ToRGB256(c)
		r, g, b = util.Darken(r, g, b, sg.colorAdjustment.DarkenPercentage)
		sg.lights.SetLed(pos, r, g, b)
	}
	_, err := port.Write(sg.lights.Buffer())
	if err != nil {
		logrus.Error("unable to send data to the serial port: ", err)
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
