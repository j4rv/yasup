package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type tmplVars struct {
	Type         string
	TypeCased    string
	SingleVal    string
	MultipleVals string
}

func main() {
	var typeStr string
	var singleVal string
	var multipleVals string
	var addTest bool
	flag.StringVar(&typeStr, "type", "string", "The string identifier of the type. Example: 'int' or 'Reader'")
	flag.StringVar(&singleVal, "singleVal", "\"foobar\"", "A single value for the type. Will be used in tests.")
	flag.StringVar(&multipleVals, "multipleVals", "\"0\", \"1\", \"2\", \"3\", \"4\", \"5\", \"6\", \"7\", \"8\", \"9\"",
		"A comma separated list of ~ten different values for the type. Will be used in tests.")
	flag.BoolVar(&addTest, "test", false, "Adds automatic tests. A one item list and ten item lists will need to be provided using their corresponding flags.")
	flag.Parse()

	fmt.Println("Building slices file for type", typeStr)
	typeToBuild := tmplVars{typeStr, strings.Title(typeStr), singleVal, multipleVals}

	srcTmpl := template.Must(template.ParseFiles("template.tmpl"))
	execTmpl(srcTmpl, typeToBuild.Type+"Slices.go", typeToBuild)

	if addTest {
		testTmpl := template.Must(template.ParseFiles("test_template.tmpl"))
		execTmpl(testTmpl, typeToBuild.Type+"Slices_test.go", typeToBuild)
	}
}

func execTmpl(t *template.Template, outputName string, typeToBuild interface{}) {
	f, err := os.Create("../" + outputName)
	defer f.Close()

	if err != nil {
		panic(fmt.Sprint("create file: ", err))
	}

	err = t.Execute(f, typeToBuild)
	if err != nil {
		panic(fmt.Sprint("execute: ", err))
	}
}
