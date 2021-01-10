package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/token"
)

var keywords = map[string]int{
	"and":    kAND,
	"do":     kDO,
	"elif":   kELIF,
	"else":   kELSE,
	"end":    kEND,
	"false":  kFALSE,
	"fn":     kFN,
	"for":    kFOR,
	"if":     kIF,
	"in":     kIN,
	"not":    kNOT,
	"or":     kOR,
	"return": kRETURN,
	"then":   kTHEN,
	"true":   kTRUE,
	"var":    kVAR,
	"while":  kWHILE,
}

type lex struct {
	Main *ast.Main

	input      *bytes.Reader
	buf        bytes.Buffer
	pos        int
	err        error
	line       int
	col        int
	lastTokPos token.Pos

	synthSemicol bool
}

func newLex(input []byte) *lex {
	return &lex{
		input: bytes.NewReader(input),
		line:  1,
		col:   0,
	}
}

func (l *lex) Lex(lval *yySymType) int {
	return l.scan(lval)
}

func (l *lex) scan(lval *yySymType) int {
	lval.tok = token.Token{}

	eolToken := false
	defer func() {
		l.synthSemicol = eolToken
		l.lastTokPos = token.Pos{Line: l.line, Col: l.col}
		lval.tok.Pos = l.lastTokPos
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
			l.line++
			l.col = 0
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

				switch tok {
				case ID, kEND:
					eolToken = true

				case kTRUE:
					eolToken = true
					lval.tok.Type = token.Bool
					lval.tok.Bool = true

				case kFALSE:
					eolToken = true
					lval.tok.Type = token.Bool
					lval.tok.Bool = false
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

	lval.tok.Str = l.buf.String()

	keyword, ok := keywords[lval.tok.Str]
	if ok {
		lval.tok.Type = token.Keyword
		return keyword
	}

	lval.tok.Type = token.ID
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

	lval.tok.Type = token.String
	lval.tok.Str = l.buf.String()
	return STRING
}

func (l *lex) scanNumber(lval *yySymType) int {
	isFloat := false

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
		isFloat = true
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

	if isFloat {
		var err error
		lval.tok.Float, err = strconv.ParseFloat(l.buf.String(), 64)
		if err != nil {
			return LEXERR
		}
		lval.tok.Type = token.Float
	} else {
		var err error
		lval.tok.Int, err = strconv.ParseInt(l.buf.String(), 10, 64)
		if err != nil {
			return LEXERR
		}
		lval.tok.Type = token.Int
	}

	return NUMBER
}

func (l *lex) read() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	l.col++
	return r
}

func (l *lex) unread() {
	l.input.UnreadRune()
	l.col--
}

func (l *lex) Error(s string) {
	l.err = fmt.Errorf("%s:%d: %s", l.lastTokPos.Filename, l.lastTokPos.Line, s)
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
