package parser2

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"unicode/utf8"
)

var escapeSeqRegex = regexp.MustCompile(
	`\\x[A-Fa-f0-9]{2}|\\u[A-Fa-f0-9]{2}|\\U[A-Fa-f0-9]{4}|\\.`)

func expandEscapeSequences(s string) (string, error) {
	var err error
	s = escapeSeqRegex.ReplaceAllStringFunc(s, func(s string) string {
		switch s[1] {
		case 'n':
			return "\n"
		case 'r':
			return "\r"
		case 't':
			return "\t"
		case '"':
			return "\""
		case '\\':
			return "\\"
		case 'x':
			bytes, _ := hex.DecodeString(s[2:])
			return string(bytes)
		default:
			if err == nil {
				err = fmt.Errorf("invalid escape sequence %s", s)
			}
			return ""
		}
	})

	if err != nil {
		return "", err
	}

	return s, nil
}

var charEscapeSeqRegex = regexp.MustCompile(`\\.`)

func expandCharEscapeSequences(s string) (rune, error) {
	var err error
	s = charEscapeSeqRegex.ReplaceAllStringFunc(s, func(s string) string {
		switch s[1] {
		case 'n':
			return "\n"
		case 'r':
			return "\r"
		case 't':
			return "\t"
		case '\'':
			return "'"
		case '\\':
			return "\\"
		default:
			if err == nil {
				err = fmt.Errorf("invalid character escape sequence %s", s)
			}
			return ""
		}
	})

	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeRuneInString(s)
	if size != len(s) {
		return 0, fmt.Errorf("character literal must be a single rune")
	}

	return r, nil
}
