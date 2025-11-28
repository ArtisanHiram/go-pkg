package __obf_558a7db000042ed1

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_711c936c8cf651bc(__obf_49265ee5f1ae335d, __obf_6e7f0359ff738fe4 any, __obf_d3950410c54f7cb7 token.Token) (any, error) {
	__obf_c9758fe35a2d2cd5 := reflect.ValueOf(__obf_49265ee5f1ae335d)
	__obf_583383b335f720f6 := reflect.ValueOf(__obf_6e7f0359ff738fe4)
	switch __obf_d3950410c54f7cb7 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_f8c2c128acb210f7, __obf_106b3ecc6336db1c float64
		var __obf_e6c13dab132e0931 error
		if __obf_f8c2c128acb210f7, __obf_e6c13dab132e0931 = __obf_c21732ed5660156f(__obf_c9758fe35a2d2cd5); __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		if __obf_106b3ecc6336db1c, __obf_e6c13dab132e0931 = __obf_c21732ed5660156f(__obf_583383b335f720f6); __obf_e6c13dab132e0931 != nil {
			return nil, __obf_e6c13dab132e0931
		}
		switch __obf_d3950410c54f7cb7 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_f8c2c128acb210f7).Add(decimal.NewFromFloat(__obf_106b3ecc6336db1c)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_f8c2c128acb210f7).Sub(decimal.NewFromFloat(__obf_106b3ecc6336db1c)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_f8c2c128acb210f7).Mul(decimal.NewFromFloat(__obf_106b3ecc6336db1c)), nil
		case token.QUO:
			if __obf_106b3ecc6336db1c == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_f8c2c128acb210f7 / __obf_106b3ecc6336db1c, nil
		case token.LSS:
			return __obf_f8c2c128acb210f7 < __obf_106b3ecc6336db1c, nil
		case token.GTR:
			return __obf_f8c2c128acb210f7 > __obf_106b3ecc6336db1c, nil
		case token.LEQ:
			return __obf_f8c2c128acb210f7 <= __obf_106b3ecc6336db1c, nil
		case token.GEQ:
			return __obf_f8c2c128acb210f7 >= __obf_106b3ecc6336db1c, nil
		case token.EQL:
			return __obf_f8c2c128acb210f7 == __obf_106b3ecc6336db1c, nil
		case token.NEQ:
			return __obf_f8c2c128acb210f7 != __obf_106b3ecc6336db1c, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_c9758fe35a2d2cd5.Kind() != reflect.Bool || __obf_583383b335f720f6.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_d3950410c54f7cb7 {
		case token.LAND:
			return __obf_c9758fe35a2d2cd5.Bool() && __obf_583383b335f720f6.Bool(), nil
		case token.LOR:
			return __obf_c9758fe35a2d2cd5.Bool() || __obf_583383b335f720f6.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_c21732ed5660156f(__obf_49265ee5f1ae335d reflect.Value) (float64, error) {
	switch __obf_49265ee5f1ae335d.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_49265ee5f1ae335d.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_49265ee5f1ae335d.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_d7cbaf5643ea200d(__obf_49265ee5f1ae335d reflect.Value) {

}
