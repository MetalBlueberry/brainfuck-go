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
	maxSteps := flag.Int("max-steps", 1000000, "limit the number of interations")
	debug := flag.Bool("debug", false, "enable debug output with the character ?")
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
