package __obf_da2d8c26779e1ac8

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

// 错误定义
var (
	ErrRuleEmpty           = errors.New("rule is empty")
	ErrUnsupportToken      = errors.New("unsupport token")
	ErrUnsupportExpr       = errors.New("unsupport expr")
	ErrUnsupportParam      = errors.New("unsupport param")
	ErrNotNumber           = errors.New("not a number")
	ErrIndexNotNumber      = errors.New("index not a number")
	ErrNotBool             = errors.New("not boolean")
	ErrKeyNotFound         = errors.New("map key not found")
	__obf_5e5c20586356ff92 = map[string]__obf_dd6483bb5f13267b{
		"contains": func(__obf_43abe842ed1cbd54 []any) (any, error) {
			fmt.Println("contains print: ", __obf_43abe842ed1cbd54)
			return nil, nil
		},
	}
)

type __obf_dd6483bb5f13267b = func(__obf_43abe842ed1cbd54 []any) (any, error)

type Rule struct {
	__obf_b38124cf9d4e552c ast.Expr
}

func (__obf_2de70ff62047a14c *Rule) SetExpr(__obf_b38124cf9d4e552c string) error {
	if len(__obf_b38124cf9d4e552c) == 0 {
		return ErrRuleEmpty
	}
	if __obf_7740656e75737ca9, __obf_5a128dc949259d7a := parser.ParseExpr(__obf_b38124cf9d4e552c); __obf_5a128dc949259d7a != nil {
		return __obf_5a128dc949259d7a
	} else {
		__obf_2de70ff62047a14c.__obf_b38124cf9d4e552c = __obf_7740656e75737ca9
	}

	return nil
}

func (__obf_2de70ff62047a14c *Rule) Bool(__obf_7e57d55c76f7632b map[string]any) (bool, error) {
	if __obf_2de70ff62047a14c.__obf_b38124cf9d4e552c != nil {
		__obf_40dee8321fd779d9, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_7e57d55c76f7632b)
		if __obf_5a128dc949259d7a != nil {
			return false, __obf_5a128dc949259d7a
		}
		if __obf_2de70ff62047a14c, __obf_a2328d0329a8f982 := __obf_40dee8321fd779d9.(bool); __obf_a2328d0329a8f982 {
			return __obf_2de70ff62047a14c, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_2de70ff62047a14c *Rule) Int(__obf_7e57d55c76f7632b map[string]any) (int64, error) {
	if __obf_2de70ff62047a14c.__obf_b38124cf9d4e552c != nil {
		__obf_40dee8321fd779d9, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_7e57d55c76f7632b)
		if __obf_5a128dc949259d7a != nil {
			return 0, __obf_5a128dc949259d7a
		}
		switch __obf_40dee8321fd779d9 := __obf_40dee8321fd779d9.(type) {
		case int64:
			return __obf_40dee8321fd779d9, nil
		case float64:
			return int64(__obf_40dee8321fd779d9), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_2de70ff62047a14c *Rule) Float(__obf_7e57d55c76f7632b map[string]any) (float64, error) {
	if __obf_2de70ff62047a14c.__obf_b38124cf9d4e552c != nil {
		__obf_40dee8321fd779d9, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_7e57d55c76f7632b)
		if __obf_5a128dc949259d7a != nil {
			return 0, __obf_5a128dc949259d7a
		}
		switch __obf_40dee8321fd779d9 := __obf_40dee8321fd779d9.(type) {
		case int64:
			return float64(__obf_40dee8321fd779d9), nil
		case float64:
			return __obf_40dee8321fd779d9, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_2de70ff62047a14c *Rule) Eval(__obf_4a9263fe8a776aa1 map[string]any) (any, error) {
	switch __obf_1fcca0fe411ce494 := __obf_2de70ff62047a14c.__obf_b38124cf9d4e552c.(type) {
	case *ast.UnaryExpr:
		__obf_2de70ff62047a14c. // 一元表达式
					__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.X
		__obf_2003a9708eba3d2b, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		__obf_aeee35061b02cb77 := reflect.ValueOf(__obf_2003a9708eba3d2b)

		switch __obf_1fcca0fe411ce494.Op {
		case token.NOT: // !
			if __obf_aeee35061b02cb77.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_aeee35061b02cb77.Bool(), nil
		case token.SUB: // -
			if __obf_8d1c7d6c4b72dec7, __obf_5a128dc949259d7a := __obf_9e412df46e486a0b(__obf_aeee35061b02cb77); __obf_5a128dc949259d7a == nil {
				return (-1.0) * __obf_8d1c7d6c4b72dec7, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_2de70ff62047a14c. // 二元表达式
					__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.X
		__obf_8d1c7d6c4b72dec7, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		__obf_2de70ff62047a14c.__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.Y
		__obf_22eaa72847d00940, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		return __obf_5f98a3add70ee4cc(__obf_8d1c7d6c4b72dec7, __obf_22eaa72847d00940,
			// 标志符（已定义变量或常量（bool））
			__obf_1fcca0fe411ce494.Op)
	case *ast.Ident:
		return __obf_e3e08f3a07ae9023(__obf_1fcca0fe411ce494.Name, __obf_4a9263fe8a776aa1)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_1fcca0fe411ce494.Kind {
		case token.STRING:
			return strings.Trim(__obf_1fcca0fe411ce494.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_1fcca0fe411ce494.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_1fcca0fe411ce494.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_2de70ff62047a14c. // 圆括号内表达式
					__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.X
		return __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
	case *ast.SelectorExpr:
		__obf_2de70ff62047a14c. // 属性或方法选择表达式
					__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.X
		__obf_925e7ae002270cd3, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		return __obf_e3e08f3a07ae9023(__obf_1fcca0fe411ce494.Sel.Name, __obf_925e7ae002270cd3.(map[string]any))
	case *ast.IndexExpr:
		__obf_2de70ff62047a14c. // 中括号内表达式——map或slice索引
					__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.X
		__obf_f4b02b67ec3b5225, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		__obf_2de70ff62047a14c.__obf_b38124cf9d4e552c = __obf_1fcca0fe411ce494.Index
		__obf_380cd3ef344b74f2, __obf_5a128dc949259d7a := __obf_2de70ff62047a14c.Eval(__obf_4a9263fe8a776aa1)
		if __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}

		switch __obf_f4b02b67ec3b5225 := __obf_f4b02b67ec3b5225.(type) {
		case map[string]any:
			if __obf_380cd3ef344b74f2, __obf_3b6c3355f3aded48 := __obf_380cd3ef344b74f2.(string); __obf_3b6c3355f3aded48 {
				return __obf_f4b02b67ec3b5225[__obf_380cd3ef344b74f2], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_380cd3ef344b74f2 := __obf_380cd3ef344b74f2.(type) {
			case int:
				return __obf_f4b02b67ec3b5225[int64(__obf_380cd3ef344b74f2)], nil
			case int64:
				return __obf_f4b02b67ec3b5225[__obf_380cd3ef344b74f2], nil
			default:
				return nil, fmt.Errorf("slice index index must be number")
			}
		default:
			return nil, fmt.Errorf("IndexExpr: unsupport data type")
		}
	case *ast.CallExpr: // 方法调用表达式
		// r.expr = t.Fun
		// f, err := r.Eval(fns)

		// if err != nil {
		// 	return nil, err
		// }
		// switch f := f.(type) {
		// case func(args []any) (any, error):
		// 	if params, err := evalArg(t.Args, datasource); err != nil {
		// 		return nil, err
		// 	} else {
		// 		return f(params)
		// 	}
		// }
		if __obf_2712bcad0ddc6c89, __obf_5a128dc949259d7a := __obf_4f12430246f2e181(__obf_1fcca0fe411ce494.Args, __obf_4a9263fe8a776aa1); __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		} else {
			return __obf_5e5c20586356ff92[__obf_1fcca0fe411ce494.Fun.(*ast.Ident).Name](__obf_2712bcad0ddc6c89)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_4f12430246f2e181(__obf_43abe842ed1cbd54 []ast.Expr, __obf_4a9263fe8a776aa1 map[string]any) ([]any, error) {
	var __obf_23893e02a0a17784 []any
	for _, __obf_aaeff6a9b07ab7e4 := range __obf_43abe842ed1cbd54 {
		switch __obf_aaeff6a9b07ab7e4 := __obf_aaeff6a9b07ab7e4.(type) {
		case *ast.BasicLit:
			__obf_23893e02a0a17784 = append(__obf_23893e02a0a17784, __obf_aaeff6a9b07ab7e4.Value)
		case *ast.Ident:
			if __obf_4f255ce446c63c19, __obf_5a128dc949259d7a := __obf_e3e08f3a07ae9023(__obf_aaeff6a9b07ab7e4.Name, __obf_4a9263fe8a776aa1); __obf_5a128dc949259d7a != nil {
				return nil, __obf_5a128dc949259d7a
			} else {
				__obf_23893e02a0a17784 = append(__obf_23893e02a0a17784, __obf_4f255ce446c63c19)
			}
		}
	}
	return __obf_23893e02a0a17784, nil
}

func __obf_e3e08f3a07ae9023(__obf_e72b611da4a59da3 string, __obf_4a9263fe8a776aa1 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_e72b611da4a59da3 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_018d54dc16e0b479, __obf_a2328d0329a8f982 := __obf_4a9263fe8a776aa1[__obf_e72b611da4a59da3]; __obf_a2328d0329a8f982 {
		return __obf_018d54dc16e0b479, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
