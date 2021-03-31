package parser2

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

var escapeSeqRegex = regexp.MustCompile(`\\x[A-Fa-f0-9]{2}|\\u[A-Fa-f0-9]{2}|\\U[A-Fa-f0-9]{4}|\\.`)

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
