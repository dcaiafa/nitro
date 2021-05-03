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

	RunSubO(t, `no_input`, `
		exec({ cmd: ["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"] }) |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `input`, `
		range(1024) | 
			map(tostring) |
			exec({ cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stdout"] }) |
				lines() |
				map(&e -> parseint(e)) |
				reduce(sum) |
				print()
  `, `523776`)

	RunSubO(t, `input_concise`, `
		range(100001) |
			map(tostring) |
			exec("go", "run", "./testexec/testexec.go", "-echo-to-stdout") |
			lines() |
			map(&e -> parseint(e)) |
			reduce(sum) |
			print()
	`, `5000050000`)

	RunSubO(t, `capture_stderr`, `
		var err_buf = buf()
		var out = range(2049) | 
			map(tostring) |
			exec({ 
				cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-range", "11", "-range-stdout"]
				stderr: err_buf
			})
    var out_sum = out | lines() | map(parseint) | reduce(sum)
		var err_sum = err_buf | lines() | map(parseint) | reduce(sum)
		print(out_sum, err_sum)
	`, `55 2098176`)

	RunSubO(t, `redirect_stderr`, `
		"hello world" |
			exec("go", "run", "./testexec/testexec.go", "-echo-to-stdout") |
			read() |
			print()
	`, `hello world`)

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
