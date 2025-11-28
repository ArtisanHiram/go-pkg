package __obf_5f51b02f59c5c56d

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_8d788ae5e55c2330(__obf_920ae112986c362e, __obf_ade37eb002f33fef any, __obf_d8b2b6e6bf68b5b1 token.Token) (any, error) {
	__obf_b4f2071542038f6a := reflect.ValueOf(__obf_920ae112986c362e)
	__obf_3b0e472a3bb58b7f := reflect.ValueOf(__obf_ade37eb002f33fef)
	switch __obf_d8b2b6e6bf68b5b1 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_71cd3470d452de59, __obf_507e33dc5361e4ff float64
		var __obf_8704f5a03e161128 error
		if __obf_71cd3470d452de59, __obf_8704f5a03e161128 = __obf_2cdd46ca83346469(__obf_b4f2071542038f6a); __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		if __obf_507e33dc5361e4ff, __obf_8704f5a03e161128 = __obf_2cdd46ca83346469(__obf_3b0e472a3bb58b7f); __obf_8704f5a03e161128 != nil {
			return nil, __obf_8704f5a03e161128
		}
		switch __obf_d8b2b6e6bf68b5b1 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_71cd3470d452de59).Add(decimal.NewFromFloat(__obf_507e33dc5361e4ff)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_71cd3470d452de59).Sub(decimal.NewFromFloat(__obf_507e33dc5361e4ff)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_71cd3470d452de59).Mul(decimal.NewFromFloat(__obf_507e33dc5361e4ff)), nil
		case token.QUO:
			if __obf_507e33dc5361e4ff == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_71cd3470d452de59 / __obf_507e33dc5361e4ff, nil
		case token.LSS:
			return __obf_71cd3470d452de59 < __obf_507e33dc5361e4ff, nil
		case token.GTR:
			return __obf_71cd3470d452de59 > __obf_507e33dc5361e4ff, nil
		case token.LEQ:
			return __obf_71cd3470d452de59 <= __obf_507e33dc5361e4ff, nil
		case token.GEQ:
			return __obf_71cd3470d452de59 >= __obf_507e33dc5361e4ff, nil
		case token.EQL:
			return __obf_71cd3470d452de59 == __obf_507e33dc5361e4ff, nil
		case token.NEQ:
			return __obf_71cd3470d452de59 != __obf_507e33dc5361e4ff, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_b4f2071542038f6a.Kind() != reflect.Bool || __obf_3b0e472a3bb58b7f.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_d8b2b6e6bf68b5b1 {
		case token.LAND:
			return __obf_b4f2071542038f6a.Bool() && __obf_3b0e472a3bb58b7f.Bool(), nil
		case token.LOR:
			return __obf_b4f2071542038f6a.Bool() || __obf_3b0e472a3bb58b7f.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_2cdd46ca83346469(__obf_920ae112986c362e reflect.Value) (float64, error) {
	switch __obf_920ae112986c362e.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_920ae112986c362e.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_920ae112986c362e.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_5b0d2352d360824f(__obf_920ae112986c362e reflect.Value) {

}
