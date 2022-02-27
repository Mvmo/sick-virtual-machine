package instructions

type Instruction struct {
	OpCode int
	Params []interface{}
}

const (
	INS_PUSH = iota
	INS_ADD
	INS_SUB
	INS_MUL
	INS_DIV
	INS_MOD
	INS_CMP
	INS_LT
	INS_GT
	INS_LTE
	INS_GTE
	INS_JMP
	INS_CJMP
	INS_DUP
	INS_GOTO
	INS_DUMP
	INS_VOID
)
