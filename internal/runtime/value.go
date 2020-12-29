package runtime

type Value interface{}

type Closure struct {
	External bool
	Fn       int
	Captures []Value
}
