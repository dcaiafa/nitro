package tests

import (
	"testing"

	"github.com/dcaiafa/nitro"
)

func TestFn(t *testing.T) {
	RunSubO(t, "no_args", `
		func f() {
      print("X")
		}
		f()
`, `X`)

	RunSubO(t, "ret1_stmt", `
		func f(a, b) {
			return a + b
		}
		f(1, 2)
		print("done")
`, `done`)

	RunSubO(t, "ret1_var_decl", `
		func f(a, b) {
			return a + b
		}
		var x = f(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "ret1_assign", `
		func f(a, b) {
			return a + b
		}
		var x
		x = f(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "ret1_expr", `
		func f(a, b) {
			return a + b
		}
		var x
		x = f(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "local_vars", `
		func f(a, b) {
			var x = a + b
			return x + 1
		}
		print(f(1, 2))
`, `4`)

	RunSubO(t, "sub_func", `
		func X(a, b) {
			func Y(x, y) {
      	return x + y
			}
			return Y(a+1, b+1)
		}
		print(X(1, 2))
`, `5`)

	RunSubO(t, "sub_func_cap", `
		func x() {
			var a = 0
			func y() {
      	a = a + 1
				return a
			}
			return y
		}
		var f1 = x()
		var f2 = x()
		print(f1(), f2(), f1(), f2())
`, `1 1 2 2`)

	RunSubO(t, "sub_func_cap_arg", `
		func x(a) {
			func y() {
      	a = a + 1
				return a
			}
			return y
		}
		var f1 = x(10)
		var f2 = x(20)
		print(f1(), f2(), f1(), f2())
`, `11 21 12 22`)

	RunSubO(t, "sub_func_multi", `
		func x() {
			var a = 0
			var f = func() {
				return func() {
        	a = a + 1
					return a
				}
			}
			a = 10
			return f
		}
		var f1 = x()()
		var f2 = x()()
		print(f1(), f2(), f1(), f2())
`, `11 11 12 12`)

	RunSubO(t, "lambda_inline_call", `
		print((&x,y->x+y)(1,2))
	`, `3`)

	RunSubO(t, "lambda_as_arg", `
		func apply(f, v) {
			return f(v)
		}
		print(apply(&x->x+1, 10))
	`, `11`)

	RunSubO(t, "lambda_with_capture", `
		var n = 3
		func apply(f, v) {
			return f(v)
		}
		print(apply(&x->x+n, 10))
		n = 4
		print(apply(&x->x+n, 10))
	`, `
13
14
`)

	RunSubO(t, "lambda_return_object", `
		print((&n->{name: n})("bob"))
	`, `{name: "bob"}`)

	RunSubO(t, "recursive", `
		func fib(n) {
    	if n <= 1 {
				return 1
			} else {
				return fib(n-2) + fib(n-1)
			}
		}
		print(fib(6))
	`, `13`)

	RunSubO(t, "less_args", `
		func f(a, b) {
			print(a, b)
		}
    f(1)
		`, `1 <nil>`)

	RunSubErr(t, "err_call_nil", `
			var a
			a(2)
		`, nitro.ErrCannotCallNil)

	RunSubO(t, "call_expand", `
			print(1, [2, 3, 4]...)
		`, `1 2 3 4`)

	RunSubO(t, "call_expand_empty", `
			print(1, []...)
		`, `1`)

	RunSubO(t, "call_expand_nil", `
			print(1, nil...)
		`, `1`)
}
