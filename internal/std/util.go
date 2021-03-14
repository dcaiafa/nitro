package std

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/runtime"
)

func getIntArg(args []runtime.Value, ndx int) (int64, error) {
	if ndx >= len(args) {
		return 0, errNotEnoughArgs
	}
	v, ok := args[ndx].(runtime.Int)
	if !ok {
		return 0, fmt.Errorf(
			"expected argument %d to be Int, but it is %v",
			ndx, args[ndx].Type())
	}
	return v.Int64(), nil
}
