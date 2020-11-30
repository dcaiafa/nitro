package parser

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

var keywords = map[string]int{
	"and":   kAND,
	"elif":  kELIF,
	"else":  kELSE,
	"end":   kEND,
	"false": kFALSE,
	"if":    kIF,
	"in":    kIN,
	"not":   kNOT,
	"or":    kOR,
	"true":  kTRUE,
	"then":  kTHEN,
	"fn":    kFN,
	"var":   kVAR,
}

type lex struct {
	//Program *ast.Program

	input *strings.Reader
	buf   bytes.Buffer
	pos   int
	err   error

	synthSemicol bool
}

func newLex(input string) *lex {
	return &lex{
		input: strings.NewReader(input),
	}
}

func (l *lex) Lex(lval *yySymType) int {
	return l.scan(lval)
}

func (l *lex) scan(lval *yySymType) int {
	eolToken := false
	defer func() {
		l.synthSemicol = eolToken
	}()

	for {
		r := l.read()
		if r == 0 {
			return 0
		}
		if isSpace(r) {
			continue
		}
		if r == '\n' {
			if l.synthSemicol {
				return ';'
			}
			continue
		}
		switch r {
		case '=':
			r = l.read()
			if r != '=' {
				l.unread()
				return '='
			}
			return EQ
		case '!':
			r = l.read()
			if r != '=' {
				return LEXERR
			}
			return NE
		case '<':
			r = l.read()
			if r != '=' {
				l.unread()
				return '<'
			}
			return LE
		case '>':
			r = l.read()
			if r != '=' {
				l.unread()
				return '>'
			}
			return GE
		case '"':
			l.unread()
			eolToken = true
			return l.scanQuotedString(lval)
		case '|', '+', '-', '*', '/', ';', '(', ',', '[', ':', '.', '{':
			return int(r)
		case ')', ']', '}':
			eolToken = true
			return int(r)
		default:
			if isNumber(r) {
				l.unread()
				eolToken = true
				return l.scanNumber(lval)
			} else if isLetter(r) || r == '_' {
				l.unread()
				tok := l.scanIdentifier(lval)
				if tok == ID || tok == kEND {
					eolToken = true
				}
				return tok
			} else {
				return LEXERR
			}
		}
	}
}

func (l *lex) scanIdentifier(lval *yySymType) int {
	l.buf.Reset()

	r := l.read()
	if !isLetter(r) && r != '_' {
		return LEXERR
	}
	l.buf.WriteRune(r)

	for {
		r := l.read()
		if !isLetter(r) && !isNumber(r) && r != '_' {
			l.unread()
			break
		}
		l.buf.WriteRune(r)
	}

	lval.str = l.buf.String()

	keyword, ok := keywords[lval.str]
	if ok {
		return keyword
	}

	return ID
}

func (l *lex) scanQuotedString(lval *yySymType) int {
	l.buf.Reset()

	if l.read() != '"' {
		return LEXERR
	}

	for {
		r := l.read()
		if r == '"' {
			break
		} else if r == '\\' {
			r = l.read()
			switch r {
			case '\\', '"':
				l.buf.WriteRune(r)
			default:
				return LEXERR
			}
		} else if r == '\n' || r == '\r' {
			return LEXERR
		} else {
			l.buf.WriteRune(r)
		}
	}

	lval.str = l.buf.String()
	return STRING
}

func (l *lex) scanNumber(lval *yySymType) int {
	l.buf.Reset()
	l.buf.WriteRune(l.read())

	var r rune
	for {
		r = l.read()
		if !isNumber(r) {
			break
		}
		l.buf.WriteRune(r)
	}

	if r == '.' {
		l.buf.WriteRune(r)
		r = l.read()
		if !isNumber(r) {
			return LEXERR
		}
		l.buf.WriteRune(r)
		for {
			r = l.read()
			if !isNumber(r) {
				l.unread()
				break
			}
			l.buf.WriteRune(r)
		}
	}

	l.unread()

	var err error
	lval.num, err = strconv.ParseFloat(l.buf.String(), 64)
	if err != nil {
		return LEXERR
	}

	return NUMBER

}

func (l *lex) read() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	return r
}

func (l *lex) unread() {
	l.input.UnreadRune()
}

func (l *lex) Error(s string) {
	l.err = errors.New(s)
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
