package tests

import "testing"

func TestWhile(t *testing.T) {
	RunSubO(t, "simple", `
		var x = 0
		while x < 3 {
    	print(x)
			x = x + 1
		}
	`, `
0
1
2
	`)
}
