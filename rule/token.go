package __obf_0e6b24ccfa510502

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_de977d59d4b9776a(__obf_0b4520301b4db6fe, __obf_dc1194f976ddf66b any, __obf_6e2451d41ecf99bd token.Token) (any, error) {
	__obf_749d2009c566df44 := reflect.ValueOf(__obf_0b4520301b4db6fe)
	__obf_80f73a2a1340c7b6 := reflect.ValueOf(__obf_dc1194f976ddf66b)
	switch __obf_6e2451d41ecf99bd {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_55dc9f87fd451939, __obf_0d8ccf61232f0629 float64
		var __obf_5e0dddb6044f6b24 error
		if __obf_55dc9f87fd451939, __obf_5e0dddb6044f6b24 = __obf_fb48b5f1c53929e3(__obf_749d2009c566df44); __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		if __obf_0d8ccf61232f0629, __obf_5e0dddb6044f6b24 = __obf_fb48b5f1c53929e3(__obf_80f73a2a1340c7b6); __obf_5e0dddb6044f6b24 != nil {
			return nil, __obf_5e0dddb6044f6b24
		}
		switch __obf_6e2451d41ecf99bd {
		case token.ADD:
			return decimal.NewFromFloat(__obf_55dc9f87fd451939).Add(decimal.NewFromFloat(__obf_0d8ccf61232f0629)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_55dc9f87fd451939).Sub(decimal.NewFromFloat(__obf_0d8ccf61232f0629)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_55dc9f87fd451939).Mul(decimal.NewFromFloat(__obf_0d8ccf61232f0629)), nil
		case token.QUO:
			if __obf_0d8ccf61232f0629 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_55dc9f87fd451939 / __obf_0d8ccf61232f0629, nil
		case token.LSS:
			return __obf_55dc9f87fd451939 < __obf_0d8ccf61232f0629, nil
		case token.GTR:
			return __obf_55dc9f87fd451939 > __obf_0d8ccf61232f0629, nil
		case token.LEQ:
			return __obf_55dc9f87fd451939 <= __obf_0d8ccf61232f0629, nil
		case token.GEQ:
			return __obf_55dc9f87fd451939 >= __obf_0d8ccf61232f0629, nil
		case token.EQL:
			return __obf_55dc9f87fd451939 == __obf_0d8ccf61232f0629, nil
		case token.NEQ:
			return __obf_55dc9f87fd451939 != __obf_0d8ccf61232f0629, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_749d2009c566df44.Kind() != reflect.Bool || __obf_80f73a2a1340c7b6.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_6e2451d41ecf99bd {
		case token.LAND:
			return __obf_749d2009c566df44.Bool() && __obf_80f73a2a1340c7b6.Bool(), nil
		case token.LOR:
			return __obf_749d2009c566df44.Bool() || __obf_80f73a2a1340c7b6.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_fb48b5f1c53929e3(__obf_0b4520301b4db6fe reflect.Value) (float64, error) {
	switch __obf_0b4520301b4db6fe.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_0b4520301b4db6fe.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_0b4520301b4db6fe.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_2dd34d8c6c0afcac(__obf_0b4520301b4db6fe reflect.Value) {

}
