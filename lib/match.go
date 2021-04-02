package lib

import (
	"context"
	"fmt"

	"github.com/dcaiafa/nitro"
)

func fnMatch(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	regex, err := getRegexArg(args, 1)
	if err != nil {
		return nil, err
	}

	fmt.Println(regex.String())

	if retN == 1 {
		ok := regex.MatchString(str)
		return []nitro.Value{nitro.NewBool(ok)}, nil
	} else {
		matches := regex.FindStringSubmatch(str)
		if matches == nil {
			return []nitro.Value{nitro.NewBool(false), nil}, nil
		}
		a := nitro.NewArrayWithCapacity(len(matches))
		for _, match := range matches {
			a.Append(nitro.NewString(match))
		}
		return []nitro.Value{nitro.NewBool(true), a}, nil
	}
}
