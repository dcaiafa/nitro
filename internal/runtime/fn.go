package runtime

type Fn struct {
	name      int
	locations []Location
	instrs    []Instr
	minArgs   int
}

func (f *Fn) Type() string   { return "Func" }
func (f *Fn) String() string { return "<func>" }
func (f *Fn) isCallable()    {}
