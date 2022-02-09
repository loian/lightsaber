package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"lightsaber/config"
	"lightsaber/mode"
	"os"
	"time"
)

func Quit(c *gin.Context) {
	c.JSON(200, gin.H{"status": "terminated"})
	go func() {
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()
	return
}

func Stop(c *gin.Context) {
	lightSaber := mode.Lightsaber{
		TerminateRenderChannel: mode.TerminateRenderChannel,
	}
	lightSaber.StopRendering()
	c.JSON(200, gin.H{"status": "stopped"})
}

func Status(c *gin.Context) {
	lightSaber := mode.Lightsaber{
		TerminateRenderChannel: mode.TerminateRenderChannel,
	}

	if (lightSaber.IsRunning()) {
		c.JSON(200, gin.H{"status": "started"})
	} else {
		c.JSON(200, gin.H{"status": "stopped"})
	}
}

func Start(c *gin.Context) {

	var conf config.Configuration

	if err := c.ShouldBindJSON(&conf); err != nil {
		logrus.Warnln(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	configuration := &serial.Config{Name: conf.Serial.Port, Baud: conf.Serial.Baud}
	serialPort, err := serial.OpenPort(configuration)
	logrus.Info("Connecting to the serial port ", conf.Serial.Port)

	if err != nil {
		logrus.Error("Unable to connect to the serial port: ", err)
		c.JSON(400, gin.H{"status": "wrong port or already busy"})
		return
	}

	lightSaber := mode.Lightsaber{
		TerminateRenderChannel: mode.TerminateRenderChannel,
	}
	
	lightSaber.StopRendering()
	lightSaber.Render(conf, serialPort)

	c.JSON(200, gin.H{"status": "started"})
}
