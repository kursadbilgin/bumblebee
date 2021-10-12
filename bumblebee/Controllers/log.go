package Controllers

import (
	"log"
	"os"
)

func LogError(message, err string) {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Printf("%s error: %s", message, err)
}

func LogInfo(message string) {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	infoLog.Printf("%s Status: %d", message, 200)
}
