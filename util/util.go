package util

import (
	"github.com/teacat/noire"
	"image/color"
)

func ToRGB256(color color.Color) (byte, byte, byte) {
	r, g, b, _ := color.RGBA()
	return byte(r / 256), byte(g / 256), byte(b / 256)
}

func Darken(r byte, g byte, b byte, percent float64) (byte, byte, byte) {
	tmp := noire.NewRGB(float64(r), float64(g), float64(b))
	tmp = tmp.Darken(percent)
	r1, g1, b1 := tmp.RGB()
	return byte(r1), byte(g1), byte(b1)
}
