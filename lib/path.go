package lib

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
)

func pathBase(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	base := filepath.Base(path)
	return []nitro.Value{nitro.NewString(base)}, nil
}

func pathClean(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	cleanPath := filepath.Clean(path)
	return []nitro.Value{nitro.NewString(cleanPath)}, nil
}

func pathDir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(path)
	return []nitro.Value{nitro.NewString(dir)}, nil
}

func pathExt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(path)
	return []nitro.Value{nitro.NewString(ext)}, nil
}

func pathFromSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	fromSlash := filepath.FromSlash(path)
	return []nitro.Value{nitro.NewString(fromSlash)}, nil
}

func pathToSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	p := filepath.ToSlash(path)
	return []nitro.Value{nitro.NewString(p)}, nil
}

func pathJoin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error
	paths := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		paths[i], err = getStringArg(args, i)
		if err != nil {
			return nil, err
		}
	}
	path := filepath.Join(paths...)
	return []nitro.Value{nitro.NewString(path)}, nil
}

func pathMatch(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	pattern, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	res, err := doublestar.PathMatch(pattern, path)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(res)}, nil
}
