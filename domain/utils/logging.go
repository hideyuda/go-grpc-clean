package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// refresh logFile if it gets too big
	if err := os.Truncate(logFile, 1000000); err != nil {
		os.Remove(logFile)
		fmt.Println("log file is removed")
	}

	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multilogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multilogfile)

}
