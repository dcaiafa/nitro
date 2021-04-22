package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
)

func emit(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func emitShort(m *nitro.Machine, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

func printSysUsage(flags *Flags) {
	p := func(s string, arg ...interface{}) {
		fmt.Fprintf(os.Stderr, s+"\n", arg...)
	}

	p("nitro <sys-opt> program.n <prog-opt>")
	p("nitro <sys-opt> -n <inline-program>")
	p("")
	p("System Options:")
	flags.Print(os.Stderr)
	os.Exit(1)
}

func printProgUsage(flags *Flags) {
	p := func(s string, arg ...interface{}) {
		fmt.Fprintf(os.Stderr, s+"\n", arg...)
	}

	p("Options:")
	flags.Print(os.Stderr)
	os.Exit(1)
}

func main() {
	var err error

	rand.Seed(time.Now().Unix())

	sysFlags := NewFlags()

	flagN := sysFlags.AddFlag(&Flag{Name: "n", Sys: true, Desc: "Inline program", Value: new(bool)})
	flagP := sysFlags.AddFlag(&Flag{Name: "p", Sys: true, Desc: "Create CPU profile", Value: new(string)})

	args := os.Args[1:]
	if len(args) == 0 {
		printSysUsage(sysFlags)
	}

	args, err = sysFlags.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	if sysFlags.Help || len(args) == 0 {
		printSysUsage(sysFlags)
	}

	target := args[0]
	args = args[1:]

	compiler := nitro.NewCompiler(nitro.NewNativeFileLoader())
	compiler.SetDiag(true)
	lib.RegisterAll(compiler)

	var compiled *nitro.Program

	if *flagN.Value.(*bool) {
		compiler.AddNativeFn("emit", emitShort)

		compiled, err = compiler.CompileShort(target, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	} else {
		compiler.AddNativeFn("emit", emit)

		compiled, err = compiler.Compile(target, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	}

	progFlags := NewFlags()
	err = progFlags.AddFlagsFromMetadata(compiled.Metadata)
	if err != nil {
		log.Fatal(err)
	}

	args, err = progFlags.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	if progFlags.Help {
		printProgUsage(progFlags)
	}

	if len(args) != 0 {
		log.Fatalf("Invalid argument %v", args[0])
	}

	cpuProfile := *flagP.Value.(*string)
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	machine := nitro.NewMachine(context.Background(), compiled)

	nitroParams := progFlags.GetNitroValues()
	for paramName, paramValue := range nitroParams {
		err := machine.SetParam(paramName, paramValue)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = machine.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
