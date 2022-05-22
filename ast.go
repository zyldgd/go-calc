package gocalc

import (
	"encoding/json"
	"fmt"
)

type expr interface {
	String() string
}

type (
	literalExpr struct {
		kind    token
		literal string
	}

	accessExpr struct {
		e      expr
		access idExpr
	}

	indexExpr struct {
		e     expr
		index expr
	}

	idExpr struct {
		name string
	}

	binaryExpr struct {
		e1 expr
		Op token
		e2 expr
	}

	parenExpr struct {
		e expr
	}

	unaryExpr struct {
		op token
		e  expr
	}
)

func (e *literalExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *accessExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *indexExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *idExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *binaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *parenExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *unaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func printAst(e expr) {
	b, _ := json.MarshalIndent(e, "", "    ")

	fmt.Printf("ast:\n%+v\n", string(b))
}
