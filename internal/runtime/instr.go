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

type Instr uint64

func MakeInstr(opcode OpCode, operand uint64) Instr {
	return Instr((uint64(opcode) << 56) | operand&0x00FFFFFFFFFFFFFF)
}

func (i Instr) Op() OpCode {
	return OpCode((i >> 56) & 0xFF)
}

func (i Instr) Operand() uint64 {
	return uint64(i) & 0x00FFFFFFFFFFFFFF
}
