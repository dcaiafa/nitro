package tests

import "testing"

func VarDeclTest(t *testing.T) {
	RunSubO(t, "global_no_init", `
		var a
		a = 1
		print(a)
`, `1`)
	RunSubO(t, "global_with_init", `
		var a = 1
		print(a)
`, `1`)
	RunSubO(t, "local_no_init", `
		var a = 1
		fn f()
			var a
			a = 2
			return a
		end
		print(f(), a)
`, `2 1`)
	RunSubO(t, "local_with_init", `
		var a = 1
		fn f()
			var a = 2
			return a
		end
		print(f(), a)
`, `2 1`)
}
