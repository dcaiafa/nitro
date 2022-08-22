package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncRegistry(t *testing.T) {
	r := NewExportRegistry()
	for _, f := range allExports {
		nativeFn := r.GetExport(f.P, f.N)
		require.True(t, nativeFn != nil, f.P+"."+f.N)
	}
	require.True(t, r.GetExport("filepath", "map") == nil)
	require.True(t, r.IsValidPackage("filepath"))
	require.False(t, r.IsValidPackage("filepatho"))
}
