package mod

type VersionedModule struct {
	ModuleID string
	Version  Version
}

type Requirement struct {
	ModuleID string
	Version  Version
}

type ModuleFile struct {
	Name string
}
