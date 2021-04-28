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
		flagRange        = flag.Int("range", 0, "")
		flagRangeStdout  = flag.Bool("range-stdout", false, "")
		flagRangeStderr  = flag.Bool("range-stderr", false, "")
		flagExitCode     = flag.Int("exit-code", 0, "")
	)

	flag.Parse()

	if *flagEchoToStdout {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Panic(err)
		}
	}

	if *flagRange != 0 {
		for i := 0; i < *flagRange; i++ {
			if *flagRangeStdout {
				fmt.Fprintln(os.Stdout, i)
			}
			if *flagRangeStderr {
				fmt.Fprintln(os.Stderr, i)
			}
		}
	}

	os.Exit(*flagExitCode)
}
