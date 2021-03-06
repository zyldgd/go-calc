package gocalc

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSkipSpace(t *testing.T) {
	scan := NewScanner(`     12+3  4.98-66 +a ('v')  <= 44 && true + "asdasd\""        `)

	for tok, lit := scan.scan(); tok != EOF && tok != Illegal; tok, lit = scan.scan() {
		t.Logf("tok:%s , lit:%s", tok, lit)
	}

}

func TestParseExpr(t *testing.T) {
	e := ParserAST(`(a.e + b.c) - "12312312" - v[c.s]`)
	PrintAst(e)
}

func TestScanNumber(t *testing.T) {
	scan := NewScanner("1231.23 ")
	tok, lit := scan.scanNumber()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)

	scan = NewScanner("12345 ")
	tok, lit = scan.scanNumber()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)

	scan = NewScanner("12345. ")
	tok, lit = scan.scanNumber()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)
}

func TestScanString(t *testing.T) {
	quote := strconv.Quote("  -- \"---A-\t-\n-asdasdadsd-")

	scan := NewScanner(quote[0:1])
	tok, lit := scan.scanString()
	t.Logf("%s:%s", tok, lit)
	t.Logf("%+v", scan)
}

func TestCalc(t *testing.T) {
	expr := NewExpression("'1' > 321")
	result := expr.Calc(map[string]interface{}{"a": 89.9, "b": 2})
	res, err := result.Int()

	fmt.Printf("result:%+v err:%+v\n", res, err)
}
