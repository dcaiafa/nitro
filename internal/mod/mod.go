package mod

import (
	"bytes"
	"errors"
	"fmt"

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
