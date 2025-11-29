package __obf_5b802ce8d9ba56d6

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

var __obf_970c42142eea13e0 []byte
var __obf_1e59f2a2903c120b []ValueType

func init() {
	__obf_970c42142eea13e0 = make([]byte, 256)
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < len(__obf_970c42142eea13e0); __obf_2deec7c38ffb6ae3++ {
		__obf_970c42142eea13e0[__obf_2deec7c38ffb6ae3] = 255
	}
	for __obf_2deec7c38ffb6ae3 := '0'; __obf_2deec7c38ffb6ae3 <= '9'; __obf_2deec7c38ffb6ae3++ {
		__obf_970c42142eea13e0[__obf_2deec7c38ffb6ae3] = byte(__obf_2deec7c38ffb6ae3 - '0')
	}
	for __obf_2deec7c38ffb6ae3 := 'a'; __obf_2deec7c38ffb6ae3 <= 'f'; __obf_2deec7c38ffb6ae3++ {
		__obf_970c42142eea13e0[__obf_2deec7c38ffb6ae3] = byte((__obf_2deec7c38ffb6ae3 - 'a') + 10)
	}
	for __obf_2deec7c38ffb6ae3 := 'A'; __obf_2deec7c38ffb6ae3 <= 'F'; __obf_2deec7c38ffb6ae3++ {
		__obf_970c42142eea13e0[__obf_2deec7c38ffb6ae3] = byte((__obf_2deec7c38ffb6ae3 - 'A') + 10)
	}
	__obf_1e59f2a2903c120b = make([]ValueType, 256)
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < len(__obf_1e59f2a2903c120b); __obf_2deec7c38ffb6ae3++ {
		__obf_1e59f2a2903c120b[__obf_2deec7c38ffb6ae3] = InvalidValue
	}
	__obf_1e59f2a2903c120b['"'] = StringValue
	__obf_1e59f2a2903c120b['-'] = NumberValue
	__obf_1e59f2a2903c120b['0'] = NumberValue
	__obf_1e59f2a2903c120b['1'] = NumberValue
	__obf_1e59f2a2903c120b['2'] = NumberValue
	__obf_1e59f2a2903c120b['3'] = NumberValue
	__obf_1e59f2a2903c120b['4'] = NumberValue
	__obf_1e59f2a2903c120b['5'] = NumberValue
	__obf_1e59f2a2903c120b['6'] = NumberValue
	__obf_1e59f2a2903c120b['7'] = NumberValue
	__obf_1e59f2a2903c120b['8'] = NumberValue
	__obf_1e59f2a2903c120b['9'] = NumberValue
	__obf_1e59f2a2903c120b['t'] = BoolValue
	__obf_1e59f2a2903c120b['f'] = BoolValue
	__obf_1e59f2a2903c120b['n'] = NilValue
	__obf_1e59f2a2903c120b['['] = ArrayValue
	__obf_1e59f2a2903c120b['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
	__obf_1fcab394164d9b41 io.Reader
	__obf_9fc06d9180f0daca []byte
	__obf_14babd6f9a55bd36 int
	__obf_3a36550914545c79 int
	__obf_e0b51eaba31d0671 int
	__obf_787a5df6d8c7d251 int
	__obf_645fd9e4de8f601b []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_dca106e1cdf85392 API) *Iterator {
	return &Iterator{__obf_dca106e1cdf85392: __obf_dca106e1cdf85392.(*__obf_5d13d7f3bd06c6cf), __obf_1fcab394164d9b41: nil, __obf_9fc06d9180f0daca: nil, __obf_14babd6f9a55bd36: 0, __obf_3a36550914545c79: 0, __obf_e0b51eaba31d0671: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_dca106e1cdf85392 API, __obf_1fcab394164d9b41 io.Reader, __obf_3f8315b01fa921cf int) *Iterator {
	return &Iterator{__obf_dca106e1cdf85392: __obf_dca106e1cdf85392.(*__obf_5d13d7f3bd06c6cf), __obf_1fcab394164d9b41: __obf_1fcab394164d9b41, __obf_9fc06d9180f0daca: make([]byte, __obf_3f8315b01fa921cf), __obf_14babd6f9a55bd36: 0, __obf_3a36550914545c79: 0, __obf_e0b51eaba31d0671: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_dca106e1cdf85392 API, __obf_afc7a7cda61d5a23 []byte) *Iterator {
	return &Iterator{__obf_dca106e1cdf85392: __obf_dca106e1cdf85392.(*__obf_5d13d7f3bd06c6cf), __obf_1fcab394164d9b41: nil, __obf_9fc06d9180f0daca: __obf_afc7a7cda61d5a23, __obf_14babd6f9a55bd36: 0, __obf_3a36550914545c79: len(__obf_afc7a7cda61d5a23), __obf_e0b51eaba31d0671: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_dca106e1cdf85392 API, __obf_afc7a7cda61d5a23 string) *Iterator {
	return ParseBytes(__obf_dca106e1cdf85392, []byte(__obf_afc7a7cda61d5a23))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_67008a6a9e5ba828 *Iterator) Pool() IteratorPool {
	return __obf_67008a6a9e5ba828.

		// Reset reuse iterator instance by specifying another reader
		__obf_dca106e1cdf85392
}

func (__obf_67008a6a9e5ba828 *Iterator) Reset(__obf_1fcab394164d9b41 io.Reader) *Iterator {
	__obf_67008a6a9e5ba828.__obf_1fcab394164d9b41 = __obf_1fcab394164d9b41
	__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = 0
	__obf_67008a6a9e5ba828.__obf_3a36550914545c79 = 0
	__obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671 = 0
	return __obf_67008a6a9e5ba828
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_67008a6a9e5ba828 *Iterator) ResetBytes(__obf_afc7a7cda61d5a23 []byte) *Iterator {
	__obf_67008a6a9e5ba828.__obf_1fcab394164d9b41 = nil
	__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca = __obf_afc7a7cda61d5a23
	__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = 0
	__obf_67008a6a9e5ba828.__obf_3a36550914545c79 = len(__obf_afc7a7cda61d5a23)
	__obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671 = 0
	return __obf_67008a6a9e5ba828
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_67008a6a9e5ba828 *Iterator) WhatIsNext() ValueType {
	__obf_9ad2e389a2a6fdb8 := __obf_1e59f2a2903c120b[__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()]
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	return __obf_9ad2e389a2a6fdb8
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_2f7f02fb7eb25964() bool {
	for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		switch __obf_dab9baaadfa7c8c2 {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
		return false
	}
	return true
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_cddc9a4e84baa583() bool {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == ',' {
		return false
	}
	if __obf_dab9baaadfa7c8c2 == '}' {
		return true
	}
	__obf_67008a6a9e5ba828.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_dab9baaadfa7c8c2}))
	return true
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_b781a59d5a0d2490() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			switch __obf_dab9baaadfa7c8c2 {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
			return __obf_dab9baaadfa7c8c2
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_67008a6a9e5ba828 *Iterator) ReportError(__obf_9e6ac1d1087b1b28 string, __obf_599756f70dc37f5f string) {
	if __obf_67008a6a9e5ba828.Error != nil {
		if __obf_67008a6a9e5ba828.Error != io.EOF {
			return
		}
	}
	__obf_7d518e38df477a75 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 10
	if __obf_7d518e38df477a75 < 0 {
		__obf_7d518e38df477a75 = 0
	}
	__obf_27015c72f401e640 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 + 10
	if __obf_27015c72f401e640 > __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
		__obf_27015c72f401e640 = __obf_67008a6a9e5ba828.__obf_3a36550914545c79
	}
	__obf_97c80978049ece93 := string(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_7d518e38df477a75:__obf_27015c72f401e640])
	__obf_aabf89460f6745cd := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 50
	if __obf_aabf89460f6745cd < 0 {
		__obf_aabf89460f6745cd = 0
	}
	__obf_242ccef47ae0b8b5 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 + 50
	if __obf_242ccef47ae0b8b5 > __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
		__obf_242ccef47ae0b8b5 = __obf_67008a6a9e5ba828.__obf_3a36550914545c79
	}
	__obf_d6482764e484c6a4 := string(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_aabf89460f6745cd:__obf_242ccef47ae0b8b5])
	__obf_67008a6a9e5ba828.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_9e6ac1d1087b1b28, __obf_599756f70dc37f5f, __obf_67008a6a9e5ba828.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_14babd6f9a55bd36-__obf_7d518e38df477a75, __obf_97c80978049ece93, __obf_d6482764e484c6a4)
}

func (__obf_67008a6a9e5ba828 *Iterator) CurrentBuffer() string {
	__obf_7d518e38df477a75 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 10
	if __obf_7d518e38df477a75 < 0 {
		__obf_7d518e38df477a75 = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36, string(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_7d518e38df477a75:__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36]), string(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[0:__obf_67008a6a9e5ba828.__obf_3a36550914545c79]))
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_ea3ebd5c6789bccb() (__obf_5dabcdfee5097ed6 byte) {
	if __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
		if __obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_5dabcdfee5097ed6 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36]
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36++
			return __obf_5dabcdfee5097ed6
		}
		return 0
	}
	__obf_5dabcdfee5097ed6 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36]
	__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36++
	return __obf_5dabcdfee5097ed6
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_538f3d3337e0395f() bool {
	if __obf_67008a6a9e5ba828.__obf_1fcab394164d9b41 == nil {
		if __obf_67008a6a9e5ba828.Error == nil {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_67008a6a9e5ba828.__obf_3a36550914545c79
			__obf_67008a6a9e5ba828.
				Error = io.EOF
		}
		return false
	}
	if __obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b != nil {
		__obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b = append(__obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b, __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_787a5df6d8c7d251:__obf_67008a6a9e5ba828.__obf_3a36550914545c79]...)
		__obf_67008a6a9e5ba828.__obf_787a5df6d8c7d251 = 0
	}
	for {
		__obf_2c0ce853cb81f014, __obf_6d071d48f3b321ad := __obf_67008a6a9e5ba828.__obf_1fcab394164d9b41.Read(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca)
		if __obf_2c0ce853cb81f014 == 0 {
			if __obf_6d071d48f3b321ad != nil {
				if __obf_67008a6a9e5ba828.Error == nil {
					__obf_67008a6a9e5ba828.
						Error = __obf_6d071d48f3b321ad
				}
				return false
			}
		} else {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = 0
			__obf_67008a6a9e5ba828.__obf_3a36550914545c79 = __obf_2c0ce853cb81f014
			return true
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_3284a1eaa2a0abb6() {
	if __obf_67008a6a9e5ba828.Error != nil {
		return
	}
	__obf_67008a6a9e5ba828.

		// Read read the next JSON element as generic some.
		__obf_14babd6f9a55bd36--
}

func (__obf_67008a6a9e5ba828 *Iterator) Read() any {
	__obf_9ad2e389a2a6fdb8 := __obf_67008a6a9e5ba828.WhatIsNext()
	switch __obf_9ad2e389a2a6fdb8 {
	case StringValue:
		return __obf_67008a6a9e5ba828.ReadString()
	case NumberValue:
		if __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_97c6754e53034071.UseNumber {
			return json.Number(__obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05())
		}
		return __obf_67008a6a9e5ba828.ReadFloat64()
	case NilValue:
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_67008a6a9e5ba828.ReadBool()
	case ArrayValue:
		__obf_0ac6bb318bf53d42 := []any{}
		__obf_67008a6a9e5ba828.
			ReadArrayCB(func(__obf_67008a6a9e5ba828 *Iterator) bool {
				var __obf_77670b1f1899744d any
				__obf_67008a6a9e5ba828.
					ReadVal(&__obf_77670b1f1899744d)
				__obf_0ac6bb318bf53d42 = append(__obf_0ac6bb318bf53d42, __obf_77670b1f1899744d)
				return true
			})
		return __obf_0ac6bb318bf53d42
	case ObjectValue:
		__obf_f9b1526531f3c024 := map[string]any{}
		__obf_67008a6a9e5ba828.
			ReadMapCB(func(Iter *Iterator, __obf_61998fb083387c3c string) bool {
				var __obf_77670b1f1899744d any
				__obf_67008a6a9e5ba828.
					ReadVal(&__obf_77670b1f1899744d)
				__obf_f9b1526531f3c024[__obf_61998fb083387c3c] = __obf_77670b1f1899744d
				return true
			})
		return __obf_f9b1526531f3c024
	default:
		__obf_67008a6a9e5ba828.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_9ad2e389a2a6fdb8))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_7eb3239d116f59eb = 10000

func (__obf_67008a6a9e5ba828 *Iterator) __obf_093e3593687574f7() (__obf_96ae77751ddfcfbd bool) {
	__obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671++
	if __obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671 <= __obf_7eb3239d116f59eb {
		return true
	}
	__obf_67008a6a9e5ba828.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_c37c04dba052f09d() (__obf_96ae77751ddfcfbd bool) {
	__obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671--
	if __obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671 >= 0 {
		return true
	}
	__obf_67008a6a9e5ba828.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
