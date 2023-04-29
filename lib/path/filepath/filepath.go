package filepath

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/file"
)

//go:generate go run ../../../internal/stub/stubgen filepath.stubgen

func abs0(vm *vm.VM, path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func base0(vm *vm.VM, path string) (string, error) {
	return filepath.Base(path), nil
}

func clean0(vm *vm.VM, path string) (string, error) {
	return filepath.Clean(path), nil
}

func dir0(vm *vm.VM, path string) (string, error) {
	return filepath.Dir(path), nil
}

func eval_symlinks0(vm *vm.VM, path string) (string, error) {
	res, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return res, nil
}

func ext0(vm *vm.VM, path string) (string, error) {
	return filepath.Ext(path), nil
}

func from_slash0(vm *vm.VM, path string) (string, error) {
	return filepath.FromSlash(path), nil
}

func is_abs0(vm *vm.VM, path string) (bool, error) {
	return filepath.IsAbs(path), nil
}

func join0(vm *vm.VM, elems []string) (string, error) {
	return filepath.Join(elems...), nil
}

func match0(vm *vm.VM, path, pattern string) (bool, error) {
	res, err := doublestar.PathMatch(pattern, path)
	if err != nil {
		return false, err
	}
	return res, nil
}

func rel0(vm *vm.VM, basePath, targPath string) (string, error) {
	res, err := filepath.Rel(basePath, targPath)
	if err != nil {
		return "", err
	}
	return res, nil
}

func split_list0(_ *vm.VM, path string) (*vm.Array, error) {
	list := filepath.SplitList(path)
	res := make([]nitro.Value, len(list))
	for i, elem := range list {
		res[i] = nitro.NewString(elem)
	}
	return vm.NewArrayWithSlice(res), nil
}

func to_slash0(vm *vm.VM, path string) (string, error) {
	return filepath.ToSlash(path), nil
}

func volume_name0(vm *vm.VM, path string) (string, error) {
	return filepath.VolumeName(path), nil
}

func is_dir0(vm *vm.VM, path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return fi.IsDir(), nil
}

func rename0(vm *vm.VM, oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func exists0(vm *vm.VM, path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
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

func ls0(vm *vm.VM, path string) (vm.Iterator, error) {
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

		return nitro.NewIterator(iter.Next, nil, 2), nil
	}

	iter := newLSDoubleStarIter(base, pattern)
	return nitro.NewIterator(iter.Next, iter.Close, 2), nil
}

func remove0(vm *vm.VM, path string) (bool, error) {
	err := os.Remove(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func remove1(vm *vm.VM, f *file.File) (bool, error) {
	f.Close()
	err := os.Remove(f.Name())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func remove_all0(vm *vm.VM, path string) error {
	return os.RemoveAll(path)
}

func mkdir0(vm *vm.VM, path string) error {
	return os.Mkdir(path, 0777)
}

func mkdir_all0(vm *vm.VM, path string) error {
	return os.MkdirAll(path, 0777)
}

func mkdir_temp0(vm *vm.VM, pattern, dir string) (string, error) {
	tempDir, err := os.MkdirTemp(dir, pattern)
	if err != nil {
		return "", err
	}
	return tempDir, nil
}
