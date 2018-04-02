package luftil

import (
	"strings"

	"gitlab.com/alexnikita/gols/gol"
)

// ColorizeStatus add color to status
// green - to 20*
// cyan - to 40*
// red - to 50*
func ColorizeStatus(status string) string {
	switch true {
	case strings.HasPrefix(status, "2"):
		return gol.Green(status)
	case strings.HasPrefix(status, "4"):
		return gol.Cyan(status)
	case strings.HasPrefix(status, "5"):
		return gol.Red(status)
	}

	return status
}
