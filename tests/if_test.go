package tests

import "testing"

func TestIf(t *testing.T) {
	RunSubO(t, "if_true", `
		print("before")
		if 1 < 2 then
			print("block1")
		end
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_false", `
		print("before")
		if 1 > 2 then
			print("block1")
		end
		print("after")
	`, `
before
after`)

	RunSubO(t, "if_else_true", `
		print("before")
		if 1 < 2 then
			print("block1")
		else
			print("else")
		end
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_else_false", `
		print("before")
		if 1 > 2 then
			print("block1")
		else
			print("else")
		end
		print("after")
	`, `
before
else
after`)

	RunSubO(t, "if_elif_else_true_true", `
		print("before")
		if true then
			print("block1")
		elif true then
			print("block2")
		else
			print("else")
		end
		print("after")
	`, `
before
block1
after`)

	RunSubO(t, "if_elif_else_false_true", `
		print("before")
		if false then
			print("block1")
		elif true then
			print("block2")
		else
			print("else")
		end
		print("after")
	`, `
before
block2
after`)

	RunSubO(t, "if_elif_else_false_false", `
		print("before")
		if false then
			print("block1")
		elif false then
			print("block2")
		else
			print("else")
		end
		print("after")
	`, `
before
else
after`)

	RunSubO(t, "fn", `
		fn evenOdd(n)
			if n % 2 == 0 then
				return "even"
			else
				return "odd"
			end
		end
		print(evenOdd(10))
		print(evenOdd(13))
`, `
even
odd`)
}
