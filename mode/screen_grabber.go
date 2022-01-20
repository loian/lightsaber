package mode

import (
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"image"
	"lightsaber/config"
	"lightsaber/hardware"
	"lightsaber/util"
	"time"
)

type ScreenGrabber struct {
	displayIndex    int
	colorAdjustment config.ColorAdjustment
	sampleAreas     []image.Rectangle
	lights          *hardware.LightsArray
}

func (sg *ScreenGrabber) Stop(port *serial.Port) {
	for i := 0; i < sg.lights.NumberOfLights(); i++ {
		sg.lights.SetLed(i, 0, 0, 0)
	}
	//TODO: implement signaling rather than this crap solution
	time.Sleep(1000 * time.Millisecond)
}

func (sg *ScreenGrabber) Render(port *serial.Port, signal chan bool) {

	screen := hardware.Screen{}
	ticker := time.NewTicker(60 * time.Millisecond)

	for {
		select {
		case <-signal:
			ticker.Stop()
			sg.Stop(port)
		case <-ticker.C:
			img, _ := screenshot.CaptureDisplay(sg.displayIndex)
			colors := screen.DominantColors(img, sg.sampleAreas)
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
	}

}

func NewScreenGrabber(
	displayIndex int,
	colorAdjustment config.ColorAdjustment,
	sampleAreas []image.Rectangle,
	lights *hardware.LightsArray) *ScreenGrabber {
	return &ScreenGrabber{
		displayIndex:    displayIndex,
		colorAdjustment: colorAdjustment,
		sampleAreas:     sampleAreas,
		lights:          lights,
	}
}
