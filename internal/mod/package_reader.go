package mod

import (
	"os"
	"path/filepath"
)

type PackageReader interface {
	Path() string
	ListUnits() ([]string, error)
	ReadUnit(unit string) ([]byte, error)
}

type NativePackageReader struct {
	path string
}

func NewNativePackageReader(path string) *NativePackageReader {
	return &NativePackageReader{
		path: path,
	}
}

func (r *NativePackageReader) Path() string {
	return r.path
}

func (r *NativePackageReader) ListUnits() ([]string, error) {
	allFiles, err := os.ReadDir(r.path)
	if err != nil {
		return nil, err
	}
	units := make([]string, 0, len(allFiles))
	for _, file := range allFiles {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".n" {
			units = append(units, filepath.Join(r.path, file.Name()))
		}
	}
	return units, nil
}

func (r *NativePackageReader) ReadUnit(unit string) ([]byte, error) {
	unitData, err := os.ReadFile(unit)
	if err != nil {
		return nil, err
	}
	return unitData, nil
}
