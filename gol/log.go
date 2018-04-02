package gol

import (
	"log"
	"os"
)

var defaultLogger *log.Logger

func init() {
	defaultLogger = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
}
