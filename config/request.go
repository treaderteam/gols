package config

import (
	"log"
	"os"
)

// Request requests config params
func Request(name string, strict bool) (result string) {
	t, ok := config.(map[string]interface{})
	if ok {
		v, exists := t[name]
		if exists {
			result, ok = v.(string)
			if ok {
				return
			}
		}
	}

	result = os.Getenv(name)
	if strict && len(result) < 1 {
		log.Fatalf("config %s requested but not specified, aborting...\n", name)
	}
	return
}
