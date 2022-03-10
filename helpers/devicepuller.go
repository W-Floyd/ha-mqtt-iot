package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/ghodss/yaml"
)

var DeviceNames = []string{
	"alarm_control_panel",
	"binary_sensor",
	"camera",
	"cover",
	"device_tracker",
	"device_trigger",
	"fan",
	"humidifier",
	"climate",
	"light",
	"lock",
	"number",
	"scene",
	"select",
	"sensor",
	"switch",
	"tag",
	"vacuum",
}

type Device struct {
	Name          string
	JSONContainer *gabs.Container
}

func DevicesInit() (retval []Device) {
	for _, name := range DeviceNames {
		d := Device{
			Name: name,
		}
		d.init()
		retval = append(retval, d)
	}
	return retval
}

func (dev *Device) init() {
	dev.JSONContainer = JsonExtractor(dev.Name)
}

func JsonExtractor(deviceName string) *gabs.Container {
	yamlString, err := splitDocument(deviceName)
	if err != nil {
		log.Fatal(err)
	}

	jsonDoc, err := yaml.YAMLToJSON([]byte(yamlString))
	if err != nil {
		log.Fatal(err)
	}

	jsonParsed, err := gabs.ParseJSON(jsonDoc)
	if err != nil {
		log.Fatal(err)
	}

	return jsonParsed
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func splitDocument(devicename string) (string, error) {

	data, err := fetchDocument(devicename)
	if err != nil {
		return "", err
	}

	match := between(string(data), "{% configuration %}", "{% endconfiguration %}")

	return match, nil

}

func fetchDocument(devicename string) ([]byte, error) {

	url := "https://raw.githubusercontent.com/home-assistant/home-assistant.io/rc/source/_integrations/" + devicename + ".mqtt.markdown"

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func Unquote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func TypeTranslator(t *gabs.Container) string {

	if len(t.Children()) > 0 {

		if len(t.Children()) == 1 {
			if Unquote(t.Children()[0].String()) == "list" {
				return "[]string"
			}
		}

		first := Unquote(t.Children()[0].String())
		second := Unquote(t.Children()[1].String())

		if (first == "list" && second == "map") || (second == "list" && first == "map") {
			return "[][2]string"
		} else if (first == "string" && second == "list") || (second == "string" && first == "list") {
			return "[]string"
		}
		fmt.Println("Don't know what to do with: ", t.Children())
		os.Exit(1)
	}
	v := Unquote(t.String())
	switch v {
	case "template", "icon":
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "list":
		return "[]string"
	case "map":
		return "map[string]string"
	case "float":
		return "float64"
	case "device_class":
		return "string"
	default:
		return v
	}
}
