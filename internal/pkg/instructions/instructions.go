package instructions

type Instruction struct {
	OpCode int
	Params []interface{}
}

const (
	INS_IPUSH   = iota // int push
	INS_SPUSH          // string push
	INS_BPUSH          // bool push
	INS_ADD            // add
	INS_SUB            // substract
	INS_MUL            // multiply
	INS_DIV            // division
	INS_MOD            // modulo
	INS_CMP            // compare
	INS_LT             // less than
	INS_GT             // greater than
	INS_LTE            // less than or equals
	INS_GTE            // greater than or equals
	INS_STORE          // stores identifier associated with head of stack
	INS_LOAD           // loads value from storage by identifier
	INS_DEL            // deletes value from storage by identififer
	INS_JMP            // jumps to instruction by instruction count
	INS_CJMP           // conditional jump
	INS_DUP            // duplicates head
	INS_PRINT          // prints head of stack
	INS_PRINTLN        // prints head of stack with newline
	INS_GOTO           // goto specified label
	INS_DUMP           // print whole stack
	INS_VOID           // do nothing
)
