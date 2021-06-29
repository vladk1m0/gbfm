package main

import (
	"fmt"
	"io/ioutil"
	"os"

	bf "github.com/vladk1m0/gbfm/brainfuck"
)

var usage string =`
usage: ./gbfm.sh [run|translate] file.bf
	run brainfuck program
	translate brainfuck program into file.bf.js
`
func printHelp() {
	fmt.Fprint(os.Stderr, usage)
	os.Exit(0)
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
	}

	cmd := os.Args[1]
	if cmd != "translate" && cmd != "run" {
		printHelp()
	}

	finName := os.Args[2]
	if _, err := os.Stat(finName); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	r, err := os.Open(finName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	srcCode, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	switch cmd {
	case "run":
		interpreter := bf.NewInterpreter(os.Stdin, os.Stdout)
		_, err = interpreter.Run(srcCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}
	case "translate":
		transpiler := bf.NewTranspiler(bf.NewJSTranspilerTarget())
		dstCode, err := transpiler.Transpile(srcCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}

		foutName := finName + ".js"
		err = ioutil.WriteFile(foutName, []byte(dstCode), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "create out file [%s]\n", foutName)
	}
}
