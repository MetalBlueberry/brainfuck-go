package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kgabis/brainfuck-go/pkg/bf"
)

func main() {
	maxSteps := flag.Int("max-steps", 100000, "limit the number of interations")
	debug := flag.Bool("debug", false, "enable debug output with the character ?")

	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		os.Exit(1)
		return
	}
	filename := args[1]
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s\n", filename)
		os.Exit(1)
		return
	}
	program, err := bf.Compile(string(fileContents))

	if err != nil {
		fmt.Println(err)
		return
	}
	buf := &bytes.Buffer{}
	exe := bf.Executor{
		Debug:    *debug,
		MaxSteps: *maxSteps,
		Reader:   bufio.NewReader(os.Stdin),
		Writer:   buf,
	}
	err = exe.Execute(program)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(buf.String())
}
