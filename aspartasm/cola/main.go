package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/zephyyrr/soda/aspartasm"
	"log"
	"os"
)

var (
	verbose   = flag.Bool("v", false, "Verbose output")
	output    = flag.String("o", "a.sc", "Sets output name of assembled binary")
	linearize = flag.Bool("l", false, "Only linearize (input is pre-parsed AST)")
	parse     = flag.Bool("p", false, "Stop after parsing. Output is AST.")
	disassemble = flag.Bool("d", false, "Disassemble the given .sc file.")
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
	defer in.Close()

	if *disassemble {
		*parse = false
		*linearize = false
	}

	if *disassemble && *output == "a.sc" {
		*output = "a.cola"
	}

	if *parse && *output == "a.sc" {
		*output = "a.ast"
	}

	out, err := os.OpenFile(*output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("Error opening output file %s for writing: %s", *output, err)
	}
	defer out.Close()

	if *verbose {
		log.Println("Input:", flag.Arg(0))
		log.Println("Output:", *output)
	}

	if *disassemble {
		if *verbose {
			log.Println("Mode:", "Disassemble-only")
		}

		ins, err := aspartasm.ReadInstructions(in)

		if err != nil {
			log.Fatalf("Error disassembling %s: %s", flag.Arg(0), err)
		}

		w := bufio.NewWriter(out)

		for _, i := range ins {
			w.WriteString(i.String())
			w.WriteString("\n")
		}

		w.Flush()

	} else if *parse {
		if *verbose {
			log.Println("Mode:", "Parse-only")
		}
		tree, err := aspartasm.Parse(in)
		if err != nil {
			log.Fatalln(err)
		}
		enc := json.NewEncoder(out)

		if err := enc.Encode(tree); err != nil {
			log.Fatalln(err)
		}

	} else if *linearize {
		if *verbose {
			log.Println("Mode:", "Linearize-only")
		}
		dec := json.NewDecoder(in)
		var tree aspartasm.AST
		if err := dec.Decode(&tree); err != nil {
			log.Fatalln(err)
		}

		if *verbose {
			log.Println(tree)
		}
		err := aspartasm.AssembleAst(tree, out)
		if err != nil {
			log.Println(err)
		}
	} else {
		if *verbose {
			log.Println("Mode:", "Assemble")
		}
		err := aspartasm.Assemble(in, out)
		if err != nil {
			log.Println(err)
		}
	}
}
