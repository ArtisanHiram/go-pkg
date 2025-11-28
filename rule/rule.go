package __obf_564c951b28a491a4

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
	__obf_de1aa98cbfa86bee = map[string]__obf_a4aa34e26a8e0253{
		"contains": func(__obf_4e71c359f950b000 []any) (any, error) {
			fmt.Println("contains print: ", __obf_4e71c359f950b000)
			return nil, nil
		},
	}
)

type __obf_a4aa34e26a8e0253 = func(__obf_4e71c359f950b000 []any) (any, error)

type Rule struct {
	__obf_0bbf24e4bce80eab ast.Expr
}

func (__obf_ebff778f327fe057 *Rule) SetExpr(__obf_0bbf24e4bce80eab string) error {
	if len(__obf_0bbf24e4bce80eab) == 0 {
		return ErrRuleEmpty
	}
	if __obf_22d028e0d1c1b40c, __obf_a10eba5810272430 := parser.ParseExpr(__obf_0bbf24e4bce80eab); __obf_a10eba5810272430 != nil {
		return __obf_a10eba5810272430
	} else {
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_22d028e0d1c1b40c
	}

	return nil
}

func (__obf_ebff778f327fe057 *Rule) Bool(__obf_e59f8f8d78d277c3 map[string]any) (bool, error) {
	if __obf_ebff778f327fe057.__obf_0bbf24e4bce80eab != nil {
		__obf_b5e9b7666394d92f, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_e59f8f8d78d277c3)
		if __obf_a10eba5810272430 != nil {
			return false, __obf_a10eba5810272430
		}
		if __obf_ebff778f327fe057, __obf_ee9be32b06fce9c3 := __obf_b5e9b7666394d92f.(bool); __obf_ee9be32b06fce9c3 {
			return __obf_ebff778f327fe057, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_ebff778f327fe057 *Rule) Int(__obf_e59f8f8d78d277c3 map[string]any) (int64, error) {
	if __obf_ebff778f327fe057.__obf_0bbf24e4bce80eab != nil {
		__obf_b5e9b7666394d92f, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_e59f8f8d78d277c3)
		if __obf_a10eba5810272430 != nil {
			return 0, __obf_a10eba5810272430
		}
		switch __obf_b5e9b7666394d92f := __obf_b5e9b7666394d92f.(type) {
		case int64:
			return __obf_b5e9b7666394d92f, nil
		case float64:
			return int64(__obf_b5e9b7666394d92f), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_ebff778f327fe057 *Rule) Float(__obf_e59f8f8d78d277c3 map[string]any) (float64, error) {
	if __obf_ebff778f327fe057.__obf_0bbf24e4bce80eab != nil {
		__obf_b5e9b7666394d92f, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_e59f8f8d78d277c3)
		if __obf_a10eba5810272430 != nil {
			return 0, __obf_a10eba5810272430
		}
		switch __obf_b5e9b7666394d92f := __obf_b5e9b7666394d92f.(type) {
		case int64:
			return float64(__obf_b5e9b7666394d92f), nil
		case float64:
			return __obf_b5e9b7666394d92f, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_ebff778f327fe057 *Rule) Eval(__obf_53a87482d9b22f35 map[string]any) (any, error) {
	switch __obf_84dff80350103e72 := __obf_ebff778f327fe057.__obf_0bbf24e4bce80eab.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.X
		__obf_4571ec77fa3095b0, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		__obf_075f6730d4058dc4 := reflect.ValueOf(__obf_4571ec77fa3095b0)

		switch __obf_84dff80350103e72.Op {
		case token.NOT: // !
			if __obf_075f6730d4058dc4.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_075f6730d4058dc4.Bool(), nil
		case token.SUB: // -
			if __obf_c8711b6013a70c03, __obf_a10eba5810272430 := __obf_5a3fa900c94b4a96(__obf_075f6730d4058dc4); __obf_a10eba5810272430 == nil {
				return (-1.0) * __obf_c8711b6013a70c03, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.X
		__obf_c8711b6013a70c03, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.Y
		__obf_b0b1735d1adfdde8, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		return __obf_decba2110d4ae088(__obf_c8711b6013a70c03, __obf_b0b1735d1adfdde8, __obf_84dff80350103e72.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_67ea18e3e64a8f34(__obf_84dff80350103e72.Name, __obf_53a87482d9b22f35)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_84dff80350103e72.Kind {
		case token.STRING:
			return strings.Trim(__obf_84dff80350103e72.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_84dff80350103e72.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_84dff80350103e72.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.X
		return __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.X
		__obf_a3b825b44d8f8baf, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		return __obf_67ea18e3e64a8f34(__obf_84dff80350103e72.Sel.Name, __obf_a3b825b44d8f8baf.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.X
		__obf_14f2b47d14d6bd54, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}

		__obf_ebff778f327fe057.__obf_0bbf24e4bce80eab = __obf_84dff80350103e72.Index
		__obf_76f149f2bfd7b1b4, __obf_a10eba5810272430 := __obf_ebff778f327fe057.Eval(__obf_53a87482d9b22f35)
		if __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}

		switch __obf_14f2b47d14d6bd54 := __obf_14f2b47d14d6bd54.(type) {
		case map[string]any:
			if __obf_76f149f2bfd7b1b4, __obf_16c63f821341d365 := __obf_76f149f2bfd7b1b4.(string); __obf_16c63f821341d365 {
				return __obf_14f2b47d14d6bd54[__obf_76f149f2bfd7b1b4], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_76f149f2bfd7b1b4 := __obf_76f149f2bfd7b1b4.(type) {
			case int:
				return __obf_14f2b47d14d6bd54[int64(__obf_76f149f2bfd7b1b4)], nil
			case int64:
				return __obf_14f2b47d14d6bd54[__obf_76f149f2bfd7b1b4], nil
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
		if __obf_5c190b7cdc7d914b, __obf_a10eba5810272430 := __obf_5a146a9b21deab41(__obf_84dff80350103e72.Args, __obf_53a87482d9b22f35); __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		} else {
			return __obf_de1aa98cbfa86bee[__obf_84dff80350103e72.Fun.(*ast.Ident).Name](__obf_5c190b7cdc7d914b)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_5a146a9b21deab41(__obf_4e71c359f950b000 []ast.Expr, __obf_53a87482d9b22f35 map[string]any) ([]any, error) {
	var __obf_15ee217070893e86 []any
	for _, __obf_8f7e54bda0f3e140 := range __obf_4e71c359f950b000 {
		switch __obf_8f7e54bda0f3e140 := __obf_8f7e54bda0f3e140.(type) {
		case *ast.BasicLit:
			__obf_15ee217070893e86 = append(__obf_15ee217070893e86, __obf_8f7e54bda0f3e140.Value)
		case *ast.Ident:
			if __obf_1fad6add7f6c08de, __obf_a10eba5810272430 := __obf_67ea18e3e64a8f34(__obf_8f7e54bda0f3e140.Name, __obf_53a87482d9b22f35); __obf_a10eba5810272430 != nil {
				return nil, __obf_a10eba5810272430
			} else {
				__obf_15ee217070893e86 = append(__obf_15ee217070893e86, __obf_1fad6add7f6c08de)
			}
		}
	}
	return __obf_15ee217070893e86, nil
}

func __obf_67ea18e3e64a8f34(__obf_c5a83148150ad65a string, __obf_53a87482d9b22f35 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_c5a83148150ad65a {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_0b26114ba56f2fd4, __obf_ee9be32b06fce9c3 := __obf_53a87482d9b22f35[__obf_c5a83148150ad65a]; __obf_ee9be32b06fce9c3 {
		return __obf_0b26114ba56f2fd4, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
