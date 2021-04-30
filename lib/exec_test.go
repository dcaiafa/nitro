package lib

import "testing"

func TestExec(t *testing.T) {
	RunSubO(t, `no_input_concise`, `
		exec("go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout") |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `input_concise`, `
		range(1024) |
			map(tostring) |
			exec("go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout") |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `input_concise_string`, `
		"hello world" |
			exec("go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout") |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `no_input`, `
		exec({ cmd: ["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"] }) |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `echo_input`, `
		range(1024) | 
			map(tostring) |
			exec({ cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stdout"] }) |
				lines() |
				map(&e -> parseint(e)) |
				reduce(sum) |
				print()
  `, `523776`)

	RunSubO(t, `process_returns_non_zero`, `
		try {
			"this will go in the error\n" |
				exec("go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-exit-code=128") |
				read()
		} catch e {
			print(e.error)
		}
	`, `
exit status 1
this will go in the error
exit status 128
`)

	/*
			RunSubO(t, `redirect-stdout`, `
				"howdy" |
				exec({ stdout: out()}, "go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout")
		  `, `523776`)
	*/

}
