package fs

type DirEntry struct {
	Name  string
	IsDir bool
}

type FileInfo struct {
	Name  string
	IsDir bool
}

type FS interface {
	Stat(path string) (*FileInfo, error)
	List(path string) ([]DirEntry, error)
	Read(path string) ([]byte, error)
}
