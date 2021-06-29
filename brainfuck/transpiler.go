package brainfuck

import (
	"strings"
)

type Transpiler struct {
	target TranspilerTarget
}

func NewTranspiler(tt TranspilerTarget) *Transpiler {
	return &Transpiler{
		target: tt,
	}
}

func (t *Transpiler) Transpile(code []byte) (string, error) {
	var targetProg strings.Builder // target program language code buffer

	bytecode, err := Parse(code) // parse brainfuck code into instructions
	if err != nil {
		return "", err
	}

	// add target program header
	prologue,_ := t.target.Prologue()
	targetProg.WriteString(prologue)

	// translate instruction to target language
	for _, inst := range bytecode {
		targetExpr, err := t.target.Translate(inst)
		if err != nil {
			return targetProg.String(), err
		}
		targetProg.WriteString(targetExpr)
	}

	// add target program footer
	epilogue, _ := t.target.Epilogue()
	targetProg.WriteString(epilogue)

	return targetProg.String(), nil
}
