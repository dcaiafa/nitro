package tests

import "testing"

func TestForStmt(t *testing.T) {
	RunSubO(t, "for_iter", `
		var i = 0
		func iter(n) {
			return func() {
				i = i + 1
				if i > n {
					return false, 0
				}
				return true, i
			}
		}
		for x in iter(3) {
    	print(x)
		}
	`, `
1
2
3
`)

	RunSubO(t, "for_array", `
		for x in ["foo", "bar"] {
    	print(x)
		}
	`, `
foo
bar
`)

	RunSubO(t, "for_array_empty", `
		print("begin")
		for x in [] {
    	print(x)
		}
		print("end")
	`, `
begin
end
`)

	RunSubO(t, "for_object", `
		var d = {
			x: 1
			y: { z: 2 }
		}
		for k, v in d {
			print(k, v)
		}
	`, `
x 1
y {z: 2}
`)

	RunSubO(t, "for_object_key_only", `
		var d = {
			x: 1
			y: { z: 2 }
		}
		for k in d {
			print(k)
		}
	`, `
x
y
`)

	RunSubO(t, "for_object_empty", `
		print("begin")
		for k, v in {} {
			print(k, v)
		}
		print("end")
	`, `
begin
end
`)

}
