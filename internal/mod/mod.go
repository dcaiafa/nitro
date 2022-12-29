package mod

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const ModuleManifestFilename = "bagl.mod"

var ErrInvalidManifest = errors.New("invalid manifest")

type VersionedModule struct {
	ModuleID string
	Version  Version
}

type Requirement struct {
	ModuleID string
	Version  Version
}

type ModuleManifest struct {
	Module string `yaml:"module"`
}

func FindModuleRoot(from string) (string, error) {
	exists := func(p string) bool {
		_, err := os.Stat(p)
		if err == nil {
			return true
		}
	}

	root, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}

	for {
		manifestPath := filepath.Join(root, ModuleManifestFilename)
		_, err := os.Stat(manifestPath)
		if err == nil {
			return root, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
	}
}

func ParseModuleManifest(buf []byte) (*ModuleManifest, error) {
	d := yaml.NewDecoder(bytes.NewReader(buf))
	d.KnownFields(true)

	m := new(ModuleManifest)
	err := d.Decode(m)
	if err != nil {
		return nil, err
	}

	if m.Module == "" {
		return nil, fmt.Errorf("%w: missing 'module'", ErrInvalidManifest)
	}

	return m, nil
}
