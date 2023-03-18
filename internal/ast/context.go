package ast

import (
	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/dcaiafa/nitro/internal/scope"
	"github.com/dcaiafa/nitro/internal/symbol"
	"github.com/dcaiafa/nitro/internal/vm"
)

type Pass int

const (
	Print Pass = iota

	Rewrite
	CreateGlobals
	Check
	Emit
)

type Context struct {
	Stack
	*errlogger.ErrLoggerWrapper

	emitter      *vm.Emitter
	deps         []*vm.CompiledPackage
	depsMap      map[string]int // packageName => deps[value]
	globalImport *symbol.Import
}

func NewContext(
	l *errlogger.ErrLoggerWrapper,
	deps []*vm.CompiledPackage,
) *Context {
	c := &Context{
		ErrLoggerWrapper: l,
		emitter:          vm.NewEmitter(),
		deps:             deps,
	}
	c.depsMap = make(map[string]int, len(deps))
	for i, dep := range deps {
		// The first dep (index 0) is reserved for "self" (this package).
		// TODO: I hate this. Reevaluate.
		if i != 0 {
			c.depsMap[dep.Name] = i
		}
	}
	globalPackageIdx, ok := c.depsMap["$global"]
	if !ok {
		panic("where is $global?")
	}
	globalPackage := c.deps[globalPackageIdx]
	c.globalImport = symbol.NewImport(globalPackage, globalPackageIdx)

	return c
}

func (c *Context) RunPassChild(parent AST, child AST, pass Pass) {
	if child == nil {
		return
	}

	c.Push(parent)
	child.RunPass(c, pass)
	c.Pop()
}

func (c *Context) Parent() AST {
	return c.stack[len(c.stack)-1]
}

func (c *Context) FindSymbol(symName string) symbol.Symbol {
	var fns []*Func
	var sym symbol.Symbol
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if scope, ok := ast.(Scope); ok {
			sym = scope.Scope().GetSymbol(symName)
			if sym != nil {
				break
			}
		}
		if fn, ok := ast.(*Func); ok && fn.IsClosure {
			fns = append(fns, fn)
		}
	}
	if sym == nil {
		return nil
	}
	if len(fns) != 0 && sym.Liftable() {
		sym.Lift()
		for i := len(fns) - 1; i >= 0; i-- {
			sym = fns[i].NewCapture(sym)
		}
	}
	return sym
}

func (c *Context) CurrentFunc() *Func {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if funcAST, ok := ast.(*Func); ok {
			return funcAST
		}
	}
	return nil
}

func (c *Context) IsInRepeatableScope() bool {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		switch ast.(type) {
		case RepeatableScope:
			return true
		case *Func:
			return false
		}
	}
	return false
}

func (c *Context) Len() int {
	return len(c.stack)
}

func (c *Context) Peek(n int) AST {
	if n >= len(c.stack) {
		return nil
	}
	return c.stack[len(c.stack)-n-1]
}

func (c *Context) GetDep(packageName string) (pkg *vm.CompiledPackage, index int) {
	index, ok := c.depsMap[packageName]
	if !ok {
		// This should never happen because all imports should have already been
		// processed by the compiler.
		panic("invalid dependency")
	}
	pkg = c.deps[index]
	return pkg, index
}

func (c *Context) Package() *Package {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if packageAST, ok := ast.(*Package); ok {
			return packageAST
		}
	}
	return nil
}

func (c *Context) GetScope(typeMask scope.Type) scope.Scope {
	for i := len(c.stack) - 1; i >= 0; i-- {
		ast := c.stack[i]
		if scopeAST, ok := ast.(Scope); ok {
			scope := scopeAST.Scope()
			if (scope.Type() & typeMask) != 0 {
				return scope
			}
		}
	}
	return nil
}

func (c *Context) Emitter() *vm.Emitter {
	return c.emitter
}

type PassRunner interface {
	RunPass(ctx *Context, pass Pass)
}
