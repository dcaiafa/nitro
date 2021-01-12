package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dcaiafa/nitro"
)

type MemoryFileSystem map[string]string

func (fs MemoryFileSystem) ReadFile(name string) ([]byte, error) {
	data, ok := fs[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return []byte(data), nil
}

func Run(t *testing.T, prog string) {
	fs := make(MemoryFileSystem)
	fs["main"] = prog

	compiler := nitro.NewCompiler(fs)
	compiler.RegisterExternalFn(
		"print",
		func(ctx context.Context, args []nitro.Value) ([]nitro.Value, error) {
			fmt.Println(args)
			return nil, nil
		})

	compiled, err := compiler.Compile("main")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = nitro.Run(context.Background(), compiled)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test1(t *testing.T) {
	Run(t, "print(1, 2, 3)")
}

func Test2(t *testing.T) {
	Run(t, `
	var a = 1
	a = 2
	var b = "foobar"
	var c = b
	var d
	d = "foobar"
	print(a, b, c, d, a, "blah")
`)
}

func Test3(t *testing.T) {
	Run(t, `
		fn foo(a) 
		  var b = a
    	return b
		end
		print(foo(1))
`)
}

func Test4(t *testing.T) {
	Run(t, `
		fn foo(a, b) 
			print(b, a)
		end
		var a = 2
		foo(1, a)
`)
}

func Test5(t *testing.T) {
	Run(t, `
		print(1 + 2)
		print(1 < 2)
		print(1000000 < 1000000)
		print(1000000 <= 1000000)
		`)
}

func Test6(t *testing.T) {
	Run(t, `
		fn add(a, b)
			return a + b
		end
		var a = add(2, 3 * 6)
		print("hello world", a)
		`)
}
