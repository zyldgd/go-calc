package gocalc

import (
	"errors"
	"fmt"
)

type Expression struct {
	Expr   Expr
	Params map[string]interface{}
}

func NewExpression(expression string) *Expression {
	expr := ParserAST(expression)
	return &Expression{
		Expr: expr,
	}
}

func (e *Expression) Calc(params map[string]interface{}) *Result {
	e.Params = params
	if e.Params == nil {
		e.Params = make(map[string]interface{}, 2)
	}

	return e.calc()
}

func (e *Expression) calc() *Result {
	result, err := e.calcExpr(e.Expr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return result
}

func (e *Expression) calcExpr(expr Expr) (*Result, error) {
	switch ex := expr.(type) {
	case *LiteralExpr:
		return e.calcLiteralExpr(ex)
	case *IdentExpr:
		return e.calcIdentExpr(ex)
	case *UnaryExpr:
		return e.calcUnaryExpr(ex)
	case *BinaryExpr:
		return e.calcBinaryExpr(ex)
	case *ParenExpr:
		return e.calcParenExpr(ex)
	default:
		return nil, fmt.Errorf("unknown express type[%s]", ex)
	}
}

func (e *Expression) calcIdentExpr(expr *IdentExpr) (*Result, error) {
	if val, find := e.Params[expr.Name]; find {
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
				kind: Bool,
				data: v,
			}
			return result, nil
			//case func():

		}
	}
	return nil, errors.New("unsupported data type")
}

func (e *Expression) calcParenExpr(expr *ParenExpr) (*Result, error) {
	return e.calcExpr(expr.E)
}

func (e *Expression) calcBinaryExpr(expr *BinaryExpr) (*Result, error) {
	l, err := e.calcExpr(expr.LE)
	if err != nil {
		return nil, err
	}
	r, err := e.calcExpr(expr.RE)
	if err != nil {
		return nil, err
	}

	switch expr.Op {
	case OpAdd:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) + r.data.(int)
				result := &Result{
					kind: Integer,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) + r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) + float32(r.data.(int))
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) + r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == String && r.kind == String {
			data := l.data.(string) + r.data.(string)
			result := &Result{
				kind: String,
				data: data,
			}
			return result, nil
		}
	case OpMinus:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) - r.data.(int)
				result := &Result{
					kind: Integer,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) + r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) - float32(r.data.(int))
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) - r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		}
	case OpMultiply:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) * r.data.(int)
				result := &Result{
					kind: Integer,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) * r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) * float32(r.data.(int))
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) * r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		}
	case OpDivide:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) / r.data.(int)
				result := &Result{
					kind: Integer,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) / r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) / float32(r.data.(int))
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) / r.data.(float32)
				result := &Result{
					kind: Float,
					data: data,
				}
				return result, nil
			}
		}
	case OpModulus:
		if l.kind == Integer && r.kind == Integer {
			data := l.data.(int) % r.data.(int)
			result := &Result{
				kind: Integer,
				data: data,
			}
			return result, nil
		}
	case OpAnd:
		if l.kind == Bool && r.kind == Bool {
			data := l.data.(bool) && r.data.(bool)
			result := &Result{
				kind: Bool,
				data: data,
			}
			return result, nil
		}
	case OpOr:
		if l.kind == Bool && r.kind == Bool {
			data := l.data.(bool) || r.data.(bool)
			result := &Result{
				kind: Bool,
				data: data,
			}
			return result, nil
		}
	case OpGte:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) >= r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) >= r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) >= float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) >= r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		}
	case OpGt:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) > r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) > r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) > float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) > r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		}
	case OpLte:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) <= r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) <= r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) <= float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) <= r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		}
	case OpLt:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) < r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) < r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) < float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) < r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		}
	case OpEq:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) == r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) == r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) == float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) == r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Bool && r.kind == Bool {
			data := l.data.(bool) == r.data.(bool)
			result := &Result{
				kind: Bool,
				data: data,
			}
			return result, nil
		}
	case OpNeq:
		if l.kind == Integer {
			if r.kind == Integer {
				data := l.data.(int) != r.data.(int)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := float32(l.data.(int)) != r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Float {
			if r.kind == Integer {
				data := l.data.(float32) != float32(r.data.(int))
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			} else if r.kind == Float {
				data := l.data.(float32) != r.data.(float32)
				result := &Result{
					kind: Bool,
					data: data,
				}
				return result, nil
			}
		} else if l.kind == Bool && r.kind == Bool {
			data := l.data.(bool) != r.data.(bool)
			result := &Result{
				kind: Bool,
				data: data,
			}
			return result, nil
		}
	}

	return nil, fmt.Errorf("wrong binary expression[%v %s %v]", l.kind, expr.Op, r.kind)
}

func (e *Expression) calcLiteralExpr(expr *LiteralExpr) (*Result, error) {
	result := &Result{
		kind: expr.Kind,
		data: expr.Date,
	}
	return result, nil
}

func (e *Expression) calcUnaryExpr(expr *UnaryExpr) (*Result, error) {
	result, err := e.calcExpr(expr.E)
	if err != nil {
		return nil, err
	}

	switch expr.Op {
	case OpAdd:
		switch result.kind {
		case Integer:
			data := +(result.data.(int))
			res := &Result{
				kind: result.kind,
				data: data,
			}
			return res, nil
		case Float:
			data := +(result.data.(float32))
			res := &Result{
				kind: result.kind,
				data: data,
			}
			return res, nil
		}
	case OpMinus:
		switch result.kind {
		case Integer:
			data := -(result.data.(int))
			res := &Result{
				kind: result.kind,
				data: data,
			}
			return res, nil
		case Float:
			data := -(result.data.(float32))
			res := &Result{
				kind: result.kind,
				data: data,
			}
			return res, nil
		}

	case OpNot:
		switch result.kind {
		case Bool:
			data := !result.data.(bool)
			res := &Result{
				kind: Bool,
				data: data,
			}
			return res, nil
		}
	case OpBitwiseXor:
		// TODO
	case OpBitwiseNot:

	}
	return nil, fmt.Errorf("wrong unary expression[%s%v]", expr.Op, result.kind)
}

// --------------------------------------------------------------------------

type Result struct {
	kind Token
	data interface{}
}

func (r Result) Int() (int, error) {
	if v, ok := r.data.(int); ok {
		return v, nil
	}
	return 0, fmt.Errorf("conversion error, %v is not int", r.data)
}

func (r Result) Float() (float32, error) {
	if v, ok := r.data.(float32); ok {
		return v, nil
	}
	return 0.0, fmt.Errorf("conversion error, %v is not float", r.data)
}

func (r Result) Bool() (bool, error) {
	if v, ok := r.data.(bool); ok {
		return v, nil
	}
	return false, fmt.Errorf("conversion error, %v is not bool", r.data)
}

func (r Result) String() (string, error) {
	if v, ok := r.data.(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("conversion error, %v is not string", r.data)
}

func (r Result) Char() (rune, error) {
	if v, ok := r.data.(int); ok {
		return rune(v), nil
	}
	return 0, fmt.Errorf("conversion error, %v is not char", r.data)
}
