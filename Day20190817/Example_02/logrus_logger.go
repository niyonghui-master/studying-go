package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 如果多个地方使用logging，可以创建一个logrus实例Logger

var log = logrus.New()

func main() {
	file, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	file.Close()
}
