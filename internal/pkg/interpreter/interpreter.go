package interpreter

import (
	"fmt"

	"mvmo.dev/sickvm/internal/pkg/instructions"
	"mvmo.dev/sickvm/internal/pkg/types"
)

type Interpreter struct {
	Instructions []instructions.Instruction
	Labels       *map[string]int
}

func NewInterpreter(instructions []instructions.Instruction, Labels *map[string]int) Interpreter {
	interpreter := new(Interpreter)
	interpreter.Instructions = instructions
	interpreter.Labels = Labels

	return *interpreter
}

func (interpreter Interpreter) Run() {
	var stack Stack
	var storage map[string]types.SickType = make(map[string]types.SickType)

	for i := 0; i < len(interpreter.Instructions); i++ {
		instruction := interpreter.Instructions[i]

		switch instruction.OpCode {
		case instructions.INS_IPUSH:
		case instructions.INS_SPUSH:
		case instructions.INS_BPUSH:
			stack.Push(types.AnyToSickType(instruction.Params[0]))
			continue
		case instructions.INS_ADD:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val1.AsInt() + val2.AsInt())
			continue
		case instructions.INS_SUB:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsInt() - val1.AsInt())
			continue
		case instructions.INS_MUL:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsInt() * val1.AsInt())
			continue
		case instructions.INS_DIV:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsInt() / val1.AsInt())
			continue
		case instructions.INS_MOD:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsInt() % val1.AsInt())
			continue
		case instructions.INS_CMP:
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val1 == val2)
			continue
		case instructions.INS_LT:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsFloat() < val1.AsFloat())
			continue
		case instructions.INS_GT:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsFloat() > val1.AsFloat())
			continue
		case instructions.INS_LTE:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsFloat() <= val1.AsFloat())
			continue
		case instructions.INS_GTE:
			val1 := stack.Pop().(types.SickNum)
			val2 := stack.Pop().(types.SickNum)
			stack.Push(val2.AsFloat() >= val1.AsFloat())
			continue
		case instructions.INS_STORE:
			identifier := instruction.Params[0].(string)
			toStore := stack.Pop()
			storage[identifier] = toStore
			continue
		case instructions.INS_LOAD:
			identifier := instruction.Params[0].(string)
			toPush := storage[identifier]
			stack.Push(toPush)
			continue
		case instructions.INS_DEL:
			identifier := instruction.Params[0].(string)
			delete(storage, identifier)
			continue
		case instructions.INS_JMP:
			whereToJump := instruction.Params[0].(int)
			i = whereToJump - 1
			continue
		case instructions.INS_CJMP: // first param is where to jump if true and second where to jump if false
			condition := stack.Pop().(types.SickBool)
			var whereToJump int
			if condition.Value {
				whereToJump = instruction.Params[0].(int) - 1
			} else {
				whereToJump = instruction.Params[1].(int) - 1
			}

			i = whereToJump
			continue
		case instructions.INS_DUP:
			head := stack.Pop()
			stack.Push(head)
			stack.Push(head)
			continue
		case instructions.INS_PRINT:
			head := stack.Pop()
			fmt.Print(head.ToHuman())
			continue
		case instructions.INS_PRINTLN:
			head := stack.Pop()
			fmt.Println(head.ToHuman())
			continue
		case instructions.INS_GOTO:
			labelName := instruction.Params[0].(string)
			i = (*interpreter.Labels)[labelName] - 1
			continue
		case instructions.INS_DUMP:
			fmt.Printf("=== Stack Dump ===\n")
			for i := len(stack); i > 0; i-- {
				var anno string
				if len(stack) == i {
					anno = "   <-- head"
				}
				fmt.Printf("%v: %v%v\n", i, stack[i-1], anno)
			}
			fmt.Printf("==================\n")
			continue
		case instructions.INS_VOID:
			continue
		default:
			fmt.Printf("Interpreter: No handling for instruction: %v\n", instruction.OpCode)
			continue
		}
	}
}
