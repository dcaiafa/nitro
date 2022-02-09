package imports

import (
	"path/filepath"

	fcopy "github.com/otiai10/copy"
)

type Fetcher interface {
	Fetch(mod ModuleRef, destPath string) error
}

type LocalFetcher struct {
	root string
}

func NewLocalFetcher(root string) *LocalFetcher {
	return &LocalFetcher{root: root}
}

func (f *LocalFetcher) Fetch(mod ModuleRef, destPath string) error {
	fromPath := filepath.Join(
		f.root,
		filepath.ToSlash(mod.Module),
		mod.Version)
	return fcopy.Copy(fromPath, destPath)
}
