package compiler

import (
	"errors"
	"fmt"
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
	fs              fs.FS
	packageResolver *mod.PackageResolver
	errLogger       *errlogger.ErrLoggerWrapper
}

func New() *Compiler {
	c := &Compiler{
		pkgMap:          make(map[string]int),
		inflight:        make(map[string]bool),
		fs:              fs.NewNative(),
		packageResolver: mod.NewPackageResolver(),
	}
	c.errLogger = errlogger.NewErrLoggerBase(
		&errlogger.ConsoleErrLogger{})
	return c
}

func (c *Compiler) RegisterBuiltins(pkgName string, exports export.Exports) {
	if _, ok := c.pkgMap[pkgName]; ok {
		panic("built-in package already registered")
	}
	pkg := &vm.CompiledPackage{
		Name:     pkgName,
		Builtin:  true,
		Literals: make([]vm.Value, len(exports)),
		Symbols:  make(map[string]int, len(exports)),
	}
	for i, export := range exports {
		pkg.Literals[i] = export.Value()
		pkg.Symbols[export.N] = i
	}
	c.pkgs = append(c.pkgs, pkg)
	pkgIndex := len(c.pkgs) - 1
	c.pkgMap[pkgName] = pkgIndex
}

func (c *Compiler) Compile(programFile string) (*vm.Program, error) {
	root := filepath.Dir(programFile)

	manifest, err := mod.ReadManifest(c.fs, root)
	if err != nil {
		return nil, err
	}

	err = c.packageResolver.AddModule(&mod.Module{
		Name: manifest.Module,
		Path: root,
	})
	if err != nil {
		return nil, err
	}

	packageName := manifest.Module

	c.inflight[packageName] = true

	pkgCompiler := packageCompiler{
		ModuleRoot:    root,
		ModuleName:    manifest.Module,
		PackageName:   manifest.Module,
		PackageGetter: c,
		ProgramPath:   programFile,
		ErrLogger:     c.errLogger,
		FS:            c.fs,
	}

	pkg, err := pkgCompiler.Compile()
	if err != nil {
		return nil, err
	}

	c.pkgs = append(c.pkgs, pkg)
	pkg.Index = len(c.pkgs) - 1
	c.pkgMap[packageName] = pkg.Index

	prog := &vm.Program{
		Packages: c.pkgs,
	}

	return prog, nil
}

func (c *Compiler) compilePackage(packageName string) (*vm.CompiledPackage, error) {
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

	return c.compilePackage(packageName)
}

func (c *Compiler) getAllDeps() []*vm.CompiledPackage {
	return c.pkgs
}
