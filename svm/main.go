package main

import (
	"flag"
	"fmt"
	"github.com/zephyyrr/soda"
	"os"
)

var (
	debug = flag.Bool("d", false, "Set to true for debug mode.")
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

	vm := soda.New(file)
	vm.Execute()
}
