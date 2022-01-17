package mode

import "github.com/tarm/serial"

type Renderer interface {
	Render(port *serial.Port, signal chan bool)
}
