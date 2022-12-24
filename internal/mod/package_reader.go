package mod

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

type ModuleReader struct {
	rootFS        fs.FS
	rootPath      string
	moduleName    string
	packagePrefix string
}

func NewModuleReader(rootFS fs.FS, rootPath string) (*ModuleReader, error) {
	manifestData, err := fs.ReadFile(rootFS, ModuleManifestFilename)
	if err != nil {
		return nil, err
	}
	manifest, err := ParseModuleManifest(manifestData)
	if err != nil {
		return nil, err
	}

	// TODO: validate manifest.
	mr := &ModuleReader{
		rootFS:        rootFS,
		rootPath:      rootPath,
		moduleName:    manifest.Module,
		packagePrefix: manifest.Module + "/",
	}

	return mr, nil
}

func (mr *ModuleReader) GetPackage(packageName string) (*PackageReader, error) {
	if !strings.HasPrefix(packageName, mr.packagePrefix) {
		return nil, fmt.Errorf(
			"package %s does not belong to module %s",
			packageName, mr.moduleName)
	}
	packagePath := packageName[len(mr.packagePrefix):]
	if packagePath == "" {
		return nil, fmt.Errorf(
			"package path without module prefix is empty")
	}
	fi, err := fs.Stat(mr.rootFS, packagePath)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("package %s is not a directory", packageName)
	}
	return &PackageReader{
		rootFS:      mr.rootFS,
		packagePath: packagePath,
	}, nil
}

func (mr *ModuleReader) ToLocalPath(path string) string {
	return filepath.Join(mr.rootPath, filepath.FromSlash(path))
}

type PackageReader struct {
	rootFS      fs.FS
	packagePath string
}

func (r *PackageReader) ListUnits() ([]string, error) {
	allFiles, err := fs.ReadDir(r.rootFS, r.packagePath)
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

func (r *PackageReader) ReadUnit(unitPath string) ([]byte, error) {
	unitData, err := fs.ReadFile(r.rootFS, unitPath)
	if err != nil {
		return nil, err
	}
	return unitData, nil
}
