package tests

import "testing"

func TestForStmt(t *testing.T) {
	RunSubO(t, "for_iter", `
		var i = 0
		fn iter(n)
			return fn()
				i = i + 1
				if i > n
					return 0, false
				end
				return i, true
			end
		end
		for x in iter(3)
    	print(x)
		end
	`, `
1
2
3
`)

	RunSubO(t, "for_array", `
		for x in ["foo", "bar"]
    	print(x)
		end
	`, `
foo
bar
`)

	RunSubO(t, "for_array_empty", `
		print("begin")
		for x in []
    	print(x)
		end
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
		for k, v in d 
			print(k, v)
		end
	`, `
x 1
y {z: 2}
`)

	RunSubO(t, "for_object_empty", `
		print("begin")
		for k, v in {}
			print(k, v)
		end
		print("end")
	`, `
begin
end
`)

}
