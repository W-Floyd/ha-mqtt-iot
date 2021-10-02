package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/W-Floyd/ha-mqtt-iot/helpers/yamlpuller"
	"github.com/eidolon/wordwrap"
	"github.com/iancoleman/strcase"
)

const ShowComments bool = true

func main() {

	// create file
	f, err := os.Create("../devices/structs.go")
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

	output := []string{"package devices", "import (", "mqtt \"github.com/eclipse/paho.mqtt.golang\"", ")"}

	configOut := []string{"type Config struct {", "Devices *struct{"}

	for _, deviceName := range deviceTypes {

		jsonParsed := yamlpuller.JsonExtractor(deviceName)

		output = append(output, generateDevice(deviceName, jsonParsed.ChildrenMap())...)

		funct, conf := generateFunctions(deviceName, jsonParsed.ChildrenMap())

		output = append(output, funct...)
		output = append(output, conf...)

		configOut = append(configOut, strcase.ToCamel(deviceName)+" []struct {", "Functions *HADevice"+strcase.ToCamel(deviceName)+"FunctionsConfig `yaml:\"functions,omitempty\"`", "Configuration *HADevice"+strcase.ToCamel(deviceName)+"`yaml:\"configuration,omitempty\"`", "} `yaml:\""+strcase.ToSnake(deviceName)+",omitempty\"`")

	}

	configOut = append(configOut, "} `yaml:\"devices,omitempty\"`", "}")

	output = append(output, configOut...)

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

	keys := make([]string, 0, len(item))

	for key := range item {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		child := item[key]
		returnlines = append(returnlines, recurseItem(key, child.ChildrenMap())...)
	}

	returnlines = append(returnlines, "}", "")
	return returnlines
}

func generateFunctions(deviceName string, item map[string]*gabs.Container) (functionlines, configlines []string) {
	functionlines = append(functionlines, "type HADevice"+strcase.ToCamel(deviceName)+"Functions struct {")
	configlines = append(configlines, "type HADevice"+strcase.ToCamel(deviceName)+"FunctionsConfig struct {")

	keys := make([]string, 0, len(item))

	for key := range item {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		child := item[key]

		funclines, conflines := recurseItemFunctions(key, child.ChildrenMap())

		functionlines = append(functionlines, funclines...)
		configlines = append(configlines, conflines...)
	}

	functionlines = append(functionlines, "}", "")
	configlines = append(configlines, "}", "")
	return functionlines, configlines
}

func recurseItemFunctions(keyname string, item map[string]*gabs.Container) (functionlines, configlines []string) {

	if keyname == "keys" {

		keys := make([]string, 0, len(item))

		for key := range item {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			child := item[key]

			funclines, conflines := recurseItemFunctions(key, child.ChildrenMap())

			functionlines = append(functionlines, funclines...)
			configlines = append(configlines, conflines...)
		}

	} else {

		camelName := strcase.ToCamel(keyname)

		hasKeys := false
		if _, ok := item["keys"]; ok {
			hasKeys = true
		}

		if !hasKeys && !((strings.HasPrefix(camelName, "Set") && strings.HasSuffix(camelName, "Topic")) || strings.HasSuffix(camelName, "CommandTopic") || strings.HasSuffix(camelName, "StateTopic") || strings.HasSuffix(camelName, "StatusTopic") || camelName == "Topic") {
			return nil, nil
		}

		var functionType string

		if strings.HasSuffix(camelName, "CommandTopic") || strings.HasPrefix(camelName, "Set") {
			functionType = "func(mqtt.Message, mqtt.Client)"
		} else {
			functionType = "func() string"
		}

		if camelName == "Topic" {
			camelName = "StateTopic"
		}

		camelName = strings.TrimSuffix(camelName, "Topic")

		if val, ok := item["keys"]; ok {
			funclines, conflines := recurseItemFunctions("keys", val.ChildrenMap())

			if len(funclines) == 0 {
				return nil, nil
			}

			configlines = append(configlines, camelName+" struct {")
			configlines = append(configlines, conflines...)
			configlines = append(configlines, "}"+"`yaml:\""+strcase.ToSnake(camelName)+",omitempty\"`")

			functionlines = append(functionlines, camelName+" struct {")
			functionlines = append(functionlines, funclines...)
			functionlines = append(functionlines, "}")
		} else {
			functionlines = append(functionlines, camelName+" "+functionType)
			configlines = append(configlines, camelName+" *[]string `yaml:\""+strcase.ToSnake(camelName)+",omitempty\"`")
		}

	}

	return functionlines, configlines
}

func recurseItem(keyname string, item map[string]*gabs.Container) (returnlines []string) {

	if keyname == "keys" {

		keys := make([]string, 0, len(item))

		for key := range item {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			child := item[key]

			returnlines = append(returnlines, recurseItem(key, child.ChildrenMap())...)
		}
	} else {

		camelName := strcase.ToCamel(keyname)

		wrapper := wordwrap.Wrapper(80, false)
		description := wrapper(yamlpuller.Unquote(item["description"].String()))
		commented := wordwrap.Indent(description, "// ", true)
		if ShowComments {
			returnlines = append(returnlines, commented)
		}

		if val, ok := item["keys"]; ok {
			returnlines = append(returnlines, camelName+" struct {")
			returnlines = append(returnlines, recurseItem("keys", val.ChildrenMap())...)
			returnlines = append(returnlines, "}")
		} else {

			localType := yamlpuller.TypeTranslator(item["type"])

			returnlines = append(returnlines, camelName+" *"+localType)
		}

		returnlines[len(returnlines)-1] = returnlines[len(returnlines)-1] + " `yaml:\"" + keyname + ",omitempty\"`"
	}

	return returnlines
}
