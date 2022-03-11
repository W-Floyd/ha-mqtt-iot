package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {

	devices := DevicesInit()
	loadKeyNames()

	filesHa := make(map[string]*jen.File)

	fileList := []string{
		"types",
		"store",
	}

	for _, v := range append(DeviceNames, fileList...) {
		filesHa[v] = jen.NewFilePathName("../hadiscovery/"+v+".go", "hadiscovery")
		filesHa[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		filesHa[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		filesHa[v].Comment("Do not modify this file, it is automatically generated")
		filesHa[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	filesDev := make(map[string]*jen.File)

	fileList = []string{}

	for _, v := range append(DeviceNames, fileList...) {
		filesDev[v] = jen.NewFilePathName("../devices/"+v+".go", "iotconfig")
		filesDev[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		filesDev[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		filesDev[v].Comment("Do not modify this file, it is automatically generated")
		filesDev[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	sort.Strings(keyNames)

	filesHa["types"].Type().Id("Device").Interface(
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

	filesHa["store"].Type().Id("store").StructFunc(
		func(g *jen.Group) {
			for _, d := range devices {
				g.Add(
					jen.Id(strcase.ToCamel(d.Name)).StructFunc(
						func(h *jen.Group) {
							for _, key := range keyNames {
								if d.JSONContainer.Exists(key) && strings.HasSuffix(key, "_topic") {
									if !IsCommand(key, d) {
										h.Add(
											jen.Id(strcase.ToCamel(strings.TrimSuffix(key, "_topic"))).Map(jen.String()).String(),
										)
									}
								}
							}
						},
					),
				)
			}
		},
	)

	filesHa["store"].Func().Id("initStore").Params().Id("store").BlockFunc(
		func(g *jen.Group) {
			g.Add(
				jen.Id("s").Op(":=").Id("store").Block(),
			)
			for _, d := range devices {
				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) && strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							g.Add(
								jen.Id("s").Dot(strcase.ToCamel(d.Name)).Dot(strcase.ToCamel(strings.TrimSuffix(key, "_topic"))).
									Op("=").
									Make(jen.Map(jen.String()).String()),
							)
						}
					}
				}
			}
			g.Add(
				jen.Return(jen.Id("s")),
			)
		},
	)

	for _, d := range devices {

		camName := strcase.ToCamel(d.Name)

		// d.GetRawID()
		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("GetRawId").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// d.AddMessageHandler()
		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("AddMessageHandler").Params().Block(
			jen.Id("d").Dot("MQTT").Dot("MessageHandler").Op("=").Id("MakeMessageHandler").Params(jen.Id("d")),
		)

		// d.GetUniqueID()
		if d.JSONContainer.Exists("unique_id") {
			filesHa[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Id("d.UniqueId")),
			)
		} else {
			filesHa[d.Name].Func().Params(
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
			filesHa[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block(
				jen.Id("d.Device.Manufacturer").Op("=").Id("Manufacturer"),
				jen.Id("d.Device.Model").Op("=").Id("SoftwareName"),
				jen.Id("d.Device.Name").Op("=").Id("InstanceName"),
				jen.Id("d.Device.SwVersion").Op("=").Id("SWVersion"),
			)
		} else {
			filesHa[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block()
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)

		// Device MQTT Struct
		filesHa[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
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

		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("UpdateState").Params().BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if !IsCommand(key, d) {
							trimmed := strings.TrimSuffix(key, "_topic")
							cam := strcase.ToCamel(key)
							camTrimmed := strcase.ToCamel(trimmed)
							g.Add(
								jen.If(
									jen.Id("d").Dot(cam).Op("!=").Lit(""),
								).Block(
									jen.Id("state").Op(":=").Id("d").Dot(strcase.ToCamel(trimmed+"_func")).Params(),
									jen.If(
										jen.Id("state").Op("!=").Id("stateStore").Dot(strcase.ToCamel(d.Name)).Dot(camTrimmed).Index(jen.Id("d").Dot("UniqueId")).
											Op("||").
											Id("d").Dot("MQTT").Dot("ForceUpdate"),
										// state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
									).Block(
										jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
										jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
											jen.Id("d").Dot(cam),
											jen.Id("qos"),
											jen.Id("retain"),
											jen.Id("state"),
										),
										jen.Id("stateStore").Dot(camName).Dot(camTrimmed).Index(jen.Id("d").Dot("UniqueId")).Op("=").Id("state"),
										jen.Id("token").Dot("Wait").Params(),
									),
								),
							)
						}
					}
				}
			},
		)

		// d.Subscribe()
		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("Subscribe").Params().BlockFunc(
			func(g *jen.Group) {

				g.Add(
					jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
				)

				g.Add(
					jen.List(jen.Id("message"), jen.Err()).Op(":=").Qual("encoding/json", "Marshal").Params(jen.Id("d")),
				)
				g.Add(
					jen.If(
						jen.Id("err").Op("!=").Id("nil"),
					).Block(
						jen.Qual("log", "Fatal").Params(jen.Err()),
					),
				)

				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "_topic") {
						if IsCommand(key, d) {
							// trimmed := strings.TrimSuffix(key, "_topic")
							cam := strcase.ToCamel(key)
							// camTrimmed := strcase.ToCamel(trimmed)

							g.Add(
								jen.If(
									jen.Id("d").Dot(cam).Op("!=").Lit(""),
								).Block(
									jen.Id("t").Op(":=").Id("c").Dot("Subscribe").Params(
										jen.Id("d").Dot(cam),
										jen.Lit(0),
										jen.Id("d").Dot("MQTT").Dot("MessageHandler"),
									),
									jen.Id("t").Dot("Wait").Params(),
									jen.If(
										jen.Id("t").Dot("Error").Params().Op("!=").Nil(),
									).Block(
										jen.Qual("log", "Fatal").Params(jen.Id("t").Dot("Error").Params()),
									),
								),
							)
						}
					}
				}

				g.Add(
					jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
						jen.Id("GetDiscoveryTopic").Params(jen.Id("d")),
						jen.Lit(0),
						jen.Lit(true),
						jen.Id("message"),
					),
				)

				g.Add(
					jen.Id("token").Dot("Wait").Params(),
				)

				g.Add(
					jen.Qual("time", "Sleep").Params(jen.Lit(500).Op("*").Qual("time", "Millisecond")),
				)

				g.Add(
					jen.Id("d").Dot("AnnounceAvailable").Params(),
				)

				g.Add(
					jen.Id("d").Dot("UpdateState").Params(),
				)

			},
		)

		// d.UnSubscribe()
		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("UnSubscribe").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)

					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Id("d").Dot("AvailabilityTopic"),
							jen.Id("qos"),
							jen.Id("retain"),
							jen.Lit("offline"),
						),
					)

					g.Add(
						jen.Id("token").Dot("Wait").Params(),
					)

					for _, key := range sortedKeys {
						if strings.HasSuffix(key, "_topic") {
							if IsCommand(key, d) {
								// trimmed := strings.TrimSuffix(key, "_topic")
								cam := strcase.ToCamel(key)
								// camTrimmed := strcase.ToCamel(trimmed)

								g.Add(
									jen.If(
										jen.Id("d").Dot(cam).Op("!=").Lit(""),
									).Block(
										jen.Id("t").Op(":=").Id("c").Dot("Unsubscribe").Params(
											jen.Id("d").Dot(cam),
										),
										jen.Id("t").Dot("Wait").Params(),
										jen.If(
											jen.Id("t").Dot("Error").Params().Op("!=").Nil(),
										).Block(
											jen.Qual("log", "Fatal").Params(jen.Id("t").Dot("Error").Params()),
										),
									),
								)
							}
						}
					}
				}
			},
		)

		// d.AnnounceAvailable()
		filesHa[d.Name].Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("AnnounceAvailable").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)
					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Id("d").Dot("AvailabilityTopic"),
							jen.Id("qos"),
							jen.Id("retain"),
							jen.Lit("online"),
						),
					)
					g.Add(
						jen.Id("token").Dot("Wait").Params(),
					)
				}
			},
		)

		filesHa[d.Name].Func().Params(jen.Id("d").Id(strcase.ToCamel(d.Name))).Id("Initialize").Params().BlockFunc(func(g *jen.Group) {
			if d.JSONContainer.Exists("retain") {
				g.Add(jen.Id("d").Dot("Retain").Op("=").Lit(false))
			}
			g.Add(jen.Id("d").Dot("PopulateDevice").Params())
			g.Add(jen.Id("d").Dot("PopulateTopics").Params())
			g.Add(jen.Id("d").Dot("AddMessageHandler").Params())
		})

		// d.PopulateTopics()
		filesHa[d.Name].Func().Params(
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

	for k, v := range filesHa {
		v.Save("../hadiscovery/" + k + ".go")
	}

	for k, v := range filesDev {
		v.Save("../devices/" + k + ".go")
	}

}
