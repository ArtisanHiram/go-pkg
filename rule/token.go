package __obf_564c951b28a491a4

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_decba2110d4ae088(__obf_c8711b6013a70c03, __obf_b0b1735d1adfdde8 any, __obf_479cf89720570702 token.Token) (any, error) {
	__obf_0d140d88064afd5f := reflect.ValueOf(__obf_c8711b6013a70c03)
	__obf_23ad708a99f67c83 := reflect.ValueOf(__obf_b0b1735d1adfdde8)
	switch __obf_479cf89720570702 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_00fb419604042279, __obf_b5e9b7666394d92f float64
		var __obf_a10eba5810272430 error
		if __obf_00fb419604042279, __obf_a10eba5810272430 = __obf_5a3fa900c94b4a96(__obf_0d140d88064afd5f); __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		if __obf_b5e9b7666394d92f, __obf_a10eba5810272430 = __obf_5a3fa900c94b4a96(__obf_23ad708a99f67c83); __obf_a10eba5810272430 != nil {
			return nil, __obf_a10eba5810272430
		}
		switch __obf_479cf89720570702 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_00fb419604042279).Add(decimal.NewFromFloat(__obf_b5e9b7666394d92f)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_00fb419604042279).Sub(decimal.NewFromFloat(__obf_b5e9b7666394d92f)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_00fb419604042279).Mul(decimal.NewFromFloat(__obf_b5e9b7666394d92f)), nil
		case token.QUO:
			if __obf_b5e9b7666394d92f == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_00fb419604042279 / __obf_b5e9b7666394d92f, nil
		case token.LSS:
			return __obf_00fb419604042279 < __obf_b5e9b7666394d92f, nil
		case token.GTR:
			return __obf_00fb419604042279 > __obf_b5e9b7666394d92f, nil
		case token.LEQ:
			return __obf_00fb419604042279 <= __obf_b5e9b7666394d92f, nil
		case token.GEQ:
			return __obf_00fb419604042279 >= __obf_b5e9b7666394d92f, nil
		case token.EQL:
			return __obf_00fb419604042279 == __obf_b5e9b7666394d92f, nil
		case token.NEQ:
			return __obf_00fb419604042279 != __obf_b5e9b7666394d92f, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_0d140d88064afd5f.Kind() != reflect.Bool || __obf_23ad708a99f67c83.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_479cf89720570702 {
		case token.LAND:
			return __obf_0d140d88064afd5f.Bool() && __obf_23ad708a99f67c83.Bool(), nil
		case token.LOR:
			return __obf_0d140d88064afd5f.Bool() || __obf_23ad708a99f67c83.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_5a3fa900c94b4a96(__obf_c8711b6013a70c03 reflect.Value) (float64, error) {
	switch __obf_c8711b6013a70c03.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_c8711b6013a70c03.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_c8711b6013a70c03.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_c12d9ca2193cc6a0(__obf_c8711b6013a70c03 reflect.Value) {

}
