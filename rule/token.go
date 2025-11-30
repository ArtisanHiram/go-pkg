package __obf_0b4da6ad90a2d4d3

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_588bd91612f08558(__obf_438ec6046ae75ad9, __obf_1ddf299c12ea5b7f any, __obf_decdc11ff15e938b token.Token) (any, error) {
	__obf_2e73a34749dd8302 := reflect.ValueOf(__obf_438ec6046ae75ad9)
	__obf_54605fa365e52311 := reflect.ValueOf(__obf_1ddf299c12ea5b7f)
	switch __obf_decdc11ff15e938b {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_a2862747a0ab8d6c, __obf_274b92d14f5bccb8 float64
		var __obf_c7400825d595bd55 error
		if __obf_a2862747a0ab8d6c, __obf_c7400825d595bd55 = __obf_d5fbfc4a5302e6e5(__obf_2e73a34749dd8302); __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		if __obf_274b92d14f5bccb8, __obf_c7400825d595bd55 = __obf_d5fbfc4a5302e6e5(__obf_54605fa365e52311); __obf_c7400825d595bd55 != nil {
			return nil, __obf_c7400825d595bd55
		}
		switch __obf_decdc11ff15e938b {
		case token.ADD:
			return decimal.NewFromFloat(__obf_a2862747a0ab8d6c).Add(decimal.NewFromFloat(__obf_274b92d14f5bccb8)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_a2862747a0ab8d6c).Sub(decimal.NewFromFloat(__obf_274b92d14f5bccb8)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_a2862747a0ab8d6c).Mul(decimal.NewFromFloat(__obf_274b92d14f5bccb8)), nil
		case token.QUO:
			if __obf_274b92d14f5bccb8 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_a2862747a0ab8d6c / __obf_274b92d14f5bccb8, nil
		case token.LSS:
			return __obf_a2862747a0ab8d6c < __obf_274b92d14f5bccb8, nil
		case token.GTR:
			return __obf_a2862747a0ab8d6c > __obf_274b92d14f5bccb8, nil
		case token.LEQ:
			return __obf_a2862747a0ab8d6c <= __obf_274b92d14f5bccb8, nil
		case token.GEQ:
			return __obf_a2862747a0ab8d6c >= __obf_274b92d14f5bccb8, nil
		case token.EQL:
			return __obf_a2862747a0ab8d6c == __obf_274b92d14f5bccb8, nil
		case token.NEQ:
			return __obf_a2862747a0ab8d6c != __obf_274b92d14f5bccb8, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_2e73a34749dd8302.Kind() != reflect.Bool || __obf_54605fa365e52311.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_decdc11ff15e938b {
		case token.LAND:
			return __obf_2e73a34749dd8302.Bool() && __obf_54605fa365e52311.Bool(), nil
		case token.LOR:
			return __obf_2e73a34749dd8302.Bool() || __obf_54605fa365e52311.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_d5fbfc4a5302e6e5(__obf_438ec6046ae75ad9 reflect.Value) (float64, error) {
	switch __obf_438ec6046ae75ad9.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_438ec6046ae75ad9.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_438ec6046ae75ad9.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_88c268411e30ee93(__obf_438ec6046ae75ad9 reflect.Value) {

}
