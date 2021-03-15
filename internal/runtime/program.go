package runtime

type Program struct {
	globals   int
	fns       []Fn
	extFns    []ExternFn
	literals  []Value
	params    map[string]*Param
	reqParamN int
}
