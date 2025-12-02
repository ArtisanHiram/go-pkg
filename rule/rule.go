package __obf_d3be831d1774c7b9

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
	__obf_b89ad499c1291a8b = map[string]__obf_998c92e34be2448b{
		"contains": func(__obf_1cc74566e725673f []any) (any, error) {
			fmt.Println("contains print: ", __obf_1cc74566e725673f)
			return nil, nil
		},
	}
)

type __obf_998c92e34be2448b = func(__obf_1cc74566e725673f []any) (any, error)

type Rule struct {
	__obf_835fd2be7cc2d9c3 ast.Expr
}

func (__obf_76f46044fa4b8bb3 *Rule) SetExpr(__obf_835fd2be7cc2d9c3 string) error {
	if len(__obf_835fd2be7cc2d9c3) == 0 {
		return ErrRuleEmpty
	}
	if __obf_bb6fd5d94b30f9ab, __obf_df5f028324461f18 := parser.ParseExpr(__obf_835fd2be7cc2d9c3); __obf_df5f028324461f18 != nil {
		return __obf_df5f028324461f18
	} else {
		__obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 = __obf_bb6fd5d94b30f9ab
	}

	return nil
}

func (__obf_76f46044fa4b8bb3 *Rule) Bool(__obf_1ed209e307b6c597 map[string]any) (bool, error) {
	if __obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 != nil {
		__obf_3b73e3235cb8d0f9, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_1ed209e307b6c597)
		if __obf_df5f028324461f18 != nil {
			return false, __obf_df5f028324461f18
		}
		if __obf_76f46044fa4b8bb3, __obf_b17963af6726dc93 := __obf_3b73e3235cb8d0f9.(bool); __obf_b17963af6726dc93 {
			return __obf_76f46044fa4b8bb3, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_76f46044fa4b8bb3 *Rule) Int(__obf_1ed209e307b6c597 map[string]any) (int64, error) {
	if __obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 != nil {
		__obf_3b73e3235cb8d0f9, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_1ed209e307b6c597)
		if __obf_df5f028324461f18 != nil {
			return 0, __obf_df5f028324461f18
		}
		switch __obf_3b73e3235cb8d0f9 := __obf_3b73e3235cb8d0f9.(type) {
		case int64:
			return __obf_3b73e3235cb8d0f9, nil
		case float64:
			return int64(__obf_3b73e3235cb8d0f9), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_76f46044fa4b8bb3 *Rule) Float(__obf_1ed209e307b6c597 map[string]any) (float64, error) {
	if __obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 != nil {
		__obf_3b73e3235cb8d0f9, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_1ed209e307b6c597)
		if __obf_df5f028324461f18 != nil {
			return 0, __obf_df5f028324461f18
		}
		switch __obf_3b73e3235cb8d0f9 := __obf_3b73e3235cb8d0f9.(type) {
		case int64:
			return float64(__obf_3b73e3235cb8d0f9), nil
		case float64:
			return __obf_3b73e3235cb8d0f9, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_76f46044fa4b8bb3 *Rule) Eval(__obf_a18974084b3d7c9e map[string]any) (any, error) {
	switch __obf_773c991dab8c5f98 := __obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3.(type) {
	case *ast.UnaryExpr:
		__obf_76f46044fa4b8bb3. // 一元表达式
					__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.X
		__obf_f636f4a29fb3301c, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		__obf_16dc3b2a0ccfa4d7 := reflect.ValueOf(__obf_f636f4a29fb3301c)

		switch __obf_773c991dab8c5f98.Op {
		case token.NOT: // !
			if __obf_16dc3b2a0ccfa4d7.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_16dc3b2a0ccfa4d7.Bool(), nil
		case token.SUB: // -
			if __obf_ca275d416555ddd0, __obf_df5f028324461f18 := __obf_55b9fda27ba97528(__obf_16dc3b2a0ccfa4d7); __obf_df5f028324461f18 == nil {
				return (-1.0) * __obf_ca275d416555ddd0, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_76f46044fa4b8bb3. // 二元表达式
					__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.X
		__obf_ca275d416555ddd0, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		__obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.Y
		__obf_84db90cea1ee5669, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		return __obf_4a4ddbfa6fc0ec8e(__obf_ca275d416555ddd0, __obf_84db90cea1ee5669,
			// 标志符（已定义变量或常量（bool））
			__obf_773c991dab8c5f98.Op)
	case *ast.Ident:
		return __obf_5a34ccf92dd52180(__obf_773c991dab8c5f98.Name, __obf_a18974084b3d7c9e)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_773c991dab8c5f98.Kind {
		case token.STRING:
			return strings.Trim(__obf_773c991dab8c5f98.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_773c991dab8c5f98.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_773c991dab8c5f98.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_76f46044fa4b8bb3. // 圆括号内表达式
					__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.X
		return __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
	case *ast.SelectorExpr:
		__obf_76f46044fa4b8bb3. // 属性或方法选择表达式
					__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.X
		__obf_adeffbb923cb1920, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		return __obf_5a34ccf92dd52180(__obf_773c991dab8c5f98.Sel.Name, __obf_adeffbb923cb1920.(map[string]any))
	case *ast.IndexExpr:
		__obf_76f46044fa4b8bb3. // 中括号内表达式——map或slice索引
					__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.X
		__obf_4fac41c60eac60a4, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		__obf_76f46044fa4b8bb3.__obf_835fd2be7cc2d9c3 = __obf_773c991dab8c5f98.Index
		__obf_b132a2660d94e198, __obf_df5f028324461f18 := __obf_76f46044fa4b8bb3.Eval(__obf_a18974084b3d7c9e)
		if __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}

		switch __obf_4fac41c60eac60a4 := __obf_4fac41c60eac60a4.(type) {
		case map[string]any:
			if __obf_b132a2660d94e198, __obf_0ee1de645f2ee7c2 := __obf_b132a2660d94e198.(string); __obf_0ee1de645f2ee7c2 {
				return __obf_4fac41c60eac60a4[__obf_b132a2660d94e198], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_b132a2660d94e198 := __obf_b132a2660d94e198.(type) {
			case int:
				return __obf_4fac41c60eac60a4[int64(__obf_b132a2660d94e198)], nil
			case int64:
				return __obf_4fac41c60eac60a4[__obf_b132a2660d94e198], nil
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
		if __obf_fde2790d8a25b976, __obf_df5f028324461f18 := __obf_d8d7e324805180b5(__obf_773c991dab8c5f98.Args, __obf_a18974084b3d7c9e); __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		} else {
			return __obf_b89ad499c1291a8b[__obf_773c991dab8c5f98.Fun.(*ast.Ident).Name](__obf_fde2790d8a25b976)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_d8d7e324805180b5(__obf_1cc74566e725673f []ast.Expr, __obf_a18974084b3d7c9e map[string]any) ([]any, error) {
	var __obf_c453b68a304b786d []any
	for _, __obf_e39fe6881aecaedf := range __obf_1cc74566e725673f {
		switch __obf_e39fe6881aecaedf := __obf_e39fe6881aecaedf.(type) {
		case *ast.BasicLit:
			__obf_c453b68a304b786d = append(__obf_c453b68a304b786d, __obf_e39fe6881aecaedf.Value)
		case *ast.Ident:
			if __obf_cc69cfda9f6630dd, __obf_df5f028324461f18 := __obf_5a34ccf92dd52180(__obf_e39fe6881aecaedf.Name, __obf_a18974084b3d7c9e); __obf_df5f028324461f18 != nil {
				return nil, __obf_df5f028324461f18
			} else {
				__obf_c453b68a304b786d = append(__obf_c453b68a304b786d, __obf_cc69cfda9f6630dd)
			}
		}
	}
	return __obf_c453b68a304b786d, nil
}

func __obf_5a34ccf92dd52180(__obf_cde9e9dbfcc48cac string, __obf_a18974084b3d7c9e map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_cde9e9dbfcc48cac {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_6a92e22c2d6c8251, __obf_b17963af6726dc93 := __obf_a18974084b3d7c9e[__obf_cde9e9dbfcc48cac]; __obf_b17963af6726dc93 {
		return __obf_6a92e22c2d6c8251, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
