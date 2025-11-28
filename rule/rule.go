package __obf_558a7db000042ed1

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
	__obf_3400fc0cde864fc1 = map[string]__obf_454237873e8fa003{
		"contains": func(__obf_8e7bfd787f3b78e8 []any) (any, error) {
			fmt.Println("contains print: ", __obf_8e7bfd787f3b78e8)
			return nil, nil
		},
	}
)

type __obf_454237873e8fa003 = func(__obf_8e7bfd787f3b78e8 []any) (any, error)

type Rule struct {
	__obf_3c74d690cb5cf7fe ast.Expr
}

func (__obf_0e0272a7cd2df0ea *Rule) SetExpr(__obf_3c74d690cb5cf7fe string) error {
	if len(__obf_3c74d690cb5cf7fe) == 0 {
		return ErrRuleEmpty
	}
	if __obf_6c30c6d3a8dab959, __obf_e6c13dab132e0931 := parser.ParseExpr(__obf_3c74d690cb5cf7fe); __obf_e6c13dab132e0931 != nil {
		return __obf_e6c13dab132e0931
	} else {
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_6c30c6d3a8dab959
	}

	return nil
}

func (__obf_0e0272a7cd2df0ea *Rule) Bool(__obf_0e6a429f06d1cd98 map[string]any) (bool, error) {
	if __obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe != nil {
		__obf_106b3ecc6336db1c, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_0e6a429f06d1cd98)
		if __obf_e6c13dab132e0931 != nil {
			return false, __obf_e6c13dab132e0931
		}
		if __obf_0e0272a7cd2df0ea, __obf_2a137c3b5014c8f7 := __obf_106b3ecc6336db1c.(bool); __obf_2a137c3b5014c8f7 {
			return __obf_0e0272a7cd2df0ea, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_0e0272a7cd2df0ea *Rule) Int(__obf_0e6a429f06d1cd98 map[string]any) (int64, error) {
	if __obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe != nil {
		__obf_106b3ecc6336db1c, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_0e6a429f06d1cd98)
		if __obf_e6c13dab132e0931 != nil {
			return 0, __obf_e6c13dab132e0931
		}
		switch __obf_106b3ecc6336db1c := __obf_106b3ecc6336db1c.(type) {
		case int64:
			return __obf_106b3ecc6336db1c, nil
		case float64:
			return int64(__obf_106b3ecc6336db1c), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_0e0272a7cd2df0ea *Rule) Float(__obf_0e6a429f06d1cd98 map[string]any) (float64, error) {
	if __obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe != nil {
		__obf_106b3ecc6336db1c, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_0e6a429f06d1cd98)
		if __obf_e6c13dab132e0931 != nil {
			return 0, __obf_e6c13dab132e0931
		}
		switch __obf_106b3ecc6336db1c := __obf_106b3ecc6336db1c.(type) {
		case int64:
			return float64(__obf_106b3ecc6336db1c), nil
		case float64:
			return __obf_106b3ecc6336db1c, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_0e0272a7cd2df0ea *Rule) Eval(__obf_be9f9e408d3837ec map[string]any) (any, error) {
	switch __obf_4adc35d9b0dfd3a0 := __obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.X
		__obf_a788789d8f610e50, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		__obf_6ffe26da8be643e5 := reflect.ValueOf(__obf_a788789d8f610e50)

		switch __obf_4adc35d9b0dfd3a0.Op {
		case token.NOT: // !
			if __obf_6ffe26da8be643e5.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_6ffe26da8be643e5.Bool(), nil
		case token.SUB: // -
			if __obf_49265ee5f1ae335d, __obf_e6c13dab132e0931 := __obf_c21732ed5660156f(__obf_6ffe26da8be643e5); __obf_e6c13dab132e0931 == nil {
				return (-1.0) * __obf_49265ee5f1ae335d, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.X
		__obf_49265ee5f1ae335d, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.Y
		__obf_6e7f0359ff738fe4, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		return __obf_711c936c8cf651bc(__obf_49265ee5f1ae335d, __obf_6e7f0359ff738fe4, __obf_4adc35d9b0dfd3a0.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_5b98d9599723eef3(__obf_4adc35d9b0dfd3a0.Name, __obf_be9f9e408d3837ec)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_4adc35d9b0dfd3a0.Kind {
		case token.STRING:
			return strings.Trim(__obf_4adc35d9b0dfd3a0.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_4adc35d9b0dfd3a0.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_4adc35d9b0dfd3a0.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.X
		return __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.X
		__obf_44d04258ba957d93, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		return __obf_5b98d9599723eef3(__obf_4adc35d9b0dfd3a0.Sel.Name, __obf_44d04258ba957d93.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.X
		__obf_aa14fbcb6893737a, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}

		__obf_0e0272a7cd2df0ea.__obf_3c74d690cb5cf7fe = __obf_4adc35d9b0dfd3a0.Index
		__obf_2ca688f83306aa8b, __obf_e6c13dab132e0931 := __obf_0e0272a7cd2df0ea.Eval(__obf_be9f9e408d3837ec)
		if __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}

		switch __obf_aa14fbcb6893737a := __obf_aa14fbcb6893737a.(type) {
		case map[string]any:
			if __obf_2ca688f83306aa8b, __obf_38f6ef2de8dbd0e4 := __obf_2ca688f83306aa8b.(string); __obf_38f6ef2de8dbd0e4 {
				return __obf_aa14fbcb6893737a[__obf_2ca688f83306aa8b], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_2ca688f83306aa8b := __obf_2ca688f83306aa8b.(type) {
			case int:
				return __obf_aa14fbcb6893737a[int64(__obf_2ca688f83306aa8b)], nil
			case int64:
				return __obf_aa14fbcb6893737a[__obf_2ca688f83306aa8b], nil
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
		if __obf_9bffa4e302d7f480, __obf_e6c13dab132e0931 := __obf_6bf6d31ebd14543b(__obf_4adc35d9b0dfd3a0.Args, __obf_be9f9e408d3837ec); __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		} else {
			return __obf_3400fc0cde864fc1[__obf_4adc35d9b0dfd3a0.Fun.(*ast.Ident).Name](__obf_9bffa4e302d7f480)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_6bf6d31ebd14543b(__obf_8e7bfd787f3b78e8 []ast.Expr, __obf_be9f9e408d3837ec map[string]any) ([]any, error) {
	var __obf_87c2df60d54f0db4 []any
	for _, __obf_34428e77ecb150b3 := range __obf_8e7bfd787f3b78e8 {
		switch __obf_34428e77ecb150b3 := __obf_34428e77ecb150b3.(type) {
		case *ast.BasicLit:
			__obf_87c2df60d54f0db4 = append(__obf_87c2df60d54f0db4, __obf_34428e77ecb150b3.Value)
		case *ast.Ident:
			if __obf_7895726caff732d7, __obf_e6c13dab132e0931 := __obf_5b98d9599723eef3(__obf_34428e77ecb150b3.Name, __obf_be9f9e408d3837ec); __obf_e6c13dab132e0931 != nil {
				return nil, __obf_e6c13dab132e0931
			} else {
				__obf_87c2df60d54f0db4 = append(__obf_87c2df60d54f0db4, __obf_7895726caff732d7)
			}
		}
	}
	return __obf_87c2df60d54f0db4, nil
}

func __obf_5b98d9599723eef3(__obf_d898ed9837d55b53 string, __obf_be9f9e408d3837ec map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_d898ed9837d55b53 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_df24342d0423f0ac, __obf_2a137c3b5014c8f7 := __obf_be9f9e408d3837ec[__obf_d898ed9837d55b53]; __obf_2a137c3b5014c8f7 {
		return __obf_df24342d0423f0ac, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
