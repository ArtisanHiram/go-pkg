package __obf_0e6b24ccfa510502

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
	__obf_af3675e23329e0c0 = map[string]__obf_86ce1dc4a9bc0a07{
		"contains": func(__obf_8784e3336259837e []any) (any, error) {
			fmt.Println("contains print: ", __obf_8784e3336259837e)
			return nil, nil
		},
	}
)

type __obf_86ce1dc4a9bc0a07 = func(__obf_8784e3336259837e []any) (any, error)

type Rule struct {
	__obf_625709991b6983d2 ast.Expr
}

func (__obf_e21facdeaa50a5b5 *Rule) SetExpr(__obf_625709991b6983d2 string) error {
	if len(__obf_625709991b6983d2) == 0 {
		return ErrRuleEmpty
	}
	if __obf_ba9967a449b71398, __obf_5e0dddb6044f6b24 := parser.ParseExpr(__obf_625709991b6983d2); __obf_5e0dddb6044f6b24 != nil {
		return __obf_5e0dddb6044f6b24
	} else {
		__obf_e21facdeaa50a5b5.__obf_625709991b6983d2 = __obf_ba9967a449b71398
	}

	return nil
}

func (__obf_e21facdeaa50a5b5 *Rule) Bool(__obf_c4e202b2b9ede99f map[string]any) (bool, error) {
	if __obf_e21facdeaa50a5b5.__obf_625709991b6983d2 != nil {
		__obf_0d8ccf61232f0629, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_c4e202b2b9ede99f)
		if __obf_5e0dddb6044f6b24 != nil {
			return false, __obf_5e0dddb6044f6b24
		}
		if __obf_e21facdeaa50a5b5, __obf_fea627f079d0e903 := __obf_0d8ccf61232f0629.(bool); __obf_fea627f079d0e903 {
			return __obf_e21facdeaa50a5b5, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_e21facdeaa50a5b5 *Rule) Int(__obf_c4e202b2b9ede99f map[string]any) (int64, error) {
	if __obf_e21facdeaa50a5b5.__obf_625709991b6983d2 != nil {
		__obf_0d8ccf61232f0629, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_c4e202b2b9ede99f)
		if __obf_5e0dddb6044f6b24 != nil {
			return 0, __obf_5e0dddb6044f6b24
		}
		switch __obf_0d8ccf61232f0629 := __obf_0d8ccf61232f0629.(type) {
		case int64:
			return __obf_0d8ccf61232f0629, nil
		case float64:
			return int64(__obf_0d8ccf61232f0629), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_e21facdeaa50a5b5 *Rule) Float(__obf_c4e202b2b9ede99f map[string]any) (float64, error) {
	if __obf_e21facdeaa50a5b5.__obf_625709991b6983d2 != nil {
		__obf_0d8ccf61232f0629, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_c4e202b2b9ede99f)
		if __obf_5e0dddb6044f6b24 != nil {
			return 0, __obf_5e0dddb6044f6b24
		}
		switch __obf_0d8ccf61232f0629 := __obf_0d8ccf61232f0629.(type) {
		case int64:
			return float64(__obf_0d8ccf61232f0629), nil
		case float64:
			return __obf_0d8ccf61232f0629, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_e21facdeaa50a5b5 *Rule) Eval(__obf_2e475ccfb4a71a27 map[string]any) (any, error) {
	switch __obf_446f602ed0da86e1 := __obf_e21facdeaa50a5b5.__obf_625709991b6983d2.(type) {
	case *ast.UnaryExpr:
		__obf_e21facdeaa50a5b5. // 一元表达式
					__obf_625709991b6983d2 = __obf_446f602ed0da86e1.X
		__obf_60c126591ded171e, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		__obf_66c0191feeb7c9ab := reflect.ValueOf(__obf_60c126591ded171e)

		switch __obf_446f602ed0da86e1.Op {
		case token.NOT: // !
			if __obf_66c0191feeb7c9ab.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_66c0191feeb7c9ab.Bool(), nil
		case token.SUB: // -
			if __obf_0b4520301b4db6fe, __obf_5e0dddb6044f6b24 := __obf_fb48b5f1c53929e3(__obf_66c0191feeb7c9ab); __obf_5e0dddb6044f6b24 == nil {
				return (-1.0) * __obf_0b4520301b4db6fe, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr:
		__obf_e21facdeaa50a5b5. // 二元表达式
					__obf_625709991b6983d2 = __obf_446f602ed0da86e1.X
		__obf_0b4520301b4db6fe, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		__obf_e21facdeaa50a5b5.__obf_625709991b6983d2 = __obf_446f602ed0da86e1.Y
		__obf_dc1194f976ddf66b, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		return __obf_de977d59d4b9776a(__obf_0b4520301b4db6fe, __obf_dc1194f976ddf66b,
			// 标志符（已定义变量或常量（bool））
			__obf_446f602ed0da86e1.Op)
	case *ast.Ident:
		return __obf_461d4009f79c363b(__obf_446f602ed0da86e1.Name, __obf_2e475ccfb4a71a27)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_446f602ed0da86e1.Kind {
		case token.STRING:
			return strings.Trim(__obf_446f602ed0da86e1.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_446f602ed0da86e1.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_446f602ed0da86e1.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr:
		__obf_e21facdeaa50a5b5. // 圆括号内表达式
					__obf_625709991b6983d2 = __obf_446f602ed0da86e1.X
		return __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
	case *ast.SelectorExpr:
		__obf_e21facdeaa50a5b5. // 属性或方法选择表达式
					__obf_625709991b6983d2 = __obf_446f602ed0da86e1.X
		__obf_ac14ad88c327d685, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		return __obf_461d4009f79c363b(__obf_446f602ed0da86e1.Sel.Name, __obf_ac14ad88c327d685.(map[string]any))
	case *ast.IndexExpr:
		__obf_e21facdeaa50a5b5. // 中括号内表达式——map或slice索引
					__obf_625709991b6983d2 = __obf_446f602ed0da86e1.X
		__obf_3fc9b5745ec3f5e1, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		__obf_e21facdeaa50a5b5.__obf_625709991b6983d2 = __obf_446f602ed0da86e1.Index
		__obf_b1565e7715252d27, __obf_5e0dddb6044f6b24 := __obf_e21facdeaa50a5b5.Eval(__obf_2e475ccfb4a71a27)
		if __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}

		switch __obf_3fc9b5745ec3f5e1 := __obf_3fc9b5745ec3f5e1.(type) {
		case map[string]any:
			if __obf_b1565e7715252d27, __obf_d5eb783e603af58c := __obf_b1565e7715252d27.(string); __obf_d5eb783e603af58c {
				return __obf_3fc9b5745ec3f5e1[__obf_b1565e7715252d27], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_b1565e7715252d27 := __obf_b1565e7715252d27.(type) {
			case int:
				return __obf_3fc9b5745ec3f5e1[int64(__obf_b1565e7715252d27)], nil
			case int64:
				return __obf_3fc9b5745ec3f5e1[__obf_b1565e7715252d27], nil
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
		if __obf_af6c627c7a943cf1, __obf_5e0dddb6044f6b24 := __obf_ecc43d8417e5a228(__obf_446f602ed0da86e1.Args, __obf_2e475ccfb4a71a27); __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		} else {
			return __obf_af3675e23329e0c0[__obf_446f602ed0da86e1.Fun.(*ast.Ident).Name](__obf_af6c627c7a943cf1)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_ecc43d8417e5a228(__obf_8784e3336259837e []ast.Expr, __obf_2e475ccfb4a71a27 map[string]any) ([]any, error) {
	var __obf_85b005ca7ea18913 []any
	for _, __obf_2934ba11088459b4 := range __obf_8784e3336259837e {
		switch __obf_2934ba11088459b4 := __obf_2934ba11088459b4.(type) {
		case *ast.BasicLit:
			__obf_85b005ca7ea18913 = append(__obf_85b005ca7ea18913, __obf_2934ba11088459b4.Value)
		case *ast.Ident:
			if __obf_98595368c7fb0a1e, __obf_5e0dddb6044f6b24 := __obf_461d4009f79c363b(__obf_2934ba11088459b4.Name, __obf_2e475ccfb4a71a27); __obf_5e0dddb6044f6b24 != nil {
				return nil, __obf_5e0dddb6044f6b24
			} else {
				__obf_85b005ca7ea18913 = append(__obf_85b005ca7ea18913, __obf_98595368c7fb0a1e)
			}
		}
	}
	return __obf_85b005ca7ea18913, nil
}

func __obf_461d4009f79c363b(__obf_00a57b82f4bae3d9 string, __obf_2e475ccfb4a71a27 map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_00a57b82f4bae3d9 {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_760e5bdd6328b4b0, __obf_fea627f079d0e903 := __obf_2e475ccfb4a71a27[__obf_00a57b82f4bae3d9]; __obf_fea627f079d0e903 {
		return __obf_760e5bdd6328b4b0, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
