package main

import (
	"context"
	"flag"
	"log"

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

	compiled, err := compiler.Compile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = nitro.Run(context.Background(), compiled, nil)
	if err != nil {
		log.Fatal(err)
	}
}
