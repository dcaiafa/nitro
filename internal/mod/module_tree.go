package mod

import (
	"fmt"
	"path/filepath"
	"strings"
)

type moduleTreeNode struct {
	Name     string
	Children map[string]*moduleTreeNode
	Module   *Module
}

type moduleTree struct {
	root *moduleTreeNode
}

func newModuleTree() *moduleTree {
	return &moduleTree{
		root: &moduleTreeNode{},
	}
}

func (t *moduleTree) AddModule(info *Module) error {
	names := strings.Split(info.Name, "/")
	node, err := t.createBranches(t.root, names)
	if err != nil {
		return fmt.Errorf("couldn't add module %q: %w", info.Name, err)
	}
	if node.Module != nil {
		return fmt.Errorf("couldn't add module %q: already exists", info.Name)
	}
	node.Module = info
	return nil
}

func (t *moduleTree) FindModuleForPackage(packageName string) (module *Module, packagePath string) {
	names := strings.Split(packageName, "/")

	node, names := t.findNode(t.root, names)
	if node.Module == nil {
		return nil, ""
	}

	module = node.Module

	// Construct the package path inside the module directory.
	packagePath = filepath.Join(module.Path, filepath.FromSlash(strings.Join(names, "/")))

	return module, packagePath
}

func (t *moduleTree) findNode(root *moduleTreeNode, names []string) (*moduleTreeNode, []string) {
	if len(names) == 0 {
		return root, nil
	}
	name := names[0]
	if root.Children == nil {
		return root, names
	}
	node := root.Children[name]
	if node == nil {
		return root, names
	}
	return t.findNode(node, names[1:])
}

func (t *moduleTree) createBranches(root *moduleTreeNode, names []string) (*moduleTreeNode, error) {
	if len(names) == 0 {
		return root, nil
	}
	if root.Module != nil {
		return nil, fmt.Errorf("conflict with module %q", root.Module.Name)
	}
	name := names[0]
	node := root.Children[name]
	if node == nil {
		if root.Children == nil {
			root.Children = make(map[string]*moduleTreeNode)
		}
		node = &moduleTreeNode{
			Name: name,
		}
		root.Children[name] = node
	}
	return t.createBranches(node, names[1:])
}
