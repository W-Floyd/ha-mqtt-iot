package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

var fileList = []string{
	"devicetypes",
	"deviceinit",
	"devicefunctions",
	"state",
	"subscribe",
}

func main() {

	devices := DevicesInit()
	loadKeyNames()

	files := make(map[string]*jen.File)

	for _, v := range fileList {
		files[v] = jen.NewFilePathName("../hadiscovery/"+v+".go", "hadiscovery")
		files[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
	}

	sort.Strings(keyNames)

	files["devicetypes"].Type().Id("Device").Interface(
		// jen.UnionFunc(
		// 	func(g *jen.Group) {
		// 		for _, d := range devices {
		// 			g.Add(jen.Id(strcase.ToCamel(d.Name)))
		// 		}
		// 	},
		// ),
		jen.Id("GetRawId").Params().String(),
		jen.Id("GetUniqueId").Params().String(),
		jen.Id("PopulateDevice").Params(),
		jen.Id("PopulateTopics").Params(),
		jen.Id("UpdateState").Params(),
		jen.Id("Subscribe").Params(),
		jen.Id("AddMessageHandler").Params(),
	)

	for _, d := range devices {

		// d.GetRawID()
		files["devicefunctions"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("GetRawId").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// // d.GetMQTTClient()
		// files["devicefunctions"].Func().Params(
		// 	jen.Id("d").Id(strcase.ToCamel(d.Name)),
		// ).Id("GetMQTTClient").Params().Qual("github.com/eclipse/paho.mqtt.golang", "Client").Block(
		// 	jen.Return(jen.Op("*").Id("d").Dot("MQTT").Dot("Client")),
		// )

		// d.AddMessageHandler()
		files["devicefunctions"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("AddMessageHandler").Params().Block(
			jen.Id("d").Dot("MQTT").Dot("MessageHandler").Op("=").Id("MakeMessageHandler").Params(jen.Id("d")),
		)

		// d.GetUniqueID()
		if d.JSONContainer.Exists("unique_id") {
			files["devicefunctions"].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Id("d.UniqueId")),
			)
		} else {
			files["devicefunctions"].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Lit("")),
			)
		}

		st := make(map[string][]*jen.Statement)

		// Add standalone base level fields
		for _, key := range keyNames {
			if d.JSONContainer.Exists(key) {
				st[key] = append(st[key], d.FieldAdder(key), d.FunctionAdder(key))
			}
		}

		if d.JSONContainer.Exists("device") {
			st["device"] = append(st["device"], jen.Id(strcase.ToCamel("device")).Struct(
				jen.Id(strcase.ToCamel("configuration_url")).String().Tag(map[string]string{"json": "configuration_url"}),
				jen.Id(strcase.ToCamel("connections")).Index().String().Tag(map[string]string{"json": "connections"}),
				jen.Id(strcase.ToCamel("identifiers")).Index().String().Tag(map[string]string{"json": "identifiers"}),
				jen.Id(strcase.ToCamel("manufacturer")).String().Tag(map[string]string{"json": "manufacturer"}),
				jen.Id(strcase.ToCamel("model")).String().Tag(map[string]string{"json": "model"}),
				jen.Id(strcase.ToCamel("name")).String().Tag(map[string]string{"json": "name"}),
				jen.Id(strcase.ToCamel("suggested_area")).String().Tag(map[string]string{"json": "suggested_area"}),
				jen.Id(strcase.ToCamel("sw_version")).String().Tag(map[string]string{"json": "sw_version"}),
				jen.Id(strcase.ToCamel("via_device")).String().Tag(map[string]string{"json": "via_device"}),
			).Tag(map[string]string{"json": "device"}))
		}

		// d.PopulateDevice()
		if d.JSONContainer.Exists("device") {
			files["deviceinit"].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block(
				jen.Id("d.Device.Manufacturer").Op("=").Id("Manufacturer"),
				jen.Id("d.Device.Model").Op("=").Id("SoftwareName"),
				jen.Id("d.Device.Name").Op("=").Id("InstanceName"),
				jen.Id("d.Device.SwVersion").Op("=").Id("SWVersion"),
			)
		} else {
			files["devicefunctions"].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block()
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)

		// Device MQTT Struct
		files["devicetypes"].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					for _, val := range v {
						g.Add(val)
					}
				}
				g.Id("MQTT").Id("MQTTFields").Tag(map[string]string{"json": "-"})
			},
		)

		files["state"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("UpdateState").Params().BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							// TODO - implement UpdateState for all state topics
						}
					}
				}
			},
		)

		// d.Subscribe()
		files["subscribe"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("Subscribe").Params().BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							// TODO - implement Subscribe to all topics
						}
					}
				}
			},
		)

		// d.UnSubscribe()
		files["devicefunctions"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("UnSubscribe").Params().BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							// TODO - implement Subscribe to all topics
						}
					}
				}
			},
		)

		// d.AnnounceAvailable()
		files["devicefunctions"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("AnnounceAvailable").Params().BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							// TODO - implement Subscribe to all topics
						}
					}
				}
			},
		)

		files["deviceinit"].Func().Params(jen.Id("d").Id(strcase.ToCamel(d.Name))).Id("Initialize").Params().BlockFunc(func(g *jen.Group) {
			if d.JSONContainer.Exists("retain") {
				g.Add(jen.Id("d").Dot("Retain").Op("=").Lit(false))
			}
			g.Add(jen.Id("d").Dot("PopulateDevice").Params())
			g.Add(jen.Id("d").Dot("PopulateTopics").Params())
			g.Add(jen.Id("d").Dot("AddMessageHandler").Params())
		})

		// d.PopulateTopics()
		files["deviceinit"].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("PopulateTopics").Params().BlockFunc(func(g *jen.Group) {
			for _, name := range keyNames {
				if strings.HasSuffix(name, "_topic") && d.JSONContainer.Exists(name) {
					lName := strcase.ToCamel(strings.TrimSuffix(name, "_topic"))
					g.Add(
						jen.If(
							jen.Id("d").Dot(lName + "Func").Op("!=").Nil(),
						).BlockFunc(
							func(g *jen.Group) {
								g.Add(jen.Id("d").Dot(strcase.ToCamel(name)).Op("=").Id("GetTopic").Params(jen.Id("d"), jen.Lit(name)))
								if IsCommand(name, d) {
									g.Add(jen.Id("topicStore").Index(
										jen.Id("d").Dot(strcase.ToCamel(name)),
									).Op("=").Id("&d").Dot(lName + "Func"))
								}
							},
						),
					)
				}
			}
		})

	}

	for k, v := range files {
		v.Save("../hadiscovery/" + k + ".go")
	}

}
