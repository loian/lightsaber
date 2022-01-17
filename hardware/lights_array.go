package hardware

import "lightsaber/config"

const (
	HEADER_COUNT = 6
	R_OFFSET     = 0
	G_OFFSET     = 1
	B_OFFSET     = 2
)

type Led struct {
	R byte
	G byte
	B byte
}

type LightsArray struct {
	positions []int
	buffer    []byte
}

func (la *LightsArray) SetLed(position int, r byte, g byte, b byte) {
	la.buffer[la.positions[position]*3+HEADER_COUNT+R_OFFSET] = r
	la.buffer[la.positions[position]*3+HEADER_COUNT+G_OFFSET] = g
	la.buffer[la.positions[position]*3+HEADER_COUNT+B_OFFSET] = b
}

func (la *LightsArray) NumberOfLights() int {
	return len(la.positions)
}

func (la *LightsArray) Buffer() []byte {
	return la.buffer
}

func NewArray(geometry config.LedGeometry) *LightsArray {

	total := geometry.Right + geometry.Top + geometry.Left + geometry.Bottom
	positions := make([]int, total)

	for i := 0; i < int(total); i++ {
		currentLed := i - geometry.Offset
		if currentLed > total {
			currentLed = currentLed - total
		}
		if currentLed < 0 {
			currentLed = total - currentLed
		}
		positions[i] = currentLed
	}

	buffer := make([]byte, total*3+HEADER_COUNT)
	buffer[0] = 'A'
	buffer[1] = 'd'
	buffer[2] = 'a'
	buffer[3] = byte((total - 1) >> 8)       // LED count high byte
	buffer[4] = byte((total - 1) & 0xff)     // LED count low byte
	buffer[5] = buffer[3] ^ buffer[4] ^ 0x55 // Checksum

	lights := &LightsArray{
		positions,
		buffer,
	}
	for i := 0; i < total; i++ {
		lights.SetLed(i, 255, 255, 255)
	}

	return lights
}
