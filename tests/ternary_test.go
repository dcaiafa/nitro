package tests

import "testing"

func TestTernaryExpr(t *testing.T) {
	RunSubO(t, "then", `print(true ? "yes" : "no")`, "yes")
	RunSubO(t, "else", `print(false ? "yes" : "no")`, "no")
	RunSubO(t, "nested", `
	func f(a) {
		return 
			a >= 20 ? "high" :
			a >= 10 ? "medium" :
			"low"
	}
	print(f(5), f(11), f(20))
	`, `low medium high`)
}
