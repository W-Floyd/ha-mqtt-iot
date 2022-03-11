package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

var keyNames = []string{
	"action_template",
	"action_topic",
	"automation_type",
	"aux_command_topic",
	"aux_state_template",
	"aux_state_topic",
	"availability_mode",
	"availability_template",
	"availability_topic",
	"battery_level_template",
	"battery_level_topic",
	"brightness_command_template",
	"brightness_command_topic",
	"brightness_scale",
	"brightness_state_topic",
	"brightness_value_template",
	"charging_template",
	"charging_topic",
	"cleaning_template",
	"cleaning_topic",
	"code",
	"code_arm_required",
	"code_disarm_required",
	"code_trigger_required",
	"color_mode_state_topic",
	"color_mode_value_template",
	"color_temp_command_template",
	"color_temp_command_topic",
	"color_temp_state_topic",
	"color_temp_value_template",
	"command_template",
	"command_topic",
	"current_temperature_template",
	"current_temperature_topic",
	"device_class",
	"devices",
	"docked_template",
	"docked_topic",
	"effect_command_template",
	"effect_command_topic",
	"effect_list",
	"effect_state_topic",
	"effect_value_template",
	"enabled_by_default",
	"encoding",
	"entity_category",
	"error_template",
	"error_topic",
	"expire_after",
	"fan_mode_command_template",
	"fan_mode_command_topic",
	"fan_modes",
	"fan_mode_state_template",
	"fan_mode_state_topic",
	"fan_speed_list",
	"fan_speed_template",
	"fan_speed_topic",
	"force_update",
	"hs_value_template",
	"icon",
	"initial",
	"json_attributes_template",
	"json_attributes_topic",
	"last_reset_value_template",
	"max",
	"max_humidity",
	"max_mireds",
	"max_temp",
	"min",
	"min_humidity",
	"min_mireds",
	"min_temp",
	"mode_command_template",
	"mode_command_topic",
	"modes",
	"mode_state_template",
	"mode_state_topic",
	"name",
	"object_id",
	"off_delay",
	"on_command_type",
	"optimistic",
	"options",
	"oscillation_command_template",
	"oscillation_command_topic",
	"oscillation_state_topic",
	"oscillation_value_template",
	"payload",
	"payload_arm_away",
	"payload_arm_custom_bypass",
	"payload_arm_home",
	"payload_arm_night",
	"payload_arm_vacation",
	"payload_available",
	"payload_clean_spot",
	"payload_close",
	"payload_disarm",
	"payload_home",
	"payload_locate",
	"payload_lock",
	"payload_not_available",
	"payload_not_home",
	"payload_off",
	"payload_on",
	"payload_open",
	"payload_oscillation_off",
	"payload_oscillation_on",
	"payload_reset",
	"payload_reset_humidity",
	"payload_reset_mode",
	"payload_reset_percentage",
	"payload_reset_preset_mode",
	"payload_return_to_base",
	"payload_start_pause",
	"payload_stop",
	"payload_trigger",
	"payload_turn_off",
	"payload_turn_on",
	"payload_unlock",
	"percentage_command_template",
	"percentage_command_topic",
	"percentage_state_topic",
	"percentage_value_template",
	"position_closed",
	"position_open",
	"position_template",
	"position_topic",
	"power_command_topic",
	"precision",
	"preset_mode_command_template",
	"preset_mode_command_topic",
	"preset_modes",
	"preset_mode_state_topic",
	"preset_mode_value_template",
	"qos",
	"retain",
	"rgb_command_template",
	"rgb_command_topic",
	"rgb_state_topic",
	"rgb_value_template",
	"schema",
	"send_command_topic",
	"set_fan_speed_topic",
	"set_position_template",
	"set_position_topic",
	"source_type",
	"speed_range_max",
	"speed_range_min",
	"state_class",
	"state_closed",
	"state_closing",
	"state_locked",
	"state_off",
	"state_on",
	"state_open",
	"state_opening",
	"state_stopped",
	"state_topic",
	"state_unlocked",
	"state_value_template",
	"step",
	"subtype",
	"supported_features",
	"swing_mode_command_template",
	"swing_mode_command_topic",
	"swing_modes",
	"swing_mode_state_template",
	"swing_mode_state_topic",
	"target_humidity_command_template",
	"target_humidity_command_topic",
	"target_humidity_state_template",
	"target_humidity_state_topic",
	"temperature_command_template",
	"temperature_command_topic",
	"temperature_high_command_template",
	"temperature_high_command_topic",
	"temperature_high_state_template",
	"temperature_high_state_topic",
	"temperature_low_command_template",
	"temperature_low_command_topic",
	"temperature_low_state_template",
	"temperature_low_state_topic",
	"temperature_state_template",
	"temperature_state_topic",
	"temperature_unit",
	"temp_step",
	"tilt_closed_value",
	"tilt_command_template",
	"tilt_command_topic",
	"tilt_max",
	"tilt_min",
	"tilt_opened_value",
	"tilt_optimistic",
	"tilt_status_template",
	"tilt_status_topic",
	"topic",
	"type",
	"unique_id",
	"unit_of_measurement",
	"value_template",
	"white_command_topic",
	"white_scale",
	"xy_command_topic",
	"xy_state_topic",
	"xy_value_template",
}

func main() {

	devices := DevicesInit()

	devicetypesfile := jen.NewFilePathName("../hadiscovery/devicetypes.go", "hadiscovery")

	deviceinitfile := jen.NewFilePathName("../hadiscovery/deviceinit.go", "hadiscovery")

	devicefunctionsfile := jen.NewFilePathName("../hadiscovery/devicefunctions.go", "hadiscovery")

	sort.Strings(keyNames)

	for _, d := range devices {

		// d.GetRawID()
		devicefunctionsfile.Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("GetRawId").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// d.GetUniqueID()
		if d.JSONContainer.Exists("unique_id") {
			devicefunctionsfile.Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Id("d.UniqueId")),
			)
		} else {
			devicefunctionsfile.Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Lit("")),
			)
		}

		st := make(map[string][]*jen.Statement)

		// Add standalone base level fields
		for _, key := range keyNames {
			if d.JSONContainer.Exists(key) {
				st[key] = append(st[key], d.FieldAdder(key))
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
			devicefunctionsfile.Func().Params(
				jen.Id("d").Id("*"+strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block(
				jen.Id("d.Device.Manufacturer").Op("=").Id("Manufacturer"),
				jen.Id("d.Device.Model").Op("=").Id("SoftwareName"),
				jen.Id("d.Device.Name").Op("=").Id("InstanceName"),
				jen.Id("d.Device.SwVersion").Op("=").Id("SWVersion"),
			)
		} else {
			devicefunctionsfile.Func().Params(
				jen.Id("d").Id("*" + strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params().Block()
		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)

		// Device MQTT Struct
		devicetypesfile.Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					for _, val := range v {
						g.Add(val)
					}
				}
				g.Id("RawId").String().Tag(map[string]string{"json": "-"})
			},
		)

		deviceinitfile.Func().Params(jen.Id("d").Id(strcase.ToCamel(d.Name))).Id("Init").Params().BlockFunc(func(g *jen.Group) {
			if d.JSONContainer.Exists("retain") {
				g.Add(jen.Id("d").Dot("Retain").Op("=").Lit(false))
			}
			g.Add(jen.Id("d").Dot("PopulateDevice").Params())
			g.Add(jen.Id("d").Dot("PopulateTopics").Params())
		})

		// d.PopulateTopics()
		deviceinitfile.Func().Params(
			jen.Id("d").Id("*" + strcase.ToCamel(d.Name)),
		).Id("PopulateTopics").Params().BlockFunc(func(g *jen.Group) {
			for _, name := range keyNames {
				if strings.HasSuffix(name, "_topic") && d.JSONContainer.Exists(name) {
					lName := strcase.ToCamel(strings.TrimSuffix(name, "_topic"))
					g.Add(
						jen.If(
							jen.Id("d").Dot(lName+"Func").Op("!=").Nil(),
						).Block(
							jen.Id("d").Dot(strcase.ToCamel(name)).Op("=").Id("GetTopic").Params(jen.Id("d"), jen.Lit(name)),
							jen.Id("topicStore").Index(
								jen.Id("d").Dot(strcase.ToCamel(name)),
							).Op("=").Id("&d").Dot(lName+"Func"),
							// topicStore[device.BrightnessCommandTopic] = &device.BrightnessCommandFunc
						),
					)
				}
			}
		})

	}

	devicetypesfile.Type().Id("Device").Interface(
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
	)

	devicetypesfile.Save("../hadiscovery/devicetypes.go")

	deviceinitfile.Save("../hadiscovery/deviceinit.go")

	devicefunctionsfile.Save("../hadiscovery/devicefunctions.go")

}
