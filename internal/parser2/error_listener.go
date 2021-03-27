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

func (l *errorListener) LogError(line, col int, msg string, args ...interface{}) {
	pos := token.Pos{
		Filename: l.filename,
		Line:     line,
		Col:      col,
	}
	renderedMsg := fmt.Sprintf(msg, args...)
	l.errLogger.Failf(pos, "%s", renderedMsg)
	if l.err == nil {
		l.err = fmt.Errorf("%v: %v", pos.String(), renderedMsg)
	}
}

func (l *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.LogError(line, column, "%s", msg)
}
