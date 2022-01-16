package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"github.com/tarm/serial"
	"image"
	"lightsaber/hardware"
	"log"
	"time"
)

func main() {

	//displays := screenshot.NumActiveDisplays()
	display := 0

	configuration := &serial.Config{Name: "/dev/tty.usbserial-1110", Baud: 115200}
	serialPort, err := serial.OpenPort(configuration)
	if err != nil {
		log.Fatal(err)
	}
	ledGeometry := hardware.LedGeometry{22, 48, 22, 48}
	lights := hardware.NewArray(0, ledGeometry)

	samplesGeometry := hardware.SamplesGeometry{
		screenshot.GetDisplayBounds(display),
		ledGeometry,
		200,
		250,
	}

	img, err := screenshot.CaptureDisplay(display)
	screen := hardware.Screen{[]image.Image{}}
	screen.Sample(img, samplesGeometry.Calculate())
	fmt.Println(screen.Elements)
	for {
		img, err := screenshot.CaptureDisplay(display)
		if err != nil {
			panic(err)
		}
		colorAt00 := img.At(1720, 720)
		for pos := 0; pos < lights.NumberOfLights(); pos++ {
			r, g, b, _ := colorAt00.RGBA()
			lights.SetLed(pos, byte(r), byte(g), byte(b))
		}
		serialPort.Write(lights.Buffer())
		time.Sleep(200 * time.Millisecond)
	}
}
