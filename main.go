package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"go.bug.st/serial.v1/enumerator"
	"lightsaber/config"
	"lightsaber/handler"
	"log"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: streamwiz [port]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func discoverPort() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("no serial ports found")
	}

	for _, port := range ports {
		c := &serial.Config{Name: port.Name, Baud: 115200, ReadTimeout: time.Second * 1}
		s, err := serial.OpenPort(c)
		if err == nil {
			buf := make([]byte, 128)
			n, _ := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			if n>=3 {
				msg := string(buf[:n])
				if msg[0:3] == "Ada" {
					s.Close()
					return port.Name, nil
				}
			}
			s.Close()
		}

	}
	return "", errors.New ("no adalight found")
}

func main() {

	adalightPort, errDiscover := discoverPort()
	if errDiscover != nil {
		log.Fatal (errDiscover)
	}
	config.DiscoveredPort = adalightPort

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
	r.POST("/status", handler.Status)

	err := r.Run("127.0.0.1:" + args[0])
	if err != nil {
		logrus.Fatal(err)
	}
	select {}
}
