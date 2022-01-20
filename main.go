package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lightsaber/config"
	"lightsaber/handler"
)

var conf = config.Configuration{
	DisplayIndex: 0,
	SelectedMode: "color_swirl",
	Serial: config.Serial{
		Port: "/dev/tty.usbserial-110",
		Baud: 115200,
	},
	ColorAdjustment: config.ColorAdjustment{
		DarkenPercentage: 0.0,
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
	Swirl: config.Swirl{
		PulseDepth:         0.2,
		PulseSpeed:         0.09,
		ColorRotationSpeed: 5,
	},
}

func main() {
	logrus.Info("Starting Lightsaber daemon.")
	r := gin.Default()
	r.POST("/start", handler.Start)
	r.POST("/off", handler.Stop)

	go r.Run("127.0.0.1:8877")

	select {}
}
