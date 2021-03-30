package tests

import "testing"

func TestForStmt(t *testing.T) {
	RunSubO(t, "for_iter", `
		var i = 0
		func iter(n)
			return func()
				i = i + 1
				if i > n
					return false, 0
				end
				return true, i
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

	RunSubO(t, "for_object_key_only", `
		var d = {
			x: 1
			y: { z: 2 }
		}
		for k in d 
			print(k)
		end
	`, `
x
y
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
