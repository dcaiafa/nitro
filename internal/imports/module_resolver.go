package imports

type importLister interface {
	ListImports(modRef ModuleRef) ([]ImportRef, error)
}

type moduleMajorRef struct {
	Name         string
	MajorVersion string
}

type moduleResolverInfo struct {
	ModuleRef ModuleRef
	Imports   []moduleMajorRef
}

type ModuleSetResolver struct {
	importLister importLister
	modules      map[moduleMajorRef]*moduleResolverInfo
}

func NewModuleSetResolver(importLister importLister) *ModuleSetResolver {
	return &ModuleSetResolver{
		importLister: importLister,
		modules:      make(map[moduleMajorRef]*moduleResolverInfo),
	}
}

func (r *ModuleSetResolver) Resolve(imports []ImportRef) {
	for _, imp := range imports {
		majorRef := moduleMajorRef{
			Name:         imp.Module.Name,
			MajorVersion: imp.Module.MajorVersion,
		}
		mod := r.modules[majorRef]

	}
}
