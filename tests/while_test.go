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

	RunSubO(t, "break", `
		var x = 0
		while x < 3 {
    	print(x)
			if x == 1 {
				break
			}
			x = x + 1
		}
	`, `
0
1
	`)

	RunSubO(t, "continue", `
		var x = 0
		while x < 3 {
			x = x + 1

			if x == 2 {
				continue
			}

    	print(x)
		}
	`, `
1
3
	`)
}
