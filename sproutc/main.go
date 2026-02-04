package main

import (
	"fmt"
	"io"
	"os"

	"github.com/loderunner/sprout/lang"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s [file]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "reads from stdin if file is omitted")
}

func main() {
	var err error
	var source []byte
	if len(os.Args) == 1 {
		source, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read from stdin: %s\n", err)
			printUsage()
			os.Exit(1)
		}
	} else if len(os.Args) == 2 {
		sourcePath := os.Args[1]
		f, err := os.Open(sourcePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open file: %s\n", err)
			printUsage()
			os.Exit(1)
		}
		defer f.Close()
		source, err = io.ReadAll(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read file: %s\n", err)
			printUsage()
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "unexpected number of arguments")
		printUsage()
		os.Exit(1)
	}

	expr, err := lang.Parse(string(source))
	if err != nil {
		fmt.Fprintf(os.Stderr, "syntax error: %s\n", err)
		os.Exit(1)
	}

	ctx := lang.NewContext()
	result, err := lang.TypeCheck(ctx, expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "type error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("type: %s\n", result.TypeName())
}
