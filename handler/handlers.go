package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/mode"
)

func Stop(c *gin.Context) {
	lightSaber := mode.Lightsaber{
		TerminateRenderChannel: mode.TerminateRenderChannel,
	}
	lightSaber.StopRendering()
	c.JSON(200, gin.H{"status": "stopped"})
}

func Start(c *gin.Context) {
	var conf = config.Configuration{
		DisplayIndex: 0,
		SelectedMode: "custom",
		Serial: config.Serial{
			Port: "/dev/tty.usbserial-14310",
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
		Vader: config.Vader{
			Speed: 40,
		},
		Backlight: config.Backlight{
			R: 255,
			G: 255,
			B: 0,
		},

		Custom: config.Custom{
			[]config.RGB{
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{0, 0, 255},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
				{255, 0, 0},
			},
		},
	}

	configuration := &serial.Config{Name: conf.Serial.Port, Baud: conf.Serial.Baud}
	serialPort, err := serial.OpenPort(configuration)
	logrus.Info("Connecting to the serial port ", conf.Serial.Port)

	if err != nil {
		logrus.Fatal("Unable to connect to the serial port: ", err)
		panic("serial connection failed")
	}

	lightSaber := mode.Lightsaber{
		TerminateRenderChannel: mode.TerminateRenderChannel,
	}
	go lightSaber.Render(conf, serialPort)

	c.JSON(200, gin.H{"status": "started"})
}
