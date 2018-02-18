// Package config responsible for providing all neccessary configuration
// data
package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	filepathFlag = flag.String("gols_filepath", "", "specify config filepath to parse")
	config       interface{}
)

const (
	CONFIG_FILEPATH = "GOLS_CONFIG_FILEPATH"
	EXT_JSON        = iota
	EXT_YAML
)

func init() {
	log.SetPrefix("GOLS CONFIG: ")
	log.Println("Gols init")
	flag.Parse()
	log.Println(*filepathFlag)

	if *filepathFlag != "" {
		parse(*filepathFlag)
	} else if filepath := os.Getenv(CONFIG_FILEPATH); filepath != "" {
		parse(filepath)
	}
}

func parse(filepath string) {
	ext, err := detectExtension(filepath)
	if err != nil {
		log.Fatal(err)
	}

	switch ext {
	case EXT_JSON:
		err = parseConfigFromJSON(filepath, config)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}

func parseConfigFromJSON(filepath string, dest interface{}) (err error) {
	var (
		file   *os.File
		entity []byte
	)

	if file, err = os.Open(filepath); err != nil {
		return
	}

	defer file.Close()

	if entity, err = ioutil.ReadAll(file); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(entity, &config); err != nil {
		return
	}

	return
}

func detectExtension(filename string) (result int, err error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		err = InvalidFormatError{}
		return
	}
	ext := parts[len(parts)-1]
	switch ext {
	case "json":
		return EXT_JSON, nil
	case "yml":
		return EXT_YAML, nil
	case "yaml":
		return EXT_YAML, nil
	}

	return
}
