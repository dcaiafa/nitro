package parser

import (
	"bytes"
	gotoken "go/token"
	"io"
)

type Token struct {
	Type TokenType
	Str  string
	Pos  gotoken.Pos
}

var keywords = map[string]TokenType{
	"!info":  INFO,
	"!param": PARAM,
	"!flag":  FLAG,

	"and":      AND,
	"break":    BREAK,
	"catch":    CATCH,
	"continue": CONTINUE,
	"defer":    DEFER,
	"else":     ELSE,
	"false":    FALSE,
	"for":      FOR,
	"func":     FUNC,
	"if":       IF,
	"import":   IMPORT,
	"nil":      NIL,
	"not":      NOT,
	"or":       OR,
	"return":   RETURN,
	"throw":    THROW,
	"true":     TRUE,
	"try":      TRY,
	"var":      VAR,
	"while":    WHILE,
	"yield":    YIELD,
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

type ErrLogger interface {
	Errorf(pos gotoken.Position, msg string, args ...any)
}

type lex struct {
	file  *gotoken.File
	input *bytes.Reader
	errs  ErrLogger
	buf   bytes.Buffer
	char  rune
}

func newLex(file *gotoken.File, input []byte, errs ErrLogger) *lex {
	l := &lex{
		file:  file,
		input: bytes.NewReader(input),
		errs:  errs,
	}
	l.advance()
	return l
}

func (l *lex) ReadToken() (Token, TokenType) {
	var tok Token
	l.readToken(&tok)
	return tok, tok.Type
}

func (l *lex) advance() {
	if l.char == '\n' {
		// Line starts at the character after the \n.
		l.file.AddLine(l.offset() + 1)
	}
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
}

func (l *lex) recover() {
	for {
		l.advance()
		ch := l.peek()
		if ch == ' ' || ch == '\n' || ch == 0 {
			return
		}
	}
}

func (l *lex) offset() int {
	offset, _ := l.input.Seek(0, io.SeekCurrent)
	return int(offset) - 1
}

func (l *lex) pos() gotoken.Pos {
	return l.file.Pos(l.offset())
}

func (l *lex) peek() rune {
	return l.char
}

func (l *lex) readToken(tok *Token) {
	tok.Type = -1
	for tok.Type == -1 {
		r := l.peek()
		if isSpace(r) {
			l.advance()
			continue
		}
		if r == '\n' {
			l.advance()
			continue
		}
		tok.Pos = l.pos()
		if r == 0 {
			tok.Type = EOF
			return
		}
		switch r {
		case '=':
			l.advance()
			switch l.peek() {
			case '=':
				l.advance()
				tok.Type = EQ // ==
			default:
				tok.Type = ASSIGN // =
			}
		case '+':
			l.advance()
			switch l.peek() {
			case '+':
				l.advance()
				tok.Type = INC // ++
			case '=':
				l.advance()
				tok.Type = ASSIGN_ADD // +=
			default:
				tok.Type = ADD // +
			}
		case '-':
			l.advance()
			switch l.peek() {
			case '-':
				l.advance()
				tok.Type = DEC
			case '=':
				l.advance()
				tok.Type = ASSIGN_SUB
			default:
				tok.Type = SUB
			}
		case '*':
			l.advance()
			switch l.peek() {
			case '=':
				l.advance()
				tok.Type = ASSIGN_MUL
			default:
				tok.Type = MUL
			}
		case '/':
			l.advance()
			switch l.peek() {
			case '/':
				l.scanComment(tok)
			case '=':
				l.advance()
				tok.Type = ASSIGN_DIV
			default:
				tok.Type = DIV
			}
		case '!':
			l.advance()
			switch l.peek() {
			case '=':
				l.advance()
				tok.Type = NE
			default:
				if isLetter(l.peek()) {
					l.scanBangKeyword(tok)
				} else {
					l.logError(l.pos(), "unexpected character: %c", r)
					l.recover()
					tok.Type = ERROR
				}
			}

		case '\'':
			l.scanLiteral(tok)
		case '@':
			l.scanKeyword(tok)
		default:
			if isLetter(r) || r == '_' {
				l.scanIdentifier(tok)
			} else if isNumber(r) {
				l.scanNum(tok)
			} else {
				l.errLogger.Errorf(
					l.file.Position(l.pos()), "unexpected character: %v", r)
				l.recover()
				tok.Type = ERROR
			}
		}
	}
}

func (l *lex) scanBangKeyword(tok *Token) {
}

func (l *lex) scanKeyword(tok *Token) {
	l.buf.Reset()

	r := l.peek()
	l.advance()
	l.buf.WriteRune(r)

	for {
		r := l.peek()
		if !isLetter(r) {
			break
		}
		l.advance()
		l.buf.WriteRune(r)
	}

	tokStr := l.buf.String()
	keyword, ok := keywords[tokStr]
	if !ok {
		l.errLogger.Errorf(l.file.Position(tok.Pos), "invalid keyword %v", tokStr)
		tok.Type = ERROR
		l.recover()
		return
	}
	tok.Type = keyword
	tok.Str = l.buf.String()
}

func (l *lex) scanComment(tok *Token) {
	l.advance()
	if l.peek() != '/' {
		l.errLogger.Errorf(
			l.file.Position(l.pos()), "unexpected character: %v", l.peek())
		l.recover()
		tok.Type = ERROR
		return
	}
	for l.peek() != '\n' {
		l.advance()
	}
}

func (l *lex) scanLiteral(tok *Token) {
	l.buf.Reset()

	// Skip the first '.
	l.advance()

	for {
		r := l.peek()
		switch r {
		case '\'':
			l.advance()
			tok.Str = l.buf.String()
			tok.Type = LITERAL
			return

		case '\\':
			l.advance()
			r = l.peek()
			switch r {
			case '\'', '\\':
				l.buf.WriteRune(r)
				l.advance()
			default:
				l.errLogger.Errorf(
					l.file.Position(l.pos()), "unexpected character %v in string literal", r)
				l.recover()
				tok.Type = ERROR
				return
			}

		case '\n':
			l.errLogger.Errorf(
				l.file.Position(l.pos()), "newline in string literal")
			// Don't recover. We will start parsing from the '\n'.
			tok.Type = ERROR
			return

		default:
			l.buf.WriteRune(r)
			l.advance()
		}
	}
}

func (l *lex) scanIdentifier(tok *Token) {
	l.buf.Reset()

	r := l.peek()
	l.advance()
	l.buf.WriteRune(r)

	for {
		r := l.peek()
		if !isLetter(r) && !isNumber(r) && r != '_' {
			break
		}
		l.advance()
		l.buf.WriteRune(r)
	}
	tok.Type = ID
	tok.Str = l.buf.String()
}

func (l *lex) scanNum(tok *Token) {
	l.buf.Reset()

	for isNumber(l.peek()) {
		l.buf.WriteRune(l.peek())
		l.advance()
	}

	tok.Type = NUM
	tok.Str = l.buf.String()
}
