package yamlpuller

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/ghodss/yaml"
)

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

func Unquote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func TypeTranslator(t string) string {
	t = Unquote(t)
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
