package parser

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"

	"mvmo.dev/sickvm/internal/pkg/instructions"
)

type Parser struct {
	paramsParseFunctionsMap map[int][]interface{}
}

func (self Parser) Parse(input string) ([]instructions.Instruction, error) {
	var returnValue []instructions.Instruction

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		opname := fields[0]
		params := fields[1:]

		var opcode int
		switch opname {
		case "push":
			opcode = instructions.INS_PUSH
			break
		case "add":
			opcode = instructions.INS_ADD
			break
		case "sub":
			opcode = instructions.INS_SUB
			break
		case "cmp":
			opcode = instructions.INS_CMP
			break
		case "lt":
			opcode = instructions.INS_LT
		case "gt":
			opcode = instructions.INS_GT
		case "dump":
			opcode = instructions.INS_DUMP
			break
		case "jmp":
			opcode = instructions.INS_JMP
			break
		case "cjmp":
			opcode = instructions.INS_CJMP
			break
		case "dup":
			opcode = instructions.INS_DUP
			break
		default:
			fmt.Printf("No instruction parsing for %v\n", opname)
			syscall.Exit(-1)
		}

		parseFunctions := self.paramsParseFunctionsMap[opcode]

		if len(params) != len(parseFunctions) {
			return nil, fmt.Errorf("Required %v parameters for %v instruction and got %v", len(parseFunctions), opname, len(params))
		}

		var parsedParams = make([]interface{}, len(parseFunctions))
		for i, parseParam := range parseFunctions {
			unparsedParam := params[i]
			parsedParams[i] = parseParam.(func(string) int)(unparsedParam)
		}

		instruction := new(instructions.Instruction)
		instruction.OpCode = opcode
		instruction.Params = parsedParams

		returnValue = append(returnValue, *instruction)
	}

	return returnValue, nil
}

func NewParser() *Parser {
	parser := new(Parser)
	parser.paramsParseFunctionsMap = map[int][]interface{}{
		instructions.INS_PUSH: {
			parseIntParam,
		},
		instructions.INS_ADD: {},
		instructions.INS_SUB: {},
		instructions.INS_CMP: {},
		instructions.INS_LT:  {},
		instructions.INS_GT:  {},
		instructions.INS_LTE: {},
		instructions.INS_GTE: {},
		instructions.INS_JMP: {
			parseIntParam,
		},
		instructions.INS_CJMP: {
			parseIntParam,
			parseIntParam,
		},
		instructions.INS_DUP:  {},
		instructions.INS_DUMP: {},
	}

	return parser
}

func parseIntParam(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		syscall.Exit(-1)
	}

	return val
}
