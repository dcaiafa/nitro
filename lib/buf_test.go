package lib

import "testing"

func TestBuf(t *testing.T) {
	RunSubO(t, ``, `
		var b = buf.new()
		["hello", "world"] | b()
		print(b)
	`, `
hello
world
	`)
}
