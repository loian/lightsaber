package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lightsaber/handler"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: streamwiz [port]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("port is missing.")
		os.Exit(1)
	}

	logrus.Info("Starting Lightsaber daemon.")
	r := gin.Default()
	r.POST("/start", handler.Start)
	r.POST("/stop", handler.Stop)
	r.POST("/quit", handler.Quit)

	err := r.Run("127.0.0.1:" + args[0])
	if err != nil {
		logrus.Fatal(err)
	}
	select {}
}
