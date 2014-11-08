package main

import (
	"flag"
	"fmt"
	"github.com/zephyyrr/soda"
	"os"
)

var (
	verbose = flag.Bool("v", false, "Force verbose printing")
	debug   = flag.Bool("d", false, "Debug mode.")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "No source file found.")
		return
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to read source file.")
		fmt.Fprintln(os.Stderr, err)
	}

	options := Options()

	vm := soda.New(file, options)
	go func() {
		for message := range vm.Messages() {
			fmt.Fprintln(os.Stderr, message)
		}
	}()
	vm.Execute()
}

func Options() soda.Options {
	return soda.Options{
		Verbose: *verbose,
		Debug:   *debug,
	}
}
