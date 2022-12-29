package mod

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

var ErrPackageNotFound = errors.New("package not found")

type FS struct {
	fs.FS
	Path string
}

func (f *FS) ExpandPath(path string) string {
	return filepath.Clean(
		filepath.Join(f.Path, filepath.FromSlash(path)))
}

func (f *FS) Sub(path string) *FS {
	subFS, err := fs.Sub(f.FS, path)
	if err != nil {
		panic(err)
	}
	return &FS{
		FS:   subFS,
		Path: filepath.Join(f.Path, path),
	}
}

type module struct {
	fs           *FS
	moduleName   string
	modulePrefix string
}

type PackageResolver struct {
	modules []*module
}

func NewPackageResolver() *PackageResolver {
	return &PackageResolver{}
}

func (d *PackageResolver) RegisterModule(moduleFS *FS, moduleName string) {
	d.modules = append(d.modules, &module{
		fs:           moduleFS,
		moduleName:   moduleName,
		modulePrefix: moduleName + "/",
	})
}

func (d *PackageResolver) GetPackage(packageName string) (*PackageReader, error) {
	mod, err := d.getModuleForPackage(packageName)
	if err != nil {
		return nil, err
	}

	packagePath := packageName[len(mod.modulePrefix):]

	fi, err := fs.Stat(mod.fs, packagePath)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		return &PackageReader{
			FS:          mod.fs.Sub(packagePath),
			module:      mod.moduleName,
			packageName: packageName,
			packagePath: path.Base(packagePath),
			isDir:       true,
		}, nil
	} else {
		return &PackageReader{
			FS:          mod.fs.Sub(path.Dir(packagePath)),
			module:      mod.moduleName,
			packageName: packageName,
			packagePath: path.Base(packagePath),
			isDir:       false,
		}, nil
	}
}

func (d *PackageResolver) getModuleForPackage(packageName string) (*module, error) {
	for _, mod := range d.modules {
		if strings.HasPrefix(packageName, mod.modulePrefix) {
			return mod, nil
		}
	}
	return nil, fmt.Errorf("%w: %v", ErrPackageNotFound, packageName)
}

type PackageReader struct {
	*FS
	module      string
	packageName string
	packagePath string
	isDir       bool
}

func (r *PackageReader) ListUnits() ([]string, error) {
	if !r.isDir {
		return []string{r.packagePath}, nil
	}

	allFiles, err := fs.ReadDir(r, r.packagePath)
	if err != nil {
		return nil, err
	}
	units := make([]string, 0, len(allFiles))
	for _, file := range allFiles {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".n" {
			units = append(units, path.Join(r.packagePath, file.Name()))
		}
	}
	return units, nil
}
