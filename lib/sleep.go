package lib

import (
	"context"
	"errors"
	"time"

	"github.com/dcaiafa/nitro"
	btime "github.com/dcaiafa/nitro/lib/time"
)

var errSleepUsage = errors.New(
	`invalid usage. Expected sleep(duration|int)`)

func sleep(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errSleepUsage
	}

	var dur time.Duration

	if durArg, ok := args[0].(btime.Duration); ok {
		dur = durArg.Duration()
	} else if intArg, ok := args[0].(nitro.Int); ok {
		dur = time.Duration(intArg.Int64()) * time.Second
	} else {
		return nil, errSleepUsage
	}

	timer := time.NewTimer(dur)
	defer timer.Stop()

	var err error
	m.Block(func(ctx context.Context) {
		select {
		case <-timer.C:
		case <-ctx.Done():
			err = ctx.Err()
		}
	})

	return nil, err
}
