package __obf_5f51b02f59c5c56d

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
	__obf_4f91bdca7b833848 = map[string]__obf_c083cf78537c3b2b{
		"contains": func(__obf_69971b7661cab3e3 []any) (any, error) {
			fmt.Println("contains print: ", __obf_69971b7661cab3e3)
			return nil, nil
		},
	}
)

type __obf_c083cf78537c3b2b = func(__obf_69971b7661cab3e3 []any) (any, error)

type Rule struct {
	__obf_e8c7d25ff9c196ed ast.Expr
}

func (__obf_1ce9e5823b0c68cc *Rule) SetExpr(__obf_e8c7d25ff9c196ed string) error {
	if len(__obf_e8c7d25ff9c196ed) == 0 {
		return ErrRuleEmpty
	}
	if __obf_60403580fe6778d2, __obf_8704f5a03e161128 := parser.ParseExpr(__obf_e8c7d25ff9c196ed); __obf_8704f5a03e161128 != nil {
		return __obf_8704f5a03e161128
	} else {
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_60403580fe6778d2
	}

	return nil
}

func (__obf_1ce9e5823b0c68cc *Rule) Bool(__obf_f182b4aa635b46d1 map[string]any) (bool, error) {
	if __obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed != nil {
		__obf_507e33dc5361e4ff, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_f182b4aa635b46d1)
		if __obf_8704f5a03e161128 != nil {
			return false, __obf_8704f5a03e161128
		}
		if __obf_1ce9e5823b0c68cc, __obf_da661f857ca82785 := __obf_507e33dc5361e4ff.(bool); __obf_da661f857ca82785 {
			return __obf_1ce9e5823b0c68cc, nil
		}
	}

	return false, errors.New("expr is nil")
}

func (__obf_1ce9e5823b0c68cc *Rule) Int(__obf_f182b4aa635b46d1 map[string]any) (int64, error) {
	if __obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed != nil {
		__obf_507e33dc5361e4ff, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_f182b4aa635b46d1)
		if __obf_8704f5a03e161128 != nil {
			return 0, __obf_8704f5a03e161128
		}
		switch __obf_507e33dc5361e4ff := __obf_507e33dc5361e4ff.(type) {
		case int64:
			return __obf_507e33dc5361e4ff, nil
		case float64:
			return int64(__obf_507e33dc5361e4ff), nil

		}
	}
	return 0, errors.New("expr is nil")

}

func (__obf_1ce9e5823b0c68cc *Rule) Float(__obf_f182b4aa635b46d1 map[string]any) (float64, error) {
	if __obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed != nil {
		__obf_507e33dc5361e4ff, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_f182b4aa635b46d1)
		if __obf_8704f5a03e161128 != nil {
			return 0, __obf_8704f5a03e161128
		}
		switch __obf_507e33dc5361e4ff := __obf_507e33dc5361e4ff.(type) {
		case int64:
			return float64(__obf_507e33dc5361e4ff), nil
		case float64:
			return __obf_507e33dc5361e4ff, nil

		}
	}
	return 0, errors.New("expr is nil")
}

func (__obf_1ce9e5823b0c68cc *Rule) Eval(__obf_b26fedb0ab6db5eb map[string]any) (any, error) {
	switch __obf_87206e4df4315c3b := __obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed.(type) {
	case *ast.UnaryExpr: // 一元表达式
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.X
		__obf_f12bcb35da55030d, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		__obf_f3fa423f055f34f1 := reflect.ValueOf(__obf_f12bcb35da55030d)

		switch __obf_87206e4df4315c3b.Op {
		case token.NOT: // !
			if __obf_f3fa423f055f34f1.Kind() != reflect.Bool {
				return false, ErrNotBool
			}
			return !__obf_f3fa423f055f34f1.Bool(), nil
		case token.SUB: // -
			if __obf_920ae112986c362e, __obf_8704f5a03e161128 := __obf_2cdd46ca83346469(__obf_f3fa423f055f34f1); __obf_8704f5a03e161128 == nil {
				return (-1.0) * __obf_920ae112986c362e, nil
			}
			return 0.0, ErrNotNumber
		}
	case *ast.BinaryExpr: // 二元表达式
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.X
		__obf_920ae112986c362e, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.Y
		__obf_ade37eb002f33fef, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		return __obf_8d788ae5e55c2330(__obf_920ae112986c362e, __obf_ade37eb002f33fef, __obf_87206e4df4315c3b.Op)
	case *ast.Ident: // 标志符（已定义变量或常量（bool））
		return __obf_3828dab82a6cd03d(__obf_87206e4df4315c3b.Name, __obf_b26fedb0ab6db5eb)
	case *ast.BasicLit: // 基本类型文字（当作字符串存储）
		switch __obf_87206e4df4315c3b.Kind {
		case token.STRING:
			return strings.Trim(__obf_87206e4df4315c3b.Value, "\""), nil
		case token.INT:
			return strconv.ParseInt(__obf_87206e4df4315c3b.Value, 10, 64)
		case token.FLOAT:
			return strconv.ParseFloat(__obf_87206e4df4315c3b.Value, 64)
		default:
			return nil, ErrUnsupportParam
		}
	case *ast.ParenExpr: // 圆括号内表达式
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.X
		return __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
	case *ast.SelectorExpr: // 属性或方法选择表达式
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.X
		__obf_12bf641073fbaf99, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		return __obf_3828dab82a6cd03d(__obf_87206e4df4315c3b.Sel.Name, __obf_12bf641073fbaf99.(map[string]any))
	case *ast.IndexExpr: // 中括号内表达式——map或slice索引
		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.X
		__obf_3dc53621b4874bf4, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}

		__obf_1ce9e5823b0c68cc.__obf_e8c7d25ff9c196ed = __obf_87206e4df4315c3b.Index
		__obf_4b84f58100d4255e, __obf_8704f5a03e161128 := __obf_1ce9e5823b0c68cc.Eval(__obf_b26fedb0ab6db5eb)
		if __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}

		switch __obf_3dc53621b4874bf4 := __obf_3dc53621b4874bf4.(type) {
		case map[string]any:
			if __obf_4b84f58100d4255e, __obf_27da0d58958a1113 := __obf_4b84f58100d4255e.(string); __obf_27da0d58958a1113 {
				return __obf_3dc53621b4874bf4[__obf_4b84f58100d4255e], nil
			} else {
				return nil, fmt.Errorf("map here index must be string")
			}
		case []any:
			switch __obf_4b84f58100d4255e := __obf_4b84f58100d4255e.(type) {
			case int:
				return __obf_3dc53621b4874bf4[int64(__obf_4b84f58100d4255e)], nil
			case int64:
				return __obf_3dc53621b4874bf4[__obf_4b84f58100d4255e], nil
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
		if __obf_d547c59da9d3d3f7, __obf_8704f5a03e161128 := __obf_e62fcc141504485d(__obf_87206e4df4315c3b.Args, __obf_b26fedb0ab6db5eb); __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		} else {
			return __obf_4f91bdca7b833848[__obf_87206e4df4315c3b.Fun.(*ast.Ident).Name](__obf_d547c59da9d3d3f7)
		}
	}
	return nil, ErrUnsupportExpr
}

func __obf_e62fcc141504485d(__obf_69971b7661cab3e3 []ast.Expr, __obf_b26fedb0ab6db5eb map[string]any) ([]any, error) {
	var __obf_c1fffd963158cca3 []any
	for _, __obf_8a9124d99610921a := range __obf_69971b7661cab3e3 {
		switch __obf_8a9124d99610921a := __obf_8a9124d99610921a.(type) {
		case *ast.BasicLit:
			__obf_c1fffd963158cca3 = append(__obf_c1fffd963158cca3, __obf_8a9124d99610921a.Value)
		case *ast.Ident:
			if __obf_fe43bc1d957c35a5, __obf_8704f5a03e161128 := __obf_3828dab82a6cd03d(__obf_8a9124d99610921a.Name, __obf_b26fedb0ab6db5eb); __obf_8704f5a03e161128 != nil {
				return nil, __obf_8704f5a03e161128
			} else {
				__obf_c1fffd963158cca3 = append(__obf_c1fffd963158cca3, __obf_fe43bc1d957c35a5)
			}
		}
	}
	return __obf_c1fffd963158cca3, nil
}

func __obf_3828dab82a6cd03d(__obf_fa02964d5f47f36f string, __obf_b26fedb0ab6db5eb map[string]any) (any, error) {
	// while bool type is Ident
	switch __obf_fa02964d5f47f36f {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	if __obf_65ef0d4c7fb2f67b, __obf_da661f857ca82785 := __obf_b26fedb0ab6db5eb[__obf_fa02964d5f47f36f]; __obf_da661f857ca82785 {
		return __obf_65ef0d4c7fb2f67b, nil
	} else {
		return nil, ErrKeyNotFound
	}
}
