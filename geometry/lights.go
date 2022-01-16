package geometry

import "fmt"

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

type Lights struct {
	Positions []int
	buffer    []byte
}

func (l *Lights) SetLed(position int, r byte, g byte, b byte) {

	fmt.Println(l.Positions)
	l.buffer[l.Positions[position]*3+HEADER_COUNT+R_OFFSET] = r
	l.buffer[l.Positions[position]*3+HEADER_COUNT+G_OFFSET] = g
	l.buffer[l.Positions[position]*3+HEADER_COUNT+B_OFFSET] = b
}

func (l *Lights) GetBuffer() []byte {
	return l.buffer
}

func New(offset int, right int, top int, left int, bottom int) *Lights {

	total := right + top + left + bottom
	positions := make([]int, total)

	for i := 0; i < int(total); i++ {
		currentLed := i - offset
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

	lights := &Lights{
		positions,
		buffer,
	}
	for i := 0; i < total; i++ {
		lights.SetLed(i, 255, 255, 255)
	}

	return lights
}
