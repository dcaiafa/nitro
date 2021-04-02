package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
)

func emit(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	stdout := lib.Stdout(ctx)
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

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("<program> required")
	}

	filename := flag.Arg(0)

	compiler := nitro.NewCompiler(nitro.NewNativeFileSystem())
	compiler.SetDiag(true)

	lib.RegisterAll(compiler)
	compiler.AddExternalFn("emit", emit)

	compiled, err := compiler.Compile(filename, nitro.NewConsoleErrLogger())
	if err != nil {
		// Error was already logged by ConsoleErrLogger.
		os.Exit(1)
	}

	machine := nitro.NewMachine(compiled)
	err = machine.Run(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
