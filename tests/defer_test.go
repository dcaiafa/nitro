package tests

import "testing"

func TestDefer(t *testing.T) {
	RunSubO(t, "main", `
		var x = 1
		defer print(x)
	`, ``)
}
