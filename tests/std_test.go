package tests

import "testing"

func TestLen(t *testing.T) {
	RunSubO(t, "nil", `var a; print(len(a))`, `0`)
	RunSubO(t, "object", `var a = {a: 1, b: 3.1415, c: false}; print(len(a), len({}))`, `3 0`)
	RunSubO(t, "array", `var a = [1, 2]; print(len(a), len([]))`, `2 0`)
	RunSubO(t, "string", `var a = "foobar"; print(len(a), len(""))`, `6 0`)
}

func TestRange(t *testing.T) {
	RunSubO(t, "end", `
		for i in range(3)
			print(i)
		end
  `, `
0
1
2
`)

	RunSubO(t, "start_lt_end", `
		for i in range(1, 3)
			print(i)
		end
  `, `
1
2
`)

	RunSubO(t, "start_lt_end_step", `
		for i in range(1, 7, 2)
			print(i)
		end
  `, `
1
3
5
`)

	RunSubO(t, "start_gt_end", `
		for i in range(3, 1)
			print(i)
		end
  `, `
3
2
`)

	RunSubO(t, "start_gt_end_step", `
		for i in range(5, -2, -2)
			print(i)
		end
  `, `
5
3
1
-1
`)
}
