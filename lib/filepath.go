package lib

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
)

func filepathAbs(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewString(abs)}, nil
}

func filepathBase(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathClean(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathDir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathEvalSymlinks(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}
	return []nitro.Value{nitro.NewString(res)}, nil
}

func filepathExt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathFromSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathIsAbs(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := filepath.IsAbs(path)
	return []nitro.Value{nitro.NewBool(res)}, nil
}

func filepathJoin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathMatch(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}

	pattern, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}

	res, err := doublestar.PathMatch(pattern, path)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(res)}, nil
}

func filepathRel(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	basePath, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	res, err := filepath.Rel(basePath, path)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(res)}, nil
}

func filepathSplit(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	list := filepath.SplitList(path)
	res := make([]nitro.Value, len(list))
	for i, part := range list {
		res[i] = nitro.NewString(part)
	}
	return []nitro.Value{nitro.NewArrayFromSlice(res)}, nil
}

func filepathToSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func filepathVolumeName(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	res := filepath.VolumeName(path)
	return []nitro.Value{nitro.NewString(res)}, nil
}
