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

func (parser Parser) Parse(input string) ([]instructions.Instruction, *map[string]int, error) {
	var returnValue []instructions.Instruction
	var labels map[string]int = make(map[string]int)

	rawLines := strings.Split(input, "\n")

	var lines []string
	for _, line := range rawLines {
		if strings.HasPrefix(line, ";") {
			continue
		}
		lines = append(lines, line)
	}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		opname := fields[0]

		if strings.HasPrefix(opname, ";") {
			continue
		}

		if strings.HasSuffix(opname, ":") {
			opname = strings.TrimSuffix(opname, ":")
			labels[opname] = i

			instruction := new(instructions.Instruction)
			instruction.OpCode = instructions.INS_VOID
			instruction.Params = nil

			returnValue = append(returnValue, *instruction)

			continue
		}

		var opcode int
		switch opname {
		case "ipush":
			opcode = instructions.INS_IPUSH
		case "bpush":
			opcode = instructions.INS_BPUSH
		case "spush":
			opcode = instructions.INS_SPUSH
		case "add":
			opcode = instructions.INS_ADD
		case "sub":
			opcode = instructions.INS_SUB
		case "mul":
			opcode = instructions.INS_MUL
		case "div":
			opcode = instructions.INS_DIV
		case "mod":
			opcode = instructions.INS_MOD
		case "cmp":
			opcode = instructions.INS_CMP
		case "lt":
			opcode = instructions.INS_LT
		case "lte":
			opcode = instructions.INS_LTE
		case "gt":
			opcode = instructions.INS_GT
		case "gte":
			opcode = instructions.INS_GTE
		case "req":
			opcode = instructions.INS_REQ
		case "store":
			opcode = instructions.INS_STORE
		case "load":
			opcode = instructions.INS_LOAD
		case "del":
			opcode = instructions.INS_DEL
		case "dump":
			opcode = instructions.INS_DUMP
		case "jmp":
			opcode = instructions.INS_JMP
		case "cjmp":
			opcode = instructions.INS_CJMP
		case "swap":
			opcode = instructions.INS_SWAP
		case "dup":
			opcode = instructions.INS_DUP
		case "print":
			opcode = instructions.INS_PRINT
		case "println":
			opcode = instructions.INS_PRINTLN
		case "goto":
			opcode = instructions.INS_GOTO
		case "call":
			opcode = instructions.INS_CALL
		default:
			fmt.Printf("Parser: No instruction parsing for %v\n", opname)
			syscall.Exit(-1)
		}
		parseFunctions := parser.paramsParseFunctionsMap[opcode]

		var params []string = make([]string, 0)
		fieldsSlice := fields[1:]
		for j := 0; j < len(fieldsSlice); j++ {
			param := fieldsSlice[j]
			if !strings.HasPrefix(param, "\"") || strings.HasSuffix(param, "\"") {
				params = append(params, param)
				continue
			}

			for k := j + 1; k < len(fieldsSlice); k++ {
				nextParam := fieldsSlice[k]
				param += " " + nextParam
				if strings.HasSuffix(nextParam, "\"") {
					params = append(params, param)
					j = k
					break
				}
			}
		}

		if len(params) != len(parseFunctions) {
			return nil, nil, fmt.Errorf("Parser: Required %v parameters for %v instruction and got %v", len(parseFunctions), opname, len(params))
		}

		var parsedParams = make([]interface{}, len(parseFunctions))
		for i, parseParam := range parseFunctions {
			unparsedParam := params[i]
			switch parseParam := parseParam.(type) {
			case func(string) int:
				parsedParams[i] = parseParam(unparsedParam)
			case func(string) string:
				parsedParams[i] = parseParam(unparsedParam)
			case func(string) bool:
				parsedParams[i] = parseParam(unparsedParam)
			}
		}

		instruction := new(instructions.Instruction)
		instruction.OpCode = opcode
		instruction.Params = parsedParams

		returnValue = append(returnValue, *instruction)
	}

	return returnValue, &labels, nil
}

func NewParser() *Parser {
	parser := new(Parser)
	parser.paramsParseFunctionsMap = map[int][]interface{}{
		instructions.INS_IPUSH: {
			parseIntParam,
		},
		instructions.INS_SPUSH: {
			parseStringParam,
		},
		instructions.INS_BPUSH: {
			parseBoolParam,
		},
		instructions.INS_ADD: {},
		instructions.INS_SUB: {},
		instructions.INS_MUL: {},
		instructions.INS_DIV: {},
		instructions.INS_MOD: {},
		instructions.INS_CMP: {},
		instructions.INS_LT:  {},
		instructions.INS_GT:  {},
		instructions.INS_LTE: {},
		instructions.INS_GTE: {},
		instructions.INS_REQ: {
			parseIdentifierParam,
		},
		instructions.INS_STORE: {
			parseIdentifierParam,
		},
		instructions.INS_LOAD: {
			parseIdentifierParam,
		},
		instructions.INS_DEL: {
			parseIdentifierParam,
		},
		instructions.INS_JMP: {
			parseIntParam,
		},
		instructions.INS_CJMP: {
			parseIntParam,
			parseIntParam,
		},
		instructions.INS_SWAP: {},
		instructions.INS_DUP:  {},
		instructions.INS_CALL: {
			parseIdentifierParam,
		},
		instructions.INS_GOTO: {
			parseIdentifierParam,
		},
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

func parseStringParam(str string) string {
	if !strings.HasPrefix(str, "\"") {
		syscall.Exit(-1)
	}

	if len(str) == 0 || str == " " {
		syscall.Exit(-1)
	}

	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}

func parseIdentifierParam(str string) string {
	if len(str) == 0 || str == " " {
		syscall.Exit(-1)
	}

	return str
}

func parseBoolParam(str string) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		syscall.Exit(-1)
	}

	return val
}
