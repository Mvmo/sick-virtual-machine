package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"mvmo.dev/sickvm/internal/pkg/interpreter"
	"mvmo.dev/sickvm/internal/pkg/parser"
)

var (
	inputFile *string
)

func init() {
	inputFile = flag.String("file", "undefined", "file to be interpreted")
}

func main() {
	flag.Parse()

	if strings.ToLower(*inputFile) == "undefined" {
		log.Fatalf("you need to provide a file to be interpreted")
		return
	}

	content, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("unable to read file: %v\n", err)
		return
	}

	parser := parser.NewParser()
	instructions, err := parser.Parse(string(content))
	if err != nil {
		fmt.Println(err)
	}
	interpreter := interpreter.NewInterpreter(instructions)
	interpreter.Run()
}
