package gocalc

import (
	"encoding/json"
)

type parser struct {
	scanner *scanner
	tok     token
	lit     string
	json.Number
}

func parserAST(expr string) expr {
	p := &parser{
		scanner: newScanner(expr),
	}
	p.next()
	e := p.parseExpr()

	return e
}

func (p *parser) parseExpr() expr {
	e := p.parseBinaryExpr(99)
	return e
}

func (p *parser) next() {
	p.tok, p.lit = p.scanner.scan()
}

func (p *parser) parseOperand() expr {
	var e expr
	switch p.tok {
	case Integer, Float, Char, String:
		e = &literalExpr{
			kind:    p.tok,
			literal: p.lit,
		}
		p.next()
	case ID:
		e = &idExpr{
			name: p.lit,
		}
		p.next()
	case OpLParen:
		p.next()
		e = p.parseExpr()
		p.next()
		e = &parenExpr{e: e}
	}

	switch p.tok {
	case OpLBracket:
		p.next()
		index := p.parseExpr()
		p.next()
		e = &indexExpr{
			e:     e,
			index: index,
		}
	case OpAccess:
		p.next()
		switch p.tok {
		case ID:
			e = &accessExpr{
				e: e,
				access: idExpr{
					name: p.lit,
				},
			}
		}
		p.next()
	}

	return e
}

func (p *parser) parseUnaryExpr() expr {
	switch p.tok {
	case OpAdd, OpMinus, OpNot, OpBitwiseXor, OpBitwiseNot:
		op := p.tok
		p.next()
		e := p.parseUnaryExpr()
		return &unaryExpr{op: op, e: e}
	}

	return p.parseOperand()
}

func (p *parser) parseBinaryExpr(pre1 precedence) expr {
	le := p.parseUnaryExpr()
	// 1 + 2 + 3
	for {
		op := p.tok
		pre := op.precedence()
		if pre == 0 || !pre.precedenceWith(pre1) {
			break
		}
		p.next()
		re := p.parseBinaryExpr(pre)
		le = &binaryExpr{e1: le, Op: op, e2: re}
	}

	return le
}
