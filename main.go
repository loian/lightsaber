package main

import (
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"

	"github.com/tarm/serial"
	"lightsaber/hardware"
)

func main() {

	logrus.Info("Starting Lightsaber daemon.")

	//displays := screenshot.NumActiveDisplays()
	display := 0
	logrus.Info("Connecting to the serial port")
	configuration := &serial.Config{Name: "/dev/tty.usbserial-1110", Baud: 115200}
	serialPort, err := serial.OpenPort(configuration)
	if err != nil {
		logrus.Fatal("Unable to conenct to the serial port: ", err)
		panic("serial connection failed")
	}

	ledGeometry := hardware.LedGeometry{Right: 22, Top: 48, Left: 22, Bottom: 48}
	margins := hardware.Margins{Right: 200, Top: 100, Left: 200, Bottom: 100}
	lights := hardware.NewArray(0, ledGeometry)

	samplesGeometry := hardware.NewSamplesGeometry(
		screenshot.GetDisplayBounds(display),
		ledGeometry,
		margins,
		150,
		100,
	)

	sampleAreas := samplesGeometry.Calculate()
	screen := hardware.Screen{}
	for {
		img, _ := screenshot.CaptureDisplay(display)
		colors := screen.Sample(img, sampleAreas)
		for pos, c := range colors {
			r, g, b, _ := c.RGBA()
			lights.SetLed(pos, byte(r), byte(g), byte(b))
		}
		_, err := serialPort.Write(lights.Buffer())
		if err != nil {
			logrus.Error("unable to send data to the serial port: ", err)
		}
	}
}
