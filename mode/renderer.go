package mode

import (
	"github.com/kbinani/screenshot"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
)

type Renderer interface {
	Render(port *serial.Port, signal chan bool)
	Stop()
}

type Lightsaber struct {
	TerminateRenderChannel chan bool
	Lights                 *hardware.LightsArray
}

func (l *Lightsaber) StopRendering() {
	renderIsrunningMutex.Lock()
	stop := renderIsRunning
	renderIsRunning = false
	renderIsrunningMutex.Unlock()
	if stop {
		l.TerminateRenderChannel <- true
	}
}

func (l *Lightsaber) Render(configuration config.Configuration, serialPort *serial.Port) {
	renderIsrunningMutex.Lock()
	start := !renderIsRunning
	renderIsRunning = true
	renderIsrunningMutex.Unlock()

	if start {
		l.Lights = hardware.NewArray(configuration.LedGeometry)
		//shut down the lights
		serialPort.Write(l.Lights.Buffer())

		switch *configuration.SelectedMode {
		case "color_swirl":
			colorAdj := ColorAdjustment{configuration.ColorAdjustment}

			swirl := NewSwirl(
				configuration.Swirl,
				colorAdj,
				configuration.LedGeometry,
				l.Lights,
			)

			go swirl.Render(serialPort, l.TerminateRenderChannel)
		case "screen_grabber":
			samplesGeometry := hardware.NewSamplesGeometry(
				screenshot.GetDisplayBounds(*configuration.DisplayIndex),
				configuration.LedGeometry,
				configuration.ScreenGrabber,
			)

			colorAdj := ColorAdjustment{configuration.ColorAdjustment}

			screenGrabber := NewScreenGrabber(
				*configuration.DisplayIndex,
				colorAdj,
				samplesGeometry.Calculate(),
				l.Lights)

			go screenGrabber.Render(serialPort, l.TerminateRenderChannel)

		case "vader":
			colorAdj := ColorAdjustment{configuration.ColorAdjustment}
			vader := NewVader(
				configuration.Vader,
				colorAdj,
				l.Lights,
			)
			go vader.Render(serialPort, l.TerminateRenderChannel)

		case "backlight":
			colorAdj := ColorAdjustment{configuration.ColorAdjustment}
			backlight := NewBacklight(configuration.Backlight, colorAdj, l.Lights)
			go backlight.Render(serialPort, l.TerminateRenderChannel)

		case "custom_scene":
			colorAdj := ColorAdjustment{configuration.ColorAdjustment}
			customScene := NewCustom(configuration.Custom, colorAdj, l.Lights)
			go customScene.Render(serialPort, l.TerminateRenderChannel)

		case "ocean":
			colorAdj := ColorAdjustment{configuration.ColorAdjustment}
			ocean := NewOcean(colorAdj, l.Lights)
			go ocean.Render(serialPort, l.TerminateRenderChannel)
		}
	}
}
