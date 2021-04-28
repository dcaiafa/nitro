package lib

import "testing"

func TestExec(t *testing.T) {
	RunSubO(t, `simple`, `
		exec("go", "run", "./testexec/testexec.go", "-range", 1024, "-range-stdout") |
			lines() |





	`, ``)

}
