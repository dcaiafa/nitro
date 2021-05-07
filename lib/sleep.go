package lib

import (
	"errors"
	"time"

	"github.com/dcaiafa/nitro"
)

var errSleepUsage = errors.New(
	`invalid usage. Expected sleep(duration|int)`)

func sleep(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errSleepUsage
	}

	var dur time.Duration

	if durArg, ok := args[0].(Duration); ok {
		dur = durArg.dur
	} else if intArg, ok := args[0].(nitro.Int); ok {
		dur = time.Duration(intArg.Int64()) * time.Second
	} else {
		return nil, errSleepUsage
	}

	time.Sleep(dur)
	return nil, nil
}
