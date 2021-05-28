package tests

import "testing"

func TestIter(t *testing.T) {
	RunSubO(t, "simple", `
		func iterable() {
			yield 1
			yield 2
		}
		for i in iterable() {
    	print(i)
		}
`, `
1
2
`)

	RunSubO(t, "simple2", `
		func iterable() {
			yield 1,10
			yield 2,20
		}
		for i,j in iterable() {
    	print(i,j)
		}
`, `
1 10
2 20
`)

	RunSubO(t, "a_little_more_complex", `
		func counter(n, m) {
			var last
			for i in range(n) {
				var res = i % m
				yield res, last
				last = res
			}
		}
		for i, l in counter(5, 3) {
			print(i, l)
		}
	`, `
0 <nil>
1 0
2 1
0 2
1 0
`)

	RunSubO(t, "return", `
		func iter(x) {
			for i in range(1000) {
				if i > x {
					return
				}
				yield i
			}
		}
		for i in iter(3) {
    	print(i)
		}
`, `
0
1
2
3
`)

	RunSubO(t, "call_again_after_ended", `
		func iter() {
			yield 1
			yield 2
		}
		var i = iter()
		for x in i {
			print(x)
		}
		for x in i {
			print(x)
		}
`, `
1
2
`)
}
