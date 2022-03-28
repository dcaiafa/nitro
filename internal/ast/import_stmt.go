package ast

import (
	"path"

	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Import struct {
	astBase

	Alias      string
	ModuleName string

	ModuleRef *symbol.ModuleRef
}

func (i *Import) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Rewrite:
		reg := ctx.ModuleRegistry()
		module := reg.GetModule(i.ModuleName)
		if module == nil {
			module = i.loadModule(ctx)
			if module == nil {
				return
			}
		}
		alias := i.Alias
		if alias == "" {
			alias = path.Base(i.ModuleName)
		}

		i.ModuleRef = &symbol.ModuleRef{}
		i.ModuleRef.SetName(alias)
		i.ModuleRef.SetPos(i.Pos())
		i.ModuleRef.Module = module

	case Check:
		if !ctx.CurrentScope().PutSymbol(ctx, i.ModuleRef) {
			return
		}
	}
}

func (i *Import) loadModule(ctx *Context) *symbol.Module {
	reg := ctx.ModuleRegistry()

	loader := reg.GetNativeLoader(i.ModuleName)
	if loader == nil {
		ctx.Failf(i.Pos(), "invalid module %v", i.ModuleName)
		return nil
	}

	module := symbol.NewModule(i.ModuleName)
	loader(&nativeModuleRegister{
		ctx:    ctx,
		module: module,
	})

	reg.AddModule(i.ModuleName, module)

	return module
}

type nativeModuleRegister struct {
	ctx    *Context
	module *symbol.Module
}

func (r *nativeModuleRegister) RegisterNativeFn(name string, natFn vm.NativeFn) {
	fn := &symbol.FuncSymbol{}
	fn.SetReadOnly(true)
	fn.SetName(name)
	fn.External = true
	fn.IdxFunc = r.ctx.Emitter().AddExternalFunc(natFn)
	r.module.Scope.PutSymbol(r.ctx, fn)
}
