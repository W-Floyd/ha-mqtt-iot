package main

import (
	"bufio"
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
	f, err := os.Create("../devices/generators.go")
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

	output := []string{"package devices", "import (", "\"github.com/clarketm/json\"", "\"github.com/W-Floyd/ha-mqtt-iot/logging\"", ")", "func DeepCopy(a, b interface{}) {", "byt, _ := json.Marshal(a)", "json.Unmarshal(byt, b)", "}"}

	for _, deviceName := range deviceTypes {

		jsonParsed := yamlpuller.JsonExtractor(deviceName)

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
	returnlines = append(returnlines, "func (entity HADevice"+strcase.ToCamel(deviceName)+") Generate"+"() (output HADevice"+strcase.ToCamel(deviceName)+") {", "DeepCopy(&entity,&output)")

	keys := make([]string, 0, len(item))

	for key := range item {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		child := item[key]
		returnlines = append(returnlines, recurseItem(key, child.ChildrenMap(), []string{})...)
	}

	returnlines = append(returnlines, "return", "}")

	return returnlines
}

func recurseItem(keyname string, item map[string]*gabs.Container, parentname []string) (returnlines []string) {

	if keyname == "keys" {

		keys := make([]string, 0, len(item))

		for key := range item {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			child := item[key]

			returnlines = append(returnlines, recurseItem(key, child.ChildrenMap(), parentname)...)
		}
	} else {

		camelName := strcase.ToCamel(keyname)

		if val, ok := item["keys"]; ok {
			parentname = append(parentname, camelName)
			returnlines = append(returnlines, recurseItem("keys", val.ChildrenMap(), parentname)...)
		} else {

			var substring string

			for _, name := range parentname {
				substring = substring + "." + name
			}

			var localType string
			if len(item["type"].Children()) > 0 {
				localType = "[]string"
			} else {
				localType = yamlpuller.TypeTranslator(item["type"].String())
			}

			returnlines = append(returnlines, generateIfEmpty(substring+"."+camelName, "output"+substring+"."+camelName, localType, item)...)
		}

	}

	return returnlines
}

func fetchSnippet(parameterName string) (error, []string) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("../helpers/generators/snippets/" + parameterName + ".go")

	if err != nil {
		return err, []string{}

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return nil, text
}

func generateIfEmpty(source, target, ty string, item map[string]*gabs.Container) (retval []string) {

	hasSnippet := true
	err, snippet := fetchSnippet(strings.ReplaceAll(source, ".", ""))

	if err != nil || len(snippet) == 0 {
		hasSnippet = false
	}

	source = "entity" + source

	isRequired := item["required"].String() == "true"

	if hasSnippet || isRequired {
		retval = append(retval, "if "+source+" == nil {")
	} else {
		return []string{}
	}

	if hasSnippet {
		retval = append(retval, snippet...)
		retval = append(retval, "}")
		return
	} else if isRequired {
		retval = append(retval, "logging.LogError(\""+source+" generator not found, but field is required!\")", "}")
		return
	}

	return retval
}
