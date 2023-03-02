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

func (c *Compiler) RegisterBuiltins(pkgName string, exports export.Exports) {
	if _, ok := c.builtins[pkgName]; ok {
		panic("package already registered")
	}
	c.builtins[pkgName] = exports
}

func (c *Compiler) Compile(pkgPath string) (*vm.Program, error) {
	_, err := c.compilePackage(pkgPath)
	if err != nil {
		return nil, err
	}
	prog := &vm.Program{
		Packages: c.pkgs,
	}
	return prog, nil
}

func (c *Compiler) compilePackage(pkgPath string) (*vm.CompiledPackage, int, error) {
	root, modManifest, err := mod.Root(c.fs, pkgPath)
	if err != nil {
		return nil, 0, err
	}

	packageName, err := filepath.Rel(root, pkgPath)
	if err != nil {
		return nil, 0, err
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
		return nil, 0, err
	}

	c.pkgs[index] = pkg

	return pkg, index, nil
}

func (c *Compiler) getPackage(packageName string) (*vm.CompiledPackage, int, error) {
	// Have we compiled this package already?
	pkgIndex, ok := c.pkgMap[packageName]
	if ok {
		pkg := c.pkgs[pkgIndex]
		if pkg == nil {
			return nil, 0, fmt.Errorf("%w: %v", ErrCircularDependency, packageName)
		}
		return pkg, pkgIndex, nil
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
		pkgIndex := len(c.pkgs) - 1
		c.pkgMap[packageName] = pkgIndex
		return pkg, pkgIndex, nil
	}

	// Just a plain, run-of-the-mill Bagl package.
	// Compile and cache it.
	pkgPath := filepath.Join(
		c.packageRoot, filepath.FromSlash(packageName))

	return c.compilePackage(pkgPath)
}

/*
func Compile(pkgPath string, opts ...Option) (*vm.Program, error) {
	compiler, err := newCompiler(opts...)
	if err != nil {
		return nil, err
	}
	return compiler.Compile(pkgPath)
}

func CompileSimple(filePath string) (*vm.Program, error) {
	panic("notimpl")
}

func CompileInline(script []byte) (*vm.Program, error) {
	panic("notimpl")
}
*/
