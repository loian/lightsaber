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
	ConfigSignal          chan config.Configuration
	TerminateRenderSignal chan bool
	config                config.Configuration
	Lights                *hardware.LightsArray
	SerialPort            *serial.Port
}

func (l *Lightsaber) Render() {
	select {
	case l.config = <-l.ConfigSignal:

		l.Lights = hardware.NewArray(l.config.LedGeometry)

		switch l.config.SelectedMode {
		case "color_swirl":
			swirl := NewSwirl(
				l.config.Swirl,
				l.config.LedGeometry,
				l.Lights,
			)

			go swirl.Render(l.SerialPort, l.TerminateRenderSignal)
		case "screen_grabber":
			samplesGeometry := hardware.NewSamplesGeometry(
				screenshot.GetDisplayBounds(l.config.DisplayIndex),
				l.config.LedGeometry,
				l.config.ScreenGrabber,
			)

			screenGrabber := NewScreenGrabber(
				l.config.DisplayIndex,
				l.config.ColorAdjustment,
				samplesGeometry.Calculate(),
				l.Lights)

			go screenGrabber.Render(l.SerialPort, l.TerminateRenderSignal)
		}
	}
}
