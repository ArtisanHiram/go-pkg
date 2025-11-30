package __obf_da2d8c26779e1ac8

import (
	"errors"
	decimal "github.com/ArtisanHiram/go-pkg/decimal"
	"go/token"
	"reflect"
)

func __obf_5f98a3add70ee4cc(__obf_8d1c7d6c4b72dec7, __obf_22eaa72847d00940 any, __obf_cc73fd613fcc4064 token.Token) (any, error) {
	__obf_761f92f380e69e60 := reflect.ValueOf(__obf_8d1c7d6c4b72dec7)
	__obf_25bae469ca95ff05 := reflect.ValueOf(__obf_22eaa72847d00940)
	switch __obf_cc73fd613fcc4064 {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.GTR, token.LEQ, token.GEQ, token.EQL, token.NEQ:
		var __obf_2f6590e78095cb9b, __obf_40dee8321fd779d9 float64
		var __obf_5a128dc949259d7a error
		if __obf_2f6590e78095cb9b, __obf_5a128dc949259d7a = __obf_9e412df46e486a0b(__obf_761f92f380e69e60); __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		if __obf_40dee8321fd779d9, __obf_5a128dc949259d7a = __obf_9e412df46e486a0b(__obf_25bae469ca95ff05); __obf_5a128dc949259d7a != nil {
			return nil, __obf_5a128dc949259d7a
		}
		switch __obf_cc73fd613fcc4064 {
		case token.ADD:
			return decimal.NewFromFloat(__obf_2f6590e78095cb9b).Add(decimal.NewFromFloat(__obf_40dee8321fd779d9)), nil
		case token.SUB:
			return decimal.NewFromFloat(__obf_2f6590e78095cb9b).Sub(decimal.NewFromFloat(__obf_40dee8321fd779d9)), nil
		case token.MUL:
			return decimal.NewFromFloat(__obf_2f6590e78095cb9b).Mul(decimal.NewFromFloat(__obf_40dee8321fd779d9)), nil
		case token.QUO:
			if __obf_40dee8321fd779d9 == 0 {
				return 0, errors.New("x/0 error")
			}
			return __obf_2f6590e78095cb9b / __obf_40dee8321fd779d9, nil
		case token.LSS:
			return __obf_2f6590e78095cb9b < __obf_40dee8321fd779d9, nil
		case token.GTR:
			return __obf_2f6590e78095cb9b > __obf_40dee8321fd779d9, nil
		case token.LEQ:
			return __obf_2f6590e78095cb9b <= __obf_40dee8321fd779d9, nil
		case token.GEQ:
			return __obf_2f6590e78095cb9b >= __obf_40dee8321fd779d9, nil
		case token.EQL:
			return __obf_2f6590e78095cb9b == __obf_40dee8321fd779d9, nil
		case token.NEQ:
			return __obf_2f6590e78095cb9b != __obf_40dee8321fd779d9, nil
		default:
			return 0, ErrUnsupportToken
		}
	case token.LAND, token.LOR:
		if __obf_761f92f380e69e60.Kind() != reflect.Bool || __obf_25bae469ca95ff05.Kind() != reflect.Bool {
			return false, ErrNotBool
		}
		switch __obf_cc73fd613fcc4064 {
		case token.LAND:
			return __obf_761f92f380e69e60.Bool() && __obf_25bae469ca95ff05.Bool(), nil
		case token.LOR:
			return __obf_761f92f380e69e60.Bool() || __obf_25bae469ca95ff05.Bool(), nil
		default:
			return false, ErrUnsupportToken
		}
	default:
		return nil, ErrUnsupportToken
	}
}

func __obf_9e412df46e486a0b(__obf_8d1c7d6c4b72dec7 reflect.Value) (float64, error) {
	switch __obf_8d1c7d6c4b72dec7.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(__obf_8d1c7d6c4b72dec7.Int()), nil
	case reflect.Float32, reflect.Float64:
		return __obf_8d1c7d6c4b72dec7.Float(), nil
	default:
		return 0, ErrNotNumber
	}
}

func __obf_eea7f74ec374b001(__obf_8d1c7d6c4b72dec7 reflect.Value) {

}
