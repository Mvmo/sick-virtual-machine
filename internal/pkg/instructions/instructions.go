package instructions

type Instruction struct {
	OpCode int
	Params []interface{}
}

const (
	INS_PUSH = iota
	INS_ADD
	INS_SUB
	INS_CMP
	INS_JMP
	INS_CJMP
	INS_DUMP
)