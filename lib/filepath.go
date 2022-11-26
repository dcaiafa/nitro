package lib

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
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

func filepathIsDir(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	fi, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.NewBool(false)}, nil
		}
		return nil, err
	}
	return []nitro.Value{nitro.NewBool(fi.IsDir())}, nil
}

func filepathRename(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 2, 2); err != nil {
		return nil, err
	}
	oldPath, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	newPath, err := getStringArg(args, 1)
	if err != nil {
		return nil, err
	}
	err = os.Rename(oldPath, newPath)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func filepathExists(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	_, err = os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.NewBool(false)}, nil
		}
		return nil, err
	}

	return []nitro.Value{nitro.NewBool(true)}, nil
}

type lsDoubleStarIterEntry struct {
	path     string
	dirEntry fs.DirEntry
}

type lsDoubleStarIter struct {
	base    string
	pattern string
	outChan chan *lsDoubleStarIterEntry
	wg      sync.WaitGroup
	cancel  context.CancelFunc
}

func newLSDoubleStarIter(base, pattern string) *lsDoubleStarIter {
	i := &lsDoubleStarIter{
		base:    base,
		pattern: pattern,
		outChan: make(chan *lsDoubleStarIterEntry, 100),
	}

	var ctx context.Context
	ctx, i.cancel = context.WithCancel(context.Background())
	i.wg.Add(1)
	go i.run(ctx)

	return i
}

func (i *lsDoubleStarIter) run(ctx context.Context) {
	defer i.wg.Done()
	defer close(i.outChan)

	fsys := os.DirFS(i.base)
	doublestar.GlobWalk(fsys, i.pattern, func(path string, d fs.DirEntry) error {
		if d.Name() == "." || d.Name() == ".." {
			return ctx.Err()
		}

		select {
		case i.outChan <- &lsDoubleStarIterEntry{path, d}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
}

func (i *lsDoubleStarIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	entry, ok := <-i.outChan
	if !ok {
		i.cancel()
		return nil, nil
	}

	return []nitro.Value{
		nitro.NewString(filepath.FromSlash(filepath.Join(i.base, entry.path))),
		nitro.NewBool(entry.dirEntry.IsDir())}, nil
}

func (i *lsDoubleStarIter) Close(vm *nitro.VM) error {
	i.cancel()
	return nil
}

type lsSimpleIter struct {
	root    string
	entries []fs.DirEntry
}

func (i *lsSimpleIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(i.entries) == 0 {
		return nil, nil
	}

	res := []nitro.Value{
		nitro.NewString(filepath.Join(i.root, i.entries[0].Name())),
		nitro.NewBool(i.entries[0].IsDir())}

	i.entries = i.entries[1:]
	return res, nil
}

func filepathLs(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	base, pattern := doublestar.SplitPattern(filepath.ToSlash(path))
	if pattern == "" || pattern == "." {
		// This is a simple path. I.e. it does not include a pattern. Using
		// os.ReadDir is simpler/faster/leaner (no goroutines).
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		iter := &lsSimpleIter{
			root:    path,
			entries: entries,
		}

		return []nitro.Value{nitro.NewIterator(iter.Next, nil, 2)}, nil
	}

	iter := newLSDoubleStarIter(base, pattern)
	return []nitro.Value{nitro.NewIterator(iter.Next, iter.Close, 2)}, nil
}

func filepathRemove(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	var err error
	switch arg := args[0].(type) {
	case nitro.String:
		err = os.Remove(arg.String())

	case *File:
		arg.Close()
		err = os.Remove(arg.Name())

	default:
		return nil, errExpectedArg(0, arg, "str", "file")
	}

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []nitro.Value{nitro.False}, nil
		}
		return nil, err
	}

	return []nitro.Value{nitro.True}, nil
}

func filepathRemoveAll(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}

	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	err = os.RemoveAll(path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func filepathMkdir(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(path, 0777)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func filepathMkdirAll(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 1, 1); err != nil {
		return nil, err
	}
	path, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type mkdirTempOptions struct {
	Dir     string `nitro:"dir"`
	Pattern string `nitro:"pattern"`
}

var mkdirTempOptionsConv core.Value2Structer

func filepathMkdirTemp(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if err := expectArgCount(args, 0, 1); err != nil {
		return nil, err
	}

	var opt mkdirTempOptions
	if len(args) > 0 {
		err := mkdirTempOptionsConv.Convert(args[0], &opt)
		if err != nil {
			return nil, err
		}
	}

	tempDir, err := os.MkdirTemp(opt.Dir, opt.Pattern)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(tempDir)}, nil
}
