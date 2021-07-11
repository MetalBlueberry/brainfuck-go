package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type scenario struct {
	name   string
	input  string
	output string
}

func Test_CesarEncode(t *testing.T) {
	scenarios := []scenario{
		{
			name:   "01step",
			input:  "01abc",
			output: "bcd",
		},
		{
			name:   "1 step from different start position",
			input:  "01mno",
			output: "nop",
		},
		{
			name:   "5 steps",
			input:  "05abc",
			output: "fgh",
		},
		{
			name:   "wrap",
			input:  "03xyz",
			output: "abc",
		},
		{
			name:   "hello world no modified",
			input:  "00hello world",
			output: "hello world",
		},
		{
			name:   "hello world encrypted",
			input:  "01hello world",
			output: "ifmmp xpsme",
		},
	}

	fileContents, err := ioutil.ReadFile("cesar_encode.hw")
	if err != nil {
		fmt.Printf("Error reading %s\n", "cesar_encode.hw")
		return
	}
	program, err := Compile(string(fileContents))
	require.NoError(t, err)

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			reader := strings.NewReader(s.input)
			writer := &bytes.Buffer{}
			exe := Executor{
				MaxIterations: 10000,
				reader:        reader,
				writer:        writer,
			}
			err := exe.Execute(program)
			require.NoError(t, err)

			require.Equal(t, s.output, writer.String())
		})
	}
}
