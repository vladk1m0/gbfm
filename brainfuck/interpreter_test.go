package brainfuck

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestIncrement(t *testing.T) {
	code := "+++++"

	interpreter := NewInterpreter(new(bytes.Buffer), new(bytes.Buffer))
	mem, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	if mem[0] != 5 {
		t.Errorf("cell not correctly incremented. got=%d", mem[0])
	}
}

func TestDecrement(t *testing.T) {
	code := "++++++++++-----"

	interpreter := NewInterpreter(new(bytes.Buffer), new(bytes.Buffer))
	mem, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	if mem[0] != 5 {
		t.Errorf("cell not correctly incremented. got=%d", mem[0])
	}
}

func TestIncrementingDataPointer(t *testing.T) {
	code := "+>++>+++"

	interpreter := NewInterpreter(new(bytes.Buffer), new(bytes.Buffer))
	mem, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	for i, expected := range []byte{1, 2, 3} {
		if mem[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d", i, expected, mem[0])
		}
	}
}

func TestDecrementDataPointer(t *testing.T) {
	code := ">>+++<++<+"

	interpreter := NewInterpreter(new(bytes.Buffer), new(bytes.Buffer))
	mem, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	for i, expected := range []byte{1, 2, 3} {
		if mem[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d", i, expected, mem[0])
		}
	}

}

func TestReadChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)

	code := ",>,>,>,>,>,>"

	interpreter := NewInterpreter(in, out)
	mem, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	expectedMemory := []byte{
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
	}

	for i, expected := range expectedMemory {
		if mem[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d", i, expected, mem[0])
		}
	}
}

func TestPutChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)
	interpreter := NewInterpreter(in, out)

	code := ",>,>,>,>,>,><<<<<<.>.>.>.>.>.>"
	_, err := interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	output := out.String()
	if output != "ABCDEF" {
		t.Errorf("output wrong. got=%q", output)
	}
}

func TestHelloWorld(t *testing.T) {
	code, err := ioutil.ReadFile("../test_data/hello_world.bf")
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}
	in := bytes.NewBuffer(code)
	out := new(bytes.Buffer)

	interpreter := NewInterpreter(in, out)
	_, err = interpreter.Run([]byte(code))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("output wrong. got=%q", output)
	}
}
