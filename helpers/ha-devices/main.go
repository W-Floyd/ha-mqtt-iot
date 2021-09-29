package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/eidolon/wordwrap"
	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
)

const ShowComments bool = false

func main() {

	// create file
	f, err := os.Create("../devices/autogen.go")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	deviceTypes := []string{
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

	output := []string{"package devices"}

	for _, deviceName := range deviceTypes {

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

		output = append(output, generateDevice(deviceName, jsonParsed.ChildrenMap())...)

	}

	for _, line := range output {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func generateDevice(deviceName string, item map[string]*gabs.Container) (returnlines []string) {
	if ShowComments {
		returnlines = append(returnlines, "// "+deviceName)
	}
	returnlines = append(returnlines, "type HADevice"+strcase.ToCamel(deviceName)+" struct {")
	for key, child := range item {
		returnlines = append(returnlines, recurseItem(key, child.ChildrenMap())...)
	}
	returnlines = append(returnlines, "}", "")
	return returnlines
}

func recurseItem(keyname string, item map[string]*gabs.Container) (returnlines []string) {

	if keyname == "keys" {
		for key, child := range item {
			returnlines = append(returnlines, recurseItem(key, child.ChildrenMap())...)
		}
	} else {

		camelName := strcase.ToCamel(keyname)

		wrapper := wordwrap.Wrapper(80, false)
		description := wrapper(unquote(item["description"].String()))
		commented := wordwrap.Indent(description, "// ", true)
		if ShowComments {
			returnlines = append(returnlines, commented)
		}

		if val, ok := item["keys"]; ok {
			returnlines = append(returnlines, camelName+" struct {")
			returnlines = append(returnlines, recurseItem("keys", val.ChildrenMap())...)
			returnlines = append(returnlines, "}")
		} else {
			var localType string
			if len(item["type"].Children()) > 0 {
				localType = "[]string"
			} else {
				localType = typeTranslator(item["type"].String())
			}
			returnlines = append(returnlines, camelName+" "+localType)
		}

		returnlines[len(returnlines)-1] = returnlines[len(returnlines)-1] + " `json:\"" + keyname + ",omitempty\"`"
	}

	return returnlines
}

func typeTranslator(t string) string {
	t = unquote(t)
	switch t {
	case "template", "icon":
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "list":
		return "[]string"
	case "map":
		return ""
	case "float":
		return "float64"
	case "device_class":
		return "string"
	default:
		return t
	}
}

func unquote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
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

	url := "https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/" + devicename + ".mqtt.markdown"

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
