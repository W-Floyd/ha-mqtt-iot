package main

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/dave/jennifer/jen"
	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
)

var DeviceNames = []string{
	"alarm_control_panel",
	"binary_sensor",
	"button",
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
	"siren",
	"switch",
	"tag",
	"vacuum",
}

type Device struct {
	Name          string
	JSONContainer *gabs.Container
}

var PullNew = false

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
		log.Println("Failed to find first match")
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		log.Println("Failed to find last match")
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		log.Println("First is after last")
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func splitDocument(devicename string) (string, error) {

	data, err := fetchDocument(devicename)
	if err != nil {
		return "", err
	}

	dat := string(data)

	if devicename == "vacuum" {
		dat = dat[strings.Index(dat, "State Configuration"):]
	}

	match := between(dat, "{% configuration %}", "{% endconfiguration %}")

	targetFile := "../helpers/cache/" + devicename + "_split.yaml"

	err = ioutil.WriteFile(targetFile, []byte(match), fs.FileMode(0644))
	if err != nil {
		return "nil", err
	}

	return match, nil

}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func fetchDocument(devicename string) ([]byte, error) {

	url := "https://raw.githubusercontent.com/home-assistant/home-assistant.io/rc/source/_integrations/" + devicename + ".mqtt.markdown"

	targetFile := "../helpers/cache/" + devicename + ".md"

	if exists(targetFile) && PullNew == false {

		data, err := ioutil.ReadFile(targetFile)
		if err == nil {
			return data, nil
		}

		log.Println("Loading " + targetFile + " failed")

	}

	log.Println("Fetching " + devicename)

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

	err = ioutil.WriteFile(targetFile, bodyBytes, fs.FileMode(0644))
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

func TypeTranslator(t string, s *jen.Statement) *jen.Statement {

	v := Unquote(t)
	switch v {
	case "string", "template", "icon", "device_class":
		return s.String()
	case "integer":
		return s.Int()
	case "boolean":
		return s.Bool()
	case "list", "[\"list\"]", "[\"string\",\"list\"]":
		return s.Index().String()
	case "map":
		return s.Map(jen.String()).String()
	case "float":
		return s.Float64()
	default:
		log.Fatalln("Unknown type " + t)
		return nil
	}
}

func (dev *Device) FieldAdder(key string) *jen.Statement {
	dat := dev.JSONContainer.ChildrenMap()

	t := Unquote(dat[key].ChildrenMap()["type"].String())

	return TypeTranslator(t, jen.Id(strcase.ToCamel(key))).Tag(map[string]string{"json": key})
}
