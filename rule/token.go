package __obf_3bb823b8917aab5c

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_2abb3b12b1ae34a3(__obf_a27a3582cc4e1e1d, __obf_4a8fe9c05037d3f5 any, __obf_09671e6d097fbe01 token.Token) (any, error) {
	__obf_108c1bd00b38312c := reflect.ValueOf(__obf_a27a3582cc4e1e1d)
	__obf_e2067a2d5f8ea9c9 := reflect.ValueOf(__obf_4a8fe9c05037d3f5)
	switch __obf_09671e6d097fbe01 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_58a1c0aea4a48991, __obf_856b8ce413ab0dcf float64
		var __obf_0057b8b996f24b37 error
		if __obf_58a1c0aea4a48991, __obf_0057b8b996f24b37 = __obf_2b3f156b123f0c7b(__obf_108c1bd00b38312c); __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		if __obf_856b8ce413ab0dcf, __obf_0057b8b996f24b37 = __obf_2b3f156b123f0c7b(__obf_e2067a2d5f8ea9c9); __obf_0057b8b996f24b37 != nil {
			return nil, __obf_0057b8b996f24b37
		}
		switch __obf_09671e6d097fbe01 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_58a1c0aea4a48991).Add(decimal.NewFromFloat(__obf_856b8ce413ab0dcf)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_58a1c0aea4a48991).Sub(decimal.NewFromFloat(__obf_856b8ce413ab0dcf)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_58a1c0aea4a48991).Mul(decimal.NewFromFloat(__obf_856b8ce413ab0dcf)), nil
		case token.QUO:
			if __obf_856b8ce413ab0dcf == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_58a1c0aea4a48991 / __obf_856b8ce413ab0dcf, nil
		case token.LSS:
			return __obf_58a1c0aea4a48991 < __obf_856b8ce413ab0dcf, nil
		case token.GTR:
			return __obf_58a1c0aea4a48991 > __obf_856b8ce413ab0dcf, nil
		case token.LEQ:
			return __obf_58a1c0aea4a48991 <= __obf_856b8ce413ab0dcf, nil
		case token.GEQ:
			return __obf_58a1c0aea4a48991 >= __obf_856b8ce413ab0dcf, nil
		case token.EQL:
			return __obf_58a1c0aea4a48991 == __obf_856b8ce413ab0dcf, nil
		case token.NEQ:
			return __obf_58a1c0aea4a48991 != __obf_856b8ce413ab0dcf, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_108c1bd00b38312c.Kind() != reflect.Bool || __obf_e2067a2d5f8ea9c9.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_09671e6d097fbe01 {
		case token.LAND:
			return __obf_108c1bd00b38312c.Bool() && __obf_e2067a2d5f8ea9c9.Bool(), nil
		case token.LOR:
			return __obf_108c1bd00b38312c.Bool() || __obf_e2067a2d5f8ea9c9.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_2b3f156b123f0c7b(__obf_a27a3582cc4e1e1d reflect.Value) (float64, error) {
	switch __obf_a27a3582cc4e1e1d.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_a27a3582cc4e1e1d.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_a27a3582cc4e1e1d.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_1323d1dc2a10a417(__obf_a27a3582cc4e1e1d reflect.Value) {

}
