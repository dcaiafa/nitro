package vm

type Fn struct {
	name      int
	locations []Location
	instrs    []Instr
	minArgs   int
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }
func (f *Fn) Traits() Traits { return TraitNone }

func (f *Fn) Call(m *VM, args []Value, nRet int) ([]Value, error) {
	// Function calls are handled directly by the VM.
	panic("not called")
}
