package errlogger

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/token"
)

type ErrLogger interface {
	Failf(pos token.Pos, msg string, args ...interface{})
	Detailf(pos token.Pos, msg string, args ...interface{})
}

type ErrLoggerWrapper struct {
	l   ErrLogger
	err error
}

func NewErrLoggerBase(l ErrLogger) *ErrLoggerWrapper {
	return &ErrLoggerWrapper{l: l}
}

func (l *ErrLoggerWrapper) Failf(pos token.Pos, msg string, args ...interface{}) {
	if l.err == nil {
		l.err = fmt.Errorf(pos.String()+": "+msg, args...)
	}
	l.l.Failf(pos, msg, args...)
}

func (l *ErrLoggerWrapper) Detailf(pos token.Pos, msg string, args ...interface{}) {
	l.l.Detailf(pos, msg, args...)
}

func (l *ErrLoggerWrapper) Error() error {
	return l.err
}
