package brainfuck

import (
	"errors"
	"fmt"

	"github.com/vladk1m0/gbfm/brainfuck/stack"
)

type InstType byte

const (
	Plus      InstType = '+'
	Minus     InstType = '-'
	Right     InstType = '>'
	Left      InstType = '<'
	PutChar   InstType = '.'
	ReadChar  InstType = ','
	LoopBegin InstType = '['
	LoopEnd   InstType = ']'
	Nop       InstType = 0
)

type Instruction struct {
	Type   InstType
	Length uint
}

func Parse(code []byte) ([]Instruction, error) {
	var (
		bytecode   []Instruction = []Instruction{} // parsed instructions sequence
		loopsStack *stack.Stack  = stack.New()     // loops optimization stack

		ln  uint = uint(len(code))
		idx uint = 0
	)

	for idx < ln {
		cmd := code[idx]            // get current token
		inst := Instruction{Nop, 0} // create default instruction

		switch cmd {
		case '+', '-', '>', '<':
			inst.Type = InstType(cmd)

			// compress tokens seequinces, example tokentokentokentokentoken -> {token,5}
			for idx < ln && cmd == code[idx] {
				inst.Length++
				idx++
			}

		case '.', ',':
			inst.Type = InstType(cmd)
			inst.Length = 1
			idx++

		case '[':
			inst.Type = InstType(cmd)
			inst.Length = 1

			loopsStack.Push(uint(len(bytecode))) // push into stack current loop begin position
			idx++

		case ']':
			begin, err := loopsStack.Pop() // find loop begin instruction position
			if err != nil {
				return nil, fmt.Errorf("unexpected end of loop [pos=%d]", idx+1)
			}
			ln := uint(len(bytecode)) - begin // loop length in instructions
			bytecode[begin].Length = ln       // jump to the loop end

			inst.Type = InstType(cmd)
			inst.Length = ln // jump to the loop begin
			idx++

		default:
			idx++
		}
		if inst.Type != Nop {
			bytecode = append(bytecode, inst)
		}
	}

	_, err := loopsStack.Pop()
	if err == nil {
		return nil, errors.New("excessive opening brackets [")
	}

	return bytecode, nil
}
