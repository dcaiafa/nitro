package lib

import "testing"

func TestExec(t *testing.T) {
	RunSubO(t, `literals`,
		"e`go run ./testexec/testexec.go -print-args ./some/path    \\other\\path 123 ~!@#$%^&*()[]{}| ` |\n"+`
			stdout
	`, `
[./some/path]
[\other\path]
[123]
[~!@#$%^&*()[]{}|]
  `)

	RunSubO(t, `exprs`, `
  var pets = [
    { name: "Chewie", alive: false },
    { name: "Deedee", alive: false },
    { name: "Ollie", alive: true },
  ]
` + "e`go run ./testexec/testexec.go -print-args {1} {[\"hello\", \"world\"] | join(\" \")} {pets | filter(&p->p.alive) | first | &p->p.name}` |\n"+`
			stdout
	`, `
[1]
[hello world]
[Ollie]
  `)

	RunSubO(t, `no_input`, `
		exec.exec(["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"]) |
			lines() |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `523776`)

	RunSubO(t, `input`, `
		range(100000) | 
			map(to_string) |
			exec.exec(["go", "run", "./testexec/testexec.go", "-echo-to-stdout"]) |
				read() |
				lines() |
				count() |
				print()
  `, `100000`)

	RunSubO(t, `pipe`, `
		exec.exec(["go", "run", "./testexec/testexec.go", "-range", "100000", "-range-stdout"]) |
		exec.exec(["go", "run", "./testexec/testexec.go", "-echo-to-stdout"]) |
		lines() |
		count() |
		print() 
	`, `100000`)

	RunSubO(t, `input2`, `
		range(100001) |
			map(to_string) |
			exec.exec(["go", "run", "./testexec/testexec.go", "-echo-to-stdout"]) |
			lines() |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `5000050000`)

	/*
		RunSubO(t, `capture_stderr`, `
			var err_buf = buf()
			var out = range(2049) |
				map(to_string) |
				exec.exec({
					cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-range", "11", "-range-stdout"]
					stderr: err_buf
				})
	    var out_sum = out | lines() | map(&l -> parse_int(l)) | reduce(sum)
			var err_sum = err_buf | lines() | map(&l -> parse_int(l)) | reduce(sum)
			print(out_sum, err_sum)
		`, `55 2098176`)
	*/

	/*
		RunSubO(t, `redirect_stderr`, `
			"hello world" |
				exec.exec("go", "run", "./testexec/testexec.go", "-echo-to-stdout") |
				read() |
				print()
		`, `hello world`)
	*/

	RunSubO(t, `process_returns_non_zero`, `
		try {
			"this will go in the error\n" |
				exec.exec(["go", "run", "./testexec/testexec.go", "-echo-to-stderr", "-exit-code=128"]) |
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
		exec.exec(["go", "run", "./testexec/testexec.go", "-range", "1024", "-range-stdout"]) |
			lines() |
			take(10) |
			map(&e -> parse_int(e)) |
			reduce(sum) |
			print()
	`, `45`)

	RunSubO(t, `stdin_is_file`, `
		var tmp = create_temp()
		defer remove_file(tmp)
		range(100000) | 
			map(to_string) |
			tmp
		seek(tmp, 0)
		tmp |
			exec.exec(["go", "run", "./testexec/testexec.go", "-echo-to-stdout"]) |
		  read |
			lines |
			count |
			print
`, `100000`)

	/*
	   	RunSubO(t, `stderr_is_file`, `
	   		var tmp = create_temp()
	   		defer remove_file(tmp)
	   		range(100000) |
	   			map(to_string) |
	   			exec.exec({ cmd: ["go", "run", "./testexec/testexec.go", "-echo-to-stderr"], stderr: tmp }) |
	   			discard
	   		seek(tmp, 0)
	   		read(tmp) |
	   			lines() |
	   			count() |
	   			print()
	   `, `100000`)
	*/

	RunSubO(t, `executable_not_found`, `
	try {
		exec.exec("./foo_bar_doesnt_exist") | read()
	} catch {
  	print("failed as expected")
	}
`, `failed as expected`)
}
