package vm

type Instr struct {
	op       OpCode
	operand1 uint32
	operand2 uint16
}

type OpCode byte

const (
	OpNop OpCode = iota
	OpJump
	OpJumpIfTrue
	OpJumpIfFalse
	OpDup
	OpPop
	OpCall
	OpNil
	OpNewClosure
	OpNewInt
	OpNewBool
	OpNewObject
	OpNewArray
	OpLoadGlobal
	OpLoadGlobalRef
	OpLoadLocal
	OpLoadLocalRef
	OpCaptureLocal
	OpLoadArg
	OpLoadArgRef
	OpCaptureArg
	OpLoadCapture
	OpLoadCaptureRef
	OpLoadFn
	OpLoadNativeFn
	OpLoadLiteral
	OpEvalBinOp
	OpNot
	OpUnaryMinus
	OpObjectPutNoPop
	OpObjectGet
	OpObjectGetRef
	OpArrayAppendNoPop
	OpRet
	OpStore
	OpInitCallFrame
	OpMakeIter
	OpBeginTry
	OpEndTry
	OpSwap
	OpThrow
	OpDefer
	OpNext
	OpSlice
	OpIterYield
	OpIterRet
	OpNewIter
)
