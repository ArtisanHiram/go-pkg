package __obf_e8080eb36ea958ff

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
	__obf_923e6e6fede40135 = map[string]__obf_419af90fe51c3015{
		"contains": func(__obf_6324600e1e794b0c []any) (any, error) {
			fmt.Println("contains print: ", __obf_6324600e1e794b0c)
			return nil, nil
		},
	}
)

type __obf_419af90fe51c3015 = func(__obf_6324600e1e794b0c []any) (any, error)

type Rule struct {
	__obf_b1f3756442e5882e ast.Expr
}

func (__obf_ab44501bf4d825f1 *Rule) SetExpr(__obf_b1f3756442e5882e string) error {
	if len(__obf_b1f3756442e5882e) == 0 {
		return ErrRuleEmpty
	}
	if __obf_e5beab22a30de53d, __obf_2ffd626706248919 := parser.ParseExpr(__obf_b1f3756442e5882e); __obf_2ffd626706248919 != nil {
		return __obf_2ffd626706248919
	} else {
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_e5beab22a30de53d
	}

	return nil
}

func (__obf_ab44501bf4d825f1 *Rule) Bool(__obf_3f84ca434e6da69d map[string]any) (bool, error) {
	if __obf_ab44501bf4d825f1.__obf_b1f3756442e5882e != nil {
		__obf_e08e221734b46a54, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_3f84ca434e6da69d)
		if __obf_2ffd626706248919 != nil {
			return false, __obf_2ffd626706248919
		}
		if __obf_ab44501bf4d825f1, __obf_73ea2ef7dfc0a7db := __obf_e08e221734b46a54.(bool); __obf_73ea2ef7dfc0a7db {
			return __obf_ab44501bf4d825f1, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_ab44501bf4d825f1 *Rule) Int(__obf_3f84ca434e6da69d map[string]any) (int64, error) {
	if __obf_ab44501bf4d825f1.__obf_b1f3756442e5882e != nil {
		__obf_e08e221734b46a54, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_3f84ca434e6da69d)
		if __obf_2ffd626706248919 != nil {
			return 0, __obf_2ffd626706248919
		}
		switch __obf_e08e221734b46a54 := __obf_e08e221734b46a54.(type) {
		case int64:
			return __obf_e08e221734b46a54, nil
		case float64:
			return int64(__obf_e08e221734b46a54), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_ab44501bf4d825f1 *Rule) Float(__obf_3f84ca434e6da69d map[string]any) (float64, error) {
	if __obf_ab44501bf4d825f1.__obf_b1f3756442e5882e != nil {
		__obf_e08e221734b46a54, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_3f84ca434e6da69d)
		if __obf_2ffd626706248919 != nil {
			return 0, __obf_2ffd626706248919
		}
		switch __obf_e08e221734b46a54 := __obf_e08e221734b46a54.(type) {
		case int64:
			return float64(__obf_e08e221734b46a54), nil
		case float64:
			return __obf_e08e221734b46a54, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_ab44501bf4d825f1 *Rule) Eval(__obf_76688cf26d5015f4 map[string]any) (any, error) {
	switch __obf_6989e7681350d5b7 := __obf_ab44501bf4d825f1.__obf_b1f3756442e5882e.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.X
		__obf_3850fcd4f8775253, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		__obf_b14b81062fd9a10a := reflect.ValueOf(__obf_3850fcd4f8775253)

		switch __obf_6989e7681350d5b7.Op {
		case token.NOT: // !
			if __obf_b14b81062fd9a10a.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_b14b81062fd9a10a.Bool(), nil
		case token.SUB: // -
			if __obf_eb0b095223325ed2, __obf_2ffd626706248919 := __obf_71e4bcca69d8b68b(__obf_b14b81062fd9a10a); __obf_2ffd626706248919 == nil {
				return (-1.0) * __obf_eb0b095223325ed2, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.X
		__obf_eb0b095223325ed2, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.Y
		__obf_ba1dededac8b8fec, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		return __obf_dc83ec0e5567c1f8(__obf_eb0b095223325ed2, __obf_ba1dededac8b8fec, __obf_6989e7681350d5b7.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_0a0f803294160e3f(__obf_6989e7681350d5b7.Name, __obf_76688cf26d5015f4)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_6989e7681350d5b7.Kind {
		case token.STRING:
			return strings.Trim(__obf_6989e7681350d5b7.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_6989e7681350d5b7.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_6989e7681350d5b7.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.X
		return __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.X
		__obf_09f581eb202ae4f7, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		return __obf_0a0f803294160e3f(__obf_6989e7681350d5b7.Sel.Name, __obf_09f581eb202ae4f7.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.X
		__obf_d22e9f5c275f7a88, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}

		__obf_ab44501bf4d825f1.__obf_b1f3756442e5882e = __obf_6989e7681350d5b7.Index
		__obf_b11da54c387213ee, __obf_2ffd626706248919 := __obf_ab44501bf4d825f1.Eval(__obf_76688cf26d5015f4)
		if __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}

		switch __obf_d22e9f5c275f7a88 := __obf_d22e9f5c275f7a88.(type) {
		case map[string]any:
			if __obf_b11da54c387213ee, __obf_221d542066e2ef42 := __obf_b11da54c387213ee.(string); __obf_221d542066e2ef42 {
				return __obf_d22e9f5c275f7a88[__obf_b11da54c387213ee], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_b11da54c387213ee := __obf_b11da54c387213ee.(type) {
			case int:
				return __obf_d22e9f5c275f7a88[int64(__obf_b11da54c387213ee)], nil
			case int64:
				return __obf_d22e9f5c275f7a88[__obf_b11da54c387213ee], nil
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
		if __obf_296b163e608156dc, __obf_2ffd626706248919 := __obf_b3c7e3edf10c0cf4(__obf_6989e7681350d5b7.Args, __obf_76688cf26d5015f4); __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		} else {
			return __obf_923e6e6fede40135[__obf_6989e7681350d5b7.Fun.(*ast.Ident).Name](__obf_296b163e608156dc)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_b3c7e3edf10c0cf4(__obf_6324600e1e794b0c []ast.Expr, __obf_76688cf26d5015f4 map[string]any) ([]any, error) {
	var __obf_c717d4181ed7e54b []any
	for _, __obf_9687a2e178ba35ab := range __obf_6324600e1e794b0c {
		switch __obf_9687a2e178ba35ab := __obf_9687a2e178ba35ab.(type) {
		case *ast.BasicLit:
			__obf_c717d4181ed7e54b = append(__obf_c717d4181ed7e54b, __obf_9687a2e178ba35ab.Value)
		case *ast.Ident:
			if __obf_4679ee4b837bfe9c, __obf_2ffd626706248919 := __obf_0a0f803294160e3f(__obf_9687a2e178ba35ab.Name, __obf_76688cf26d5015f4); __obf_2ffd626706248919 != nil {
				return nil, __obf_2ffd626706248919
			} else {
				__obf_c717d4181ed7e54b = append(__obf_c717d4181ed7e54b, __obf_4679ee4b837bfe9c)
			}
		}
	}
	return __obf_c717d4181ed7e54b, nil
}

func __obf_0a0f803294160e3f(__obf_0065f27d5ca0e1c7 string, __obf_76688cf26d5015f4 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_0065f27d5ca0e1c7 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_ac45a926eaaf029b, __obf_73ea2ef7dfc0a7db := __obf_76688cf26d5015f4[__obf_0065f27d5ca0e1c7]; __obf_73ea2ef7dfc0a7db {
		return __obf_ac45a926eaaf029b, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
