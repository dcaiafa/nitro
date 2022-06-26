package mod

import "sort"

type BuildListGraph interface {
	GetDependencies(vmod VersionedModule) ([]VersionedModule, error)
}

type BuildList map[VersionedModule]bool

func (s BuildList) ToSlice() []VersionedModule {
	res := make([]VersionedModule, 0, len(s))
	for vmod := range s {
		res = append(res, vmod)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ModuleID < res[j].ModuleID
	})
	return res
}

func ConstructBuildList(g BuildListGraph, root VersionedModule) (BuildList, error) {
	buildList := make(BuildList)
	buildList[root] = true
	err := constructRoughBuildList(g, root, buildList)
	if err != nil {
		return nil, err
	}
	simplifyBuildList(&buildList)
	return buildList, nil
}

func constructRoughBuildList(g BuildListGraph, vmod VersionedModule, set BuildList) error {
	deps, err := g.GetDependencies(vmod)
	if err != nil {
		return err
	}
	for _, dep := range deps {
		if !set[dep] {
			set[dep] = true
			err = constructRoughBuildList(g, dep, set)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func simplifyBuildList(bl *BuildList) {
	m := make(map[string]VersionedModule, len(*bl))
	for vmod := range *bl {
		if greatest, ok := m[vmod.ModuleID]; !ok || vmod.Version.Compare(greatest.Version) > 0 {
			m[vmod.ModuleID] = vmod
		}
	}
	*bl = make(BuildList, len(m))
	for _, vmod := range m {
		(*bl)[vmod] = true
	}
}
