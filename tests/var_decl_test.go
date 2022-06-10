package tests

import "testing"

func TestVarDecl(t *testing.T) {
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
		func f() {
			var a
			a = 2
			return a
		}
		print(f(), a)
`, `2 1`)
	RunSubO(t, "local_with_init", `
		var a = 1
		func f() {
			var a = 2
			return a
		}
		print(f(), a)
`, `2 1`)

	RunSubO(t, "init_multi", `
		func foo(x) {
			return x + 1
		}
  	var a, b, c  = 1, foo(1), 3
		print(a, b, c)
	`, `1 2 3`)

	RunSubO(t, "init_func_single", `
		func foo(x) {
			return x + 1
		}
  	var a = foo(1)
		print(a)
	`, `2`)

	RunSubO(t, "init_func_multi", `
		func foo(x) {
			return x, x+1, x+2
		}
  	var a, b, c = foo(1)
		print(a, b, c)
	`, `1 2 3`)

	RunSubO(t, "init_func_multi_discard", `
		func foo(x) {
			return x, x+1, x+2
		}
  	var a = foo(1)
		print(a)
	`, `1`)

	RunSubO(t, "shadow", `
		var a = 1
		func f() {
			var a = a + 1
			return a
		}
		print(f(), a)
`, `2 1`)

	RunSubO(t, "scope", `
for i in range(5) {
  var state
  if i % 2 == 0 {
    state = i
  }
  print(state)
}
`, `
0
<nil>
2
<nil>
4
`)

}
