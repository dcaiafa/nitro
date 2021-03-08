package errlogger

import (
	"fmt"
	"os"

	"github.com/dcaiafa/nitro/internal/token"
)

type ConsoleErrLogger struct{}

func (l *ConsoleErrLogger) Failf(pos token.Pos, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pos.String()+": "+msg+"\n", args...)
}

func (l *ConsoleErrLogger) Detailf(pos token.Pos, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pos.String()+": "+msg+"\n", args...)
}
