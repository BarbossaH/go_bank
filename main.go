package main

import (
	"bankserver/app"
	"bankserver/logger"
)



func main() {
	// log.Println("Starting or application...")
	// logger.Log.Info("Starting the application...")
	logger.Info("Starting the app...")
	app.Start()
}


