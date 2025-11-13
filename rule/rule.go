package __obf_7954377982ec187d

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
	__obf_5436c978efbb7b0f = map[string]__obf_d5fa14392ee5c774{
		"contains": func(__obf_e16cf7bc9fa09c56 []any) (any, error) {
			fmt.Println("contains print: ", __obf_e16cf7bc9fa09c56)
			return nil, nil
		},
	}
)

type __obf_d5fa14392ee5c774 = func(__obf_e16cf7bc9fa09c56 []any) (any, error)

type Rule struct {
	__obf_43c095893757824c ast.Expr
}

func (__obf_291aa397b25d23bf *Rule) SetExpr(__obf_43c095893757824c string) error {
	if len(__obf_43c095893757824c) == 0 {
		return ErrRuleEmpty
	}
	if __obf_4f6d60523eb6b0ab, __obf_c5de6e299373359f := parser.ParseExpr(__obf_43c095893757824c); __obf_c5de6e299373359f != nil {
		return __obf_c5de6e299373359f
	} else {
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_4f6d60523eb6b0ab
	}

	return nil
}

func (__obf_291aa397b25d23bf *Rule) Bool(__obf_e0bf1c50352af4be map[string]any) (bool, error) {
	if __obf_291aa397b25d23bf.__obf_43c095893757824c != nil {
		__obf_c2138e1959af3b45, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_e0bf1c50352af4be)
		if __obf_c5de6e299373359f != nil {
			return false, __obf_c5de6e299373359f
		}
		if __obf_291aa397b25d23bf, __obf_8d5c5709ab4062b5 := __obf_c2138e1959af3b45.(bool); __obf_8d5c5709ab4062b5 {
			return __obf_291aa397b25d23bf, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_291aa397b25d23bf *Rule) Int(__obf_e0bf1c50352af4be map[string]any) (int64, error) {
	if __obf_291aa397b25d23bf.__obf_43c095893757824c != nil {
		__obf_c2138e1959af3b45, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_e0bf1c50352af4be)
		if __obf_c5de6e299373359f != nil {
			return 0, __obf_c5de6e299373359f
		}
		switch __obf_c2138e1959af3b45 := __obf_c2138e1959af3b45.(type) {
		case int64:
			return __obf_c2138e1959af3b45, nil
		case float64:
			return int64(__obf_c2138e1959af3b45), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_291aa397b25d23bf *Rule) Float(__obf_e0bf1c50352af4be map[string]any) (float64, error) {
	if __obf_291aa397b25d23bf.__obf_43c095893757824c != nil {
		__obf_c2138e1959af3b45, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_e0bf1c50352af4be)
		if __obf_c5de6e299373359f != nil {
			return 0, __obf_c5de6e299373359f
		}
		switch __obf_c2138e1959af3b45 := __obf_c2138e1959af3b45.(type) {
		case int64:
			return float64(__obf_c2138e1959af3b45), nil
		case float64:
			return __obf_c2138e1959af3b45, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_291aa397b25d23bf *Rule) Eval(__obf_4686552f3da0d21b map[string]any) (any, error) {
	switch __obf_22b6121fb760897a := __obf_291aa397b25d23bf.__obf_43c095893757824c.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.X
		__obf_e1df3df49cc30a5a, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		__obf_c138dd88a8fbddb9 := reflect.ValueOf(__obf_e1df3df49cc30a5a)

		switch __obf_22b6121fb760897a.Op {
		case token.NOT: // !
			if __obf_c138dd88a8fbddb9.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_c138dd88a8fbddb9.Bool(), nil
		case token.SUB: // -
			if __obf_a010806d3a9f83e4, __obf_c5de6e299373359f := __obf_7b5ef065fa0c8295(__obf_c138dd88a8fbddb9); __obf_c5de6e299373359f == nil {
				return (-1.0) * __obf_a010806d3a9f83e4, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.X
		__obf_a010806d3a9f83e4, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.Y
		__obf_a5e9fafdb18bcd6d, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		return __obf_5a240750dc1e1fd3(__obf_a010806d3a9f83e4, __obf_a5e9fafdb18bcd6d, __obf_22b6121fb760897a.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_3b9e1bb9493b6819(__obf_22b6121fb760897a.Name, __obf_4686552f3da0d21b)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_22b6121fb760897a.Kind {
		case token.STRING:
			return strings.Trim(__obf_22b6121fb760897a.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_22b6121fb760897a.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_22b6121fb760897a.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.X
		return __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.X
		__obf_328be5328f216335, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		return __obf_3b9e1bb9493b6819(__obf_22b6121fb760897a.Sel.Name, __obf_328be5328f216335.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.X
		__obf_a3d06eedabdb6ef2, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}

		__obf_291aa397b25d23bf.__obf_43c095893757824c = __obf_22b6121fb760897a.Index
		__obf_1d567492e2bab91e, __obf_c5de6e299373359f := __obf_291aa397b25d23bf.Eval(__obf_4686552f3da0d21b)
		if __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}

		switch __obf_a3d06eedabdb6ef2 := __obf_a3d06eedabdb6ef2.(type) {
		case map[string]any:
			if __obf_1d567492e2bab91e, __obf_0e56ff44e3f4844a := __obf_1d567492e2bab91e.(string); __obf_0e56ff44e3f4844a {
				return __obf_a3d06eedabdb6ef2[__obf_1d567492e2bab91e], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_1d567492e2bab91e := __obf_1d567492e2bab91e.(type) {
			case int:
				return __obf_a3d06eedabdb6ef2[int64(__obf_1d567492e2bab91e)], nil
			case int64:
				return __obf_a3d06eedabdb6ef2[__obf_1d567492e2bab91e], nil
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
		if __obf_101a7fb072d3010a, __obf_c5de6e299373359f := __obf_984cb88fade43c2d(__obf_22b6121fb760897a.Args, __obf_4686552f3da0d21b); __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		} else {
			return __obf_5436c978efbb7b0f[__obf_22b6121fb760897a.Fun.(*ast.Ident).Name](__obf_101a7fb072d3010a)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_984cb88fade43c2d(__obf_e16cf7bc9fa09c56 []ast.Expr, __obf_4686552f3da0d21b map[string]any) ([]any, error) {
	var __obf_a8d8ddf0cc07b0ad []any
	for _, __obf_3935067fd16b8b24 := range __obf_e16cf7bc9fa09c56 {
		switch __obf_3935067fd16b8b24 := __obf_3935067fd16b8b24.(type) {
		case *ast.BasicLit:
			__obf_a8d8ddf0cc07b0ad = append(__obf_a8d8ddf0cc07b0ad, __obf_3935067fd16b8b24.Value)
		case *ast.Ident:
			if __obf_a3f48f2c92666a2b, __obf_c5de6e299373359f := __obf_3b9e1bb9493b6819(__obf_3935067fd16b8b24.Name, __obf_4686552f3da0d21b); __obf_c5de6e299373359f != nil {
				return nil, __obf_c5de6e299373359f
			} else {
				__obf_a8d8ddf0cc07b0ad = append(__obf_a8d8ddf0cc07b0ad, __obf_a3f48f2c92666a2b)
			}
		}
	}
	return __obf_a8d8ddf0cc07b0ad, nil
}

func __obf_3b9e1bb9493b6819(__obf_663ea927697bb7ad string, __obf_4686552f3da0d21b map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_663ea927697bb7ad {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_56d3073b75df22bb, __obf_8d5c5709ab4062b5 := __obf_4686552f3da0d21b[__obf_663ea927697bb7ad]; __obf_8d5c5709ab4062b5 {
		return __obf_56d3073b75df22bb, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
