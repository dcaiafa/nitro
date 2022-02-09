package imports

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

type ModuleRef struct {
	Name         string
	Version      string
	MajorVersion string
}

type ImportRef struct {
	Path   string
	Module ModuleRef
	Native bool
}

func ParseImportPath(importPath string) (ref ImportRef, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("invalid import path %q: %w", importPath, err)
		}
	}()

	comps := strings.Split(importPath, "/")

	if len(comps) == 1 || comps[0] != "github.com" {
		ref = ImportRef{
			Path:   importPath,
			Native: true,
		}
		return
	}

	if len(comps) < 3 {
		err = fmt.Errorf("invalid number of components")
		return
	}

	last := comps[len(comps)-1]

	atIdx := strings.Index(last, "@")
	if atIdx == -1 {
		err = fmt.Errorf("version is missing")
		return
	}

	version := last[atIdx+1:]

	if !semver.IsValid(version) {
		err = fmt.Errorf("version is not valid semver")
		return
	}

	comps[len(comps)-1] = last[:atIdx]

	ref = ImportRef{
		Path: strings.Join(comps, "/"),
		Module: ModuleRef{
			Module:       strings.Join(comps[:3], "/"),
			Version:      version,
			MajorVersion: semver.Major(version),
		},
		Native: false,
	}

	return
}

type Importer struct {
}
