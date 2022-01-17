package main

import (
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/hardware"
	"lightsaber/mode"
	"time"
)

func main() {
	logrus.Info("Starting Lightsaber daemon.")

	//displays := screenshot.NumActiveDisplays()
	conf := config.Configuration{
		DisplayIndex: 0,
		Serial: config.Serial{
			Port: "/dev/tty.usbserial-110",
			Baud: 115200,
		},
		ColorAdjustment: config.ColorAdjustment{
			DarkenPercentage: 0.2,
		},
		LedGeometry: config.LedGeometry{
			Offset: 0,
			Right:  22,
			Top:    48,
			Left:   22,
			Bottom: 48,
		},
		ScreenGrabber: config.ScreenGrabber{
			Margin: config.Margins{
				Right:  150,
				Top:    100,
				Left:   150,
				Bottom: 100,
			},
			Size: config.Size{
				Width:  450,
				Height: 400,
			},
		},
	}

	logrus.Info("Connecting to the serial port")
	configuration := &serial.Config{Name: conf.Serial.Port, Baud: conf.Serial.Baud}
	serialPort, err := serial.OpenPort(configuration)
	if err != nil {
		logrus.Fatal("Unable to connect to the serial port: ", err)
		panic("serial connection failed")
	}

	lights := hardware.NewArray(conf.LedGeometry)

	samplesGeometry := hardware.NewSamplesGeometry(
		screenshot.GetDisplayBounds(conf.DisplayIndex),
		conf.LedGeometry,
		conf.ScreenGrabber,
	)

	screenGrabber := mode.NewScreenGrabber(
		conf.DisplayIndex,
		conf.ColorAdjustment,
		samplesGeometry,
		lights)

	ticker := time.NewTicker(60 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				screenGrabber.Render(serialPort)
			}
		}
	}()

	select {}
}
