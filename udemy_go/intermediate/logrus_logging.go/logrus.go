package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	//set log level
	log.SetLevel(logrus.InfoLevel)

	//set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//logging examples
	log.Info("This is an Info Messsage")
	log.Warn("This is an Warm Messsage")
	log.Error("This is an Info Messsage")

	log.WithFields(logrus.Fields{
		"username" : "John Doe",
		"method" : "GET",
	}).Info("User logged in.")
}