package interpreter

import (
	"fmt"

	"mvmo.dev/sickvm/internal/pkg/instructions"
)

type Interpreter struct {
	Instructions []instructions.Instruction
}

func NewInterpreter(instructions []instructions.Instruction) Interpreter {
	interpreter := new(Interpreter)
	interpreter.Instructions = instructions

	return *interpreter
}

func (self Interpreter) Run() {
	var stack Stack

	for i := 0; i < len(self.Instructions); i++ {
		instruction := self.Instructions[i]

		switch instruction.OpCode {
		case instructions.INS_PUSH:
			stack.Push(instruction.Params[0])
			break
		case instructions.INS_ADD:
			val1 := stack.Pop().(int)
			val2 := stack.Pop().(int)
			stack.Push(val1 + val2)
			break
		case instructions.INS_SUB:
			val1 := stack.Pop().(int)
			val2 := stack.Pop().(int)
			stack.Push(val1 - val2)
			break
		case instructions.INS_CMP:
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val1 == val2)
			break
		case instructions.INS_LT:
			val1 := stack.Pop().(int)
			val2 := stack.Pop().(int)
			stack.Push(val1 < val2)
			break
		case instructions.INS_GT:
			val1 := stack.Pop().(int)
			val2 := stack.Pop().(int)
			stack.Push(val1 > val2)
			break
		case instructions.INS_JMP:
			whereToJump := instruction.Params[0].(int)
			i = whereToJump - 1
			continue
		case instructions.INS_CJMP: // first param is where to jump if true and second where to jump if false
			condition := stack.Pop().(bool)
			var whereToJump int
			if condition {
				whereToJump = instruction.Params[0].(int) - 1
			} else {
				whereToJump = instruction.Params[1].(int) - 1
			}

			i = whereToJump
			continue
		case instructions.INS_DUP:
			val1 := stack.Pop()
			val2 := stack.Pop()

			for i := 0; i < 2; i++ {
				stack.Push(val1)
				stack.Push(val2)
			}
		case instructions.INS_DUMP:
			head := stack.Peek()
			fmt.Printf("%v\n", head)
			break
		default:
			fmt.Printf("Interpreter: No handling for instruction: %v\n", instruction.OpCode)
			break
		}
	}
}
