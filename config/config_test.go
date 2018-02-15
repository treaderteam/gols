package config

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type doubleExpectation struct {
	Value interface{}
	Err   error
}

type testCase struct {
	Testable    interface{}
	Expectation doubleExpectation
}

func TestDetectExt(t *testing.T) {
	cases := []testCase{
		testCase{
			Testable: "test.json",
			Expectation: doubleExpectation{
				Value: EXT_JSON,
				Err:   nil,
			},
		},
		testCase{
			Testable: "test.yml",
			Expectation: doubleExpectation{
				Value: EXT_YAML,
				Err:   nil,
			},
		},
		testCase{
			Testable: "test.yaml",
			Expectation: doubleExpectation{
				Value: EXT_YAML,
				Err:   nil,
			},
		},
		testCase{
			Testable: "test",
			Expectation: doubleExpectation{
				Value: 0,
				Err:   InvalidFormatError{},
			},
		},
	}

	for _, v := range cases {
		s := v.Testable.(string)
		value, err := detectExtension(s)
		assert.Equal(t, v.Expectation.Value, value)
		assert.Equal(t, v.Expectation.Err, err)
	}
}

func TestParseConfigFromJSON(t *testing.T) {
	var (
		file     *os.File
		filename = "test.json"
		filedata = map[string]string{
			"cpus": "2",
			"pi":   "3.14",
		}
		err error
	)

	defer func(f string) {
		os.Remove(f)
	}(filename)

	if file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		log.Panic(err)
	}

	defer file.Close()

	if err = json.NewEncoder(file).Encode(&filedata); err != nil {
		log.Panic(err)
	}

	if err = parseConfigFromJSON(filename, &config); err != nil {
		log.Panic(err)
	}

	assert.Equal(t, filedata["cpus"], Request("cpus", false))
	assert.Equal(t, filedata["pi"], Request("pi", false))
}
