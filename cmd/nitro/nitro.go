package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
	"github.com/dcaiafa/nitro/lib/nitromath"
	"github.com/fatih/color"
)

func emit(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	stdout := lib.Stdout(m)
	if len(args) < 1 {
		return nil, fmt.Errorf("not enough arguments")
	}
	json, err := lib.ToJSON(args[0], "", "  ")
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

func printSysUsage(flags *Flags) {
	bold := color.New(color.Bold)
	p := func(s string, arg ...interface{}) {
		fmt.Fprintf(os.Stderr, s+"\n", arg...)
	}

	bold.Fprintln(os.Stderr, "USAGE")
	p("  nitro <sys-flags> program.n <prog-flags>")
	p("  nitro <sys-flags> -n <inline-program>")
	p("")

	bold.Fprintln(os.Stderr, "FLAGS")
	flags.Print(os.Stderr)
	os.Exit(1)
}

func printProgUsage(flags *Flags) {
	bold := color.New(color.Bold)

	bold.Fprintln(os.Stderr, "FLAGS")
	flags.Print(os.Stderr)
	os.Exit(1)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(2)
}

func main() {
	var err error

	var wg sync.WaitGroup
	defer wg.Wait()

	rand.Seed(time.Now().Unix())

	sysFlags := NewFlags()

	flagN := sysFlags.AddFlag(&Flag{Name: "n", Desc: "Inline program", Value: new(bool)})
	flagP := sysFlags.AddFlag(&Flag{Name: "p", Desc: "Create CPU profile", Value: new(string)})
	flagD := sysFlags.AddFlag(&Flag{Name: "d", Desc: "Enable parser diagnostics", Value: new(bool)})

	args := os.Args[1:]
	if len(args) == 0 {
		printSysUsage(sysFlags)
	}

	args, err = sysFlags.Parse(args)

	if sysFlags.Help {
		printSysUsage(sysFlags)
	}

	if err != nil {
		fatal(err)
	}

	if len(args) == 0 {
		printSysUsage(sysFlags)
	}

	target := args[0]
	args = args[1:]

	compiler := nitro.NewCompiler(nitro.NewNativeFileLoader())
	compiler.SetDiag(*flagD.Value.(*bool))
	compiler.AddNativeFn("emit", emit)

	lib.RegisterAll(compiler)
	nitromath.RegisterNativePackage(compiler)

	var compiled *nitro.Program
	if *flagN.Value.(*bool) {
		compiled, err = compiler.CompileInline(target, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	} else {
		compiled, err = compiler.Compile(target, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	}

	progFlags := NewFlags()
	err = progFlags.AddFlagsFromMetadata(compiled.Metadata)
	if err != nil {
		fatal(err)
	}

	args, err = progFlags.Parse(args)

	if progFlags.Help {
		printProgUsage(progFlags)
	}

	if err != nil {
		fatal(err)
	}

	cpuProfile := *flagP.Value.(*string)
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	vm := nitro.NewVM(compiled)

	signalCh := make(chan os.Signal)
	stopCh := make(chan struct{})
	defer close(stopCh)

	signal.Notify(signalCh, os.Interrupt)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-signalCh:
				vm.Interrupt(errors.New("SIGINT"))
			case <-stopCh:
				return
			}
		}
	}()

	nitroParams := progFlags.GetNitroValues()
	for paramName, paramValue := range nitroParams {
		err := vm.SetParam(paramName, paramValue)
		if err != nil {
			fatal(err)
		}
	}

	var programArgs []nitro.Value
	if len(args) != 0 {
		programArgs = make([]nitro.Value, len(args))
		for i, arg := range args {
			programArgs[i] = nitro.NewString(arg)
		}
	}

	err = vm.Run(programArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
