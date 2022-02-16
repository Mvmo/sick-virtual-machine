package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"mvmo.dev/sickvm/internal/pkg/interpreter"
	"mvmo.dev/sickvm/internal/pkg/parser"
)

var (
	inputFile *string
	timing    *bool
)

func init() {
	inputFile = flag.String("file", "undefined", "file to be interpreted")
	timing = flag.Bool("timing", false, "enable timing for compiler")
}

func main() {
	flag.Parse()

	if strings.ToLower(*inputFile) == "undefined" {
		log.Fatalf("you need to provide a file to be interpreted")
		return
	}

	startTime := time.Now()

	content, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("unable to read file: %v\n", err)
		return
	}

	parser := parser.NewParser()
	instructions, labels, err := parser.Parse(string(content))
	if err != nil {
		fmt.Println(err)
	}

	interpreter := interpreter.NewInterpreter(instructions, labels)
	interpreter.Run()

	elapsed := time.Since(startTime)

	if *timing {
		log.Printf("took: %v", elapsed)
	}
}
