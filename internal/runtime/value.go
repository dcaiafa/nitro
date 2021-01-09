package runtime

import "context"

type Value interface{}

type Closure struct {
	Fn       *Fn
	Captures []Value
}

type ExternFn func(
	ctx context.Context,
	args []Value,
) ([]Value, error)

type Fn struct {
	instrs []Instr
}
