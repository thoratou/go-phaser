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

	//adapt data
	//additional import
	additionalImports := map[string]string{
		"dom": "honnef.co/go/js/dom",
	}

	additionalPackages := map[string]string{
		"HTMLCanvasElement":        "dom",
		"CanvasRenderingContext2D": "dom",
		"HTMLVideoElement":         "dom",
	}

	wrappers := map[string]string{
		"WebGLContext":                 "WrapWebGLContext",
		"DOMElement":                   "WrapDOMElement",
		"dom.HTMLCanvasElement":        "WrapHTMLCanvasElement",
		"dom.HTMLVideoElement":         "WrapHTMLVideoElement",
		"dom.CanvasRenderingContext2D": "WrapCanvasRenderingContext2D",
	}

	classes := []Class{}
	for _, class := range root.Classes {
		//remove duplicates
		class.Members = RemoveDuplicateMembers(class.Members, class.Name)
		class.Functions = RemoveDuplicateFunctions(class.Functions, class.Name)

		//additional imports
		class.Imports = map[string]string{}
		class.Members = AddMemberImports(&class, class.Members, additionalPackages, additionalImports)
		class.Functions = AddFunctionImports(&class, class.Functions, additionalPackages, additionalImports)

		//add type wrappers
		class.Members = AddMemberWrappers(class.Members, wrappers)
		class.Functions = AddFunctionWrappers(class.Functions, wrappers)

		//generate various parameter handling functions
		class.Functions = AddParameteredFunctions(class.Functions)

		classes = append(classes, class)
	}
	root.Classes = classes

	//write Go files
	t := template.New("GoFileGen.tmpl")
	t, e = t.ParseFiles("./GoFileGen.tmpl")
	if e != nil {
		fmt.Printf("Read template error: %v\n", e)
		os.Exit(1)
	}

	repositoryPath := os.Getenv("GOPATH") + "/src/" + parsedArgs.Repository + "/generated/"
	fmt.Println("Output path: " + repositoryPath)

	e = os.RemoveAll(repositoryPath)
	if e != nil {
		fmt.Printf("Remove directory error: %v\n", e)
		os.Exit(1)
	}
	fmt.Println("Cleaned-up ouput path: " + repositoryPath)

	workaroundPath := os.Getenv("GOPATH") + "/src/" + parsedArgs.Repository + "/workaround/"
	e = CopyDir(workaroundPath, repositoryPath)
	if e != nil {
		fmt.Printf("Copy directory error: %v\n", e)
		os.Exit(1)
	}
	fmt.Println("Workaround files copied: " + workaroundPath)

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
	fmt.Println("All files generated")
}
