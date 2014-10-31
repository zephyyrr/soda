package main

import (
	"flag"
	"github.com/zephyyrr/soda/aspartasme"
	"log"
	"os"
)

var (
	output = flag.String("o", "a.sc", "Sets output name of assembled binary")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalln("Missing input file.")
	}

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Error opening input file %s for reading: %s", flag.Arg(0), err)
	}

	out, err := os.OpenFile(*output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("Error opening output file %s for writing: %s", *output, err)
	}

	aspartasm.Assemble(in, out)
}
