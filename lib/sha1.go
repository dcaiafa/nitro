package lib

import (
	cryptosha1 "crypto/sha1"
	"errors"
	"io"

	"github.com/dcaiafa/nitro"
)

var errSha1Usage = errors.New(
	`invalid usage. Expected sha1(reader)`)

func sha1(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errSha1Usage
	}

	input, err := nitro.MakeReader(vm, args[0])
	if err != nil {
		return nil, errSha1Usage
	}

	h := cryptosha1.New()
	_, err = io.Copy(h, input)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(h.Sum(nil)))}, nil
}
