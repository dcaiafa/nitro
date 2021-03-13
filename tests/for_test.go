package tests

import "testing"

func TestForStmt(t *testing.T) {
	RunSubO(t, "for", `
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
}
