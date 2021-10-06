package coordinator

import (
	"errors"
	"io/ioutil"

	"github.com/W-Floyd/ha-mqtt-iot/devices"
	"github.com/W-Floyd/ha-mqtt-iot/logging"
	"gopkg.in/yaml.v3"
)

var Config devices.Config

func IngestConfigFile(filename string) error {

	err := UnmarshalYAML(filename)

	if err != nil {
		return err
	}

	err = GenerateAll()

	if err != nil {
		logging.LogError(err)
	}

	return nil
}

func UnmarshalYAML(filename string) error {

	// read file
	rawData, err := ioutil.ReadFile(filename)
	if err != nil {
		logging.LogError("Error reading config", err)
	}

	if Config != *new(devices.Config) {
		return errors.New("config is not empty")
	}

	err = yaml.Unmarshal(rawData, &Config)
	if err != nil {
		logging.LogError("Error unmashalling config", err)
	}

	return nil
}

//go:generate go run ../helpers/coordinator/main.go
//go:generate gofmt -w generateAll.go
