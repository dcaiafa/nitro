package fs

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type node struct {
	Name     string
	IsDir    bool
	Data     []byte
	Children map[string]*node
}

type MemFS struct {
	root *node
}

func NewMem() *MemFS {
	m := &MemFS{}
	m.root = &node{
		Name:     ".",
		IsDir:    true,
		Children: make(map[string]*node),
	}
	return m
}

func (m *MemFS) mkdirAll(path string) *node {
	path = filepath.Clean(path)
	if path == "." {
		return m.root
	}

	parts := strings.Split(path, string(filepath.Separator))

	parent := m.root

	for len(parts) != 0 {
		if !parent.IsDir {
			panic(os.ErrInvalid)
		}
		name := parts[0]
		parts = parts[1:]

		child := parent.Children[name]
		if child == nil {
			child = &node{
				Name:     name,
				IsDir:    true,
				Children: make(map[string]*node),
			}
			parent.Children[child.Name] = child
		}
		parent = child
	}

	return parent
}

func (m *MemFS) Put(path string, data []byte) {
	dir, name := filepath.Split(path)
	parent := m.mkdirAll(dir)
	parent.Children[name] = &node{
		Name:  name,
		IsDir: false,
		Data:  []byte(data),
	}
}

var regexCombinedSep = regexp.MustCompile(`^--- (.*) ---$`)

func (m *MemFS) PutCombined(all string) {
	var name string
	var data []byte
	scanner := bufio.NewScanner(strings.NewReader(all))
	for scanner.Scan() {
		l := scanner.Bytes()
		match := regexCombinedSep.FindSubmatch(l)
		if match != nil {
			if name != "" {
				m.Put(name, data)
			}
			name = filepath.FromSlash(string(match[1]))
			data = nil
		} else {
			data = append(data, l...)
			data = append(data, '\n')
		}
	}
	m.Put(name, data)
}

func (m *MemFS) Stat(path string) (*FileInfo, error) {
	node := m.getNode(path)
	if node == nil {
		return nil, os.ErrNotExist
	}
	return &FileInfo{
		Name:  node.Name,
		IsDir: node.IsDir,
	}, nil
}

func (m *MemFS) List(path string) ([]DirEntry, error) {
	node := m.getNode(path)
	if node == nil {
		return nil, os.ErrNotExist
	}
	if !node.IsDir {
		return nil, os.ErrInvalid
	}
	entries := make([]DirEntry, 0, len(node.Children))
	for _, child := range node.Children {
		entries = append(entries, DirEntry{
			Name:  child.Name,
			IsDir: child.IsDir,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
	return entries, nil
}

func (m *MemFS) Read(path string) ([]byte, error) {
	node := m.getNode(path)
	if node == nil {
		return nil, os.ErrNotExist
	}
	if node.IsDir {
		return nil, os.ErrInvalid
	}
	b := make([]byte, len(node.Data))
	copy(b, node.Data)
	return b, nil
}

func (m *MemFS) getNode(path string) *node {
	path = filepath.Clean(path)
	if path == "." {
		return m.root
	}

	parts := strings.Split(path, string(filepath.Separator))

	node := m.root
	for len(parts) != 0 {
		if !node.IsDir {
			return nil
		}
		name := parts[0]
		parts = parts[1:]
		node = node.Children[name]
		if node == nil {
			return nil
		}
	}

	return node
}
