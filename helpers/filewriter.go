package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

type Function struct {
	Name             string
	OperatesUponType string
	Returns          string
	Contents         []string
}

type GoFile struct {
	Name      string
	heading   []string
	Functions []Function
	footing   []string
}

func (f *GoFile) Write() error {
	f.heading = []string{"package devices"}

	f.heading = append(f.heading, "type HAEntity interface {",
		"GetDiscoveryTopic() string",
		"GetAvailabilityTopic() string",
		"}",
		"",
	)

	// f.Contents = append(f.Contents, "")
	// f.footing=append(f.footing, "")

	lines := []string{}
	lines = append(lines, f.heading...)

	for _, fn := range f.Functions {
		lines = append(lines, fn.convert()...)
	}

	lines = append(lines, f.footing...)

	file, err := os.Create(f.path())
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()

}

func (f *GoFile) path() string {
	return "../devices/" + f.Name + ".go"
}

func (fn Function) convert() (retval []string) {

	objectName := "obj"
	returnValueName := "retval"

	l := "func"

	if fn.OperatesUponType != "" {
		l = l + " (" + objectName + " " + fn.OperatesUponType + ")"
	}

	l = l + " " + strcase.ToCamel(fn.Name)

	if fn.Returns != "" {
		l = l + " (" + returnValueName + " " + fn.Returns + " )"
	}

	l = l + " {"

	retval = append(retval, l)
	retval = append(retval, fn.Contents...)
	retval = append(retval, "}", "")

	return
}
