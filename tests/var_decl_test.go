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

	RunSubO(t, "init_multi", `
		fn foo(x)
			return x + 1
		end
  	var a, b, c  = 1, foo(1), 3
		print(a, b, c)
	`, `1 2 3`)

	RunSubO(t, "init_func_single", `
		fn foo(x)
			return x + 1
		end
  	var a = foo(1)
		print(a)
	`, `2`)

	RunSubO(t, "init_func_multi", `
		fn foo(x)
			return x, x+1, x+2
		end
  	var a, b, c = foo(1)
		print(a, b, c)
	`, `1 2 3`)

	RunSubO(t, "init_func_multi_discard", `
		fn foo(x)
			return x, x+1, x+2
		end
  	var a = foo(1)
		print(a)
	`, `1`)
}
