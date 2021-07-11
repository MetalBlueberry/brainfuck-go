/*
 Brainfuck-Go ( http://github.com/kgabis/brainfuck-go )
 Copyright (c) 2013 Krzysztof Gabis

 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included in
 all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 THE SOFTWARE.
*/
package bf

import (
	"errors"
	"fmt"
	"io"
)

type Instruction struct {
	operator uint16
	operand  uint16
}

const (
	op_inc_dp = iota
	op_dec_dp
	op_inc_val
	op_dec_val
	op_out
	op_in
	op_jmp_fwd
	op_jmp_bck
	op_debug
)

const data_size int = 65535

func Compile(input string) (program []Instruction, err error) {
	var pc, jmp_pc uint16 = 0, 0
	jmp_stack := make([]uint16, 0)
	comment := false
	for _, c := range input {
		if comment && c != '\n' {
			continue
		}
		switch c {
		case '\n':
			comment = false
			pc--
		case '#':
			comment = true
			pc--
		case '>':
			program = append(program, Instruction{op_inc_dp, 0})
		case '<':
			program = append(program, Instruction{op_dec_dp, 0})
		case '+':
			program = append(program, Instruction{op_inc_val, 0})
		case '-':
			program = append(program, Instruction{op_dec_val, 0})
		case '.':
			program = append(program, Instruction{op_out, 0})
		case ',':
			program = append(program, Instruction{op_in, 0})
		case '[':
			program = append(program, Instruction{op_jmp_fwd, 0})
			jmp_stack = append(jmp_stack, pc)
		case ']':
			if len(jmp_stack) == 0 {
				return nil, errors.New("Compilation error.")
			}
			jmp_pc = jmp_stack[len(jmp_stack)-1]
			jmp_stack = jmp_stack[:len(jmp_stack)-1]
			program = append(program, Instruction{op_jmp_bck, jmp_pc})
			program[jmp_pc].operand = pc
		case '?':
			program = append(program, Instruction{op_debug, 0})
		default:
			pc--
		}
		pc++
	}
	if len(jmp_stack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

type Executor struct {
	MaxSteps int
	Reader   ByteReader
	Writer   io.Writer
	Debug    bool
}

type ByteReader interface {
	ReadByte() (byte, error)
}

func (e *Executor) Execute(program []Instruction) error {
	data := make([]int16, data_size)
	var ptr uint16 = uint16(data_size) / 2

	max := ptr
	min := ptr

	steps := 0
	for pc := 0; pc < len(program); pc++ {
		steps++
		if steps > e.MaxSteps {
			return errors.New("Max iterations reached")
		}
		switch program[pc].operator {
		case op_inc_dp:
			ptr++
			if max < ptr {
				max = ptr
			}
		case op_dec_dp:
			ptr--
			if min > ptr {
				min = ptr
			}
		case op_inc_val:
			data[ptr]++
		case op_dec_val:
			data[ptr]--
		case op_out:
			fmt.Fprintf(e.Writer, "%c", data[ptr])
		case op_in:
			readVal, err := e.Reader.ReadByte()
			switch {
			case err == io.EOF:
				data[ptr] = int16(0)
			case err != nil:
				return fmt.Errorf("Unable to read input, %w", err)
			default:
				data[ptr] = int16(readVal)
			}
		case op_jmp_fwd:
			if data[ptr] == 0 {
				pc = int(program[pc].operand)
			}
		case op_jmp_bck:
			if data[ptr] > 0 {
				pc = int(program[pc].operand)
			}
		case op_debug:
			if e.Debug {
				fmt.Printf("step: %d\n", steps)
				fmt.Printf("position [")
				for i := min; i <= max; i++ {
					if ptr == i {
						fmt.Print("  >+ ")
					} else {
						fmt.Print("     ")
					}
				}
				fmt.Printf("]\n")
				fmt.Printf("index    [")
				for i := min; i <= max; i++ {
					fmt.Printf(" %3d ", int(i)-int(data_size)/2)
				}
				fmt.Printf("]\n")
				fmt.Printf("values   [")
				for i := min; i <= max; i++ {
					fmt.Printf(" %3d ", data[i])
				}
				fmt.Printf("]\n")
			}

		default:
			panic("Unknown operator.")
		}
	}
	return nil
}
