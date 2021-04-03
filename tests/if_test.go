package tests

import "testing"

func TestIf(t *testing.T) {
	RunSubO(t, "if_true", `
		print("before")
		if 1 < 2 {
			print("block1")
		}
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_false", `
		print("before")
		if 1 > 2 {
			print("block1")
		}
		print("after")
	`, `
before
after`)

	RunSubO(t, "if_else_true", `
		print("before")
		if 1 < 2 {
			print("block1")
		} else {
			print("else")
		}
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_else_false", `
		print("before")
		if 1 > 2 {
			print("block1")
		} else {
			print("else")
		}
		print("after")
	`, `
before
else
after`)

	RunSubO(t, "if_elif_else_true_true", `
		print("before")
		if true {
			print("block1")
		} else if true {
			print("block2")
		} else {
			print("else")
		}
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_elif_else_false_true", `
		print("before")
		if false {
			print("block1")
		} else if true {
			print("block2")
		} else {
			print("else")
		}
		print("after")
	`, `
before
block2
after`)

	RunSubO(t, "if_elif_else_false_false", `
		print("before")
		if false {
			print("block1")
		} else if false {
			print("block2")
		} else {
			print("else")
		}
		print("after")
	`, `
before
else
after`)

	RunSubO(t, "func", `
		func evenOdd(n) {
			if n % 2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}
		print(evenOdd(10))
		print(evenOdd(13))
`, `
even
odd`)
}
