package lib

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/dcaiafa/nitro"
)

func readall(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	var input io.Reader
	switch arg := args[0].(type) {
	case *reader:
		input = arg

	case nitro.String:
		f, err := os.Open(arg.String())
		if err != nil {
			return nil, err
		}
		input = f

	default:
		return nil, fmt.Errorf(
			"invalid argument %q. Expected Reader or String.",
			nitro.TypeName(arg))
	}

	defer CloseReader(input)

	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}
