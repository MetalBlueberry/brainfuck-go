package caesarcipher_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kgabis/brainfuck-go/pkg/bf"
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
	program, err := bf.Compile(string(fileContents))
	require.NoError(t, err)

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			reader := strings.NewReader(s.input)
			writer := &bytes.Buffer{}
			exe := bf.Executor{
				MaxSteps: 100000,
				Reader:   reader,
				Writer:   writer,
			}
			err := exe.Execute(program)
			require.NoError(t, err)

			require.Equal(t, s.output, writer.String())
		})
	}
}

func Test_CesarDecode(t *testing.T) {
	scenarios := []scenario{
		{
			name:   "01step",
			input:  "01bcd",
			output: "abc",
		},
		{
			name:   "1 step from different start position",
			input:  "01nop",
			output: "mno",
		},
		{
			name:   "5 steps",
			input:  "05fgh",
			output: "abc",
		},
		{
			name:   "wrap",
			input:  "03abc",
			output: "xyz",
		},
		{
			name:   "hello world no modified",
			input:  "00hello world",
			output: "hello world",
		},
		{
			name:   "hello world encrypted",
			input:  "01ifmmp xpsme",
			output: "hello world",
		},
	}

	fileContents, err := ioutil.ReadFile("cesar_decode.hw")
	if err != nil {
		fmt.Printf("Error reading %s\n", "cesar_decode.hw")
		return
	}
	program, err := bf.Compile(string(fileContents))
	require.NoError(t, err)

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			reader := strings.NewReader(s.input)
			writer := &bytes.Buffer{}
			exe := bf.Executor{
				MaxSteps: 100000,
				Reader:   reader,
				Writer:   writer,
			}
			err := exe.Execute(program)
			require.NoError(t, err)

			require.Equal(t, s.output, writer.String())
		})
	}
}
