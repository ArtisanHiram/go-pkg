package __obf_3bb823b8917aab5c

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
	__obf_99d8d22445cb0330 = map[string]__obf_c2917a70c5022469{
		"contains": func(__obf_421e26a81b6c2203 []any) (any, error) {
			fmt.Println("contains print: ", __obf_421e26a81b6c2203)
			return nil, nil
		},
	}
)

type __obf_c2917a70c5022469 = func(__obf_421e26a81b6c2203 []any) (any, error)

type Rule struct {
	__obf_69625f838b72b86c ast.Expr
}

func (__obf_e813c30e01f42201 *Rule) SetExpr(__obf_69625f838b72b86c string) error {
	if len(__obf_69625f838b72b86c) == 0 {
		return ErrRuleEmpty
	}
	if __obf_2a95ed2ebff05c3a, __obf_0057b8b996f24b37 := parser.ParseExpr(__obf_69625f838b72b86c); __obf_0057b8b996f24b37 != nil {
		return __obf_0057b8b996f24b37
	} else {
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_2a95ed2ebff05c3a
	}

	return nil
}

func (__obf_e813c30e01f42201 *Rule) Bool(__obf_19e4867cc71ae295 map[string]any) (bool, error) {
	if __obf_e813c30e01f42201.__obf_69625f838b72b86c != nil {
		__obf_856b8ce413ab0dcf, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_19e4867cc71ae295)
		if __obf_0057b8b996f24b37 != nil {
			return false, __obf_0057b8b996f24b37
		}
		if __obf_e813c30e01f42201, __obf_af26dac4852a5931 := __obf_856b8ce413ab0dcf.(bool); __obf_af26dac4852a5931 {
			return __obf_e813c30e01f42201, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_e813c30e01f42201 *Rule) Int(__obf_19e4867cc71ae295 map[string]any) (int64, error) {
	if __obf_e813c30e01f42201.__obf_69625f838b72b86c != nil {
		__obf_856b8ce413ab0dcf, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_19e4867cc71ae295)
		if __obf_0057b8b996f24b37 != nil {
			return 0, __obf_0057b8b996f24b37
		}
		switch __obf_856b8ce413ab0dcf := __obf_856b8ce413ab0dcf.(type) {
		case int64:
			return __obf_856b8ce413ab0dcf, nil
		case float64:
			return int64(__obf_856b8ce413ab0dcf), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_e813c30e01f42201 *Rule) Float(__obf_19e4867cc71ae295 map[string]any) (float64, error) {
	if __obf_e813c30e01f42201.__obf_69625f838b72b86c != nil {
		__obf_856b8ce413ab0dcf, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_19e4867cc71ae295)
		if __obf_0057b8b996f24b37 != nil {
			return 0, __obf_0057b8b996f24b37
		}
		switch __obf_856b8ce413ab0dcf := __obf_856b8ce413ab0dcf.(type) {
		case int64:
			return float64(__obf_856b8ce413ab0dcf), nil
		case float64:
			return __obf_856b8ce413ab0dcf, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_e813c30e01f42201 *Rule) Eval(__obf_3989bac34858bb11 map[string]any) (any, error) {
	switch __obf_512a4c937773992f := __obf_e813c30e01f42201.__obf_69625f838b72b86c.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.X
		__obf_51799ad778c95580, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		__obf_91f8078d0306e81a := reflect.ValueOf(__obf_51799ad778c95580)

		switch __obf_512a4c937773992f.Op {
		case token.NOT: // !
			if __obf_91f8078d0306e81a.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_91f8078d0306e81a.Bool(), nil
		case token.SUB: // -
			if __obf_a27a3582cc4e1e1d, __obf_0057b8b996f24b37 := __obf_2b3f156b123f0c7b(__obf_91f8078d0306e81a); __obf_0057b8b996f24b37 == nil {
				return (-1.0) * __obf_a27a3582cc4e1e1d, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.X
		__obf_a27a3582cc4e1e1d, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.Y
		__obf_4a8fe9c05037d3f5, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		return __obf_2abb3b12b1ae34a3(__obf_a27a3582cc4e1e1d, __obf_4a8fe9c05037d3f5, __obf_512a4c937773992f.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_e60abd8066e6afb5(__obf_512a4c937773992f.Name, __obf_3989bac34858bb11)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_512a4c937773992f.Kind {
		case token.STRING:
			return strings.Trim(__obf_512a4c937773992f.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_512a4c937773992f.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_512a4c937773992f.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.X
		return __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.X
		__obf_f15bc952fa827cf7, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		return __obf_e60abd8066e6afb5(__obf_512a4c937773992f.Sel.Name, __obf_f15bc952fa827cf7.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.X
		__obf_9aee5ce4a4684146, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}

		__obf_e813c30e01f42201.__obf_69625f838b72b86c = __obf_512a4c937773992f.Index
		__obf_96ae5f9a6fc5e7ab, __obf_0057b8b996f24b37 := __obf_e813c30e01f42201.Eval(__obf_3989bac34858bb11)
		if __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}

		switch __obf_9aee5ce4a4684146 := __obf_9aee5ce4a4684146.(type) {
		case map[string]any:
			if __obf_96ae5f9a6fc5e7ab, __obf_e087be3d89228104 := __obf_96ae5f9a6fc5e7ab.(string); __obf_e087be3d89228104 {
				return __obf_9aee5ce4a4684146[__obf_96ae5f9a6fc5e7ab], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_96ae5f9a6fc5e7ab := __obf_96ae5f9a6fc5e7ab.(type) {
			case int:
				return __obf_9aee5ce4a4684146[int64(__obf_96ae5f9a6fc5e7ab)], nil
			case int64:
				return __obf_9aee5ce4a4684146[__obf_96ae5f9a6fc5e7ab], nil
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
		if __obf_78cd8679dd463529, __obf_0057b8b996f24b37 := __obf_57b44f718bec262a(__obf_512a4c937773992f.Args, __obf_3989bac34858bb11); __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		} else {
			return __obf_99d8d22445cb0330[__obf_512a4c937773992f.Fun.(*ast.Ident).Name](__obf_78cd8679dd463529)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_57b44f718bec262a(__obf_421e26a81b6c2203 []ast.Expr, __obf_3989bac34858bb11 map[string]any) ([]any, error) {
	var __obf_8d3d0673ab63bf33 []any
	for _, __obf_b034d8d21847c172 := range __obf_421e26a81b6c2203 {
		switch __obf_b034d8d21847c172 := __obf_b034d8d21847c172.(type) {
		case *ast.BasicLit:
			__obf_8d3d0673ab63bf33 = append(__obf_8d3d0673ab63bf33, __obf_b034d8d21847c172.Value)
		case *ast.Ident:
			if __obf_5cca100195f97733, __obf_0057b8b996f24b37 := __obf_e60abd8066e6afb5(__obf_b034d8d21847c172.Name, __obf_3989bac34858bb11); __obf_0057b8b996f24b37 != nil {
				return nil, __obf_0057b8b996f24b37
			} else {
				__obf_8d3d0673ab63bf33 = append(__obf_8d3d0673ab63bf33, __obf_5cca100195f97733)
			}
		}
	}
	return __obf_8d3d0673ab63bf33, nil
}

func __obf_e60abd8066e6afb5(__obf_8ef6abf0e9c78634 string, __obf_3989bac34858bb11 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_8ef6abf0e9c78634 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_c039240184e0549b, __obf_af26dac4852a5931 := __obf_3989bac34858bb11[__obf_8ef6abf0e9c78634]; __obf_af26dac4852a5931 {
		return __obf_c039240184e0549b, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
