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

type Module struct {
	Name string
	Path string
}

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

func ReadManifest(f fs.FS, rootDir string) (*Manifest, error) {
	manifestPath := filepath.Join(rootDir, ModuleManifestFilename)

	manifestData, err := f.Read(manifestPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Manifest{
				Module: "main",
			}, nil
		}
	}

	manifest, err := ParseManifest(manifestData)
	if err != nil {
		return nil, err
	}

	return manifest, nil
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
