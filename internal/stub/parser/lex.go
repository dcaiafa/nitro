package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/dcaiafa/nitro/internal/stub/ast"
	"github.com/dcaiafa/nitro/internal/token"
)

var keywords = map[string]int{
	"const":   kCONST,
	"func":    kFUNC,
	"import":  kIMPORT,
	"nil":     kNIL,
	"package": kPACKAGE,
	"struct":  kSTRUCT,
	"type":    kTYPE,
}

type lex struct {
	Unit *ast.Unit

	input      *bytes.Reader
	buf        bytes.Buffer
	pos        int
	err        error
	line       int
	col        int
	lastTokPos token.Pos
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

	defer func() {
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
			continue
		}
		switch r {
		case '"':
			l.unread()
			return l.scanQuotedString(lval)

		case '`':
			l.unread()
			return l.scanBackQuotedString(lval)

		// Parse '.' and '...'
		case '.':
			r = l.read()
			if r != '.' {
				l.unread()
				return int('.')
			}
			r = l.read()
			if r != '.' {
				return LEXERR
			}
			return ELLIPSIS
		case '=', ':', ';', '(', ')', '[', ']', '*', ',', '?', '-', '{', '}':
			return int(r)
		default:
			if isNumber(r) {
				l.unread()
				return l.scanNumber(lval)
			} else if isLetter(r) || r == '_' {
				l.unread()
				tok := l.scanIdentifier(lval)
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

func (l *lex) scanBackQuotedString(lval *yySymType) int {
	l.buf.Reset()

	if l.read() != '`' {
		return LEXERR
	}

	for {
		r := l.read()
		if r == '`' {
			break
		}
		l.buf.WriteRune(r)
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
		return FLOAT
	} else {
		var err error
		lval.tok.Int, err = strconv.ParseInt(l.buf.String(), 10, 64)
		if err != nil {
			return LEXERR
		}
		lval.tok.Type = token.Int
		return INT
	}
}

func (l *lex) read() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	if r == '\n' {
		l.col = 0
		l.line++
	} else {
		l.col++
	}
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
