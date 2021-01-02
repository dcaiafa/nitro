package errlogger

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/token"
)

type ErrLogger interface {
	Failf(pos token.Pos, msg string, args ...interface{})
	Detailf(pos token.Pos, msg string, args ...interface{})
}

type ErrLoggerBase struct {
	l   ErrLogger
	err error
}

func NewErrLoggerBase(l ErrLogger) *ErrLoggerBase {
	return &ErrLoggerBase{l: l}
}

func (l *ErrLoggerBase) Failf(pos token.Pos, msg string, args ...interface{}) {
	if l.err == nil {
		l.err = fmt.Errorf(pos.String()+": "+msg, args...)
	}
	l.l.Failf(pos, msg, args...)
}

func (l *ErrLoggerBase) Detailf(pos token.Pos, msg string, args ...interface{}) {
	l.l.Detailf(pos, msg, args...)
}

func (l *ErrLoggerBase) Error() error {
	return l.err
}
