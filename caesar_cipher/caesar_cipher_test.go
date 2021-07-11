package caesarcipher_test

import (
	"bytes"
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
			input:  "e01abc",
			output: "bcd",
		},
		{
			name:   "1 step from different start position",
			input:  "e01mno",
			output: "nop",
		},
		{
			name:   "5 steps",
			input:  "e05abc",
			output: "fgh",
		},
		{
			name:   "wrap",
			input:  "e03xyz",
			output: "abc",
		},
		{
			name:   "hello world no modified",
			input:  "e00hello world",
			output: "hello world",
		},
		{
			name:   "hello world encrypted",
			input:  "e01hello world",
			output: "ifmmp xpsme",
		},
	}

	ExecuteScenario(t, scenarios)
}

func Test_CesarDecode(t *testing.T) {
	scenarios := []scenario{
		{
			name:   "01step",
			input:  "d01bcd",
			output: "abc",
		},
		{
			name:   "1 step from different start position",
			input:  "d01nop",
			output: "mno",
		},
		{
			name:   "5 steps",
			input:  "d05fgh",
			output: "abc",
		},
		{
			name:   "wrap",
			input:  "d03abc",
			output: "xyz",
		},
		{
			name:   "hello world no modified",
			input:  "d00hello world",
			output: "hello world",
		},
		{
			name:   "hello world encrypted",
			input:  "d01ifmmp xpsme",
			output: "hello world",
		},
	}

	ExecuteScenario(t, scenarios)
}

func Test_CesarError(t *testing.T) {
	scenarios := []scenario{
		{
			name:   "Invalid first character",
			input:  "x01bcd",
			output: "First character must be either \"e\" for encode or \"d\" for decode\n2nd and 3rd chars must be alphanumeric characters indicating the password",
		},
	}

	ExecuteScenario(t, scenarios)
}

func ExecuteScenario(t *testing.T, scenarios []scenario) {
	fileContents, err := ioutil.ReadFile("caesar.bf")
	require.NoError(t, err)
	program, err := bf.Compile(string(fileContents))
	require.NoError(t, err)

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			reader := strings.NewReader(s.input)
			writer := &bytes.Buffer{}
			exe := bf.Executor{
				MaxSteps: 1000000,
				Reader:   reader,
				Writer:   writer,
			}
			err := exe.Execute(program)
			require.NoError(t, err)

			require.Equal(t, s.output, writer.String())
		})
	}
}
