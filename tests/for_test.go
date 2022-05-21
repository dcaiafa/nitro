package tests

import "testing"

func TestForStmt(t *testing.T) {
	RunSubO(t, "iter", `
		for x in range(1,4) {
    	print(x)
		}
	`, `
1
2
3
`)

	RunSubO(t, "break", `
		for x in range(1,10) {
			if x == 5 {
				break
			}
    	print(x)
		}
	`, `
1
2
3
4
`)

	RunSubO(t, "continue", `
		for x in range(1,5) {
			if x == 2 {
				continue
			}
    	print(x)
		}
	`, `
1
3
4
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

	RunSubO(t, "closure", `
		for i in [1,2] {
			print((&x -> x+i)(i))
		}
`, `
2
4
`)
}
