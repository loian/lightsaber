package main

import (
	"github.com/tarm/serial"
	"lightsaber/geometry"
	"log"
)

func main() {
	configuration := &serial.Config{Name: "/dev/tty.usbserial-1110", Baud: 115200}
	serialPort, err := serial.OpenPort(configuration)
	if err != nil {
		log.Fatal(err)
	}

	lights := geometry.New(0, 22, 48, 22, 48)

	for {
		serialPort.Write(lights.GetBuffer())
	}
}
