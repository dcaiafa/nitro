package mod

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dcaiafa/nitro/internal/fs"
	"gopkg.in/yaml.v3"
)

const ModuleManifestFilename = "bagl.mod"

var ErrInvalidManifest = errors.New("invalid manifest")
var ErrModuleRootNotFound = fmt.Errorf("not a in module: %v not in the current directory on in any parent directory", ModuleManifestFilename)

type VersionedModule struct {
	ModuleID string
	Version  Version
}

type Requirement struct {
	ModuleID string
	Version  Version
}

type Manifest struct {
	Module string `yaml:"module"`
}

func Root(f fs.FS, path string) (root string, manifest *Manifest, err error) {
	fileExists := func(p string) bool {
		fi, err := f.Stat(p)
		return err == nil && !fi.IsDir
	}

	root, err = filepath.Abs(path)
	if err != nil {
		return
	}

	var manPath string

	for {
		manPath = filepath.Join(root, ModuleManifestFilename)
		if fileExists(manPath) {
			break
		}

		newRoot := filepath.Dir(root)
		if newRoot == root {
			err = ErrModuleRootNotFound
			return
		}

		root = newRoot
	}

	manifestData, err := os.ReadFile(manPath)
	if err != nil {
		return
	}

	manifest, err = ParseManifest(manifestData)
	if err != nil {
		return
	}

	return root, manifest, nil
}

func ParseManifest(buf []byte) (*Manifest, error) {
	d := yaml.NewDecoder(bytes.NewReader(buf))
	d.KnownFields(true)

	m := new(Manifest)
	err := d.Decode(m)
	if err != nil {
		return nil, err
	}

	if m.Module == "" {
		return nil, fmt.Errorf("%w: missing 'module'", ErrInvalidManifest)
	}

	return m, nil
}
