package __obf_f3901fe9624df715

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
	__obf_de1780faa2d8865c = map[string]__obf_a8ca01e470d02718{
		"contains": func(__obf_cd0b00dbd067a135 []any) (any, error) {
			fmt.Println("contains print: ", __obf_cd0b00dbd067a135)
			return nil, nil
		},
	}
)

type __obf_a8ca01e470d02718 = func(__obf_cd0b00dbd067a135 []any) (any, error)

type Rule struct {
	__obf_8d3fdbf6cf997977 ast.Expr
}

func (__obf_4a5b75ea4cbc2990 *Rule) SetExpr(__obf_8d3fdbf6cf997977 string) error {
	if len(__obf_8d3fdbf6cf997977) == 0 {
		return ErrRuleEmpty
	}
	if __obf_0c66b490ee97206f, __obf_0868a2b24f90315b := parser.ParseExpr(__obf_8d3fdbf6cf997977); __obf_0868a2b24f90315b != nil {
		return __obf_0868a2b24f90315b
	} else {
		__obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 = __obf_0c66b490ee97206f
	}

	return nil
}

func (__obf_4a5b75ea4cbc2990 *Rule) Bool(__obf_dda270530dee4000 map[string]any) (bool, error) {
	if __obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 != nil {
		__obf_1ee3a5135b452f80, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_dda270530dee4000)
		if __obf_0868a2b24f90315b != nil {
			return false, __obf_0868a2b24f90315b
		}
		if __obf_4a5b75ea4cbc2990, __obf_34aaa00f7f8fa378 := __obf_1ee3a5135b452f80.(bool); __obf_34aaa00f7f8fa378 {
			return __obf_4a5b75ea4cbc2990, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_4a5b75ea4cbc2990 *Rule) Int(__obf_dda270530dee4000 map[string]any) (int64, error) {
	if __obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 != nil {
		__obf_1ee3a5135b452f80, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_dda270530dee4000)
		if __obf_0868a2b24f90315b != nil {
			return 0, __obf_0868a2b24f90315b
		}
		switch __obf_1ee3a5135b452f80 := __obf_1ee3a5135b452f80.(type) {
		case int64:
			return __obf_1ee3a5135b452f80, nil
		case float64:
			return int64(__obf_1ee3a5135b452f80), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_4a5b75ea4cbc2990 *Rule) Float(__obf_dda270530dee4000 map[string]any) (float64, error) {
	if __obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 != nil {
		__obf_1ee3a5135b452f80, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_dda270530dee4000)
		if __obf_0868a2b24f90315b != nil {
			return 0, __obf_0868a2b24f90315b
		}
		switch __obf_1ee3a5135b452f80 := __obf_1ee3a5135b452f80.(type) {
		case int64:
			return float64(__obf_1ee3a5135b452f80), nil
		case float64:
			return __obf_1ee3a5135b452f80, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_4a5b75ea4cbc2990 *Rule) Eval(__obf_712a4da154d62d99 map[string]any) (any, error) {
	switch __obf_774a6b130c9410ff := __obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977.(type) {
	case *ast.UnaryExpr:
		__obf_4a5b75ea4cbc2990. // 一元表达式
					__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.X
		__obf_d4a65e2f32b7c76d, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		__obf_74eea4af84c5df70 := reflect.ValueOf(__obf_d4a65e2f32b7c76d)

		switch __obf_774a6b130c9410ff.Op {
		case token.NOT: // !
			if __obf_74eea4af84c5df70.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_74eea4af84c5df70.Bool(), nil
		case token.SUB: // -
			if __obf_58db4a39c61276eb, __obf_0868a2b24f90315b := __obf_b94cf9ae4814ea68(__obf_74eea4af84c5df70); __obf_0868a2b24f90315b == nil {
				return (-1.0) * __obf_58db4a39c61276eb, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_4a5b75ea4cbc2990. // 二元表达式
					__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.X
		__obf_58db4a39c61276eb, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		__obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.Y
		__obf_d6eb3223122662eb, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		return __obf_96558f9406e98de4(__obf_58db4a39c61276eb, __obf_d6eb3223122662eb,
			// 标志符（已定义变量或常量（bool））
			__obf_774a6b130c9410ff.Op)
	case *ast.Ident:
		return __obf_501e1dade6110a40(__obf_774a6b130c9410ff.Name, __obf_712a4da154d62d99)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_774a6b130c9410ff.Kind {
		case token.STRING:
			return strings.Trim(__obf_774a6b130c9410ff.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_774a6b130c9410ff.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_774a6b130c9410ff.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_4a5b75ea4cbc2990. // 圆括号内表达式
					__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.X
		return __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
	case *ast.SelectorExpr:
		__obf_4a5b75ea4cbc2990. // 属性或方法选择表达式
					__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.X
		__obf_8e2893629deb2423, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		return __obf_501e1dade6110a40(__obf_774a6b130c9410ff.Sel.Name, __obf_8e2893629deb2423.(map[string]any))
	case *ast.IndexExpr:
		__obf_4a5b75ea4cbc2990. // 中括号内表达式——map或slice索引
					__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.X
		__obf_65c5859eca4c7c17, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		__obf_4a5b75ea4cbc2990.__obf_8d3fdbf6cf997977 = __obf_774a6b130c9410ff.Index
		__obf_bd8ef40da22b2f25, __obf_0868a2b24f90315b := __obf_4a5b75ea4cbc2990.Eval(__obf_712a4da154d62d99)
		if __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}

		switch __obf_65c5859eca4c7c17 := __obf_65c5859eca4c7c17.(type) {
		case map[string]any:
			if __obf_bd8ef40da22b2f25, __obf_155099edc421adff := __obf_bd8ef40da22b2f25.(string); __obf_155099edc421adff {
				return __obf_65c5859eca4c7c17[__obf_bd8ef40da22b2f25], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_bd8ef40da22b2f25 := __obf_bd8ef40da22b2f25.(type) {
			case int:
				return __obf_65c5859eca4c7c17[int64(__obf_bd8ef40da22b2f25)], nil
			case int64:
				return __obf_65c5859eca4c7c17[__obf_bd8ef40da22b2f25], nil
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
		if __obf_3f9af233c9920767, __obf_0868a2b24f90315b := __obf_871399955878c3f8(__obf_774a6b130c9410ff.Args, __obf_712a4da154d62d99); __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		} else {
			return __obf_de1780faa2d8865c[__obf_774a6b130c9410ff.Fun.(*ast.Ident).Name](__obf_3f9af233c9920767)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_871399955878c3f8(__obf_cd0b00dbd067a135 []ast.Expr, __obf_712a4da154d62d99 map[string]any) ([]any, error) {
	var __obf_78d326614e627bd9 []any
	for _, __obf_1b49e64587a0b574 := range __obf_cd0b00dbd067a135 {
		switch __obf_1b49e64587a0b574 := __obf_1b49e64587a0b574.(type) {
		case *ast.BasicLit:
			__obf_78d326614e627bd9 = append(__obf_78d326614e627bd9, __obf_1b49e64587a0b574.Value)
		case *ast.Ident:
			if __obf_a655f407e6f44f05, __obf_0868a2b24f90315b := __obf_501e1dade6110a40(__obf_1b49e64587a0b574.Name, __obf_712a4da154d62d99); __obf_0868a2b24f90315b != nil {
				return nil, __obf_0868a2b24f90315b
			} else {
				__obf_78d326614e627bd9 = append(__obf_78d326614e627bd9, __obf_a655f407e6f44f05)
			}
		}
	}
	return __obf_78d326614e627bd9, nil
}

func __obf_501e1dade6110a40(__obf_63b05e81f86376fb string, __obf_712a4da154d62d99 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_63b05e81f86376fb {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_5a08dfef6b19f279, __obf_34aaa00f7f8fa378 := __obf_712a4da154d62d99[__obf_63b05e81f86376fb]; __obf_34aaa00f7f8fa378 {
		return __obf_5a08dfef6b19f279, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
