package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
)

func emit(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	stdout := lib.Stdout(m)
	if len(args) < 1 {
		return nil, fmt.Errorf("not enough arguments")
	}
	json, err := lib.ToJSON(args[0], "  ")
	if err != nil {
		return nil, err
	}
	_, err = stdout.Write(json)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintln(stdout, "")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func emitShort(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	e, err := nitro.MakeIterator(m, args[0])
	if err == nil {
		for {
			v, ok, err := nitro.Next(m, e, 1)
			if err != nil {
				return nil, err
			}
			if !ok {
				break
			}
			fmt.Println(v[0])
		}
		return nil, nil
	}

	r, ok := args[0].(io.Reader)
	if ok {
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	fmt.Println(args[0])

	return nil, nil
}

func main() {
	var (
		flagE = flag.String("e", "", "")
	)
	flag.Parse()

	if flag.NArg() == 0 && *flagE == "" {
		log.Fatalf("<program> required")
	}

	rand.Seed(time.Now().Unix())

	filename := flag.Arg(0)
	compiler := nitro.NewCompiler(nitro.NewNativeFileSystem())
	compiler.SetDiag(true)

	var err error
	var compiled *nitro.Program

	if *flagE != "" {
		lib.RegisterAll(compiler)
		compiler.AddNativeFn("emit", emitShort)

		compiled, err = compiler.CompileShort(*flagE, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	} else {
		lib.RegisterAll(compiler)
		compiler.AddNativeFn("emit", emit)

		compiled, err = compiler.Compile(filename, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	}

	machine := nitro.NewMachine(context.Background(), compiled)
	err = machine.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
