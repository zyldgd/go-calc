package gocalc

import (
	"errors"
	"fmt"
	"strconv"
)

type Expression struct {
	Expr   expr
	Params map[string]interface{}
}

func NewExpression(expression string) *Expression {
	expr := parserAST(expression)
	return &Expression{
		Expr: expr,
	}
}

func (e *Expression) Calc(params map[string]interface{}) *Result {
	e.Params = params
	if e.Params == nil {
		e.Params = make(map[string]interface{}, 2)
	}
	e.Params["true"] = true
	e.Params["false"] = false

	return e.calc()
}

func (e *Expression) calc() *Result {
	result, err := e.calcExpr(e.Expr)
	if err != nil {
		return nil
	}

	return result
}

func (e *Expression) calcExpr(expr expr) (*Result, error) {
	switch ex := expr.(type) {
	case *literalExpr:
		return e.calcLiteralExpr(ex)
	case *idExpr:
		return e.calcIDExpr(ex)
	case *unaryExpr:
		return e.calcUnaryExpr(ex)
	case *binaryExpr:
		return e.calcBinaryExpr(ex)
	case *parenExpr:
		return e.calcParenExpr(ex)
	}
	return nil, errors.New("token error")
}

func (e *Expression) calcIDExpr(expr *idExpr) (*Result, error) {
	if val, find := e.Params[expr.name]; find {
		switch v := val.(type) {
		case int:
			result := &Result{
				kind: Integer,
				data: v,
			}
			return result, nil
		case int8:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case int16:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case int32:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case uint:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case uint8:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case uint16:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case uint32:
			result := &Result{
				kind: Integer,
				data: int(v),
			}
			return result, nil
		case float64:
			result := &Result{
				kind: Float,
				data: float32(v),
			}
			return result, nil
		case float32:
			result := &Result{
				kind: Float,
				data: v,
			}
			return result, nil
		case string:
			result := &Result{
				kind: String,
				data: v,
			}
			return result, nil
		case bool:
			result := &Result{
				kind: ID,
				data: v,
			}
			return result, nil
			//case func():

		}
	}
	return nil, errors.New("token error")
}

func (e *Expression) calcParenExpr(expr *parenExpr) (*Result, error) {
	return e.calcExpr(expr.e)
}

func (e *Expression) calcBinaryExpr(expr *binaryExpr) (*Result, error) {
	switch expr.Op {
	case OpAdd:
		l, err := e.calcExpr(expr.e1)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.e2)
		if err != nil {
			return nil, err
		}

		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) + r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Float {
			data := l.data.(float32) + r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Integer && r.kind == Float {
			data := float32(l.data.(int)) + r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Integer {
			data := l.data.(float32) + float32(r.data.(int))
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == String && r.kind == String {
			data := l.data.(string) + r.data.(string)
			result := &Result{
				kind: String,
				data: data,
			}
			return result, nil
		}
	case OpMinus:
		l, err := e.calcExpr(expr.e1)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.e2)
		if err != nil {
			return nil, err
		}

		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) - r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Float {
			data := l.data.(float32) - r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Integer && r.kind == Float {
			data := float32(l.data.(int)) + r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Integer {
			data := l.data.(float32) - float32(r.data.(int))
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		}
	case OpMultiply:
		l, err := e.calcExpr(expr.e1)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.e2)
		if err != nil {
			return nil, err
		}

		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) * r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Float {
			data := l.data.(float32) * r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Integer && r.kind == Float {
			data := float32(l.data.(int)) * r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Integer {
			data := l.data.(float32) * float32(r.data.(int))
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		}
	case OpDivide:
		l, err := e.calcExpr(expr.e1)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.e2)
		if err != nil {
			return nil, err
		}

		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) / r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Float {
			data := l.data.(float32) / r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Integer && r.kind == Float {
			data := float32(l.data.(int)) / r.data.(float32)
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		} else if l.kind == Float && r.kind == Integer {
			data := l.data.(float32) / float32(r.data.(int))
			result := &Result{
				kind: Float,
				data: data,
			}
			return result, nil
		}
	case OpModulus:
		l, err := e.calcExpr(expr.e1)
		if err != nil {
			return nil, err
		}
		r, err := e.calcExpr(expr.e2)
		if err != nil {
			return nil, err
		}
		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) % r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		}
	}

	return nil, errors.New("token error")
}

func (e *Expression) calcLiteralExpr(expr *literalExpr) (*Result, error) {
	switch expr.kind {
	case Integer:
		if data, err := strconv.Atoi(expr.literal); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.kind,
				data: data,
			}
			return result, nil
		}
	case Float:
		if data, err := strconv.ParseFloat(expr.literal, 32); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.kind,
				data: float32(data),
			}
			return result, nil
		}
	case Char:
		if len(expr.literal) == 3 {
			result := &Result{
				kind: expr.kind,
				data: rune(expr.literal[1]),
			}
			return result, nil
		} else {
			return nil, errors.New("parse char error")
		}
	case String:
		if data, err := strconv.Unquote(expr.literal); err != nil {
			return nil, err
		} else {
			result := &Result{
				kind: expr.kind,
				data: data,
			}
			return result, nil
		}
	}

	return nil, errors.New("token error")
}

func (e *Expression) calcUnaryExpr(expr *unaryExpr) (*Result, error) {
	switch expr.op {
	case OpAdd:
		if result, err := e.calcExpr(expr.e); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case Integer:
				data := +(result.data.(int))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			case Float:
				data := +(result.data.(float32))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			default:
				return nil, errors.New("token error")
			}
		}
	case OpMinus:
		if result, err := e.calcExpr(expr.e); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case Integer:
				data := -(result.data.(int))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			case Float:
				data := -(result.data.(float32))
				result := &Result{
					kind: result.kind,
					data: data,
				}
				return result, nil
			default:
				return nil, errors.New("token error")
			}
		}

	case OpNot:
		if result, err := e.calcExpr(expr.e); err != nil {
			return nil, err
		} else {
			switch result.kind {
			case ID:
				if data, ok := result.data.(bool); ok {
					result := &Result{
						kind: result.kind,
						data: data,
					}
					return result, nil
				}
			default:
				return nil, errors.New("token error")
			}
		}
	case OpBitwiseXor:
		// TODO
	case OpBitwiseNot:
		// TODO
	}
	return nil, errors.New("token error")
}

// --------------------------------------------------------------------------

type Result struct {
	kind token
	data interface{}
}

//
//func (e result) Int() (int, error) {
//
//}
//
//func (e result) Float() (float32, error) {
//
//}
//
//func (e result) Bool() (bool, error) {
//
//}
//
//func (e result) String() (string, error) {
//
//}
//
//func (e result) Char() (rune, error) {
//
//}

func main() {
	expr := NewExpression("a + 1 * b")
	result := expr.Calc(map[string]interface{}{"a": 89.9, "b": 2})
	fmt.Printf("%+v \n", result)
}
