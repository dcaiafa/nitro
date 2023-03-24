package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/compiler"
	"github.com/dcaiafa/nitro/internal/fs"
	"github.com/dcaiafa/nitro/lib"
	"github.com/fatih/color"
)

func printSysUsage(flags *Flags) {
	bold := color.New(color.Bold)
	p := func(s string, arg ...interface{}) {
		fmt.Fprintf(os.Stderr, s+"\n", arg...)
	}

	bold.Fprintln(os.Stderr, "USAGE")
	p("  nitro program.n <prog-flags>")
	p("  nitro -c <command>")
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

	flagC := sysFlags.AddFlag(&Flag{Name: "c", Desc: "Specify command to execute", Value: new(string)})
	flagP := sysFlags.AddFlag(&Flag{Name: "p", Desc: "Create CPU profile", Value: new(string)})

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

	if *flagC.Value.(*string) == "" && len(args) == 0 {
		printSysUsage(sysFlags)
	}

	compiler := compiler.New()
	lib.RegisterAll(compiler)

	var progName string
	var scriptPath string
	var compiled *nitro.Program

	if *flagC.Value.(*string) != "" {
		scriptPath = "<inline>"
		memFS := fs.NewMem()
		memFS.Put(scriptPath, []byte(*flagC.Value.(*string)))
		compiler.SetFS(memFS)
	} else {
		scriptPath = args[0]
		args = args[1:]
	}

	compiled, err = compiler.Compile(scriptPath)
	if err != nil {
		fatal(err)
	}

	progFlags := NewFlags()
	err = progFlags.AddFlagsFromMetadata(compiled.Metadata())
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

	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

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
