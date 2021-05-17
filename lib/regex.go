package lib

import (
	"errors"
	"regexp"

	"github.com/dcaiafa/nitro"
)

var errRegexUsage = errors.New(
	`invalid usage. Expected regex(string)`)

func regex(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errRegexUsage
	}

	reStr, ok := args[0].(nitro.String)
	if !ok {
		return nil, errRegexUsage
	}

	re, err := regexp.Compile(reStr.String())
	if err != nil {
		return nil, errRegexUsage
	}

	return []nitro.Value{nitro.NewRegex(re)}, nil
}
