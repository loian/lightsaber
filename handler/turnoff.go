package handler

import (
	"github.com/gin-gonic/gin"
	"lightsaber/mode"
)

func TurnOff(c *gin.Context) {
	mode.TerminateRenderChannel <- true
	c.JSON(200, gin.H{"ternimated": true})
}
