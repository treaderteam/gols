package config

import (
	"log"
	"os"
)

const loggerOUT = "LOGGER_OUT"

// Logout get logout env variable and
// targets log messages to it
func Logout() {
	loggerOut := Request(loggerOUT, false)

	if loggerOut != "" {
		file, err := os.OpenFile(loggerOut, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		} else {
			log.SetOutput(file)
		}
	}
}
