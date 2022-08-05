package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncRegistry(t *testing.T) {
	r := NewExportRegistry()
	for _, f := range allExports {
		nativeFn := r.GetExport(f.P, f.N)
		require.True(t, nativeFn != nil)
	}
	require.True(t, r.GetExport("path", "map") == nil)
	require.True(t, r.IsValidPackage("path"))
	require.False(t, r.IsValidPackage("patho"))
}
