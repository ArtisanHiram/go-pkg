package __obf_d3be831d1774c7b9

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_4a4ddbfa6fc0ec8e(__obf_ca275d416555ddd0, __obf_84db90cea1ee5669 any, __obf_4bbe6d91a49b9474 token.Token) (any, error) {
	__obf_a8793cf5a96cd7bb := reflect.ValueOf(__obf_ca275d416555ddd0)
	__obf_5132e961e446dfb8 := reflect.ValueOf(__obf_84db90cea1ee5669)
	switch __obf_4bbe6d91a49b9474 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_cf1e08e94279a589, __obf_3b73e3235cb8d0f9 float64
		var __obf_df5f028324461f18 error
		if __obf_cf1e08e94279a589, __obf_df5f028324461f18 = __obf_55b9fda27ba97528(__obf_a8793cf5a96cd7bb); __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		if __obf_3b73e3235cb8d0f9, __obf_df5f028324461f18 = __obf_55b9fda27ba97528(__obf_5132e961e446dfb8); __obf_df5f028324461f18 != nil {
			return nil, __obf_df5f028324461f18
		}
		switch __obf_4bbe6d91a49b9474 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_cf1e08e94279a589).Add(decimal.NewFromFloat(__obf_3b73e3235cb8d0f9)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_cf1e08e94279a589).Sub(decimal.NewFromFloat(__obf_3b73e3235cb8d0f9)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_cf1e08e94279a589).Mul(decimal.NewFromFloat(__obf_3b73e3235cb8d0f9)), nil
		case token.QUO:
			if __obf_3b73e3235cb8d0f9 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_cf1e08e94279a589 / __obf_3b73e3235cb8d0f9, nil
		case token.LSS:
			return __obf_cf1e08e94279a589 < __obf_3b73e3235cb8d0f9, nil
		case token.GTR:
			return __obf_cf1e08e94279a589 > __obf_3b73e3235cb8d0f9, nil
		case token.LEQ:
			return __obf_cf1e08e94279a589 <= __obf_3b73e3235cb8d0f9, nil
		case token.GEQ:
			return __obf_cf1e08e94279a589 >= __obf_3b73e3235cb8d0f9, nil
		case token.EQL:
			return __obf_cf1e08e94279a589 == __obf_3b73e3235cb8d0f9, nil
		case token.NEQ:
			return __obf_cf1e08e94279a589 != __obf_3b73e3235cb8d0f9, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_a8793cf5a96cd7bb.Kind() != reflect.Bool || __obf_5132e961e446dfb8.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_4bbe6d91a49b9474 {
		case token.LAND:
			return __obf_a8793cf5a96cd7bb.Bool() && __obf_5132e961e446dfb8.Bool(), nil
		case token.LOR:
			return __obf_a8793cf5a96cd7bb.Bool() || __obf_5132e961e446dfb8.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_55b9fda27ba97528(__obf_ca275d416555ddd0 reflect.Value) (float64, error) {
	switch __obf_ca275d416555ddd0.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_ca275d416555ddd0.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_ca275d416555ddd0.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_8b99141d6a04d090(__obf_ca275d416555ddd0 reflect.Value) {

}
