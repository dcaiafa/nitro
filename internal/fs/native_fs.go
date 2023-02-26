package fs

import "os"

func NewNative() FS {
	return new(nativeFS)
}

type nativeFS struct{}

func (f *nativeFS) Stat(path string) (*FileInfo, error) {
	osFI, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	fi := &FileInfo{
		Name:  osFI.Name(),
		IsDir: osFI.IsDir(),
	}
	return fi, nil
}

func (f *nativeFS) List(path string) ([]DirEntry, error) {
	osEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	entries := make([]DirEntry, len(osEntries))
	for i, ose := range osEntries {
		entries[i] = DirEntry{
			Name:  ose.Name(),
			IsDir: ose.IsDir(),
		}
	}

	return entries, nil
}

func (f *nativeFS) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}
