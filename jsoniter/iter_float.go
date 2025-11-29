package __obf_91620b895eeff9ed

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_0cd2940545352aa6 []int8

const __obf_17acfd2d0998ee00 = int8(-1)
const __obf_9c5885b637c648b7 = int8(-2)
const __obf_e388812e12ada1d0 = int8(-3)

func init() {
	__obf_0cd2940545352aa6 = make([]int8, 256)
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < len(__obf_0cd2940545352aa6); __obf_5aa5c8829b97f182++ {
		__obf_0cd2940545352aa6[__obf_5aa5c8829b97f182] = __obf_17acfd2d0998ee00
	}
	for __obf_5aa5c8829b97f182 := int8('0'); __obf_5aa5c8829b97f182 <= int8('9'); __obf_5aa5c8829b97f182++ {
		__obf_0cd2940545352aa6[__obf_5aa5c8829b97f182] = __obf_5aa5c8829b97f182 - int8('0')
	}
	__obf_0cd2940545352aa6[','] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6[']'] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6['}'] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6[' '] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6['\t'] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6['\n'] = __obf_9c5885b637c648b7
	__obf_0cd2940545352aa6['.'] = __obf_e388812e12ada1d0
}

// ReadBigFloat read big.Float
func (__obf_1bb30e8a74ed8233 *Iterator) ReadBigFloat() (__obf_e46f5fe3db5036fe *big.Float) {
	__obf_e91bd2feb751e4f1 := __obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return nil
	}
	__obf_c44299137777f29c := 64
	if len(__obf_e91bd2feb751e4f1) > __obf_c44299137777f29c {
		__obf_c44299137777f29c = len(__obf_e91bd2feb751e4f1)
	}
	__obf_bbfd2ac8618a6f0c, _, __obf_62967739bca1d11d := big.ParseFloat(__obf_e91bd2feb751e4f1, 10, uint(__obf_c44299137777f29c), big.ToZero)
	if __obf_62967739bca1d11d != nil {
		__obf_1bb30e8a74ed8233.
			Error = __obf_62967739bca1d11d
		return nil
	}
	return __obf_bbfd2ac8618a6f0c
}

// ReadBigInt read big.Int
func (__obf_1bb30e8a74ed8233 *Iterator) ReadBigInt() (__obf_e46f5fe3db5036fe *big.Int) {
	__obf_e91bd2feb751e4f1 := __obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return nil
	}
	__obf_e46f5fe3db5036fe = big.NewInt(0)
	var __obf_f7c62f2fa8cc8f03 bool
	__obf_e46f5fe3db5036fe, __obf_f7c62f2fa8cc8f03 = __obf_e46f5fe3db5036fe.SetString(__obf_e91bd2feb751e4f1, 10)
	if !__obf_f7c62f2fa8cc8f03 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_e46f5fe3db5036fe
}

// ReadFloat32 read float32
func (__obf_1bb30e8a74ed8233 *Iterator) ReadFloat32() (__obf_e46f5fe3db5036fe float32) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		return -__obf_1bb30e8a74ed8233.__obf_4eba0dee5b9a6823()
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	return __obf_1bb30e8a74ed8233.__obf_4eba0dee5b9a6823()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_4eba0dee5b9a6823() (__obf_e46f5fe3db5036fe float32) {
	__obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.
		// first char
		__obf_a657fb48fcb34e21

	if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
		return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
	}
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
	__obf_5aa5c8829b97f182++
	__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
	switch __obf_2f4419eff5cde989 {
	case __obf_17acfd2d0998ee00:
		return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
	case __obf_9c5885b637c648b7:
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat32", "empty number")
		return
	case __obf_e388812e12ada1d0:
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
			return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
		}
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		switch __obf_f16b4157911bc9af {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_1bb30e8a74ed8233.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_4724d1b596d6a0c0 := uint64(__obf_2f4419eff5cde989)
__obf_102f485d94a4ea34: // chars before dot

	for ; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
		switch __obf_2f4419eff5cde989 {
		case __obf_17acfd2d0998ee00:
			return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
		case __obf_9c5885b637c648b7:
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			return float32(__obf_4724d1b596d6a0c0)
		case __obf_e388812e12ada1d0:
			break __obf_102f485d94a4ea34
		}
		if __obf_4724d1b596d6a0c0 > __obf_fea4c6f6b5068026 {
			return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
		}
		__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_f16b4157911bc9af == '.' {
		__obf_5aa5c8829b97f182++
		__obf_aec89de3b4506cda := 0
		if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
			return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
		}
		for ; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
			switch __obf_2f4419eff5cde989 {
			case __obf_9c5885b637c648b7:
				if __obf_aec89de3b4506cda > 0 && __obf_aec89de3b4506cda < len(__obf_20d949f9e88953b9) {
					__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
					return float32(float64(__obf_4724d1b596d6a0c0) / float64(__obf_20d949f9e88953b9[__obf_aec89de3b4506cda]))
				}
				// too many decimal places
				return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
			case __obf_17acfd2d0998ee00, __obf_e388812e12ada1d0:
				return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
			}
			__obf_aec89de3b4506cda++
			if __obf_4724d1b596d6a0c0 > __obf_fea4c6f6b5068026 {
				return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
			}
			__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989)
		}
	}
	return __obf_1bb30e8a74ed8233.__obf_2188a84873fd6708()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_9df99ebed87959cd() (__obf_e46f5fe3db5036fe string) {
	__obf_908b6e44ce700caa := [16]byte{}
	__obf_e91bd2feb751e4f1 := __obf_908b6e44ce700caa[0:0]
__obf_601698e611ef17c1:
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			switch __obf_f16b4157911bc9af {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, __obf_f16b4157911bc9af)
				continue
			default:
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				break __obf_601698e611ef17c1
			}
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return
	}
	if len(__obf_e91bd2feb751e4f1) == 0 {
		__obf_1bb30e8a74ed8233.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_e91bd2feb751e4f1))
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_2188a84873fd6708() (__obf_e46f5fe3db5036fe float32) {
	__obf_e91bd2feb751e4f1 := __obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return
	}
	__obf_6f7cc7192cca2c04 := __obf_cff7c5c361e2abed(__obf_e91bd2feb751e4f1)
	if __obf_6f7cc7192cca2c04 != "" {
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat32SlowPath", __obf_6f7cc7192cca2c04)
		return
	}
	__obf_bbfd2ac8618a6f0c, __obf_62967739bca1d11d := strconv.ParseFloat(__obf_e91bd2feb751e4f1, 32)
	if __obf_62967739bca1d11d != nil {
		__obf_1bb30e8a74ed8233.
			Error = __obf_62967739bca1d11d
		return
	}
	return float32(__obf_bbfd2ac8618a6f0c)
}

// ReadFloat64 read float64
func (__obf_1bb30e8a74ed8233 *Iterator) ReadFloat64() (__obf_e46f5fe3db5036fe float64) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		return -__obf_1bb30e8a74ed8233.__obf_d4fa3ba5dc0f9205()
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	return __obf_1bb30e8a74ed8233.__obf_d4fa3ba5dc0f9205()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_d4fa3ba5dc0f9205() (__obf_e46f5fe3db5036fe float64) {
	__obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.
		// first char
		__obf_a657fb48fcb34e21

	if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
		return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
	}
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
	__obf_5aa5c8829b97f182++
	__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
	switch __obf_2f4419eff5cde989 {
	case __obf_17acfd2d0998ee00:
		return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
	case __obf_9c5885b637c648b7:
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat64", "empty number")
		return
	case __obf_e388812e12ada1d0:
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
			return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
		}
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		switch __obf_f16b4157911bc9af {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_1bb30e8a74ed8233.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_4724d1b596d6a0c0 := uint64(__obf_2f4419eff5cde989)
__obf_102f485d94a4ea34: // chars before dot

	for ; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
		switch __obf_2f4419eff5cde989 {
		case __obf_17acfd2d0998ee00:
			return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
		case __obf_9c5885b637c648b7:
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			return float64(__obf_4724d1b596d6a0c0)
		case __obf_e388812e12ada1d0:
			break __obf_102f485d94a4ea34
		}
		if __obf_4724d1b596d6a0c0 > __obf_fea4c6f6b5068026 {
			return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
		}
		__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_f16b4157911bc9af == '.' {
		__obf_5aa5c8829b97f182++
		__obf_aec89de3b4506cda := 0
		if __obf_5aa5c8829b97f182 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
			return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
		}
		for ; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			__obf_2f4419eff5cde989 := __obf_0cd2940545352aa6[__obf_f16b4157911bc9af]
			switch __obf_2f4419eff5cde989 {
			case __obf_9c5885b637c648b7:
				if __obf_aec89de3b4506cda > 0 && __obf_aec89de3b4506cda < len(__obf_20d949f9e88953b9) {
					__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
					return float64(__obf_4724d1b596d6a0c0) / float64(__obf_20d949f9e88953b9[__obf_aec89de3b4506cda])
				}
				// too many decimal places
				return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
			case __obf_17acfd2d0998ee00, __obf_e388812e12ada1d0:
				return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
			}
			__obf_aec89de3b4506cda++
			if __obf_4724d1b596d6a0c0 > __obf_fea4c6f6b5068026 {
				return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
			}
			__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989)
			if __obf_4724d1b596d6a0c0 > __obf_397fd3fe0942535f {
				return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
			}
		}
	}
	return __obf_1bb30e8a74ed8233.__obf_b6ef8e9a2b817269()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_b6ef8e9a2b817269() (__obf_e46f5fe3db5036fe float64) {
	__obf_e91bd2feb751e4f1 := __obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return
	}
	__obf_6f7cc7192cca2c04 := __obf_cff7c5c361e2abed(__obf_e91bd2feb751e4f1)
	if __obf_6f7cc7192cca2c04 != "" {
		__obf_1bb30e8a74ed8233.
			ReportError("readFloat64SlowPath", __obf_6f7cc7192cca2c04)
		return
	}
	__obf_bbfd2ac8618a6f0c, __obf_62967739bca1d11d := strconv.ParseFloat(__obf_e91bd2feb751e4f1, 64)
	if __obf_62967739bca1d11d != nil {
		__obf_1bb30e8a74ed8233.
			Error = __obf_62967739bca1d11d
		return
	}
	return __obf_bbfd2ac8618a6f0c
}

func __obf_cff7c5c361e2abed(__obf_e91bd2feb751e4f1 string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_e91bd2feb751e4f1) == 0 {
		return "empty number"
	}
	if __obf_e91bd2feb751e4f1[0] == '-' {
		return "-- is not valid"
	}
	__obf_321aec92f4cda877 := strings.IndexByte(__obf_e91bd2feb751e4f1, '.')
	if __obf_321aec92f4cda877 != -1 {
		if __obf_321aec92f4cda877 == len(__obf_e91bd2feb751e4f1)-1 {
			return "dot can not be last character"
		}
		switch __obf_e91bd2feb751e4f1[__obf_321aec92f4cda877+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_1bb30e8a74ed8233 *Iterator) ReadNumber() (__obf_e46f5fe3db5036fe json.Number) {
	return json.Number(__obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd())
}
