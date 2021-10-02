package main

import (
	"log"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/W-Floyd/ha-mqtt-iot/helpers/yamlpuller"
	"github.com/iancoleman/strcase"
)

func main() {

	// create file
	f, err := os.Create("../devices/topics.go")
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

	output := []string{"package devices", "import (", "\"github.com/W-Floyd/ha-mqtt-iot/common\"", "\"github.com/iancoleman/strcase\"", ")"}

	for _, deviceName := range deviceTypes {

		jsonParsed := yamlpuller.JsonExtractor(deviceName)

		output = append(output, generate(deviceName, jsonParsed.ChildrenMap())...)

	}

	for _, line := range output {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func generate(componentName string, item map[string]*gabs.Container) (returnlines []string) {

	returnlines = append(returnlines, "func (component HADevice"+strcase.ToCamel(componentName)+") GetTopicBase() (out string) {", "out = common.DiscoveryPrefix + \"/"+strcase.ToSnake(componentName)+"/\" + common.NodeID + \"/\"")

	if componentName == "device_trigger" || componentName == "device_tracker" {
		returnlines = append(returnlines, "return", "}")
		returnlines = append(returnlines, "func (component HADevice"+strcase.ToCamel(componentName)+") CanGenerateTopic() bool {", "if  common.NodeID == \"\"  || common.DiscoveryPrefix == \"\" {", "return false", "}", "return true", "}")
	} else if componentName == "scene" || componentName == "vacuum" {
		returnlines = append(returnlines, "out += strcase.ToSnake(*component.Name) + \"/\"", "return", "}")
		returnlines = append(returnlines, "func (component HADevice"+strcase.ToCamel(componentName)+") CanGenerateTopic() bool {", "if component.Name == nil || common.NodeID == \"\"  || common.DiscoveryPrefix == \"\" {", "return false", "}", "return true", "}")
	} else if componentName == "tag" {
		returnlines = append(returnlines, "out += strcase.ToSnake(*component.Device.Name) + \"/\"", "return", "}")
		returnlines = append(returnlines, "func (component HADevice"+strcase.ToCamel(componentName)+") CanGenerateTopic() bool {", "if component.Device.Name == nil || common.NodeID == \"\"  || common.DiscoveryPrefix == \"\" {", "return false", "}", "return true", "}")
	} else {
		returnlines = append(returnlines, "out += strcase.ToSnake(*component.Name) + \"/\"", "return", "}")
		returnlines = append(returnlines, "func (component HADevice"+strcase.ToCamel(componentName)+") CanGenerateTopic() bool {", "if ( component.Device.Name == nil && component.Name == nil ) || common.DiscoveryPrefix == \"\" || common.NodeID == \"\" {", "return false", "}", "return true", "}")
	}

	return returnlines
}
