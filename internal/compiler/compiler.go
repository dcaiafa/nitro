package compiler

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/internal/mod"
	"github.com/dcaiafa/nitro/internal/vm"
)

var ErrCircularDependency = errors.New("circular dependency")

type Compiler struct {
	pkgs            []*vm.CompiledPackage
	pkgMap          map[string]int  // package_name => pkgs index
	inflight        map[string]bool // package_name
	builtins        map[string]export.Exports
	fs              fs.FS
	packageResolver *mod.PackageResolver
	errLogger       *errlogger.ErrLoggerWrapper
}

func New() *Compiler {
	c := &Compiler{
		pkgMap:          make(map[string]int),
		inflight:        make(map[string]bool),
		builtins:        make(map[string]export.Exports),
		fs:              fs.NewNative(),
		packageResolver: mod.NewPackageResolver(),
	}
	c.errLogger = errlogger.NewErrLoggerBase(
		&errlogger.ConsoleErrLogger{})
	return c
}

func (c *Compiler) RegisterBuiltins(pkgName string, exports export.Exports) {
	if _, ok := c.builtins[pkgName]; ok {
		panic("package already registered")
	}
	c.builtins[pkgName] = exports
}

func (c *Compiler) Compile(packagePath string) (*vm.Program, error) {
	root, modManifest, err := mod.Root(c.fs, packagePath)
	if err != nil {
		return nil, err
	}

	err = c.packageResolver.AddModule(&mod.Module{
		Name: modManifest.Module,
		Path: root,
	})
	if err != nil {
		return nil, err
	}

	packageName, err := filepath.Rel(root, packagePath)
	if err != nil {
		return nil, err
	}

	if packageName == "." {
		// The root package's name is the module's name.
		packageName = modManifest.Module
	} else {
		packageName = filepath.ToSlash(packageName)
		packageName = path.Join(modManifest.Module, packageName)
	}

	_, err = c.compilePackage(packageName, true)
	if err != nil {
		return nil, err
	}
	prog := &vm.Program{
		Packages: c.pkgs,
	}
	return prog, nil
}

func (c *Compiler) compilePackage(packageName string, isMain bool) (*vm.CompiledPackage, error) {
	if _, ok := c.pkgMap[packageName]; ok {
		panic("package already compiled")
	}
	if c.inflight[packageName] {
		return nil, fmt.Errorf("%w: %v", ErrCircularDependency, packageName)
	}

	packageInfo, err := c.packageResolver.ResolvePackage(packageName)
	if err != nil {
		return nil, err
	}

	c.inflight[packageName] = true

	pkgCompiler := packageCompiler{
		ModuleRoot:    packageInfo.Module.Path,
		ModuleName:    packageInfo.Module.Name,
		PackageName:   packageName,
		PackageGetter: c,
		PackagePath:   packageInfo.Path,
		IsMain:        true,
		ErrLogger:     c.errLogger,
		FS:            c.fs,
	}

	pkg, err := pkgCompiler.Compile()
	if err != nil {
		return nil, err
	}

	delete(c.inflight, packageName)
	c.pkgs = append(c.pkgs, pkg)
	pkg.Index = len(c.pkgs) - 1
	c.pkgMap[packageName] = pkg.Index

	return pkg, nil
}

func (c *Compiler) getPackage(packageName string) (*vm.CompiledPackage, error) {
	// Have we compiled this package already?
	pkgIndex, ok := c.pkgMap[packageName]
	if ok {
		return c.pkgs[pkgIndex], nil
	}

	// Is this a builtin package?
	if exports, ok := c.builtins[packageName]; ok {
		pkg := &vm.CompiledPackage{
			Name:     packageName,
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
		return pkg, nil
	}

	return c.compilePackage(packageName, false)
}

func (c *Compiler) getAllDeps() []*vm.CompiledPackage {
	return c.pkgs
}
