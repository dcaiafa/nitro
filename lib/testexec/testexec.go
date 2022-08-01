package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		flagEchoToStdout = flag.Bool("echo-to-stdout", false, "")
		flagEchoToStderr = flag.Bool("echo-to-stderr", false, "")
		flagRange        = flag.Int("range", 0, "")
		flagRangeStdout  = flag.Bool("range-stdout", false, "")
		flagRangeStderr  = flag.Bool("range-stderr", false, "")
		flagRangeAlt     = flag.Bool("range-alt", false, "")
		flagExitCode     = flag.Int("exit-code", 0, "")
	)

	if len(os.Args) > 1 && os.Args[1] == "-print-args" {
		for _, arg := range os.Args[2:] {
			fmt.Println("[" + arg + "]")
		}
		return
	}

	flag.Parse()

	if *flagEchoToStdout {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Panic(err)
		}
	} else if *flagEchoToStderr {
		_, err := io.Copy(os.Stderr, os.Stdin)
		if err != nil {
			log.Panic(err)
		}
	}

	if *flagRange != 0 {
		for i := 0; i < *flagRange; i++ {
			if *flagRangeStdout || (*flagRangeAlt && i%2 == 0) {
				fmt.Fprintln(os.Stdout, i)
			} else if *flagRangeStderr || (*flagRangeAlt && i%2 == 1) {
				fmt.Fprintln(os.Stderr, i)
			}
		}
	}

	os.Exit(*flagExitCode)
}
