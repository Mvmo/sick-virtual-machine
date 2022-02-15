package main

import (
	"fmt"

	"mvmo.dev/sickvm/internal/pkg/interpreter"
	"mvmo.dev/sickvm/internal/pkg/parser"
)

func main() {
	parser := parser.NewParser()
	instructions, err := parser.Parse("push 200\npush 200\ncmp\ncjmp 4 7\npush 1337\ndump\njmp 9\npush 1889\ndump\npush 1000\ndump\npush 200\npush 200\njmp 14\n")
	if err != nil {
		fmt.Println(err)
	}
	interpreter := interpreter.NewInterpreter(instructions)
	interpreter.Run()
}
