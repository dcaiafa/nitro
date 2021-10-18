package vm

type Instr struct {
	opc OpCode
	op1 uint32
	op2 uint16
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
	OpLoadGlobalDeref
	OpLoadLocal
	OpLoadLocalRef
	OpLoadLocalDeref
	OpCaptureLocal
	OpLoadArg
	OpLoadArgRef
	OpLoadArgDeref
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
	OpLiftArg
	OpInitLiftedLocal
	OpInitLiftedGlobal
)

const (
	CallArgCountMask uint32 = 0x3FFFFFFF
	CallExpandFlag   uint32 = 0x80000000
	CallPipelineFlag uint32 = 0x40000000
)
