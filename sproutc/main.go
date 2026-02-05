package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/loderunner/sprout/lang"
)

func printUsage() {
	fmt.Fprintln(os.Stderr, "usage: sproutc [flags] [file]")
	fmt.Fprintln(os.Stderr, "flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "reads from stdin if file is omitted")
}

func main() {
	printASTFlag := flag.Bool("print-ast", false, "print the AST after parsing and exit")
	flag.Parse()

	var err error
	var source []byte
	if flag.NArg() == 0 {
		source, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read from stdin: %s\n", err)
			// printUsage()
			os.Exit(1)
		}
	} else if flag.NArg() == 1 {
		sourcePath := flag.Arg(0)
		f, err := os.Open(sourcePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open file: %s\n", err)
			// printUsage()
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

	expr, err := Parse(string(source))
	if err != nil {
		fmt.Fprintf(os.Stderr, "syntax error: %s\n", err)
		os.Exit(1)
	}

	if *printASTFlag {
		printAST(expr)
		return
	}

	ctx := lang.NewContext()
	result, err := lang.TypeCheck(ctx, expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "type error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("type: %s\n", result.TypeName())
}
