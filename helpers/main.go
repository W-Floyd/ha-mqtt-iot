package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {

	devices := DevicesInit()

	devicetypesfile := jen.NewFilePathName("../hadiscovery/devicetypes.go", "hadiscovery")

	topicsfile := jen.NewFilePathName("../hadiscovery/topics.go", "hadiscovery")

	genericsfile := jen.NewFilePathName("../hadiscovery/generics.go", "hadiscovery")

	genericsfile.Type().Id("Device")

	genericsfile.Add()

	u := []*jen.Statement{}

	state := []*jen.Statement{}

	for _, d := range devices {

		// d.GetTopicPrefix()
		topicsfile.Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("GetTopicPrefix").Params().String().Block(
			jen.Return(jen.Id("NodeID").Op("+").Lit("/" + d.Name + "/").Op("+").Id("d.UniqueID()").Op("+").Lit("/")),
		)

		// d.GetDiscoveryTopic()
		topicsfile.Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("GetDiscoveryTopic").Params().String().Block(
			jen.Return(jen.Id("DiscoveryPrefix").Op("+").Lit("/" + d.Name + "/").Op("+").Id("NodeID").Op("+").Lit("/").Op("+").Id("d.UniqueID()").Op("+").Lit("/")),
		)

		// Device Struct
		devicetypesfile.Type().Id(strcase.ToCamel(d.Name)).Struct(
			jen.Id("uniqueID").String(),
			jen.Id("retain").Bool(),
		)

		// d.RawID()
		devicetypesfile.Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("RawID").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// d.UniqueID()
		devicetypesfile.Func().Params(
			jen.Id("d").Id(strcase.ToCamel(d.Name)),
		).Id("UniqueID").Params().String().Block(
			jen.Return(jen.Id("d.uniqueID")),
		)

		// Devices interface
		u = append(u, jen.Id(strcase.ToCamel(d.Name)))

		if d.JSONContainer.Exists("state_topic") {
			state = append(state, jen.Id(strcase.ToCamel(d.Name)))
		}

	}

	devicetypesfile.Type().Id("Device").Interface(
		jen.UnionFunc(
			func(g *jen.Group) {
				for _, v := range u {
					g.Add(v)
				}
			},
		),
		jen.Id("GetDiscoveryTopic").Params().String(),
		jen.Id("RawID").Params().String(),
		jen.Id("UniqueID").Params().String(),
	)

	devicetypesfile.Type().Id("WithState").Interface(
		jen.UnionFunc(
			func(g *jen.Group) {
				for _, v := range state {
					g.Add(v)
				}
			},
		),
	)

	topicsfile.Save("../hadiscovery/topics.go")

	devicetypesfile.Save("../hadiscovery/devicetypes.go")

	genericsfile.Save("../hadiscovery/generics.go")

}
