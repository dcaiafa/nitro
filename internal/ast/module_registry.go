package ast

import (
	"fmt"

	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type NativeModuleContext interface {
	RegisterNativeFn(name string, natFn func(m *vm.VM, args []vm.Value, nRet int) ([]vm.Value, error))
}

type ModuleRegistry struct {
	modules       map[string]*symbol.Module
	nativeLoaders map[string]func(r NativeModuleContext)
}

func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules:       make(map[string]*symbol.Module),
		nativeLoaders: make(map[string]func(r NativeModuleContext)),
	}
}

func (r *ModuleRegistry) RegisterNativeModuleLoader(
	name string,
	regFn func(r NativeModuleContext),
) error {
	if r.nativeLoaders[name] != nil {
		return fmt.Errorf("duplicate native module %v", name)
	}
	r.nativeLoaders[name] = regFn
	return nil
}

func (r *ModuleRegistry) GetNativeLoader(name string) func(r NativeModuleContext) {
	return r.nativeLoaders[name]
}

func (r *ModuleRegistry) AddModule(name string, m *symbol.Module) {
	if r.modules[name] != nil {
		panic("duplicate module")
	}
	r.modules[name] = m
}

func (r *ModuleRegistry) GetModule(name string) *symbol.Module {
	return r.modules[name]
}
