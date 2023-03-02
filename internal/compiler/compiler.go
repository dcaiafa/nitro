package compiler

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/internal/mod"
	"github.com/dcaiafa/nitro/internal/vm"
)

var ErrCircularDependency = errors.New("circular dependency")

type Compiler struct {
	pkgs        []*vm.CompiledPackage
	pkgMap      map[string]int // package_name => pkgs index
	builtins    map[string]export.Exports
	fs          fs.FS
	packageRoot string
}

func New() (*Compiler, error) {
	c := &Compiler{
		pkgMap: make(map[string]int),
		fs:     fs.NewNative(),
	}
	return c, nil
}

func (c *Compiler) Compile(packagePath string) (*vm.Program, error) {
	_, err := c.compilePackage(packagePath)
	if err != nil {
		return nil, err
	}
	prog := &vm.Program{
		Packages: c.pkgs,
	}
	return prog, nil
}

func (c *Compiler) compilePackage(packagePath string) (*vm.CompiledPackage, error) {
	root, modManifest, err := mod.Root(c.fs, packagePath)
	if err != nil {
		return nil, err
	}

	packageName, err := filepath.Rel(root, packagePath)
	if err != nil {
		return nil, err
	}
	packageName = filepath.ToSlash(packageName)
	packageName = path.Join(modManifest.Module, packageName)

	if _, ok := c.pkgMap[packageName]; ok {
		panic("package already compiled")
	}

	c.pkgs = append(c.pkgs, nil)
	index := len(c.pkgs) - 1
	c.pkgMap[packageName] = index

	pkg, err := compilePackage(
		&compilePackageOptions{
			Module:           modManifest.Module,
			ModuleRoot:       root,
			PackageName:      packageName,
			PackageDirectory: c,
		})

	if err != nil {
		return nil, err
	}

	c.pkgs[index] = pkg

	return pkg, nil
}

func (c *Compiler) getPackage(packageName string) (*vm.CompiledPackage, error) {
	// Have we compiled this package already?
	pkgIndex, ok := c.pkgMap[packageName]
	if ok {
		pkg := c.pkgs[pkgIndex]
		if pkg == nil {
			return nil, fmt.Errorf("%w: %v", ErrCircularDependency, packageName)
		}
		return pkg, nil
	}

	// Is this a builtin package?
	if exports, ok := c.builtins[packageName]; ok {
		pkg := &vm.CompiledPackage{
			Literals: make([]vm.Value, len(exports)),
			Symbols:  make(map[string]int, len(exports)),
		}
		for i, export := range exports {
			pkg.Literals[i] = export.Value()
			pkg.Symbols[export.N] = i
		}
		c.pkgs = append(c.pkgs, pkg)
		c.pkgMap[packageName] = len(c.pkgs) - 1
		return pkg, nil
	}

	// This must be a regular Bagl package.
	// Compile and cache it.
	packagePath := filepath.Join(
		c.packageRoot, filepath.FromSlash(packageName))

	return c.compilePackage(packagePath)
}

/*
func Compile(packagePath string, opts ...Option) (*vm.Program, error) {
	compiler, err := newCompiler(opts...)
	if err != nil {
		return nil, err
	}
	return compiler.Compile(packagePath)
}

func CompileSimple(filePath string) (*vm.Program, error) {
	panic("notimpl")
}

func CompileInline(script []byte) (*vm.Program, error) {
	panic("notimpl")
}
*/
