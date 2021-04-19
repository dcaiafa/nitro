package runtime

func copyArgs(args []Value, minArgs int) []Value {
	n := len(args)
	if n < minArgs {
		n = minArgs
	}
	cargs := make([]Value, n)
	copy(cargs, args)
	return cargs
}
