package gol

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// GetDefaultLogger returns basic logger
func GetDefaultLogger() *log.Logger {
	return defaultLogger
}

// CreateLogger wraps basic logger creation
func CreateLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+":", log.Lshortfile|log.LstdFlags)
}

// Red return string written in red color
func Red(v string) string {
	red := color.New(color.FgRed).SprintfFunc()
	return red(v)
}

// Cyan return string written in Cyan color
func Cyan(v string) string {
	cyan := color.New(color.FgCyan).SprintfFunc()
	return cyan(v)
}

// Green return string written in Green color
func Green(v string) string {
	green := color.New(color.FgGreen).SprintfFunc()
	return green(v)
}
