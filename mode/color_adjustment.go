package mode

import (
	"github.com/teacat/noire"
	"lightsaber/config"
	"lightsaber/hardware"
)

type ColorAdjustment struct {
	config config.ColorAdjustment
}

func (ca ColorAdjustment) Adjust(led hardware.Led) hardware.Led {
	return ca.Darken(
		ca.Brighten(
			ca.Desaturate(
				ca.Saturate(
					ca.Hue(led),
				),
			),
		),
	)

}
func (ca ColorAdjustment) Darken(led hardware.Led) hardware.Led {
	tmp := noire.NewRGB(float64(led.R), float64(led.G), float64(led.B))
	tmp = tmp.Darken(*ca.config.Darken)
	r1, g1, b1 := tmp.RGB()
	return hardware.Led{byte(r1), byte(g1), byte(b1)}
}

func (ca ColorAdjustment) Brighten(led hardware.Led) hardware.Led {
	tmp := noire.NewRGB(float64(led.R), float64(led.G), float64(led.B))
	tmp = tmp.Brighten(*ca.config.Brighten)
	r1, g1, b1 := tmp.RGB()
	return hardware.Led{byte(r1), byte(g1), byte(b1)}
}

func (ca ColorAdjustment) Saturate(led hardware.Led) hardware.Led {
	tmp := noire.NewRGB(float64(led.R), float64(led.G), float64(led.B))
	tmp = tmp.Saturate(*ca.config.Saturate)
	r1, g1, b1 := tmp.RGB()
	return hardware.Led{byte(r1), byte(g1), byte(b1)}
}

func (ca ColorAdjustment) Desaturate(led hardware.Led) hardware.Led {
	tmp := noire.NewRGB(float64(led.R), float64(led.G), float64(led.B))
	tmp = tmp.Desaturate(*ca.config.Desaturate)
	r1, g1, b1 := tmp.RGB()
	return hardware.Led{byte(r1), byte(g1), byte(b1)}
}

func (ca ColorAdjustment) Hue(led hardware.Led) hardware.Led {
	tmp := noire.NewRGB(float64(led.R), float64(led.G), float64(led.B))
	tmp = tmp.AdjustHue(*ca.config.Hue)
	r1, g1, b1 := tmp.RGB()
	return hardware.Led{byte(r1), byte(g1), byte(b1)}
}
