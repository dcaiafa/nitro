package mod

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var vmodRegex = regexp.MustCompile(`^([A-Z]+)(\d+)\.(\d+)$`)

func parseVersionedModule(vmod string) VersionedModule {
	matches := vmodRegex.FindStringSubmatch(vmod)
	if matches == nil {
		panic("invalid VersionedModule")
	}

	atoi := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}

	return VersionedModule{
		ModuleID: matches[1],
		Version: Version{
			Major: atoi(matches[2]),
			Minor: atoi(matches[3]),
			Patch: 0,
		},
	}

}

type simpleGraph struct {
	g map[VersionedModule][]VersionedModule
}

func newSimpleGraph() *simpleGraph {
	return &simpleGraph{
		g: make(map[VersionedModule][]VersionedModule),
	}
}

func (g *simpleGraph) GetDependencies(vmod VersionedModule) ([]VersionedModule, error) {
	return g.g[vmod], nil
}

func (g *simpleGraph) Add(from, to string) {
	fromVMod := parseVersionedModule(from)
	toVMod := parseVersionedModule(to)
	g.g[fromVMod] = append(g.g[fromVMod], toVMod)
}

func expectBuildList(t *testing.T, bl BuildList, vmods ...string) {
	list := make([]string, 0, len(bl))
	for _, vmod := range bl.ToSlice() {
		list = append(
			list,
			fmt.Sprintf(
				"%v%v.%v", vmod.ModuleID, vmod.Version.Major, vmod.Version.Minor))
	}
	require.Equal(t, vmods, list)
}

func TestConstructBuildList(t *testing.T) {
	g := newSimpleGraph()
	g.Add("A1.0", "B1.2")
	g.Add("A1.0", "C1.2")
	g.Add("B1.1", "D1.1")
	g.Add("B1.2", "D1.3")
	g.Add("C1.2", "D1.4")
	g.Add("C1.3", "F1.1")
	g.Add("D1.1", "E1.1")
	g.Add("D1.2", "E1.2")
	g.Add("D1.3", "E1.2")
	g.Add("D1.4", "E1.2")
	g.Add("F1.1", "G1.1")
	g.Add("G1.1", "F1.1")

	bl, err := ConstructBuildList(g, parseVersionedModule("A1.0"))
	require.NoError(t, err)

	expectBuildList(
		t,
		bl,
		"A1.0",
		"B1.2",
		"C1.2",
		"D1.4",
		"E1.2")
}
