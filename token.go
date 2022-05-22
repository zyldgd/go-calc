package gocalc

import (
	"fmt"
)

type token int
type precedence int

// ------------------------------------------------------------------

const (
	Illegal         token = iota
	EOF                   // eof
	ID                    // Identifier
	Integer               // 12345
	Float                 // 123.45
	Char                  // 'a'
	String                // "abc"
	OpLParen              // (
	OpRParen              // )
	OpLBracket            // [
	OpRBracket            // ]
	OpNot                 // !
	OpEq                  // ==
	OpNeq                 // !=
	OpGt                  // >
	OpLt                  // <
	OpGte                 // >=
	OpLte                 // <=
	OpAnd                 // &&
	OpOr                  // ||
	OpAdd                 // +
	OpMinus               // -
	OpMultiply            // *
	OpDivide              // /
	OpModulus             // %
	OpBitwiseAnd          // &
	OpBitwiseOr           // |
	OpBitwiseXor          // ^
	OpBitwiseLShift       // <<
	OpBitwiseRShift       // >>
	OpBitwiseNot          // ~
	OpAccess              // .
	OpSeparate            // ,
)

var operatorMap = map[string]token{
	"(":  OpLParen,
	")":  OpRParen,
	"[":  OpLBracket,
	"]":  OpRBracket,
	"!":  OpNot,
	"==": OpEq,
	"!=": OpNeq,
	">":  OpGt,
	"<":  OpLt,
	">=": OpGte,
	"<=": OpLte,
	"&&": OpAnd,
	"||": OpOr,
	"+":  OpAdd,
	"-":  OpMinus,
	"*":  OpMultiply,
	"/":  OpDivide,
	"%":  OpModulus,
	"&":  OpBitwiseAnd,
	"|":  OpBitwiseOr,
	"^":  OpBitwiseXor,
	"<<": OpBitwiseLShift,
	">>": OpBitwiseRShift,
	"~":  OpBitwiseNot,
	".":  OpAccess,
	",":  OpSeparate,
}

func getOperator(str string) token {
	return operatorMap[str]
}

func (tok token) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", tok.String())
	return []byte(s), nil
}

func (tok token) String() string {
	switch tok {
	case EOF:
		return "EOF"
	case ID:
		return "Identifier"
	case Illegal:
		return "Illegal"
	case Integer:
		return "INTEGER"
	case Float:
		return "FLOAT"
	case Char:
		return "CHAR"
	case String:
		return "STRING"
	case OpLParen:
		return "("
	case OpRParen:
		return ")"
	case OpLBracket:
		return "["
	case OpRBracket:
		return "]"
	case OpNot:
		return "!"
	case OpEq:
		return "=="
	case OpNeq:
		return "!="
	case OpGt:
		return ">"
	case OpLt:
		return "<"
	case OpGte:
		return ">="
	case OpLte:
		return "<="
	case OpAnd:
		return "&&"
	case OpOr:
		return "||"
	case OpAdd:
		return "+"
	case OpMinus:
		return "-"
	case OpMultiply:
		return "*"
	case OpDivide:
		return "/"
	case OpModulus:
		return "%"
	case OpBitwiseAnd:
		return "&"
	case OpBitwiseOr:
		return "|"
	case OpBitwiseXor:
		return "^"
	case OpBitwiseLShift:
		return "<<"
	case OpBitwiseRShift:
		return ">>"
	case OpBitwiseNot:
		return "~"
	case OpAccess:
		return "."
	case OpSeparate:
		return ","
	}
	return ""
}

func (tok token) precedence() precedence {
	return precedenceOf(tok)
}

func (tok token) precedenceWith(Op2 token) bool {
	return precedenceOf(tok) < precedenceOf(Op2)
}

// ------------------------------------------------------------------

func precedenceOf(Op token) precedence {
	switch Op {
	//case OpLParen, OpRParen, OpAccess:
	//	return 1
	case OpNot, OpBitwiseNot:
		return 2
	case OpMultiply, OpDivide, OpModulus:
		return 3
	case OpAdd, OpMinus:
		return 4
	case OpBitwiseLShift, OpBitwiseRShift:
		return 5
	case OpGt, OpLt, OpGte, OpLte:
		return 6
	case OpEq, OpNeq:
		return 7
	case OpBitwiseAnd:
		return 8
	case OpBitwiseXor:
		return 9
	case OpBitwiseOr:
		return 10
	case OpAnd:
		return 11
	case OpOr:
		return 12
	case OpSeparate:
		return 15
	}
	return 0
}

func (p precedence) precedenceWith(p2 precedence) bool {
	return p < p2
}

// ------------------------------------------------------------------

func isDecimal(c rune) bool {
	return '0' <= c && c <= '9'
}

func isLetter(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func isSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\r', '\f', '\v', 0x85, 0xA0:
		return true
	}
	return false
}
