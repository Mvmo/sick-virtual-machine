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

func (interpreter Interpreter) Run() error {
	var objectStack SickObjectStack
	var referenceStack Stack
	var storage map[string]types.SickObject = make(map[string]types.SickObject)

	for i := 0; i < len(interpreter.Instructions); i++ {
		instruction := interpreter.Instructions[i]

		switch instruction.OpCode {
		case instructions.INS_IPUSH:
			objectStack.Push(types.AnyToSickObject(instruction.Params[0]))
			continue
		case instructions.INS_SPUSH:
			objectStack.Push(types.AnyToSickObject(instruction.Params[0]))
			continue
		case instructions.INS_BPUSH:
			objectStack.Push(types.AnyToSickObject(instruction.Params[0]))
			continue
		case instructions.INS_ADD:
			val1 := objectStack.Pop()
			val2 := objectStack.Pop()

			switch val1.(type) {
			case types.SickString:
				switch val2.(type) {
				case types.SickString:
					objectStack.Push(val2.(types.SickString).Value + val1.(types.SickString).Value)
				case types.SickInt:
					objectStack.Push(val2.(types.SickInt).ToHuman() + val1.(types.SickString).Value)
				case types.SickBool:
					objectStack.Push(val2.(types.SickBool).ToHuman() + val1.(types.SickString).Value)
				default:
					return fmt.Errorf("can't invoke Add-Instruction with %v(%v) and %v(%v)", val1.ToHuman(), val1.TypeName(), val2.ToHuman(), val2.TypeName())
				}
			case types.SickInt:
				if val1.TypeName() != val2.TypeName() {
					return fmt.Errorf("can't invoke Add-Instruction with %v(%v) and %v(%v)", val1.ToHuman(), val1.TypeName(), val2.ToHuman(), val2.TypeName())
				}
				objectStack.Push(val1.(types.SickInt).Value + val2.(types.SickInt).Value)
			default:
				return fmt.Errorf("can't invoke Add-Instruction with %vs", val1.TypeName())
			}

			continue
		case instructions.INS_SUB:
			val1 := objectStack.Pop()
			val2 := objectStack.Pop()

			switch val1 := val1.(type) {
			case types.SickInt:
				switch val2 := val2.(type) {
				case types.SickString:
					value := val2.Value
					objectStack.Push(value[:len(value)-val1.Value])
				case types.SickInt:
					objectStack.Push(val2.AsInt() - val1.AsInt())
				default:
					return fmt.Errorf("can't invoke Sub-Instruction with [%v(%v) - %v(%v)]", val2.ToHuman(), val2.TypeName(), val1.ToHuman(), val1.TypeName())
				}
			default:
				return fmt.Errorf("can't invoke Sub-Instruction with [%v(%v) - %v(%v)]", val2.ToHuman(), val2.TypeName(), val1.ToHuman(), val1.TypeName())
			}
			continue
		case instructions.INS_MUL:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsInt() * val1.AsInt())
			continue
		case instructions.INS_DIV:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsInt() / val1.AsInt())
			continue
		case instructions.INS_MOD:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsInt() % val1.AsInt())
			continue
		case instructions.INS_CMP:
			val1 := objectStack.Pop()
			val2 := objectStack.Pop()
			objectStack.Push(val1 == val2)
			continue
		case instructions.INS_LT:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsFloat() < val1.AsFloat())
			continue
		case instructions.INS_GT:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsFloat() > val1.AsFloat())
			continue
		case instructions.INS_LTE:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsFloat() <= val1.AsFloat())
			continue
		case instructions.INS_GTE:
			val1 := objectStack.Pop().(types.SickNum)
			val2 := objectStack.Pop().(types.SickNum)
			objectStack.Push(val2.AsFloat() >= val1.AsFloat())
			continue
		case instructions.INS_REQ:
			requiredType := instruction.Params[0].(string)
			typeOfSickObjectStackHead := objectStack.Peek().TypeName()

			if typeOfSickObjectStackHead != requiredType {
				return fmt.Errorf("required type %v and got %v", requiredType, typeOfSickObjectStackHead)
			}

			continue
		case instructions.INS_STORE:
			identifier := instruction.Params[0].(string)
			toStore := objectStack.Pop()
			storage[identifier] = toStore
			continue
		case instructions.INS_LOAD:
			identifier := instruction.Params[0].(string)
			toPush := storage[identifier]
			objectStack.Push(toPush)
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
			condition := objectStack.Pop().(types.SickBool)
			var whereToJump int
			if condition.Value {
				whereToJump = instruction.Params[0].(int) - 1
			} else {
				whereToJump = instruction.Params[1].(int) - 1
			}

			i = whereToJump
			continue
		case instructions.INS_SIZEOF:
			head := objectStack.Pop()
			switch head := head.(type) {
			case types.SickString:
				objectStack.Push(len(head.Value))
			default:
				return fmt.Errorf("can't use sizeof on %v", head.TypeName())
			}
			continue
		case instructions.INS_SWAP:
			a := objectStack.Pop()
			b := objectStack.Pop()
			objectStack.Push(b)
			objectStack.Push(a)
		case instructions.INS_DUP:
			head := objectStack.Pop()
			objectStack.Push(head)
			objectStack.Push(head)
			continue
		case instructions.INS_DROP:
			objectStack.Pop()
			continue
		case instructions.INS_PRINT:
			head := objectStack.Pop()
			fmt.Print(head.ToHuman())
			continue
		case instructions.INS_PRINTLN:
			head := objectStack.Pop()
			fmt.Println(head.ToHuman())
			continue
		case instructions.INS_CALL:
			labelName := instruction.Params[0].(string)
			referenceStack.Push(i + 1)
			i = (*interpreter.Labels)[labelName] - 1
			continue
		case instructions.INS_GOTO:
			labelName := instruction.Params[0].(string)
			if labelName == "$" {
				i = referenceStack.Pop().(int) - 1
				continue
			}
			i = (*interpreter.Labels)[labelName] - 1
			continue
		case instructions.INS_DUMP:
			fmt.Printf("=== SickObjectStack Dump ===\n")
			for i := len(objectStack); i > 0; i-- {
				var anno string
				if len(objectStack) == i {
					anno = "   <-- head"
				}
				fmt.Printf("%v: %v%v\n", i, objectStack[i-1].ToHuman(), anno)
			}
			fmt.Printf("==================\n")
			continue
		case instructions.INS_VOID:
			continue
		default:
			return fmt.Errorf("Interpreter: No handling for instruction: %v", instruction.OpCode)
		}
	}
	return nil
}
