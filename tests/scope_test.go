package tests

import "testing"

func TestScope(t *testing.T) {
	RunSubErr(t, "", `
		print(a)
		var a = 1
		print(a)`,
		nil)
}
