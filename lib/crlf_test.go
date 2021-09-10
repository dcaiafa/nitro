package lib

import (
	"testing"
)

func TestFromCRLF(t *testing.T) {
	run := func(name, input, output string) {
		RunSubO(t, name, `"`+input+`" | fromcrlf | &r->"["+r+"]" | print`, "["+output+"]")
	}

	run("crlf", `abc\r\ndef\r\n`, "abc\ndef\n")
	run("crcrlf", `abc\r\r\ndef\r\n`, "abc\r\ndef\n")
	run("crlflf", `abc\r\n\ndef\r\n`, "abc\n\ndef\n")
	run("lf", `abc\ndef\r\n`, "abc\ndef\n")
	run("lf_end", `abc\ndef\n`, "abc\ndef\n")
	run("cr", `abc\rdef\r\n`, "abc\rdef\n")
	run("cr_end", `abc\r\ndef\r`, "abc\ndef\r")

	RunSubO(t, "reader", `
	var b = buf()
	"abc\r\ndef\r\n" | b
	b = b | fromcrlf

	print(nonl(), "[")
	print(nonl(), b | read(4))
	print(nonl(), b | read())
	print(nonl(), "]")
`, "[abc\ndef\n]")
}

func TestToCRLF(t *testing.T) {
	run := func(name, input, output string) {
		RunSubO(t, name, `"`+input+`" | tocrlf(true) | &r->"["+r+"]" | print`, "["+output+"]")
	}

	run("crlf", `abc\r\ndef\r\n`, "abc\r\ndef\r\n")
	run("crcrlf", `abc\r\r\ndef\r\n`, "abc\r\r\ndef\r\n")
	run("crlflf", `abc\r\n\ndef\r\n`, "abc\r\n\r\ndef\r\n")
	run("lf", `abc\ndef\r\n`, "abc\r\ndef\r\n")
	run("lf_end", `abc\ndef\n`, "abc\r\ndef\r\n")
	run("cr", `abc\rdef\n`, "abc\rdef\r\n")
	run("cr_end", `abc\ndef\r`, "abc\r\ndef\r")

	RunSubO(t, "reader", `
	var b = buf()
	"abc\ndef\n" | b
	b = b | tocrlf(true)

	print(nonl(), "[")
	print(nonl(), b | read(4))
	print(nonl(), b | read())
	print(nonl(), "]")
`, "[abc\r\ndef\r\n]")
}
