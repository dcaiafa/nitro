package parser2

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/token"
)

type errorListener struct {
	antlr.DefaultErrorListener

	filename  string
	errLogger errlogger.ErrLogger
	err       error
}

func newErrorListener(filename string, errLogger errlogger.ErrLogger) *errorListener {
	return &errorListener{
		filename:  filename,
		errLogger: errLogger,
	}
}

func (l *errorListener) Error() error {
	return l.err
}

func (l *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	pos := token.Pos{
		Filename: l.filename,
		Line:     line,
		Col:      column,
	}
	l.errLogger.Failf(pos, "%s", msg)
	if l.err == nil {
		l.err = fmt.Errorf("%v: %v", pos.String(), msg)
	}
}
