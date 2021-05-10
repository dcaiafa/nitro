package lib

import (
	"encoding/base64"
	"errors"

	"github.com/dcaiafa/nitro"
)

var errParseBase64Usage = errors.New(
	`invalid usage. Expected parsebase64(string)`)

func parsebase64(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errParseBase64Usage
	}
	enc, ok := args[0].(nitro.String)
	if !ok {
		return nil, errParseBase64Usage
	}
	dec, err := base64.StdEncoding.DecodeString(enc.String())
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewString(string(dec))}, nil
}

var errUsageToBase64 = errors.New(
	`invalid usage. Expected tobase64(string)`)

func tobase64(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errUsageToBase64
	}
	dec, ok := args[0].(nitro.String)
	if !ok {
		return nil, errUsageToBase64
	}
	enc := base64.StdEncoding.EncodeToString([]byte(dec.String()))
	return []nitro.Value{nitro.NewString(enc)}, nil
}
