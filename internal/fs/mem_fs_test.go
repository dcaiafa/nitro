package fs

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemFS(t *testing.T) {
	m := NewMem()

	m.PutCombined(`
--- a.txt ---
Hello
World
--- x/b.txt ---
This is fine.
--- x/c.txt ---
Here is
another
file.
--- x/y/d.txt ---
Last one.`)

	entries, err := m.List(".")
	require.NoError(t, err)

	expectedEntries := []DirEntry{
		{Name: "a.txt", IsDir: false},
		{Name: "x", IsDir: true},
	}
	require.Equal(t, expectedEntries, entries)

	entries, err = m.List("x")
	require.NoError(t, err)

	expectedEntries = []DirEntry{
		{Name: "b.txt", IsDir: false},
		{Name: "c.txt", IsDir: false},
		{Name: "y", IsDir: true},
	}
	require.Equal(t, expectedEntries, entries)

	entries, err = m.List(filepath.Join("x", "y"))
	require.NoError(t, err)

	expectedEntries = []DirEntry{
		{Name: "d.txt", IsDir: false},
	}
	require.Equal(t, expectedEntries, entries)

	_, err = m.List("foo")
	require.True(t, errors.Is(err, os.ErrNotExist))

	_, err = m.List("x/foo")
	require.True(t, errors.Is(err, os.ErrNotExist))

	data, err := m.Read("a.txt")
	require.NoError(t, err)
	require.Equal(t, "Hello\nWorld\n", string(data))

	data, err = m.Read("x/c.txt")
	require.NoError(t, err)
	require.Equal(t, "Here is\nanother\nfile.\n", string(data))
}
