package __obf_f3901fe9624df715

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_96558f9406e98de4(__obf_58db4a39c61276eb, __obf_d6eb3223122662eb any, __obf_a4499acc29ab5e0b token.Token) (any, error) {
	__obf_17dc9220ceb6675d := reflect.ValueOf(__obf_58db4a39c61276eb)
	__obf_1eba3685a3e5dfeb := reflect.ValueOf(__obf_d6eb3223122662eb)
	switch __obf_a4499acc29ab5e0b {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_99e10e109bb3ebfd, __obf_1ee3a5135b452f80 float64
		var __obf_0868a2b24f90315b error
		if __obf_99e10e109bb3ebfd, __obf_0868a2b24f90315b = __obf_b94cf9ae4814ea68(__obf_17dc9220ceb6675d); __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		if __obf_1ee3a5135b452f80, __obf_0868a2b24f90315b = __obf_b94cf9ae4814ea68(__obf_1eba3685a3e5dfeb); __obf_0868a2b24f90315b != nil {
			return nil, __obf_0868a2b24f90315b
		}
		switch __obf_a4499acc29ab5e0b {
		case token.ADD:
			return decimal.NewFromFloat(__obf_99e10e109bb3ebfd).Add(decimal.NewFromFloat(__obf_1ee3a5135b452f80)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_99e10e109bb3ebfd).Sub(decimal.NewFromFloat(__obf_1ee3a5135b452f80)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_99e10e109bb3ebfd).Mul(decimal.NewFromFloat(__obf_1ee3a5135b452f80)), nil
		case token.QUO:
			if __obf_1ee3a5135b452f80 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_99e10e109bb3ebfd / __obf_1ee3a5135b452f80, nil
		case token.LSS:
			return __obf_99e10e109bb3ebfd < __obf_1ee3a5135b452f80, nil
		case token.GTR:
			return __obf_99e10e109bb3ebfd > __obf_1ee3a5135b452f80, nil
		case token.LEQ:
			return __obf_99e10e109bb3ebfd <= __obf_1ee3a5135b452f80, nil
		case token.GEQ:
			return __obf_99e10e109bb3ebfd >= __obf_1ee3a5135b452f80, nil
		case token.EQL:
			return __obf_99e10e109bb3ebfd == __obf_1ee3a5135b452f80, nil
		case token.NEQ:
			return __obf_99e10e109bb3ebfd != __obf_1ee3a5135b452f80, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_17dc9220ceb6675d.Kind() != reflect.Bool || __obf_1eba3685a3e5dfeb.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_a4499acc29ab5e0b {
		case token.LAND:
			return __obf_17dc9220ceb6675d.Bool() && __obf_1eba3685a3e5dfeb.Bool(), nil
		case token.LOR:
			return __obf_17dc9220ceb6675d.Bool() || __obf_1eba3685a3e5dfeb.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_b94cf9ae4814ea68(__obf_58db4a39c61276eb reflect.Value) (float64, error) {
	switch __obf_58db4a39c61276eb.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_58db4a39c61276eb.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_58db4a39c61276eb.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_2c0632f820ca03ad(__obf_58db4a39c61276eb reflect.Value) {

}
