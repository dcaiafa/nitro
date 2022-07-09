package lib

import (
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
)

var errPathBaseUsage = nitro.NewInvalidUsageError("base(string)")

func pathBase(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathBaseUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathBaseUsage
	}

	base := filepath.Base(path.String())

	return []nitro.Value{nitro.NewString(base)}, nil
}

var errPathCleanUsage = nitro.NewInvalidUsageError("clean(string)")

func pathClean(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathCleanUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathCleanUsage
	}

	cleanPath := filepath.Clean(path.String())
	return []nitro.Value{nitro.NewString(cleanPath)}, nil
}

var errPathDirUsage = nitro.NewInvalidUsageError("dir(string)")

func pathDir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathDirUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathDirUsage
	}

	dir := filepath.Dir(path.String())
	return []nitro.Value{nitro.NewString(dir)}, nil
}

var errPathExtUsage = nitro.NewInvalidUsageError("ext(string)")

func pathExt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathExtUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathExtUsage
	}

	ext := filepath.Ext(path.String())
	return []nitro.Value{nitro.NewString(ext)}, nil
}

var errPathFromSlashUsage = nitro.NewInvalidUsageError("from_slash(string)")

func pathFromSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathFromSlashUsage
	}

	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathFromSlashUsage
	}

	fromSlash := filepath.FromSlash(path.String())
	return []nitro.Value{nitro.NewString(fromSlash)}, nil
}

var errPathToSlashUsage = nitro.NewInvalidUsageError("to_slash(string)")

func pathToSlash(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPathToSlashUsage
	}
	path, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathToSlashUsage
	}

	p := filepath.ToSlash(path.String())
	return []nitro.Value{nitro.NewString(p)}, nil
}

var errPathJoin = nitro.NewInvalidUsageError("join(string*)")

func pathJoin(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	paths := make([]string, len(args))

	for i := 0; i < len(args); i++ {
		path, ok := args[i].(nitro.String)
		if !ok {
			return nil, errPathJoin
		}
		paths[i] = path.String()
	}

	path := filepath.Join(paths...)
	return []nitro.Value{nitro.NewString(path)}, nil
}

var errPathMatchUsage = nitro.NewInvalidUsageError(`match(string, string)`)

func pathMatch(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 2 {
		return nil, errPathMatchUsage
	}

	pattern, ok := args[0].(nitro.String)
	if !ok {
		return nil, errPathMatchUsage
	}

	path, ok := args[1].(nitro.String)
	if !ok {
		return nil, errPathMatchUsage
	}

	res, err := doublestar.PathMatch(pattern.String(), path.String())
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(res)}, nil
}
