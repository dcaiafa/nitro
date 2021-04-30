package lib

import "testing"

func TestExec(t *testing.T) {
	RunSubO(t, `no-input`, `
		exec({ cmd: ["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"] }) |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `echo-input`, `
		range(1024) | 
			map(tostring) |
			exec({ cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stdout"] }) |
				lines() |
				map(&e -> parseint(e)) |
				reduce(sum) |
				print()
  `, `523776`)

	/*
			RunSubO(t, `redirect-stdout`, `
				"howdy" |
				exec({ stdout: out()}, "go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout")
		  `, `523776`)
	*/

}
