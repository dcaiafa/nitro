package lib

import (
	"errors"
	"path/filepath"

	"github.com/dcaiafa/nitro"
)

func pathbase(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	base := filepath.Base(path)

	return []nitro.Value{nitro.NewString(base)}, nil
}

func pathclean(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	cleanPath := filepath.Clean(path)

	return []nitro.Value{nitro.NewString(cleanPath)}, nil
}

func pathdir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)

	return []nitro.Value{nitro.NewString(dir)}, nil
}

func pathext(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)

	return []nitro.Value{nitro.NewString(ext)}, nil
}

func pathfromslash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	fromSlash := filepath.FromSlash(path)

	return []nitro.Value{nitro.NewString(fromSlash)}, nil
}

func pathjoin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	paths := make([]string, len(args))

	var err error
	for i := 0; i < len(args); i++ {
		paths[i], err = getStringArg(args, i)
		if err != nil {
			return nil, err
		}
	}

	path := filepath.Join(paths...)
	return []nitro.Value{nitro.NewString(path)}, nil
}

var errPathMatchUsage = errors.New(
	`invalid usage. Expected pathmatch(string, string)`)

func pathmatch(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errPathMatchUsage
	}

	name, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathMatchUsage
	}
	pattern, ok := args[1].(nitro.String)
	if !ok {
		return nil, errPathMatchUsage
	}

	res, err := filepath.Match(pattern.String(), name.String())
	if err != nil {
		return nil, errPathMatchUsage
	}

	return []nitro.Value{nitro.NewBool(res)}, nil
}
