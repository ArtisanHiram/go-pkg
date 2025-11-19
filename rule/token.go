package __obf_e8080eb36ea958ff

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_dc83ec0e5567c1f8(__obf_eb0b095223325ed2, __obf_ba1dededac8b8fec any, __obf_6cdf57a410c1975a token.Token) (any, error) {
	__obf_4d89f165dcefeed6 := reflect.ValueOf(__obf_eb0b095223325ed2)
	__obf_9a61f187fa1ace5f := reflect.ValueOf(__obf_ba1dededac8b8fec)
	switch __obf_6cdf57a410c1975a {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_0d19a8bf68c112c4, __obf_e08e221734b46a54 float64
		var __obf_2ffd626706248919 error
		if __obf_0d19a8bf68c112c4, __obf_2ffd626706248919 = __obf_71e4bcca69d8b68b(__obf_4d89f165dcefeed6); __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		if __obf_e08e221734b46a54, __obf_2ffd626706248919 = __obf_71e4bcca69d8b68b(__obf_9a61f187fa1ace5f); __obf_2ffd626706248919 != nil {
			return nil, __obf_2ffd626706248919
		}
		switch __obf_6cdf57a410c1975a {
		case token.ADD:
			return decimal.NewFromFloat(__obf_0d19a8bf68c112c4).Add(decimal.NewFromFloat(__obf_e08e221734b46a54)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_0d19a8bf68c112c4).Sub(decimal.NewFromFloat(__obf_e08e221734b46a54)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_0d19a8bf68c112c4).Mul(decimal.NewFromFloat(__obf_e08e221734b46a54)), nil
		case token.QUO:
			if __obf_e08e221734b46a54 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_0d19a8bf68c112c4 / __obf_e08e221734b46a54, nil
		case token.LSS:
			return __obf_0d19a8bf68c112c4 < __obf_e08e221734b46a54, nil
		case token.GTR:
			return __obf_0d19a8bf68c112c4 > __obf_e08e221734b46a54, nil
		case token.LEQ:
			return __obf_0d19a8bf68c112c4 <= __obf_e08e221734b46a54, nil
		case token.GEQ:
			return __obf_0d19a8bf68c112c4 >= __obf_e08e221734b46a54, nil
		case token.EQL:
			return __obf_0d19a8bf68c112c4 == __obf_e08e221734b46a54, nil
		case token.NEQ:
			return __obf_0d19a8bf68c112c4 != __obf_e08e221734b46a54, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_4d89f165dcefeed6.Kind() != reflect.Bool || __obf_9a61f187fa1ace5f.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_6cdf57a410c1975a {
		case token.LAND:
			return __obf_4d89f165dcefeed6.Bool() && __obf_9a61f187fa1ace5f.Bool(), nil
		case token.LOR:
			return __obf_4d89f165dcefeed6.Bool() || __obf_9a61f187fa1ace5f.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_71e4bcca69d8b68b(__obf_eb0b095223325ed2 reflect.Value) (float64, error) {
	switch __obf_eb0b095223325ed2.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_eb0b095223325ed2.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_eb0b095223325ed2.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_b6021b8ce7083ffc(__obf_eb0b095223325ed2 reflect.Value) {

}
