package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("This is used to log")

	log.SetPrefix("INFO: ")	//in each terminal output, it prefixes INFO: 
	log.Println("This is logger.")

	//log flags
	//each terminal output will have date, time & file_name.
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is log message with date and time")
	log.Println("This is log message with date and time2")

	fmt.Println("--------------------------------------------------------------------")
	infoLogger.Println("This is info logger.")
	errorLogger.Println("This is error logger.")
	warnLogger.Println("This is warn logger.")

	fmt.Println("---------------------------------------------------------------")
	//let's now print all these in a file not in terminal.
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open the logfile: %v", err)
	}
	defer logFile.Close()

	warnLogger1 := log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1 := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger1 := log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	warnLogger1.Println("This is a warning messasge.")
	infoLogger1.Println("This is an Info messasge.")
	errorLogger1.Println("This is an error messasge.")
	debugLogger1.Println("This is a debug messasge.")
}
//creating few logger that will log in terminal
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
)
