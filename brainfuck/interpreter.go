package brainfuck

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MemSize uint = 30000

type Interpreter struct {
	input  bufio.Reader
	output io.Writer
}

func NewInterpreter(in io.Reader, out io.Writer) *Interpreter {
	interpreter := &Interpreter{
		input:  *bufio.NewReader(in), // buffered reader for ReadChar token
		output: out,
	}
	return interpreter
}

func (i *Interpreter) Run(code []byte) ([]byte, error) {
	var (
		mem []byte = make([]byte, MemSize) // allocate memory with size
		ptr uint   = 0                     // memory position pointer
	)

	bytecode, err := Parse(code) // parse brainfuck code into instructions
	if err != nil {
		return mem, err
	}

	plen := uint(len(bytecode)) // instructions length
	idx := uint(0)              // current instruction position

	for idx < plen {
		inst := bytecode[idx]
		switch inst.Type {
		case Plus: // increment mem cell at current position
			mem[ptr] += byte(inst.Length)
		case Minus: // decrement mem cell at current position
			mem[ptr] -= byte(inst.Length)
		case Right: // move ptr to forward position
			if ptr == MemSize-1 {
				ptr = 0
			} else {
				ptr += inst.Length
			}
		case Left: // move ptr to backward position
			if ptr == 0 {
				ptr = MemSize - 1
			} else {
				ptr -= inst.Length
			}
		case PutChar: // output value of current position
			fmt.Fprintf(i.output, "%c", mem[ptr])
		case ReadChar: // read value into current position
			if mem[ptr], err = i.input.ReadByte(); err != nil {
				os.Exit(1)
			}
		case LoopBegin: // if current position is false, skip to ]
			if mem[ptr] == 0 {
				idx += inst.Length
			}
		case LoopEnd: // if at current position true, return to [
			if mem[ptr] != 0 {
				idx -= inst.Length
			}
		}
		idx++
	}
	return mem, nil
}
