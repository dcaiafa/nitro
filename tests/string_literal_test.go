package tests

import "testing"

func TestStringLiteral(t *testing.T) {
	// TODO: unicode sequences
	RunSubO(t, "simple_escape_seq", `print("Hello\n\r\t\\World")`, "Hello\n\r\t\\World")
	RunSubO(t, "hex", `print("Hello\x2cWorld")`, "Hello,World")
	// TODO: unicode sequences
	RunSubO(t, "char", `print('A', '\n', '\'')`, `65 10 39`)
}

func TestStringSlice(t *testing.T) {
	RunSubO(t, "full", `
		print("hello"[0:4])
`, `hell`)
	/*
	     // TODO: fix
	   	RunSubO(t, "implicit_start", `
	   		print("hello"[:4])
	   `, `hell`)
	*/
}
