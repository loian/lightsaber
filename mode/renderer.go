package mode

import (
	"fmt"
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
		fmt.Println("rendering")

		l.Lights = hardware.NewArray(configuration.LedGeometry)
		switch configuration.SelectedMode {
		case "color_swirl":
			swirl := NewSwirl(
				configuration.Swirl,
				configuration.LedGeometry,
				l.Lights,
			)

			go swirl.Render(serialPort, l.TerminateRenderChannel)
		case "screen_grabber":
			samplesGeometry := hardware.NewSamplesGeometry(
				screenshot.GetDisplayBounds(configuration.DisplayIndex),
				configuration.LedGeometry,
				configuration.ScreenGrabber,
			)

			screenGrabber := NewScreenGrabber(
				configuration.DisplayIndex,
				configuration.ColorAdjustment,
				samplesGeometry.Calculate(),
				l.Lights)

			go screenGrabber.Render(serialPort, l.TerminateRenderChannel)

		case "vader":
			vader := NewVader(
				configuration.Vader,
				l.Lights,
			)
			go vader.Render(serialPort, l.TerminateRenderChannel)

		}
	}
}
