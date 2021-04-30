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

	// TODO: fix
	RunSubO(t, "implicit_start", `
		print("hello"[:4])
	`, `hell`)

	RunSubO(t, "implicit_end", `
		print("hello"[1:])
	`, `ello`)

	RunSubO(t, "relative_end", `
		print("hello"[:-1])
	`, `hell`)

	RunSubO(t, "relative_end2", `
		print("hello"[1:-1])
	`, `ell`)

	RunSubO(t, "relative_end3", `
		print("hello"[1:-2])
	`, `el`)

	RunSubO(t, "large_start", `
		print("hello"[100:])
	`, ``)

	RunSubO(t, "large_end", `
		print("hello"[:100])
	`, `hello`)

	RunSubO(t, "intersect", `
	  print("hello"[3:2])
	`, ``)

	RunSubErr(t, "err_negative_start", `
		print("hello"[-1:])
	`, nil)
}

func TestStringIter(t *testing.T) {
	RunSubO(t, "simple", `
		for c, i in "hello" {
    	print(c, i)
		}
	`, `
104 0
101 1
108 2
108 3
111 4
`)
	RunSubO(t, "empty", `
		for c, i in "" {
    	print(c, i)
		}
	`, ``)
}
