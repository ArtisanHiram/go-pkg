package __obf_4632c1c5949287c8

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_a9c7fcdfcce65ec0(__obf_559c9a6e4ad26b2a, __obf_81f0fb6882d5851e any, __obf_94db32b5a90917da token.Token) (any, error) {
	__obf_98be454b13a1b6e6 := reflect.ValueOf(__obf_559c9a6e4ad26b2a)
	__obf_40a3fc7670971fff := reflect.ValueOf(__obf_81f0fb6882d5851e)
	switch __obf_94db32b5a90917da {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_f25d95ab2b7a6b52, __obf_79d6e23a7e30df98 float64
		var __obf_f1d0e26b3a5fdda0 error
		if __obf_f25d95ab2b7a6b52, __obf_f1d0e26b3a5fdda0 = __obf_a146566d6dfec01e(__obf_98be454b13a1b6e6); __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		if __obf_79d6e23a7e30df98, __obf_f1d0e26b3a5fdda0 = __obf_a146566d6dfec01e(__obf_40a3fc7670971fff); __obf_f1d0e26b3a5fdda0 != nil {
			return nil, __obf_f1d0e26b3a5fdda0
		}
		switch __obf_94db32b5a90917da {
		case token.ADD:
			return decimal.NewFromFloat(__obf_f25d95ab2b7a6b52).Add(decimal.NewFromFloat(__obf_79d6e23a7e30df98)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_f25d95ab2b7a6b52).Sub(decimal.NewFromFloat(__obf_79d6e23a7e30df98)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_f25d95ab2b7a6b52).Mul(decimal.NewFromFloat(__obf_79d6e23a7e30df98)), nil
		case token.QUO:
			if __obf_79d6e23a7e30df98 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_f25d95ab2b7a6b52 / __obf_79d6e23a7e30df98, nil
		case token.LSS:
			return __obf_f25d95ab2b7a6b52 < __obf_79d6e23a7e30df98, nil
		case token.GTR:
			return __obf_f25d95ab2b7a6b52 > __obf_79d6e23a7e30df98, nil
		case token.LEQ:
			return __obf_f25d95ab2b7a6b52 <= __obf_79d6e23a7e30df98, nil
		case token.GEQ:
			return __obf_f25d95ab2b7a6b52 >= __obf_79d6e23a7e30df98, nil
		case token.EQL:
			return __obf_f25d95ab2b7a6b52 == __obf_79d6e23a7e30df98, nil
		case token.NEQ:
			return __obf_f25d95ab2b7a6b52 != __obf_79d6e23a7e30df98, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_98be454b13a1b6e6.Kind() != reflect.Bool || __obf_40a3fc7670971fff.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_94db32b5a90917da {
		case token.LAND:
			return __obf_98be454b13a1b6e6.Bool() && __obf_40a3fc7670971fff.Bool(), nil
		case token.LOR:
			return __obf_98be454b13a1b6e6.Bool() || __obf_40a3fc7670971fff.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_a146566d6dfec01e(__obf_559c9a6e4ad26b2a reflect.Value) (float64, error) {
	switch __obf_559c9a6e4ad26b2a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_559c9a6e4ad26b2a.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_559c9a6e4ad26b2a.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_ee3ce88765e98f04(__obf_559c9a6e4ad26b2a reflect.Value) {

}
