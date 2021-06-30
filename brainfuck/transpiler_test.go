package brainfuck

import (
	"bytes"
	"io/ioutil"
	"testing"
	"os/exec"
)

func transpileTest(finName string, t *testing.T) {
	srcCode, err := ioutil.ReadFile(finName)
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}
	in := bytes.NewBuffer(srcCode)
	out := new(bytes.Buffer)

	interpreter := NewInterpreter(in, out)
	_, err = interpreter.Run([]byte(srcCode))
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}
	out1 := out.String()

	transpiler := NewTranspiler(NewJSTranspilerTarget())
	transpiler.Transpile(srcCode)

	dstCode, err := transpiler.Transpile(srcCode)
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	foutName := finName + ".js"
	err = ioutil.WriteFile(foutName, []byte(dstCode), 0644)
	if err != nil {
		t.Fatalf("parsing error %+v", err)
	}

	_, err = exec.LookPath("node")
	if err != nil {
		t.Logf("lookup error %+v", err)
		return
	}

	tmp, err := exec.Command("node", foutName).Output()
	if err != nil {
        t.Fatalf("parsing error %+v", err)
    }

	out2 := string(tmp)
	if out1 != out2 {
		t.Errorf("output wrong. want=%+v, got=%+v", out1, out2)
	}
}

func TestTranspileHelloWorld(t *testing.T) {
	transpileTest("../test_data/hello_world.bf", t)
}

func TestTranspile99bottles(t *testing.T) {
	transpileTest("../test_data/99_bottles.bf", t)
}

func TestTranspileFizzBuzz(t *testing.T) {
	transpileTest("../test_data/fizz_buzz.bf", t)
}

// func TestTranspileMandelbrot(t *testing.T) {
// 	transpile("../test_data/mandelbrot.bf", t)
// }
