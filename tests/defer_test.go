package tests

import "testing"

func TestDefer(t *testing.T) {
	RunSubO(t, "from_main", `
		var x = 1
		defer print(x)
	`, `1`)

	RunSubO(t, "from_func", `
		func f() 
			var x = 1
      defer print(x)
			print(x)
			x = x + 1
		end
		f()
	`, `
1
2
`)

	RunSubO(t, "multi", `
		func f()
			var x = 1
			defer print("bye")
			defer print(x)
			print(x)
			x = x + 1
		end
		f()
`, `
1
2
bye
`)

	RunSubO(t, "exception", `
		func f()
			var x = 1
			defer func()
				print("bye from", x)
			end()
			
			print("bomb")
			throw "boom"
		end

		try
			f()
		catch e
			print("it did go", e.error)
		end
`, `
bomb
bye from 1
it did go boom
`)
}
