package __obf_7954377982ec187d

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_5a240750dc1e1fd3(__obf_a010806d3a9f83e4, __obf_a5e9fafdb18bcd6d any, __obf_3b1d321f9680073f token.Token) (any, error) {
	__obf_c7fee5cb87102103 := reflect.ValueOf(__obf_a010806d3a9f83e4)
	__obf_ea54a09a4d1db298 := reflect.ValueOf(__obf_a5e9fafdb18bcd6d)
	switch __obf_3b1d321f9680073f {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_5ed5804535b2288a, __obf_c2138e1959af3b45 float64
		var __obf_c5de6e299373359f error
		if __obf_5ed5804535b2288a, __obf_c5de6e299373359f = __obf_7b5ef065fa0c8295(__obf_c7fee5cb87102103); __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		if __obf_c2138e1959af3b45, __obf_c5de6e299373359f = __obf_7b5ef065fa0c8295(__obf_ea54a09a4d1db298); __obf_c5de6e299373359f != nil {
			return nil, __obf_c5de6e299373359f
		}
		switch __obf_3b1d321f9680073f {
		case token.ADD:
			return decimal.NewFromFloat(__obf_5ed5804535b2288a).Add(decimal.NewFromFloat(__obf_c2138e1959af3b45)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_5ed5804535b2288a).Sub(decimal.NewFromFloat(__obf_c2138e1959af3b45)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_5ed5804535b2288a).Mul(decimal.NewFromFloat(__obf_c2138e1959af3b45)), nil
		case token.QUO:
			if __obf_c2138e1959af3b45 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_5ed5804535b2288a / __obf_c2138e1959af3b45, nil
		case token.LSS:
			return __obf_5ed5804535b2288a < __obf_c2138e1959af3b45, nil
		case token.GTR:
			return __obf_5ed5804535b2288a > __obf_c2138e1959af3b45, nil
		case token.LEQ:
			return __obf_5ed5804535b2288a <= __obf_c2138e1959af3b45, nil
		case token.GEQ:
			return __obf_5ed5804535b2288a >= __obf_c2138e1959af3b45, nil
		case token.EQL:
			return __obf_5ed5804535b2288a == __obf_c2138e1959af3b45, nil
		case token.NEQ:
			return __obf_5ed5804535b2288a != __obf_c2138e1959af3b45, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_c7fee5cb87102103.Kind() != reflect.Bool || __obf_ea54a09a4d1db298.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_3b1d321f9680073f {
		case token.LAND:
			return __obf_c7fee5cb87102103.Bool() && __obf_ea54a09a4d1db298.Bool(), nil
		case token.LOR:
			return __obf_c7fee5cb87102103.Bool() || __obf_ea54a09a4d1db298.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_7b5ef065fa0c8295(__obf_a010806d3a9f83e4 reflect.Value) (float64, error) {
	switch __obf_a010806d3a9f83e4.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_a010806d3a9f83e4.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_a010806d3a9f83e4.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_25b386b3e191b71b(__obf_a010806d3a9f83e4 reflect.Value) {

}
