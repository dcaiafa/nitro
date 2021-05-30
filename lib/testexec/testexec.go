package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
		flagExitCode     = flag.Int("exit-code", 0, "")
	)

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

	var buf *bufio.Writer
	if *flagRange != 0 {
		dest := ioutil.Discard
		if *flagRangeStdout {
			dest = os.Stdout
		} else if *flagRangeStderr {
			dest = os.Stderr
		}
		buf = bufio.NewWriterSize(dest, 10240)

		for i := 0; i < *flagRange; i++ {
			fmt.Fprintln(buf, i)
		}
	}

	if buf != nil {
		buf.Flush()
	}

	os.Exit(*flagExitCode)
}
