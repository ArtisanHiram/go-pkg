package __obf_91620b895eeff9ed

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

var __obf_7b9593b00421fdba []byte
var __obf_4203f5be1e352cc4 []ValueType

func init() {
	__obf_7b9593b00421fdba = make([]byte, 256)
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < len(__obf_7b9593b00421fdba); __obf_5aa5c8829b97f182++ {
		__obf_7b9593b00421fdba[__obf_5aa5c8829b97f182] = 255
	}
	for __obf_5aa5c8829b97f182 := '0'; __obf_5aa5c8829b97f182 <= '9'; __obf_5aa5c8829b97f182++ {
		__obf_7b9593b00421fdba[__obf_5aa5c8829b97f182] = byte(__obf_5aa5c8829b97f182 - '0')
	}
	for __obf_5aa5c8829b97f182 := 'a'; __obf_5aa5c8829b97f182 <= 'f'; __obf_5aa5c8829b97f182++ {
		__obf_7b9593b00421fdba[__obf_5aa5c8829b97f182] = byte((__obf_5aa5c8829b97f182 - 'a') + 10)
	}
	for __obf_5aa5c8829b97f182 := 'A'; __obf_5aa5c8829b97f182 <= 'F'; __obf_5aa5c8829b97f182++ {
		__obf_7b9593b00421fdba[__obf_5aa5c8829b97f182] = byte((__obf_5aa5c8829b97f182 - 'A') + 10)
	}
	__obf_4203f5be1e352cc4 = make([]ValueType, 256)
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < len(__obf_4203f5be1e352cc4); __obf_5aa5c8829b97f182++ {
		__obf_4203f5be1e352cc4[__obf_5aa5c8829b97f182] = InvalidValue
	}
	__obf_4203f5be1e352cc4['"'] = StringValue
	__obf_4203f5be1e352cc4['-'] = NumberValue
	__obf_4203f5be1e352cc4['0'] = NumberValue
	__obf_4203f5be1e352cc4['1'] = NumberValue
	__obf_4203f5be1e352cc4['2'] = NumberValue
	__obf_4203f5be1e352cc4['3'] = NumberValue
	__obf_4203f5be1e352cc4['4'] = NumberValue
	__obf_4203f5be1e352cc4['5'] = NumberValue
	__obf_4203f5be1e352cc4['6'] = NumberValue
	__obf_4203f5be1e352cc4['7'] = NumberValue
	__obf_4203f5be1e352cc4['8'] = NumberValue
	__obf_4203f5be1e352cc4['9'] = NumberValue
	__obf_4203f5be1e352cc4['t'] = BoolValue
	__obf_4203f5be1e352cc4['f'] = BoolValue
	__obf_4203f5be1e352cc4['n'] = NilValue
	__obf_4203f5be1e352cc4['['] = ArrayValue
	__obf_4203f5be1e352cc4['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_892632c77e859037 *__obf_972c2293804d141c
	__obf_fe90e88a3ba69cf3 io.Reader
	__obf_184433571fa55237 []byte
	__obf_a657fb48fcb34e21 int
	__obf_15d837671d2809ae int
	__obf_1802821f4ca77d0e int
	__obf_f69c5dadb6d744d6 int
	__obf_7a95a5ed5bcae59e []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_892632c77e859037 API) *Iterator {
	return &Iterator{__obf_892632c77e859037: __obf_892632c77e859037.(*__obf_972c2293804d141c), __obf_fe90e88a3ba69cf3: nil, __obf_184433571fa55237: nil, __obf_a657fb48fcb34e21: 0, __obf_15d837671d2809ae: 0, __obf_1802821f4ca77d0e: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_892632c77e859037 API, __obf_fe90e88a3ba69cf3 io.Reader, __obf_0355f91790eba345 int) *Iterator {
	return &Iterator{__obf_892632c77e859037: __obf_892632c77e859037.(*__obf_972c2293804d141c), __obf_fe90e88a3ba69cf3: __obf_fe90e88a3ba69cf3, __obf_184433571fa55237: make([]byte, __obf_0355f91790eba345), __obf_a657fb48fcb34e21: 0, __obf_15d837671d2809ae: 0, __obf_1802821f4ca77d0e: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_892632c77e859037 API, __obf_8cff789a732c8734 []byte) *Iterator {
	return &Iterator{__obf_892632c77e859037: __obf_892632c77e859037.(*__obf_972c2293804d141c), __obf_fe90e88a3ba69cf3: nil, __obf_184433571fa55237: __obf_8cff789a732c8734, __obf_a657fb48fcb34e21: 0, __obf_15d837671d2809ae: len(__obf_8cff789a732c8734), __obf_1802821f4ca77d0e: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_892632c77e859037 API, __obf_8cff789a732c8734 string) *Iterator {
	return ParseBytes(__obf_892632c77e859037, []byte(__obf_8cff789a732c8734))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_1bb30e8a74ed8233 *Iterator) Pool() IteratorPool {
	return __obf_1bb30e8a74ed8233.

		// Reset reuse iterator instance by specifying another reader
		__obf_892632c77e859037
}

func (__obf_1bb30e8a74ed8233 *Iterator) Reset(__obf_fe90e88a3ba69cf3 io.Reader) *Iterator {
	__obf_1bb30e8a74ed8233.__obf_fe90e88a3ba69cf3 = __obf_fe90e88a3ba69cf3
	__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = 0
	__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae = 0
	__obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e = 0
	return __obf_1bb30e8a74ed8233
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_1bb30e8a74ed8233 *Iterator) ResetBytes(__obf_8cff789a732c8734 []byte) *Iterator {
	__obf_1bb30e8a74ed8233.__obf_fe90e88a3ba69cf3 = nil
	__obf_1bb30e8a74ed8233.__obf_184433571fa55237 = __obf_8cff789a732c8734
	__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = 0
	__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae = len(__obf_8cff789a732c8734)
	__obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e = 0
	return __obf_1bb30e8a74ed8233
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_1bb30e8a74ed8233 *Iterator) WhatIsNext() ValueType {
	__obf_f47701e28c5effac := __obf_4203f5be1e352cc4[__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()]
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	return __obf_f47701e28c5effac
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_f5ecc5d0ca8981a7() bool {
	for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		switch __obf_f16b4157911bc9af {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
		return false
	}
	return true
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_40d85895d419ee4a() bool {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == ',' {
		return false
	}
	if __obf_f16b4157911bc9af == '}' {
		return true
	}
	__obf_1bb30e8a74ed8233.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_f16b4157911bc9af}))
	return true
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_684faa48ae8c5049() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			switch __obf_f16b4157911bc9af {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
			return __obf_f16b4157911bc9af
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_1bb30e8a74ed8233 *Iterator) ReportError(__obf_c1c9252fd17989f1 string, __obf_6457f9c7d28f805d string) {
	if __obf_1bb30e8a74ed8233.Error != nil {
		if __obf_1bb30e8a74ed8233.Error != io.EOF {
			return
		}
	}
	__obf_67a281237a10cdf7 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 10
	if __obf_67a281237a10cdf7 < 0 {
		__obf_67a281237a10cdf7 = 0
	}
	__obf_0a7ed98b973e9178 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 + 10
	if __obf_0a7ed98b973e9178 > __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
		__obf_0a7ed98b973e9178 = __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae
	}
	__obf_00fca4a097ec3605 := string(__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_67a281237a10cdf7:__obf_0a7ed98b973e9178])
	__obf_62c8363367dd7d3c := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 50
	if __obf_62c8363367dd7d3c < 0 {
		__obf_62c8363367dd7d3c = 0
	}
	__obf_a885eedffa22960d := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 + 50
	if __obf_a885eedffa22960d > __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
		__obf_a885eedffa22960d = __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae
	}
	__obf_7063a37aff766b84 := string(__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_62c8363367dd7d3c:__obf_a885eedffa22960d])
	__obf_1bb30e8a74ed8233.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_c1c9252fd17989f1, __obf_6457f9c7d28f805d, __obf_1bb30e8a74ed8233.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_a657fb48fcb34e21-__obf_67a281237a10cdf7, __obf_00fca4a097ec3605, __obf_7063a37aff766b84)
}

func (__obf_1bb30e8a74ed8233 *Iterator) CurrentBuffer() string {
	__obf_67a281237a10cdf7 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 10
	if __obf_67a281237a10cdf7 < 0 {
		__obf_67a281237a10cdf7 = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21, string(__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_67a281237a10cdf7:__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21]), string(__obf_1bb30e8a74ed8233.__obf_184433571fa55237[0:__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae]))
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_9617ab9cc89bcddc() (__obf_e46f5fe3db5036fe byte) {
	if __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
		if __obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_e46f5fe3db5036fe = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21]
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21++
			return __obf_e46f5fe3db5036fe
		}
		return 0
	}
	__obf_e46f5fe3db5036fe = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21]
	__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21++
	return __obf_e46f5fe3db5036fe
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_e927802c539fc09d() bool {
	if __obf_1bb30e8a74ed8233.__obf_fe90e88a3ba69cf3 == nil {
		if __obf_1bb30e8a74ed8233.Error == nil {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae
			__obf_1bb30e8a74ed8233.
				Error = io.EOF
		}
		return false
	}
	if __obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e != nil {
		__obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e = append(__obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e, __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_f69c5dadb6d744d6:__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae]...)
		__obf_1bb30e8a74ed8233.__obf_f69c5dadb6d744d6 = 0
	}
	for {
		__obf_c80a670e818798d8, __obf_62967739bca1d11d := __obf_1bb30e8a74ed8233.__obf_fe90e88a3ba69cf3.Read(__obf_1bb30e8a74ed8233.__obf_184433571fa55237)
		if __obf_c80a670e818798d8 == 0 {
			if __obf_62967739bca1d11d != nil {
				if __obf_1bb30e8a74ed8233.Error == nil {
					__obf_1bb30e8a74ed8233.
						Error = __obf_62967739bca1d11d
				}
				return false
			}
		} else {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = 0
			__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae = __obf_c80a670e818798d8
			return true
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_a163df67f9bb1c4b() {
	if __obf_1bb30e8a74ed8233.Error != nil {
		return
	}
	__obf_1bb30e8a74ed8233.

		// Read read the next JSON element as generic some.
		__obf_a657fb48fcb34e21--
}

func (__obf_1bb30e8a74ed8233 *Iterator) Read() any {
	__obf_f47701e28c5effac := __obf_1bb30e8a74ed8233.WhatIsNext()
	switch __obf_f47701e28c5effac {
	case StringValue:
		return __obf_1bb30e8a74ed8233.ReadString()
	case NumberValue:
		if __obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_e05f98c9e7cc9089.UseNumber {
			return json.Number(__obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd())
		}
		return __obf_1bb30e8a74ed8233.ReadFloat64()
	case NilValue:
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_1bb30e8a74ed8233.ReadBool()
	case ArrayValue:
		__obf_6fc61a171ba497a3 := []any{}
		__obf_1bb30e8a74ed8233.
			ReadArrayCB(func(__obf_1bb30e8a74ed8233 *Iterator) bool {
				var __obf_97d7079fd160f22c any
				__obf_1bb30e8a74ed8233.
					ReadVal(&__obf_97d7079fd160f22c)
				__obf_6fc61a171ba497a3 = append(__obf_6fc61a171ba497a3, __obf_97d7079fd160f22c)
				return true
			})
		return __obf_6fc61a171ba497a3
	case ObjectValue:
		__obf_35136ef2ac0f94e4 := map[string]any{}
		__obf_1bb30e8a74ed8233.
			ReadMapCB(func(Iter *Iterator, __obf_7e01b5b4651074d4 string) bool {
				var __obf_97d7079fd160f22c any
				__obf_1bb30e8a74ed8233.
					ReadVal(&__obf_97d7079fd160f22c)
				__obf_35136ef2ac0f94e4[__obf_7e01b5b4651074d4] = __obf_97d7079fd160f22c
				return true
			})
		return __obf_35136ef2ac0f94e4
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_f47701e28c5effac))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_449d173e828a6c2b = 10000

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_4361ed1341bd481b() (__obf_f7c62f2fa8cc8f03 bool) {
	__obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e++
	if __obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e <= __obf_449d173e828a6c2b {
		return true
	}
	__obf_1bb30e8a74ed8233.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_e792aa2f6324acea() (__obf_f7c62f2fa8cc8f03 bool) {
	__obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e--
	if __obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e >= 0 {
		return true
	}
	__obf_1bb30e8a74ed8233.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
