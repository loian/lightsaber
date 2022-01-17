package main

import (
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/teacat/noire"
	"image/color"
	"lightsaber/hardware"
)

func toRGB256(color color.Color) (byte, byte, byte) {
	r, g, b, _ := color.RGBA()
	return byte(r / 256), byte(g / 256), byte(b / 256)
}

func Darken(r byte, g byte, b byte, percent float64) (byte, byte, byte) {
	tmp := noire.NewRGB(float64(r), float64(g), float64(b))
	tmp = tmp.Darken(percent)
	r1, g1, b1 := tmp.RGB()
	return byte(r1), byte(g1), byte(b1)
}

func main() {
	logrus.Info("Starting Lightsaber daemon.")

	//displays := screenshot.NumActiveDisplays()
	display := 0
	darkenPercent := 0.2
	logrus.Info("Connecting to the serial port")
	configuration := &serial.Config{Name: "/dev/tty.usbserial-1110", Baud: 115200}
	serialPort, err := serial.OpenPort(configuration)
	if err != nil {
		logrus.Fatal("Unable to connect to the serial port: ", err)
		panic("serial connection failed")
	}

	ledGeometry := hardware.LedGeometry{Right: 22, Top: 48, Left: 22, Bottom: 48}
	margins := hardware.Margins{Right: 500, Top: 100, Left: 500, Bottom: 100}
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
			r, g, b := toRGB256(c)
			r, g, b = Darken(r, g, b, darkenPercent)
			lights.SetLed(pos, r, g, b)
		}
		_, err := serialPort.Write(lights.Buffer())
		if err != nil {
			logrus.Error("unable to send data to the serial port: ", err)
		}
	}
}
