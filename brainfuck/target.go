package brainfuck

import (
	"fmt"
)

type TranspilerTarget interface {
	Prologue() (string, error)
	Translate(inst Instruction) (string, error)
	Epilogue() (string, error)
}

type DstInstTpl struct {
	Code         string
	IsFormatable bool
}

type JSTranspilerTarget struct {
	dstInstMap map[InstType]DstInstTpl
}

func NewJSTranspilerTarget() JSTranspilerTarget {
	return JSTranspilerTarget{
		dstInstMap: map[InstType]DstInstTpl{
			Plus:      {"m[p]+=%d;", true},
			Minus:     {"m[p]-=%d;", true},
			Right:     {"p+=%d;while(p>=l)p-=l;", true},
			Left:      {"p-=%d;while(p<0)p+=l;", true},
			PutChar:   {"o(m[p]);", false},
			ReadChar:  {"i();", false},
			LoopBegin: {"while(m[p]){", false},
			LoopEnd:   {"}", false},
			Nop:       {"", false},
		},
	}
}

func (t JSTranspilerTarget) Prologue() (string, error) {
	return `
let p=0,l=30000,m=new Uint8Array(l);
let i = () => {}
let o = (c) => {
	if (typeof process != "undefined") {
		process.stdout.write(String.fromCharCode(c))
	} else if (typeof console.log === "function") {
		console.log(String.fromCharCode(c));
	}
}

`, nil
}

func (t JSTranspilerTarget) Translate(inst Instruction) (string, error) {
	dstInst, ok := t.dstInstMap[inst.Type]
	if !ok {
		return "", fmt.Errorf("unknown instruction [type=%d]", inst.Type)
	}

	dstExpr := dstInst.Code
	if dstInst.IsFormatable {
		dstExpr = fmt.Sprintf(dstInst.Code, inst.Length)
	}
	return dstExpr, nil
}

func (t JSTranspilerTarget) Epilogue() (string, error) {
	return ``, nil
}
