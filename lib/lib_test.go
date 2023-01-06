package lib

import (
	gosort "sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncRegistry(t *testing.T) {
	require.True(t, gosort.IsSorted(&Exports))

	for _, f := range Exports {
		index := Exports.GetExportIndex(f.P, f.N)
		require.True(t, index != -1, f.P+"."+f.N)
		v := Exports.GetExportValue(index)
		require.True(t, v != nil, f.P+"."+f.N)
	}

	require.True(t, Exports.GetExportIndex("filepath", "map") == -1)
	require.True(t, Exports.IsValidPackage("filepath"))
	require.False(t, Exports.IsValidPackage("filepatho"))
}
