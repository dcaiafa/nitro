package mod

import "fmt"

type Package struct {
	Name   string
	Path   string
	Module *Module
}

type PackageResolver struct {
	modTree *moduleTree
}

func NewPackageResolver() *PackageResolver {
	return &PackageResolver{
		modTree: newModuleTree(),
	}
}

func (r *PackageResolver) AddModule(module *Module) error {
	return r.modTree.AddModule(module)
}

func (r *PackageResolver) ResolvePackage(packageName string) (*Package, error) {
	module, packagePath := r.modTree.FindModuleForPackage(packageName)
	if module == nil {
		return nil, fmt.Errorf("couldn't find module for package %q", packageName)
	}
	return &Package{
		Name:   packageName,
		Path:   packagePath,
		Module: module,
	}, nil
}
