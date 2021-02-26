package tests

import "testing"

func TestFn(t *testing.T) {
	RunSubO(t, "no_args", `
		fn func() 
      print("X")
		end
		func()
`, `X`)

	RunSubO(t, "ret1_stmt", `
		fn func(a, b) 
			return a + b
		end
		func(1, 2)
		print("done")
`, `done`)

	RunSubO(t, "ret1_var_decl", `
		fn func(a, b) 
			return a + b
		end
		var x = func(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "ret1_assign", `
		fn func(a, b) 
			return a + b
		end
		var x
		x = func(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "ret1_expr", `
		fn func(a, b) 
			return a + b
		end
		var x
		x = func(1, 2)
		print(x)
`, `3`)

	RunSubO(t, "local_vars", `
		fn func(a, b) 
			var x = a + b
			return x + 1
		end
		print(func(1, 2))
`, `4`)

	RunSubO(t, "sub_func", `
		fn X(a, b) 
			fn Y(x, y)
      	return x + y
			end
			return Y(a+1, b+1)
		end
		print(X(1, 2))
`, `5`)

	RunSubO(t, "sub_func_cap", `
		fn x() 
			var a = 0
			fn y() 
      	a = a + 1
				return a
			end
			return y
		end
		var f1 = x()
		var f2 = x()
		print(f1(), f2(), f1(), f2())
`, `1 1 2 2`)

	RunSubO(t, "sub_func_multi", `
		fn x() 
			var a = 0
			var f = fn() 
				return fn()
        	a = a + 1
					return a
				end
			end
			a = 10
			return f
		end
		var f1 = x()()
		var f2 = x()()
		print(f1(), f2(), f1(), f2())
`, `11 11 12 12`)

	RunSubO(t, "stmt_obj_literal", `
		fn x(o)
			o.stuff = "yay"
			print(o)
		end
		x{other: 123}`,
		`{other: 123, stuff: "yay"}`)

	RunSubO(t, "expr_obj_literal", `
		fn x(o)
			o.stuff = "yay"
			return o
		end
		print(x{other: 123})`,
		`{other: 123, stuff: "yay"}`)
}
