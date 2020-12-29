package runtime

type OpCode byte

const (
	OpNop OpCode = iota
	OpCall
	OpMakeClosure
	OpPushInt
	OpPushLocal
	OpPushLocalRef
	OpRet
	OpStore
)

type Instr struct {
	Op       OpCode
	Operand2 byte
	Operand1 uint16
}
