package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/W-Floyd/ha-mqtt-iot/helpers/yamlpuller"
	"github.com/iancoleman/strcase"
)

const ShowComments bool = true

func main() {

	// create file
	f, err := os.Create("../devices/translators.go")
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

		jsonParsed := yamlpuller.JsonExtractor(deviceName)

		output = append(output, generateTranslator(deviceName, jsonParsed.ChildrenMap())...)

	}

	for _, line := range output {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func generateTranslator(deviceName string, item map[string]*gabs.Container) (functionlines []string) {
	functionlines = append(functionlines, "func (entity HADevice"+strcase.ToCamel(deviceName)+"FunctionsConfig) Translate() (functions HADevice"+strcase.ToCamel(deviceName)+"Functions) {")

	keys := make([]string, 0, len(item))

	for key := range item {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		child := item[key]

		funclines := recurseTranslator(key, child.ChildrenMap(), []string{})

		functionlines = append(functionlines, funclines...)
	}

	functionlines = append(functionlines, "return", "}", "")
	return functionlines
}

func recurseTranslator(keyname string, item map[string]*gabs.Container, parentname []string) (functionlines []string) {

	if keyname == "keys" {

		keys := make([]string, 0, len(item))

		for key := range item {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			child := item[key]

			funclines := recurseTranslator(key, child.ChildrenMap(), parentname)

			functionlines = append(functionlines, funclines...)
		}

	} else {

		camelName := strcase.ToCamel(keyname)

		hasKeys := false
		if _, ok := item["keys"]; ok {
			hasKeys = true
		}

		if !hasKeys && !((strings.HasPrefix(camelName, "Set") && strings.HasSuffix(camelName, "Topic")) || strings.HasSuffix(camelName, "CommandTopic") || strings.HasSuffix(camelName, "StateTopic") || strings.HasSuffix(camelName, "StatusTopic") || camelName == "Topic") {
			return nil
		}

		var functionType string

		if strings.HasSuffix(camelName, "CommandTopic") || strings.HasPrefix(camelName, "Set") {
			functionType = "AddCommandFunction"
		} else {
			functionType = "AddStateFunction"
		}

		if camelName == "Topic" {
			camelName = "StateTopic"
		}

		camelName = strings.TrimSuffix(camelName, "Topic")

		if val, ok := item["keys"]; ok {

			parentname = append(parentname, camelName)

			funclines := recurseTranslator("keys", val.ChildrenMap(), parentname)

			if len(funclines) == 0 {
				return nil
			}

			functionlines = append(functionlines, funclines...)
		} else {
			var subname string
			for _, val := range parentname {
				subname = subname + "." + val
			}
			subname = subname + "." + camelName
			functionlines = append(functionlines, functionType+"(entity"+subname+", &functions"+subname+")")

		}

	}

	return functionlines
}
