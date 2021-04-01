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

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("<program> required")
	}

	filename := flag.Arg(0)

	compiler := nitro.NewCompiler(nitro.NewNativeFileSystem())
	compiler.SetDiag(true)

	lib.RegisterAll(compiler)

	compiled, err := compiler.Compile(filename, nitro.NewConsoleErrLogger())
	if err != nil {
		// Error was already logged by ConsoleErrLogger.
		os.Exit(1)
	}

	err = nitro.Run(context.Background(), compiled, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
