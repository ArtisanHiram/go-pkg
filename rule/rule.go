package __obf_0b4da6ad90a2d4d3

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
	__obf_6f944f1a1a99f934 = map[string]__obf_a867896048d3545f{
		"contains": func(__obf_d2efae6b9b1fbfa1 []any) (any, error) {
			fmt.Println("contains print: ", __obf_d2efae6b9b1fbfa1)
			return nil, nil
		},
	}
)

type __obf_a867896048d3545f = func(__obf_d2efae6b9b1fbfa1 []any) (any, error)

type Rule struct {
	__obf_12bd2c64d4c36270 ast.Expr
}

func (__obf_82892e6a91203334 *Rule) SetExpr(__obf_12bd2c64d4c36270 string) error {
	if len(__obf_12bd2c64d4c36270) == 0 {
		return ErrRuleEmpty
	}
	if __obf_cf5b893703796b91, __obf_c7400825d595bd55 := parser.ParseExpr(__obf_12bd2c64d4c36270); __obf_c7400825d595bd55 != nil {
		return __obf_c7400825d595bd55
	} else {
		__obf_82892e6a91203334.__obf_12bd2c64d4c36270 = __obf_cf5b893703796b91
	}

	return nil
}

func (__obf_82892e6a91203334 *Rule) Bool(__obf_9e040d08fbf7bd95 map[string]any) (bool, error) {
	if __obf_82892e6a91203334.__obf_12bd2c64d4c36270 != nil {
		__obf_274b92d14f5bccb8, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_9e040d08fbf7bd95)
		if __obf_c7400825d595bd55 != nil {
			return false, __obf_c7400825d595bd55
		}
		if __obf_82892e6a91203334, __obf_070d6069f409afc5 := __obf_274b92d14f5bccb8.(bool); __obf_070d6069f409afc5 {
			return __obf_82892e6a91203334, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_82892e6a91203334 *Rule) Int(__obf_9e040d08fbf7bd95 map[string]any) (int64, error) {
	if __obf_82892e6a91203334.__obf_12bd2c64d4c36270 != nil {
		__obf_274b92d14f5bccb8, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_9e040d08fbf7bd95)
		if __obf_c7400825d595bd55 != nil {
			return 0, __obf_c7400825d595bd55
		}
		switch __obf_274b92d14f5bccb8 := __obf_274b92d14f5bccb8.(type) {
		case int64:
			return __obf_274b92d14f5bccb8, nil
		case float64:
			return int64(__obf_274b92d14f5bccb8), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_82892e6a91203334 *Rule) Float(__obf_9e040d08fbf7bd95 map[string]any) (float64, error) {
	if __obf_82892e6a91203334.__obf_12bd2c64d4c36270 != nil {
		__obf_274b92d14f5bccb8, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_9e040d08fbf7bd95)
		if __obf_c7400825d595bd55 != nil {
			return 0, __obf_c7400825d595bd55
		}
		switch __obf_274b92d14f5bccb8 := __obf_274b92d14f5bccb8.(type) {
		case int64:
			return float64(__obf_274b92d14f5bccb8), nil
		case float64:
			return __obf_274b92d14f5bccb8, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_82892e6a91203334 *Rule) Eval(__obf_8ec3fbaa1fb82cbd map[string]any) (any, error) {
	switch __obf_36ac70e6d1f6b8d6 := __obf_82892e6a91203334.__obf_12bd2c64d4c36270.(type) {
	case *ast.UnaryExpr:
		__obf_82892e6a91203334. // 一元表达式
					__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.X
		__obf_8a881106d432e99a, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		__obf_0bf28cfc3033ff73 := reflect.ValueOf(__obf_8a881106d432e99a)

		switch __obf_36ac70e6d1f6b8d6.Op {
		case token.NOT: // !
			if __obf_0bf28cfc3033ff73.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_0bf28cfc3033ff73.Bool(), nil
		case token.SUB: // -
			if __obf_438ec6046ae75ad9, __obf_c7400825d595bd55 := __obf_d5fbfc4a5302e6e5(__obf_0bf28cfc3033ff73); __obf_c7400825d595bd55 == nil {
				return (-1.0) * __obf_438ec6046ae75ad9, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_82892e6a91203334. // 二元表达式
					__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.X
		__obf_438ec6046ae75ad9, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		__obf_82892e6a91203334.__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.Y
		__obf_1ddf299c12ea5b7f, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		return __obf_588bd91612f08558(__obf_438ec6046ae75ad9, __obf_1ddf299c12ea5b7f,
			// 标志符（已定义变量或常量（bool））
			__obf_36ac70e6d1f6b8d6.Op)
	case *ast.Ident:
		return __obf_98ba08426b4f4b39(__obf_36ac70e6d1f6b8d6.Name, __obf_8ec3fbaa1fb82cbd)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_36ac70e6d1f6b8d6.Kind {
		case token.STRING:
			return strings.Trim(__obf_36ac70e6d1f6b8d6.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_36ac70e6d1f6b8d6.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_36ac70e6d1f6b8d6.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_82892e6a91203334. // 圆括号内表达式
					__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.X
		return __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
	case *ast.SelectorExpr:
		__obf_82892e6a91203334. // 属性或方法选择表达式
					__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.X
		__obf_e15e078ccb3f02f9, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		return __obf_98ba08426b4f4b39(__obf_36ac70e6d1f6b8d6.Sel.Name, __obf_e15e078ccb3f02f9.(map[string]any))
	case *ast.IndexExpr:
		__obf_82892e6a91203334. // 中括号内表达式——map或slice索引
					__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.X
		__obf_2c6ae1bb8dc81770, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		__obf_82892e6a91203334.__obf_12bd2c64d4c36270 = __obf_36ac70e6d1f6b8d6.Index
		__obf_270ca97dad171aa9, __obf_c7400825d595bd55 := __obf_82892e6a91203334.Eval(__obf_8ec3fbaa1fb82cbd)
		if __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}

		switch __obf_2c6ae1bb8dc81770 := __obf_2c6ae1bb8dc81770.(type) {
		case map[string]any:
			if __obf_270ca97dad171aa9, __obf_3d214b63a2b7d509 := __obf_270ca97dad171aa9.(string); __obf_3d214b63a2b7d509 {
				return __obf_2c6ae1bb8dc81770[__obf_270ca97dad171aa9], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_270ca97dad171aa9 := __obf_270ca97dad171aa9.(type) {
			case int:
				return __obf_2c6ae1bb8dc81770[int64(__obf_270ca97dad171aa9)], nil
			case int64:
				return __obf_2c6ae1bb8dc81770[__obf_270ca97dad171aa9], nil
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
		if __obf_e8ad441849e966f9, __obf_c7400825d595bd55 := __obf_b86905c56ed5441c(__obf_36ac70e6d1f6b8d6.Args, __obf_8ec3fbaa1fb82cbd); __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		} else {
			return __obf_6f944f1a1a99f934[__obf_36ac70e6d1f6b8d6.Fun.(*ast.Ident).Name](__obf_e8ad441849e966f9)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_b86905c56ed5441c(__obf_d2efae6b9b1fbfa1 []ast.Expr, __obf_8ec3fbaa1fb82cbd map[string]any) ([]any, error) {
	var __obf_b88483b17dae7c80 []any
	for _, __obf_7adab497e85db29d := range __obf_d2efae6b9b1fbfa1 {
		switch __obf_7adab497e85db29d := __obf_7adab497e85db29d.(type) {
		case *ast.BasicLit:
			__obf_b88483b17dae7c80 = append(__obf_b88483b17dae7c80, __obf_7adab497e85db29d.Value)
		case *ast.Ident:
			if __obf_6b02e49e2583a961, __obf_c7400825d595bd55 := __obf_98ba08426b4f4b39(__obf_7adab497e85db29d.Name, __obf_8ec3fbaa1fb82cbd); __obf_c7400825d595bd55 != nil {
				return nil, __obf_c7400825d595bd55
			} else {
				__obf_b88483b17dae7c80 = append(__obf_b88483b17dae7c80, __obf_6b02e49e2583a961)
			}
		}
	}
	return __obf_b88483b17dae7c80, nil
}

func __obf_98ba08426b4f4b39(__obf_8101cb4b80ab2979 string, __obf_8ec3fbaa1fb82cbd map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_8101cb4b80ab2979 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_9c3baba90902c642, __obf_070d6069f409afc5 := __obf_8ec3fbaa1fb82cbd[__obf_8101cb4b80ab2979]; __obf_070d6069f409afc5 {
		return __obf_9c3baba90902c642, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
