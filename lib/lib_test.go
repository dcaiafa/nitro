package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncRegistry(t *testing.T) {
	r := NewFuncRegistry()
	for _, f := range allFuncs {
		nativeFn := r.GetNativeFn(f.Package, f.Name)
		require.True(t, nativeFn != nil)
	}
	require.True(t, r.GetNativeFn("path", "map") == nil)
	require.True(t, r.IsValidPackage("path"))
	require.False(t, r.IsValidPackage("patho"))
}
