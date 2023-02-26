package compiler

import (
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Option func(*compiler)

type packageCompiler interface {
}

type compiler struct {
	pkgs   []*vm.CompiledPackage
	pkgMap map[string]int // package_name => pkgs index
	fs     fs.FS
}

func newCompiler(opts ...Option) (*compiler, error) {
	c := new(compiler)
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func Compile(packagePath string) (*vm.Program, error) {
	panic("notimpl")
}

func CompileSimple(filePath string) (*vm.Program, error) {
	panic("notimpl")
}

func CompileInline(script []byte) (*vm.Program, error) {
	panic("notimpl")
}
