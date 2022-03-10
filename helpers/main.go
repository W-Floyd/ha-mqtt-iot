package main

import (
	"log"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {

	devices := DevicesInit()

	devicetypesfile := jen.NewFilePathName("../hadiscovery/devicetypes.go", "hadiscovery")

	topicsfile := jen.NewFilePathName("../hadiscovery/topics.go", "hadiscovery")

	u := jen.Block()

	for _, d := range devices {

		topicsfile.Func().Params(
			jen.Id("d").Id("Device" + strcase.ToCamel(d.Name)),
		).Id("GetTopicPrefix").Params().String().Block(
			jen.Return(jen.Id("NodeID").Op("+").Lit("/" + d.Name + "/").Op("+").Id("d.UniqueID").Op("+").Lit("/")),
		)

		topicsfile.Func().Params(
			jen.Id("d").Id("Device" + strcase.ToCamel(d.Name)),
		).Id("GetDiscoveryTopic").Params().String().Block(
			jen.Return(jen.Id("DiscoveryPrefix").Op("+").Lit("/" + d.Name + "/").Op("+").Id("NodeID").Op("+").Lit("/").Op("+").Id("d.UniqueID").Op("+").Lit("/")),
		)

		devicetypesfile.Type().Id("Device" + strcase.ToCamel(d.Name)).Struct(
			jen.Id("UniqueID").String(),
		)

		u = u.Union(jen.Int())

	}

	devicetypesfile.Type().Id("Device").Interface(
		u,
	)

	log.Println("Writing topics")

	topicsfile.Save("../hadiscovery/topics.go")

	devicetypesfile.Save("../hadiscovery/devicetypes.go")

}
