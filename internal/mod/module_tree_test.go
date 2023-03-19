package mod

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModuleTree_FindModuleForPackage(t *testing.T) {
	tree := newModuleTree()
	modFoobar := &Module{
		Name: "github.com/dcaiafa/foobar",
		Path: "/mods/github.com!dcaiafa!foobar",
	}
	modFoobaz := &Module{
		Name: "github.com/dcaiafa/others/foobaz",
		Path: "/mods/github.com!dcaiafa!others!foobaz",
	}
	modTestor := &Module{
		Name: "testor",
		Path: "/local/testor",
	}

	err := tree.AddModule(modFoobar)
	require.NoError(t, err)
	err = tree.AddModule(modFoobaz)
	require.NoError(t, err)
	err = tree.AddModule(modTestor)
	require.NoError(t, err)

	testFind := func(packageName string, expectedMod *Module, expectedPath string) {
		t.Run(packageName, func(t *testing.T) {
			mod, path := tree.FindModuleForPackage(packageName)
			if expectedMod == nil {
				require.Nil(t, mod)
			} else {
				require.Equal(t, expectedMod, mod)
				require.Equal(t, filepath.ToSlash(expectedPath), path)
			}
		})
	}

	testFind("github.com/dcaiafa/foobar", modFoobar, "/mods/github.com!dcaiafa!foobar")
	testFind("github.com/dcaiafa/foobar/util/hashtable", modFoobar, "/mods/github.com!dcaiafa!foobar/util/hashtable")
	testFind("github.com/dcaiafa", nil, "")
	testFind("github.com/dcaiafa/others", nil, "")
	testFind("github.com/dcaiafa/others/foobaz", modFoobaz, "/mods/github.com!dcaiafa!others!foobaz")
	testFind("github.com/dcaiafa/others/foobaz/util/hashtable", modFoobaz, "/mods/github.com!dcaiafa!others!foobaz/util/hashtable")
}
