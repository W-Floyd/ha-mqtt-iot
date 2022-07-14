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

	external := make(map[string]*jen.File)

	fileList := []string{
		"types",
		"store",
	}

	for _, v := range append(DeviceNames, fileList...) {
		external[v] = jen.NewFilePathName("./devices/externaldevice/"+v+".go", "ExternalDevice")
		external[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		external[v].Comment("Do not modify this file, it is automatically generated")
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	internal := make(map[string]*jen.File)

	fileList = []string{"types"}

	for _, v := range append(DeviceNames, fileList...) {
		internal[v] = jen.NewFilePathName("./devices/internaldevice/"+v+".go", "InternalDevice")
		internal[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		internal[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		internal[v].Comment("Do not modify this file, it is automatically generated")
		internal[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		internal[v].Comment("")
	}

	sort.Strings(keyNames)

	internal["types"].Type().Id("Device").Interface(
		jen.Id("Translate").Params().String(),
	)

	external["types"].Type().Id("Device").Interface(
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
		jen.Id("UnSubscribe").Params(),
		jen.Id("AddMessageHandler").Params(),
		jen.Id("SetMQTTFields").Params(
			jen.Id("MQTTFields"),
		),
		jen.Id("GetMQTTFields").Params().Id("MQTTFields"),
	)

	external["store"].Type().Id("StateStore").StructFunc(
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

	external["store"].Func().Id("initStore").Params().Id("StateStore").BlockFunc(
		func(g *jen.Group) {
			g.Add(
				jen.Id("s").Op(":=").Id("StateStore").Block(),
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
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("GetRawId").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// d.AddMessageHandler()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("AddMessageHandler").Params().Block(
			jen.Id("d").Dot("MQTT").Dot("MessageHandler").Op("=").Id("MakeMessageHandler").Params(jen.Id("d")),
		)

		// d.GetUniqueID()
		if d.JSONContainer.Exists("unique_id") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Op("*").Id("d").Dot("UniqueId")),
			)
		} else {
			external[d.Name].Func().Params(
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
			st["device"] = append(st["device"], jen.Id(strcase.ToCamel("device")).StructFunc(
				func(g *jen.Group) {
					for _, v := range []string{
						"configuration_url",
						"connections",
						"identifiers",
						"manufacturer",
						"model",
						"name",
						"suggested_area",
						"sw_version",
						"viadevice",
					} {
						g.Add(
							jen.Id(strcase.ToCamel(v)).Op("*").String().Tag(map[string]string{"json": v + ",omitempty"}).Comment(d.JSONContainer.Path("device.keys." + v + ".description").String()),
						)
					}
				},
			).Tag(map[string]string{"json": "device,omitempty"}))
		}

		// d.PopulateDevice()
		if d.JSONContainer.Exists("device") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block(
				jen.Id("d.Device.Manufacturer").Op("=&").Id("Manufacturer"),
				jen.Id("d.Device.Model").Op("=&").Id("SoftwareName"),
				jen.Id("d.Device.Name").Op("=&").Id("InstanceName"),
				jen.Id("d.Device.SwVersion").Op("=&").Id("SWVersion"),
				jen.Id("d.Device.Identifiers").Op("=&").Id("common").Dot("MachineID"),
			)
		} else {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block()
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)

		// Device MQTT Struct
		external[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					for _, val := range v {
						g.Add(val)
					}
				}
				g.Id("MQTT").Op("*").Id("MQTTFields").Tag(map[string]string{"json": "-"})
			},
		)

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
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
									jen.Id("d").Dot(cam).Op("!=").Nil(),
								).Block(
									jen.Id("state").Op(":=").Id("d").Dot(strcase.ToCamel(trimmed+"_func")).Params(),
									jen.If(
										jen.Id("state").Op("!=").Id("stateStore").Dot(strcase.ToCamel(d.Name)).Dot(camTrimmed).Index(jen.Op("*").Id("d").Dot("UniqueId")).
											Op("||").
											Params(jen.Id("d").Dot("MQTT").Dot("ForceUpdate").Op("!=").Nil().Op("&&").Op("*").Id("d").Dot("MQTT").Dot("ForceUpdate")),
										// state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
									).Block(
										jen.Id("token").Op(":=").Params(jen.Op("*").Id("d").Dot("MQTT").Dot("Client")).Dot("Publish").ParamsFunc(
											func(g *jen.Group) {
												g.Add(jen.Op("*").Id("d").Dot(cam))

												if d.JSONContainer.Exists("qos") {
													g.Add(jen.Byte().Params(jen.Op("*").Id("d").Dot("Qos")))
												} else {
													g.Add(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "QoS"))
												}

												if d.JSONContainer.Exists("retain") {
													g.Add(jen.Op("*").Id("d").Dot("Retain"))
												} else {
													g.Add(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "Retain"))
												}

												g.Add(jen.Id("state"))
											},
										),
										jen.Id("stateStore").Dot(camName).Dot(camTrimmed).Index(jen.Op("*").Id("d").Dot("UniqueId")).Op("=").Id("state"),
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
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
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
									jen.Op("*").Id("d").Dot(cam).Op("!=").Lit(""),
								).Block(
									jen.Id("t").Op(":=").Id("c").Dot("Subscribe").Params(
										jen.Op("*").Id("d").Dot(cam),
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
					jen.Qual("time", "Sleep").Params(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "HADiscoveryDelay")),
				)

				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) {
						if key == "availability_topic" {
							g.Add(
								jen.Id("d").Dot("AvailabilityFunc").Params(),
							)
						}
					}
				}

				g.Add(
					jen.Id("d").Dot("UpdateState").Params(),
				)

			},
		)

		// d.UnSubscribe()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("UnSubscribe").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)

					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Op("*").Id("d").Dot("AvailabilityTopic"),
							jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "QoS"),
							jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "Retain"),
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
										jen.Op("*").Id("d").Dot(cam).Op("!=").Lit(""),
									).Block(
										jen.Id("t").Op(":=").Id("c").Dot("Unsubscribe").Params(
											jen.Op("*").Id("d").Dot(cam),
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
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("AnnounceAvailable").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)
					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Op("*").Id("d").Dot("AvailabilityTopic"),
							jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "QoS"),
							jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "Retain"),
							jen.Lit("online"),
						),
					)
					g.Add(
						jen.Id("token").Dot("Wait").Params(),
					)
				}
			},
		)

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("Initialize").Params().BlockFunc(func(g *jen.Group) {
			if d.JSONContainer.Exists("qos") {
				g.Add(
					jen.If(
						jen.Id("d").Dot("Qos").Op("==").Nil(),
					).Block(
						jen.Id("d").Dot("Qos").Op("=").New(jen.Int()),
						jen.Op("*").Id("d").Dot("Qos").Op("=").Int().Params(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/common", "QoS")),
					),
				)
			}
			if d.JSONContainer.Exists("retain") {
				g.Add(
					jen.If(
						jen.Id("d").Dot("Retain").Op("==").Nil(),
					).Block(
						jen.Id("d").Dot("Retain").Op("=").New(jen.Bool()),
						jen.Op("*").Id("d").Dot("Retain").Op("=").Qual("github.com/W-Floyd/ha-mqtt-iot/common", "Retain"),
					),
				)
			}
			g.Add(jen.Id("d").Dot("PopulateDevice").Params())
			g.Add(jen.Id("d").Dot("AddMessageHandler").Params())
			g.Add(jen.Id("d").Dot("PopulateTopics").Params())
		})

		// d.PopulateTopics()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("PopulateTopics").Params().BlockFunc(func(g *jen.Group) {
			for _, name := range keyNames {
				if strings.HasSuffix(name, "_topic") && d.JSONContainer.Exists(name) {
					lName := strcase.ToCamel(strings.TrimSuffix(name, "_topic"))
					g.Add(
						jen.If(
							jen.Id("d").Dot(lName + "Func").Op("!=").Nil(),
						).BlockFunc(
							func(g *jen.Group) {
								g.Add(jen.Id("d").Dot(strcase.ToCamel(name)).Op("=").New(jen.String()))
								g.Add(jen.Op("*").Id("d").Dot(strcase.ToCamel(name)).Op("=").Id("GetTopic").Params(jen.Id("d"), jen.Lit(name)))
								if IsCommand(name, d) {
									g.Add(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/store", "TopicStore").Index(
										jen.Op("*").Id("d").Dot(strcase.ToCamel(name)),
									).Op("=").Id("&d").Dot(lName + "Func"))
								}
							},
						),
					)
				}
			}
		})

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("SetMQTTFields").Params(
			jen.Id("fields").Id("MQTTFields"),
		).BlockFunc(
			func(g *jen.Group) {
				g.Add(
					jen.Op("*").Id("d").Dot("MQTT").Op("=").Id("fields"),
				)
			},
		)

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("GetMQTTFields").Params().Params(
			jen.Id("fields").Id("MQTTFields"),
		).BlockFunc(
			func(g *jen.Group) {
				g.Add(
					jen.Return(jen.Op("*").Id("d").Dot("MQTT")),
				)
			},
		)

		// d.Translate() ExternalDevice.Light
		internal[d.Name].Func().Params(jen.Id("iDevice").Id(strcase.ToCamel(d.Name))).Id("Translate").Params().Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", strcase.ToCamel(d.Name)).BlockFunc(
			func(g *jen.Group) {
				g.Add(
					jen.Id("eDevice").Op(":=").Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", strcase.ToCamel(d.Name)).Block(),
				)

				g.Add(
					jen.Id("eDevice").Dot("MQTT").Op("=").New(jen.Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", "MQTTFields")),
				)

				g.Add(
					jen.If(jen.Id("iDevice").Dot("MQTT").Dot("ForceUpdate").Op("!=").Nil()).Block(
						jen.Id("eDevice").Dot("MQTT").Dot("ForceUpdate").Op("=").Id("iDevice").Dot("MQTT").Dot("ForceUpdate"),
					),
				)
				g.Add(
					jen.If(jen.Id("iDevice").Dot("MQTT").Dot("UpdateInterval").Op("!=").Nil()).Block(
						jen.Id("eDevice").Dot("MQTT").Dot("UpdateInterval").Op("=").Id("iDevice").Dot("MQTT").Dot("UpdateInterval"),
					),
				)

				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) {
						if !strings.HasSuffix(key, "_topic") {
							cam := strcase.ToCamel(key)
							g.Add(
								jen.If(jen.Id("iDevice").Dot(cam).Op("!=").Nil()).Block(
									jen.Id("eDevice").Dot(cam).Op("=").Id("iDevice").Dot(cam),
								),
							)
						} else {
							lName := strcase.ToCamel(strings.TrimSuffix(key, "_topic") + "_func")
							g.Add(
								jen.If(jen.Id("iDevice").Dot(strcase.ToCamel(strings.TrimSuffix(key, "_topic"))).Op("!=").Nil()).Block(
									jen.Id("eDevice").Dot(lName).Op("=").Qual("github.com/W-Floyd/ha-mqtt-iot/devices/common", "Construct"+func() string {
										if IsCommand(key, d) {
											return "Command"
										} else {
											return "State"
										}
									}()+"Func").Params(jen.Op("*").Id("iDevice").Dot(strcase.ToCamel(strings.TrimSuffix(key, "_topic")))),
								),
							)
						}
					}
				}

				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) {
						if key == "availability_topic" {
							g.Add(
								jen.If(
									jen.Id("iDevice").Dot("Availability").Op("==").Nil(),
								).Block(
									jen.Id("eDevice").Dot("AvailabilityFunc").Op("=").Qual("github.com/W-Floyd/ha-mqtt-iot/devices/common", "AvailabilityFunc"),
								),
							)
						}
					}
				}

				g.Add(
					jen.Id("eDevice").Dot("Initialize").Params(),
				)

				g.Add(
					jen.Return(
						jen.Id("eDevice"),
					),
				)
			},
		)

		internal[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) {
						if strings.HasSuffix(key, "_topic") {
							lName := strcase.ToCamel(strings.TrimSuffix(key, "_topic"))
							g.Add(
								jen.Id(lName).Op("*").Params(jen.Index().String()).Tag(map[string]string{"json": strings.TrimSuffix(key, "_topic") + ",omitempty"}),
							)
						} else {
							g.Add(
								d.FieldAdder(key),
							)
						}
					}
				}

				g.Add(
					jen.Id("MQTT").Struct(
						jen.Id("UpdateInterval").Op("*").Float64().Tag(map[string]string{"json": "update_interval,omitempty"}),
						jen.Id("ForceUpdate").Op("*").Bool().Tag(map[string]string{"json": "force_update,omitempty"}),
					).Tag(map[string]string{"json": "mqtt"}),
				)
			},
		)

	}

	////////////////////////////////////////////////////////////////////////////////

	config := jen.NewFilePathName("./config/config.go", "config")
	config.ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
	config.ImportAlias("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", "ExternalDevice")
	config.Comment("////////////////////////////////////////////////////////////////////////////////")
	config.Comment("Do not modify this file, it is automatically generated")
	config.Comment("////////////////////////////////////////////////////////////////////////////////")

	config.Type().Id("Config").StructFunc(
		func(g *jen.Group) {
			g.Add(
				jen.Id("MQTT").StructFunc(
					func(h *jen.Group) {
						p := []string{"broker", "username", "password", "node_id", "instance_name"}
						for _, n := range p {
							h.Add(jen.Id(strcase.ToCamel(n)).String().Tag(map[string]string{"json": n + ",omitempty"}))
						}
					},
				),
			)
			g.Add(
				jen.Id("Logging").StructFunc(
					func(h *jen.Group) {
						p := []string{"critical", "debug", "error", "warn", "mqtt"}
						for _, n := range p {
							h.Add(jen.Id(strcase.ToCamel(n)).Bool().Tag(map[string]string{"json": n + ",omitempty"}))
						}
					},
				),
			)
			for _, d := range devices {
				g.Add(
					jen.Id(strcase.ToCamel(d.Name)).Index().Qual("github.com/W-Floyd/ha-mqtt-iot/devices/internaldevice", strcase.ToCamel(d.Name)).Tag(map[string]string{"json": d.Name + ",omitempty"}),
				)
			}
		},
	)

	config.Func().Params(
		jen.Id("c").Id("Config"),
	).Id("Translate").Params().Params(
		jen.Id("output").Index().Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", "Device"),
	).BlockFunc(
		func(g *jen.Group) {
			// g.Add(
			// 	jen.Id("output").Op(":=").Index().Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", "Device").Values(),
			// )
			for _, d := range devices {
				g.Add(
					jen.For(
						jen.List(
							jen.Id("_"),
							jen.Id("d"),
						),
					).Op(":=").Range().Id("c").Dot(strcase.ToCamel(d.Name)).Block(
						jen.Id("new"+strcase.ToCamel(d.Name)).Op(":=").Id("d").Dot("Translate").Params(),
						jen.Id("newDevice").Op(":=").Qual("github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice", "Device").Params(jen.Op("&").Id("new"+strcase.ToCamel(d.Name))),
						jen.Id("output").Op("=").Append(jen.Id("output"), jen.Id("newDevice")),
					),
					// jen.Id(strcase.ToCamel(d.Name)).Index().Qual("github.com/W-Floyd/ha-mqtt-iot/devices/internaldevice", strcase.ToCamel(d.Name)),
				)
			}
			g.Add(
				jen.Return(),
			)
		},
	)

	////////////////////////////////////////////////////////////////////////////////

	for k, v := range external {
		v.Save("./devices/externaldevice/" + k + ".go")
	}

	for k, v := range internal {
		v.Save("./devices/internaldevice/" + k + ".go")
	}

	config.Save("./config/config.go")

}
