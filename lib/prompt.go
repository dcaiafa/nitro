package lib

import (
	"errors"
	"fmt"

	"github.com/dcaiafa/nitro"
)

var errPromptUsage = errors.New(
	`invalid usage. Expected prompt(string?)`)

func prompt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 && len(args) != 1 {
		return nil, errPromptUsage
	}

	if len(args) == 1 {
		promptMessage, ok := args[0].(nitro.String)
		if !ok {
			return nil, errPromptUsage
		}

		fmt.Fprint(Stdout(m), promptMessage.String()+" ")
	}

	var resp string
	_, err := fmt.Scanln(&resp)
	if err != nil {
		resp = ""
	}

	return []nitro.Value{nitro.NewString(resp)}, nil
}
