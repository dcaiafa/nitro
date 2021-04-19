package runtime

type nilValue struct{}

func (v nilValue) String() string { return "<nil>" }
func (v nilValue) Type() string   { return "nil" }

func wrapNil(v Value) Value {
	if v == nil {
		return nilValue{}
	}
	return v
}
