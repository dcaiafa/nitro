package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib"
	"github.com/fatih/color"
)

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
	flags.PrintFlags(os.Stderr)

	os.Exit(1)
}

func printProgUsage(name string, flags *Flags) {
	bold := color.New(color.Bold)

	usage := name
	if flags.HasFlags() {
		usage += " [flags]"
	}
	if flags.HasArgs() {
		for _, arg := range flags.GetArgs() {
			usage += " " + arg.FullName()
		}
	}

	bold.Fprintln(os.Stderr, "USAGE")
	fmt.Fprintf(os.Stderr, "  %v\n", usage)

	if flags.HasArgs() {
		fmt.Fprintln(os.Stderr, "")
		bold.Fprintln(os.Stderr, "ARGS")
		flags.PrintArgs(os.Stderr)
	}

	if flags.HasFlags() {
		fmt.Fprintln(os.Stderr, "")
		bold.Fprintln(os.Stderr, "FLAGS")
		flags.PrintFlags(os.Stderr)
	}
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

	flagN := sysFlags.AddFlag(&Flag{Name: "n", Desc: "Inline program", Value: new(string)})
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

	if *flagN.Value.(*string) == "" && len(args) == 0 {
		printSysUsage(sysFlags)
	}

	compiler := nitro.NewCompiler()
	compiler.SetDiag(*flagD.Value.(*bool))
	compiler.AddFuncRegistry(lib.NewExportRegistry())

	var progName string
	var scriptPath string
	var progData []byte
	var compiled *nitro.Program

	if *flagN.Value.(*string) != "" {
		progName = "<inline>"
		scriptPath = progName
		progData = []byte(*flagN.Value.(*string))

		compiled, err = compiler.CompileSimple(
			scriptPath, progData, nitro.NewConsoleErrLogger())
		if err != nil {
			// Error was already logged by ConsoleErrLogger.
			os.Exit(1)
		}
	} else {
		scriptPath = args[0]
		args = args[1:]

		scriptFileInfo, err := os.Stat(scriptPath)
		if err != nil {
			fatal(fmt.Errorf("failed to read %q: %w", scriptPath, err))
		}

		progName = filepath.Base(scriptPath)

		if scriptFileInfo.IsDir() {
			compiled, err = compiler.CompilePackage(
				nitro.NewNativePackageReader(scriptPath),
				nitro.NewConsoleErrLogger())
			if err != nil {
				// Error was already logged by ConsoleErrLogger.
				os.Exit(1)
			}
		} else {
			progData, err = os.ReadFile(scriptPath)
			if err != nil {
				fatal(fmt.Errorf("failed to read %q: %w", scriptPath, err))
			}
			compiled, err = compiler.CompileSimple(
				scriptPath, progData, nitro.NewConsoleErrLogger())
			if err != nil {
				// Error was already logged by ConsoleErrLogger.
				os.Exit(1)
			}
		}
	}

	/*
		fmt.Println("Symbols:")
		for symName, sym := range compiled.Symbols {
			fmt.Printf(" %v: %+v\n", symName, sym)
		}
	*/

	progFlags := NewFlags()
	err = progFlags.AddFlagsFromMetadata(compiled.Metadata)
	if err != nil {
		fatal(err)
	}

	args, err = progFlags.Parse(args)

	if progFlags.Help {
		printProgUsage(progName, progFlags)
	}

	if err != nil {
		fatal(err)
	}

	if len(args) != 0 {
		fatal(fmt.Errorf("unknown argument %v", args[0]))
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
				vm.SignalError(errors.New("SIGINT"))
			case <-stopCh:
				return
			}
		}
	}()

	nitroParams := progFlags.GetNitroValues()
	for paramName, paramValue := range nitroParams {
		err = vm.SetParam(paramName, paramValue)
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
		// Error was already logged by error handler.
		os.Exit(1)
	}
}
