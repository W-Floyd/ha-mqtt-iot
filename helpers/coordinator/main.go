package main

import (
	"log"
	"os"

	"github.com/eidolon/wordwrap"
	"github.com/iancoleman/strcase"
)

func main() {

	// create file
	f, err := os.Create("../coordinator/generateAll.go")
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

	output := []string{"package coordinator", "import (", "\"errors\"", "\"strconv\"", "\"github.com/fatih/structs\"", ")"}

	wrapper := wordwrap.Wrapper(80, false)
	description := wrapper("GenerateAll iterates through all types of devices, as declared in the config, and runs the Generate() function on each of them. This fills in any gaps that these devices have (topics, for example), and raises errors if required.")
	comment := wordwrap.Indent(description, "// ", true)

	output = append(output, comment)

	output = append(output, "func GenerateAll() error {", "n := 0")

	output = append(output, "if !structs.IsStruct(Config.Devices) {", "return errors.New(\"devices does not exist\")", "}")

	for _, deviceName := range deviceTypes {

		output = append(output, "if len(Config.Devices."+strcase.ToCamel(deviceName)+") > 0 {")

		output = append(output, "for k := range Config.Devices."+strcase.ToCamel(deviceName)+" {")

		output = append(output, "if structs.IsStruct(Config.Devices."+strcase.ToCamel(deviceName)+"[k].Configuration) {")

		output = append(output, "Config.Devices."+strcase.ToCamel(deviceName)+"[k].Generate()")

		output = append(output, "n+=1")

		output = append(output, "} else {", "return errors.New(\"device Config.Devices."+strcase.ToCamel(deviceName)+"[\"+strconv.Itoa(k)+\"].Configuration does not exist\")")

		output = append(output, "}", "}", "}")

	}
	output = append(output, "if n == 0 {", "return errors.New(\"no devices generated.\")", "}", "return nil", "}")

	for _, line := range output {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
