package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Input      string `cli:"*i,input" usage:"Input JSON file path for phase.io generated JsDoc data"`
	Repository string `cli:"*r,repository" usage:"Repository root for generated data"`
}

func main() {

	//parse arguments
	parsedArgs := new(argT)
	cli.Run(parsedArgs, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("input=%v\n", argv.Input)
		return nil
	})

	//open JsDoc Json file
	file, e := os.Open(parsedArgs.Input)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	root, e := Decode(file)
	if e != nil {
		fmt.Printf("Decode error: %v\n", e)
		os.Exit(1)
	}

	//debug
	//fmt.Println(root)

	//todo adapt data

	//write Go files
	t := template.New("GoFileGen.tmpl")
	t, e = t.ParseFiles("./GoFileGen.tmpl")
	if e != nil {
		fmt.Printf("Read template error: %v\n", e)
		os.Exit(1)
	}

	repositoryPath := os.Getenv("GOPATH") + "/src/" + parsedArgs.Repository + "/generated/"
	fmt.Println("Output path:" + repositoryPath)

	for _, class := range root.Classes {
		path := repositoryPath + strings.Replace(class.GetNamespace(), ".", "/", -1)
		e = os.MkdirAll(path, 0755)
		if e != nil {
			fmt.Printf("Create directories error: %v\n", e)
			os.Exit(1)
		}

		goFilePath := path + "/" + class.GetNameNoNamespace() + ".go"
		file, e := os.Create(goFilePath)
		if e != nil {
			fmt.Printf("Create file error: %v\n", e)
			os.Exit(1)
		}

		e = t.Execute(file, &class)
		if e != nil {
			fmt.Printf("Write from template error: %v\n", e)
			os.Exit(1)
		}

		e = file.Close()
		if e != nil {
			fmt.Printf("Close error: %v\n", e)
			os.Exit(1)
		}
	}
}
