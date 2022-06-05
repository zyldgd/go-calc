package gocalc

import (
	"encoding/json"
	"fmt"
)

type Expr interface {
	String() string
}

type (
	LiteralExpr struct {
		Kind    Token
		Literal string
		Date    interface{}
	}

	AccessExpr struct {
		E      Expr
		Access IdentExpr
	}

	IndexExpr struct {
		E     Expr
		Index Expr
	}

	IdentExpr struct {
		Name string
	}

	BinaryExpr struct {
		LE Expr
		Op Token
		RE Expr
	}

	ParenExpr struct {
		E Expr
	}

	UnaryExpr struct {
		Op Token
		E  Expr
	}
)

func (e *LiteralExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *AccessExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *IndexExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *IdentExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *BinaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *ParenExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e *UnaryExpr) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func PrintAst(e Expr) {
	b, _ := json.MarshalIndent(e, "", "    ")

	fmt.Printf("ast:\n%+v\n", string(b))
}

// -----------------------------------------------------------------------------------

func scan(expr string) {
	exprChars := []rune(expr)
	for i := 0; i < len(exprChars); i++ {
		//c := exprChars[i]

	}
}
