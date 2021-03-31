package tests

import "testing"

func TestStringLiteral(t *testing.T) {
	RunSubO(t, "simple_escape_seq", `print("Hello\n\r\t\\World")`, "Hello\n\r\t\\World")
	RunSubO(t, "hex", `print("Hello\x2cWorld")`, "Hello,World")
}
