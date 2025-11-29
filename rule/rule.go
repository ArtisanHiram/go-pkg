package __obf_4632c1c5949287c8

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
	__obf_3564b102286d1564 = map[string]__obf_cc682e3fe35d1f7a{
		"contains": func(__obf_f9c9672b7b834099 []any) (any, error) {
			fmt.Println("contains print: ", __obf_f9c9672b7b834099)
			return nil, nil
		},
	}
)

type __obf_cc682e3fe35d1f7a = func(__obf_f9c9672b7b834099 []any) (any, error)

type Rule struct {
	__obf_a37a777cf8109bda ast.Expr
}

func (__obf_51ffc2b4f9cecd25 *Rule) SetExpr(__obf_a37a777cf8109bda string) error {
	if len(__obf_a37a777cf8109bda) == 0 {
		return ErrRuleEmpty
	}
	if __obf_9725a1955720dc33, __obf_f1d0e26b3a5fdda0 := parser.ParseExpr(__obf_a37a777cf8109bda); __obf_f1d0e26b3a5fdda0 != nil {
		return __obf_f1d0e26b3a5fdda0
	} else {
		__obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda = __obf_9725a1955720dc33
	}

	return nil
}

func (__obf_51ffc2b4f9cecd25 *Rule) Bool(__obf_dd3a757f6147591a map[string]any) (bool, error) {
	if __obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda != nil {
		__obf_79d6e23a7e30df98, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_dd3a757f6147591a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return false, __obf_f1d0e26b3a5fdda0
		}
		if __obf_51ffc2b4f9cecd25, __obf_ca94ed43b7b19a38 := __obf_79d6e23a7e30df98.(bool); __obf_ca94ed43b7b19a38 {
			return __obf_51ffc2b4f9cecd25, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_51ffc2b4f9cecd25 *Rule) Int(__obf_dd3a757f6147591a map[string]any) (int64, error) {
	if __obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda != nil {
		__obf_79d6e23a7e30df98, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_dd3a757f6147591a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return 0, __obf_f1d0e26b3a5fdda0
		}
		switch __obf_79d6e23a7e30df98 := __obf_79d6e23a7e30df98.(type) {
		case int64:
			return __obf_79d6e23a7e30df98, nil
		case float64:
			return int64(__obf_79d6e23a7e30df98), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_51ffc2b4f9cecd25 *Rule) Float(__obf_dd3a757f6147591a map[string]any) (float64, error) {
	if __obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda != nil {
		__obf_79d6e23a7e30df98, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_dd3a757f6147591a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return 0, __obf_f1d0e26b3a5fdda0
		}
		switch __obf_79d6e23a7e30df98 := __obf_79d6e23a7e30df98.(type) {
		case int64:
			return float64(__obf_79d6e23a7e30df98), nil
		case float64:
			return __obf_79d6e23a7e30df98, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_51ffc2b4f9cecd25 *Rule) Eval(__obf_eba11748592ab33a map[string]any) (any, error) {
	switch __obf_7e1dc10f5620b5e5 := __obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda.(type) {
	case *ast.UnaryExpr:
		__obf_51ffc2b4f9cecd25. // 一元表达式
					__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.X
		__obf_40266e744d09afb9, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		__obf_a1978119b657a474 := reflect.ValueOf(__obf_40266e744d09afb9)

		switch __obf_7e1dc10f5620b5e5.Op {
		case token.NOT: // !
			if __obf_a1978119b657a474.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_a1978119b657a474.Bool(), nil
		case token.SUB: // -
			if __obf_559c9a6e4ad26b2a, __obf_f1d0e26b3a5fdda0 := __obf_a146566d6dfec01e(__obf_a1978119b657a474); __obf_f1d0e26b3a5fdda0 == nil {
				return (-1.0) * __obf_559c9a6e4ad26b2a, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_51ffc2b4f9cecd25. // 二元表达式
					__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.X
		__obf_559c9a6e4ad26b2a, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		__obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.Y
		__obf_81f0fb6882d5851e, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		return __obf_a9c7fcdfcce65ec0(__obf_559c9a6e4ad26b2a, __obf_81f0fb6882d5851e,
			// 标志符（已定义变量或常量（bool））
			__obf_7e1dc10f5620b5e5.Op)
	case *ast.Ident:
		return __obf_9ffb63398bb34147(__obf_7e1dc10f5620b5e5.Name, __obf_eba11748592ab33a)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_7e1dc10f5620b5e5.Kind {
		case token.STRING:
			return strings.Trim(__obf_7e1dc10f5620b5e5.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_7e1dc10f5620b5e5.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_7e1dc10f5620b5e5.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_51ffc2b4f9cecd25. // 圆括号内表达式
					__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.X
		return __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
	case *ast.SelectorExpr:
		__obf_51ffc2b4f9cecd25. // 属性或方法选择表达式
					__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.X
		__obf_2a0921cfa246cfc7, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		return __obf_9ffb63398bb34147(__obf_7e1dc10f5620b5e5.Sel.Name, __obf_2a0921cfa246cfc7.(map[string]any))
	case *ast.IndexExpr:
		__obf_51ffc2b4f9cecd25. // 中括号内表达式——map或slice索引
					__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.X
		__obf_7c423a8c2d5026f0, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		__obf_51ffc2b4f9cecd25.__obf_a37a777cf8109bda = __obf_7e1dc10f5620b5e5.Index
		__obf_2bc02e8c1742344e, __obf_f1d0e26b3a5fdda0 := __obf_51ffc2b4f9cecd25.Eval(__obf_eba11748592ab33a)
		if __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}

		switch __obf_7c423a8c2d5026f0 := __obf_7c423a8c2d5026f0.(type) {
		case map[string]any:
			if __obf_2bc02e8c1742344e, __obf_97dd4932633b53ef := __obf_2bc02e8c1742344e.(string); __obf_97dd4932633b53ef {
				return __obf_7c423a8c2d5026f0[__obf_2bc02e8c1742344e], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_2bc02e8c1742344e := __obf_2bc02e8c1742344e.(type) {
			case int:
				return __obf_7c423a8c2d5026f0[int64(__obf_2bc02e8c1742344e)], nil
			case int64:
				return __obf_7c423a8c2d5026f0[__obf_2bc02e8c1742344e], nil
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
		if __obf_16611959a7eb841f, __obf_f1d0e26b3a5fdda0 := __obf_10fe8f4624b5829e(__obf_7e1dc10f5620b5e5.Args, __obf_eba11748592ab33a); __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		} else {
			return __obf_3564b102286d1564[__obf_7e1dc10f5620b5e5.Fun.(*ast.Ident).Name](__obf_16611959a7eb841f)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_10fe8f4624b5829e(__obf_f9c9672b7b834099 []ast.Expr, __obf_eba11748592ab33a map[string]any) ([]any, error) {
	var __obf_af480e6ac9ce9afd []any
	for _, __obf_1dc82a4265cc3499 := range __obf_f9c9672b7b834099 {
		switch __obf_1dc82a4265cc3499 := __obf_1dc82a4265cc3499.(type) {
		case *ast.BasicLit:
			__obf_af480e6ac9ce9afd = append(__obf_af480e6ac9ce9afd, __obf_1dc82a4265cc3499.Value)
		case *ast.Ident:
			if __obf_fbe4c11ca4374da1, __obf_f1d0e26b3a5fdda0 := __obf_9ffb63398bb34147(__obf_1dc82a4265cc3499.Name, __obf_eba11748592ab33a); __obf_f1d0e26b3a5fdda0 != nil {
				return nil, __obf_f1d0e26b3a5fdda0
			} else {
				__obf_af480e6ac9ce9afd = append(__obf_af480e6ac9ce9afd, __obf_fbe4c11ca4374da1)
			}
		}
	}
	return __obf_af480e6ac9ce9afd, nil
}

func __obf_9ffb63398bb34147(__obf_7f3b4d04cbfda917 string, __obf_eba11748592ab33a map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_7f3b4d04cbfda917 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_e9690b2b8adad6d7, __obf_ca94ed43b7b19a38 := __obf_eba11748592ab33a[__obf_7f3b4d04cbfda917]; __obf_ca94ed43b7b19a38 {
		return __obf_e9690b2b8adad6d7, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
