package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lightsaber/handler"
)

func main() {
	logrus.Info("Starting Lightsaber daemon.")
	r := gin.Default()
	r.POST("/start", handler.Start)
	r.POST("/stop", handler.Stop)

	go r.Run("127.0.0.1:8877")

	select {}
}
