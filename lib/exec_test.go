package lib

import "testing"

func TestExec(t *testing.T) {
	RunSubO(t, `no_input_concise`, `
		exec("go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout") |
			lines() |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `no_input`, `
		exec({ cmd: ["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"] }) |
			lines() |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `input`, `
		range(100000) | 
			map(to_string) |
			exec({ cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stdout"] }) |
				read() |
				lines() |
				count() |
				print()
  `, `100000`)

	RunSubO(t, `pipe`, `
		exec("go", "run", "./testexec/testexec.go", "-range", "100000", "-range-stdout") |
		exec("go", "run", "./testexec/testexec.go", "-echo-to-stdout") |
		lines() |
		count() |
		print() 
	`, `100000`)

	RunSubO(t, `input_concise`, `
		range(100001) |
			map(to_string) |
			exec("go", "run", "./testexec/testexec.go", "-echo-to-stdout") |
			lines() |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `5000050000`)

	RunSubO(t, `capture_stderr`, `
		var err_buf = buf()
		var out = range(2049) | 
			map(to_string) |
			exec({ 
				cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-range", "11", "-range-stdout"]
				stderr: err_buf
			})
    var out_sum = out | lines() | map(&l -> parse_int(l)) | reduce(sum)
		var err_sum = err_buf | lines() | map(&l -> parse_int(l)) | reduce(sum)
		print(out_sum, err_sum)
	`, `55 2098176`)

	RunSubO(t, `switch_output`, `
		var err_buf = buf()
		var out = range(2049) | 
			map(to_string) |
			exec({ 
				cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-range", "11", "-range-stdout"]
				stderr: err_buf
				switchoutput: true
			})
    var out_sum = out | lines() | map(&l -> parse_int(l)) | reduce(sum)
		var err_sum = err_buf | lines() | map(&l -> parse_int(l)) | reduce(sum)
		print(out_sum, err_sum)
	`, `2098176 55`)

	/*
		// TODO: flacky test
		RunSubO(t, `combine_output`, `
				exec({
					cmd: ["go", "run", "./testexec/testexec.go", "-range", "2049", "-range-alt"]
					combineoutput: true
				}) |
					lines() |
					map(&l -> parse_int(l)) |
					reduce(sum) |
					print()
		`, `2098176`)
	*/

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

	RunSubO(t, `abort`, `
		exec("go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout") |
			lines() |
			take(10) |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `45`)
}
