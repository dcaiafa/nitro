package tests

import "testing"

func TestTryCatch(t *testing.T) {
	RunSubO(t, "same_frame", `
	print(1)
	try {
		print(2)
		("throw")()
		print(3)
	} catch {
		print(4)
	}
  print(5)
	`, `
1
2
4
5
`)

	RunSubO(t, "diff_frame", `
	func g() {
		print("g1")
		("throw")()
		print("g2")
	}
	func f() {
		print("f1")
		try {
			print("f2")
			print("from g:" + g())
			print("f3")
		} catch {
			print("f4")
		}
		print("f5")
	}
	f()
	`, `
f1
f2
g1
f4
f5
`)

	RunSubO(t, "cascading", `
	func g() {
		print("g1")
		throw "boom"
		print("g2")
	}
	func f() {
		var x = 0
		print("f1")
		try {
			print("f2")
			x = "f4"
			print("from g:" + call(g))
			print("f3")
		} catch e {
			print(x)
			throw e
		}
		print("f5")
	}
	func x() {
		var i = 0
		try {
			i = 1
			f()
			i = 2
		} catch e {
			print(e)
		}
	}
	x()
	`, `
f1
f2
g1
f4
boom
 main.n:4  g
 <builtin>  tests.run.func3
 main.n:13  f
 main.n:25  x
 main.n:31  $main
`)

	RunSubO(t, "catch_error", `
	var x
	try {
		x = "hi"
		throw "nope"
		x = "bye"
	} catch e {
		print(e.error, x)
	}`, `
nope hi
`)
}
