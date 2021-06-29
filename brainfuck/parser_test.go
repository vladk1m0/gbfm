package brainfuck

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := `
	+++++
	-----
	+++++
	>>>>>
	<<<<<
	`
	expected := []Instruction{
		{Plus, 5},
		{Minus, 5},
		{Plus, 5},
		{Right, 5},
		{Left, 5},
	}

	bytecode, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v", len(expected), len(bytecode))
	}

	for i, inst := range expected {
		if bytecode[i] != inst {
			t.Errorf("wrong instruction want=%+v, got=%+v", inst, bytecode[i])
		}
	}
}

func TestParseLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []Instruction{
		{Plus, 1},
		{LoopBegin, 6},
		{Plus, 1},
		{LoopBegin, 2},
		{Plus, 1},
		{LoopEnd, 2},
		{Plus, 1},
		{LoopEnd, 6},
		{Plus, 1},
	}

	bytecode, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, inst := range expected {
		if bytecode[i] != inst {
			t.Errorf("wrong instruction. want=%+v, got=%+v", inst, bytecode[i])
		}
	}
}

func TestParseEverything(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []Instruction{
		{Plus, 3},
		{LoopBegin, 6},
		{Minus, 3},
		{LoopBegin, 2},
		{Plus, 1},
		{LoopEnd, 2},
		{Right, 3},
		{LoopEnd, 6},
		{Left, 3},
	}

	bytecode, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, inst := range expected {
		if bytecode[i] != inst {
			t.Errorf("wrong instruction. want=%+v, got=%+v", inst, bytecode[i])
		}
	}
}
