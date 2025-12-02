package __obf_c7b79b12b33d8238

import (
	"encoding/json"
	"fmt"
	"io"
)

// ValueType the type for JSON element
type ValueType int

const (
	// InvalidValue invalid JSON element
	InvalidValue ValueType = iota
	// StringValue JSON element "string"
	StringValue
	// NumberValue JSON element 100 or 0.10
	NumberValue
	// NilValue JSON element null
	NilValue
	// BoolValue JSON element true or false
	BoolValue
	// ArrayValue JSON element []
	ArrayValue
	// ObjectValue JSON element {}
	ObjectValue
)

var __obf_943cde3a7c81fae6 []byte
var __obf_6f335b1ee3f1b0aa []ValueType

func init() {
	__obf_943cde3a7c81fae6 = make([]byte, 256)
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < len(__obf_943cde3a7c81fae6); __obf_a051269af3a541bb++ {
		__obf_943cde3a7c81fae6[__obf_a051269af3a541bb] = 255
	}
	for __obf_a051269af3a541bb := '0'; __obf_a051269af3a541bb <= '9'; __obf_a051269af3a541bb++ {
		__obf_943cde3a7c81fae6[__obf_a051269af3a541bb] = byte(__obf_a051269af3a541bb - '0')
	}
	for __obf_a051269af3a541bb := 'a'; __obf_a051269af3a541bb <= 'f'; __obf_a051269af3a541bb++ {
		__obf_943cde3a7c81fae6[__obf_a051269af3a541bb] = byte((__obf_a051269af3a541bb - 'a') + 10)
	}
	for __obf_a051269af3a541bb := 'A'; __obf_a051269af3a541bb <= 'F'; __obf_a051269af3a541bb++ {
		__obf_943cde3a7c81fae6[__obf_a051269af3a541bb] = byte((__obf_a051269af3a541bb - 'A') + 10)
	}
	__obf_6f335b1ee3f1b0aa = make([]ValueType, 256)
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < len(__obf_6f335b1ee3f1b0aa); __obf_a051269af3a541bb++ {
		__obf_6f335b1ee3f1b0aa[__obf_a051269af3a541bb] = InvalidValue
	}
	__obf_6f335b1ee3f1b0aa['"'] = StringValue
	__obf_6f335b1ee3f1b0aa['-'] = NumberValue
	__obf_6f335b1ee3f1b0aa['0'] = NumberValue
	__obf_6f335b1ee3f1b0aa['1'] = NumberValue
	__obf_6f335b1ee3f1b0aa['2'] = NumberValue
	__obf_6f335b1ee3f1b0aa['3'] = NumberValue
	__obf_6f335b1ee3f1b0aa['4'] = NumberValue
	__obf_6f335b1ee3f1b0aa['5'] = NumberValue
	__obf_6f335b1ee3f1b0aa['6'] = NumberValue
	__obf_6f335b1ee3f1b0aa['7'] = NumberValue
	__obf_6f335b1ee3f1b0aa['8'] = NumberValue
	__obf_6f335b1ee3f1b0aa['9'] = NumberValue
	__obf_6f335b1ee3f1b0aa['t'] = BoolValue
	__obf_6f335b1ee3f1b0aa['f'] = BoolValue
	__obf_6f335b1ee3f1b0aa['n'] = NilValue
	__obf_6f335b1ee3f1b0aa['['] = ArrayValue
	__obf_6f335b1ee3f1b0aa['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
	__obf_801a7eebc4303617 io.Reader
	__obf_684d738bc3ac851a []byte
	__obf_6a63c9ac34fe84e2 int
	__obf_840246080559848c int
	__obf_5e89e9b75fd360f4 int
	__obf_df8f35c95ef68415 int
	__obf_228f49de295a387e []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_c52e0bcfad4b8b71 API) *Iterator {
	return &Iterator{__obf_c52e0bcfad4b8b71: __obf_c52e0bcfad4b8b71.(*__obf_30fe5c95cabd69c0), __obf_801a7eebc4303617: nil, __obf_684d738bc3ac851a: nil, __obf_6a63c9ac34fe84e2: 0, __obf_840246080559848c: 0, __obf_5e89e9b75fd360f4: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_c52e0bcfad4b8b71 API, __obf_801a7eebc4303617 io.Reader, __obf_90520bf89c30ab5c int) *Iterator {
	return &Iterator{__obf_c52e0bcfad4b8b71: __obf_c52e0bcfad4b8b71.(*__obf_30fe5c95cabd69c0), __obf_801a7eebc4303617: __obf_801a7eebc4303617, __obf_684d738bc3ac851a: make([]byte, __obf_90520bf89c30ab5c), __obf_6a63c9ac34fe84e2: 0, __obf_840246080559848c: 0, __obf_5e89e9b75fd360f4: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_c52e0bcfad4b8b71 API, __obf_ed76dc67aa23a4da []byte) *Iterator {
	return &Iterator{__obf_c52e0bcfad4b8b71: __obf_c52e0bcfad4b8b71.(*__obf_30fe5c95cabd69c0), __obf_801a7eebc4303617: nil, __obf_684d738bc3ac851a: __obf_ed76dc67aa23a4da, __obf_6a63c9ac34fe84e2: 0, __obf_840246080559848c: len(__obf_ed76dc67aa23a4da), __obf_5e89e9b75fd360f4: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_c52e0bcfad4b8b71 API, __obf_ed76dc67aa23a4da string) *Iterator {
	return ParseBytes(__obf_c52e0bcfad4b8b71, []byte(__obf_ed76dc67aa23a4da))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_0da8c843dad13139 *Iterator) Pool() IteratorPool {
	return __obf_0da8c843dad13139.

		// Reset reuse iterator instance by specifying another reader
		__obf_c52e0bcfad4b8b71
}

func (__obf_0da8c843dad13139 *Iterator) Reset(__obf_801a7eebc4303617 io.Reader) *Iterator {
	__obf_0da8c843dad13139.__obf_801a7eebc4303617 = __obf_801a7eebc4303617
	__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = 0
	__obf_0da8c843dad13139.__obf_840246080559848c = 0
	__obf_0da8c843dad13139.__obf_5e89e9b75fd360f4 = 0
	return __obf_0da8c843dad13139
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_0da8c843dad13139 *Iterator) ResetBytes(__obf_ed76dc67aa23a4da []byte) *Iterator {
	__obf_0da8c843dad13139.__obf_801a7eebc4303617 = nil
	__obf_0da8c843dad13139.__obf_684d738bc3ac851a = __obf_ed76dc67aa23a4da
	__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = 0
	__obf_0da8c843dad13139.__obf_840246080559848c = len(__obf_ed76dc67aa23a4da)
	__obf_0da8c843dad13139.__obf_5e89e9b75fd360f4 = 0
	return __obf_0da8c843dad13139
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_0da8c843dad13139 *Iterator) WhatIsNext() ValueType {
	__obf_2cd064ad0a3f5c35 := __obf_6f335b1ee3f1b0aa[__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()]
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	return __obf_2cd064ad0a3f5c35
}

func (__obf_0da8c843dad13139 *Iterator) __obf_1e77e55a706a1586() bool {
	for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		switch __obf_12577c03a12f0d2c {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
		return false
	}
	return true
}

func (__obf_0da8c843dad13139 *Iterator) __obf_5059d0524ee888e7() bool {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == ',' {
		return false
	}
	if __obf_12577c03a12f0d2c == '}' {
		return true
	}
	__obf_0da8c843dad13139.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_12577c03a12f0d2c}))
	return true
}

func (__obf_0da8c843dad13139 *Iterator) __obf_2b436fcb1c0ca36d() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			switch __obf_12577c03a12f0d2c {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
			return __obf_12577c03a12f0d2c
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_0da8c843dad13139 *Iterator) ReportError(__obf_b2f2a59d70a9c2d9 string, __obf_e2c48e9bb87a8f24 string) {
	if __obf_0da8c843dad13139.Error != nil {
		if __obf_0da8c843dad13139.Error != io.EOF {
			return
		}
	}
	__obf_977f26db222f53cf := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 10
	if __obf_977f26db222f53cf < 0 {
		__obf_977f26db222f53cf = 0
	}
	__obf_a72812885959d979 := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 + 10
	if __obf_a72812885959d979 > __obf_0da8c843dad13139.__obf_840246080559848c {
		__obf_a72812885959d979 = __obf_0da8c843dad13139.__obf_840246080559848c
	}
	__obf_d123e916c3076d2b := string(__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_977f26db222f53cf:__obf_a72812885959d979])
	__obf_bc8a025a338025c9 := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 50
	if __obf_bc8a025a338025c9 < 0 {
		__obf_bc8a025a338025c9 = 0
	}
	__obf_0ae4333cc7fe523c := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 + 50
	if __obf_0ae4333cc7fe523c > __obf_0da8c843dad13139.__obf_840246080559848c {
		__obf_0ae4333cc7fe523c = __obf_0da8c843dad13139.__obf_840246080559848c
	}
	__obf_9d1d1126e41e269c := string(__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_bc8a025a338025c9:__obf_0ae4333cc7fe523c])
	__obf_0da8c843dad13139.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_b2f2a59d70a9c2d9, __obf_e2c48e9bb87a8f24, __obf_0da8c843dad13139.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_6a63c9ac34fe84e2-__obf_977f26db222f53cf, __obf_d123e916c3076d2b, __obf_9d1d1126e41e269c)
}

func (__obf_0da8c843dad13139 *Iterator) CurrentBuffer() string {
	__obf_977f26db222f53cf := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 10
	if __obf_977f26db222f53cf < 0 {
		__obf_977f26db222f53cf = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2, string(__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_977f26db222f53cf:__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2]), string(__obf_0da8c843dad13139.__obf_684d738bc3ac851a[0:__obf_0da8c843dad13139.__obf_840246080559848c]))
}

func (__obf_0da8c843dad13139 *Iterator) __obf_fe749dd9dadf32f8() (__obf_9a8638dac9e1c083 byte) {
	if __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 == __obf_0da8c843dad13139.__obf_840246080559848c {
		if __obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_9a8638dac9e1c083 = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2]
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2++
			return __obf_9a8638dac9e1c083
		}
		return 0
	}
	__obf_9a8638dac9e1c083 = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2]
	__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2++
	return __obf_9a8638dac9e1c083
}

func (__obf_0da8c843dad13139 *Iterator) __obf_baaf768e13842431() bool {
	if __obf_0da8c843dad13139.__obf_801a7eebc4303617 == nil {
		if __obf_0da8c843dad13139.Error == nil {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_0da8c843dad13139.__obf_840246080559848c
			__obf_0da8c843dad13139.
				Error = io.EOF
		}
		return false
	}
	if __obf_0da8c843dad13139.__obf_228f49de295a387e != nil {
		__obf_0da8c843dad13139.__obf_228f49de295a387e = append(__obf_0da8c843dad13139.__obf_228f49de295a387e, __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_df8f35c95ef68415:__obf_0da8c843dad13139.__obf_840246080559848c]...)
		__obf_0da8c843dad13139.__obf_df8f35c95ef68415 = 0
	}
	for {
		__obf_ff4330e73e137d5c, __obf_ea0680f8b461a85b := __obf_0da8c843dad13139.__obf_801a7eebc4303617.Read(__obf_0da8c843dad13139.__obf_684d738bc3ac851a)
		if __obf_ff4330e73e137d5c == 0 {
			if __obf_ea0680f8b461a85b != nil {
				if __obf_0da8c843dad13139.Error == nil {
					__obf_0da8c843dad13139.
						Error = __obf_ea0680f8b461a85b
				}
				return false
			}
		} else {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = 0
			__obf_0da8c843dad13139.__obf_840246080559848c = __obf_ff4330e73e137d5c
			return true
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_a675366c80290b83() {
	if __obf_0da8c843dad13139.Error != nil {
		return
	}
	__obf_0da8c843dad13139.

		// Read read the next JSON element as generic some.
		__obf_6a63c9ac34fe84e2--
}

func (__obf_0da8c843dad13139 *Iterator) Read() any {
	__obf_2cd064ad0a3f5c35 := __obf_0da8c843dad13139.WhatIsNext()
	switch __obf_2cd064ad0a3f5c35 {
	case StringValue:
		return __obf_0da8c843dad13139.ReadString()
	case NumberValue:
		if __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc.UseNumber {
			return json.Number(__obf_0da8c843dad13139.__obf_a678d14dd84cbb6b())
		}
		return __obf_0da8c843dad13139.ReadFloat64()
	case NilValue:
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_0da8c843dad13139.ReadBool()
	case ArrayValue:
		__obf_3e7cc2c62bf39a94 := []any{}
		__obf_0da8c843dad13139.
			ReadArrayCB(func(__obf_0da8c843dad13139 *Iterator) bool {
				var __obf_43cefbbf0648e3be any
				__obf_0da8c843dad13139.
					ReadVal(&__obf_43cefbbf0648e3be)
				__obf_3e7cc2c62bf39a94 = append(__obf_3e7cc2c62bf39a94, __obf_43cefbbf0648e3be)
				return true
			})
		return __obf_3e7cc2c62bf39a94
	case ObjectValue:
		__obf_d6e2df4782353c64 := map[string]any{}
		__obf_0da8c843dad13139.
			ReadMapCB(func(Iter *Iterator, __obf_ad34f8a6de2b01cb string) bool {
				var __obf_43cefbbf0648e3be any
				__obf_0da8c843dad13139.
					ReadVal(&__obf_43cefbbf0648e3be)
				__obf_d6e2df4782353c64[__obf_ad34f8a6de2b01cb] = __obf_43cefbbf0648e3be
				return true
			})
		return __obf_d6e2df4782353c64
	default:
		__obf_0da8c843dad13139.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_2cd064ad0a3f5c35))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_dc0bbe086b01c6bc = 10000

func (__obf_0da8c843dad13139 *Iterator) __obf_b638f6394ae6a28c() (__obf_b145180ce2c394bf bool) {
	__obf_0da8c843dad13139.__obf_5e89e9b75fd360f4++
	if __obf_0da8c843dad13139.__obf_5e89e9b75fd360f4 <= __obf_dc0bbe086b01c6bc {
		return true
	}
	__obf_0da8c843dad13139.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_0da8c843dad13139 *Iterator) __obf_aed83e308a23ff9c() (__obf_b145180ce2c394bf bool) {
	__obf_0da8c843dad13139.__obf_5e89e9b75fd360f4--
	if __obf_0da8c843dad13139.__obf_5e89e9b75fd360f4 >= 0 {
		return true
	}
	__obf_0da8c843dad13139.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
