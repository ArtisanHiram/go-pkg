package __obf_030d39f7a8de96c6

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

var __obf_6ef1311452044327 []byte
var __obf_c47d6682a0c4b75b []ValueType

func init() {
	__obf_6ef1311452044327 = make([]byte, 256)
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < len(__obf_6ef1311452044327); __obf_82c6e05b9d226c58++ {
		__obf_6ef1311452044327[__obf_82c6e05b9d226c58] = 255
	}
	for __obf_82c6e05b9d226c58 := '0'; __obf_82c6e05b9d226c58 <= '9'; __obf_82c6e05b9d226c58++ {
		__obf_6ef1311452044327[__obf_82c6e05b9d226c58] = byte(__obf_82c6e05b9d226c58 - '0')
	}
	for __obf_82c6e05b9d226c58 := 'a'; __obf_82c6e05b9d226c58 <= 'f'; __obf_82c6e05b9d226c58++ {
		__obf_6ef1311452044327[__obf_82c6e05b9d226c58] = byte((__obf_82c6e05b9d226c58 - 'a') + 10)
	}
	for __obf_82c6e05b9d226c58 := 'A'; __obf_82c6e05b9d226c58 <= 'F'; __obf_82c6e05b9d226c58++ {
		__obf_6ef1311452044327[__obf_82c6e05b9d226c58] = byte((__obf_82c6e05b9d226c58 - 'A') + 10)
	}
	__obf_c47d6682a0c4b75b = make([]ValueType, 256)
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < len(__obf_c47d6682a0c4b75b); __obf_82c6e05b9d226c58++ {
		__obf_c47d6682a0c4b75b[__obf_82c6e05b9d226c58] = InvalidValue
	}
	__obf_c47d6682a0c4b75b['"'] = StringValue
	__obf_c47d6682a0c4b75b['-'] = NumberValue
	__obf_c47d6682a0c4b75b['0'] = NumberValue
	__obf_c47d6682a0c4b75b['1'] = NumberValue
	__obf_c47d6682a0c4b75b['2'] = NumberValue
	__obf_c47d6682a0c4b75b['3'] = NumberValue
	__obf_c47d6682a0c4b75b['4'] = NumberValue
	__obf_c47d6682a0c4b75b['5'] = NumberValue
	__obf_c47d6682a0c4b75b['6'] = NumberValue
	__obf_c47d6682a0c4b75b['7'] = NumberValue
	__obf_c47d6682a0c4b75b['8'] = NumberValue
	__obf_c47d6682a0c4b75b['9'] = NumberValue
	__obf_c47d6682a0c4b75b['t'] = BoolValue
	__obf_c47d6682a0c4b75b['f'] = BoolValue
	__obf_c47d6682a0c4b75b['n'] = NilValue
	__obf_c47d6682a0c4b75b['['] = ArrayValue
	__obf_c47d6682a0c4b75b['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_679611dc56ff465b *__obf_81639538813814ff
	__obf_7582c70e81d83895 io.Reader
	__obf_0b1656d7ef4bd9df []byte
	__obf_5e22d6270d27491f int
	__obf_7c17042d9b4e73cc int
	__obf_66f73b1f1e6b88e2 int
	__obf_85ab55dc89939a18 int
	__obf_4541c922f6acdb7d []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_679611dc56ff465b API) *Iterator {
	return &Iterator{__obf_679611dc56ff465b: __obf_679611dc56ff465b.(*__obf_81639538813814ff), __obf_7582c70e81d83895: nil, __obf_0b1656d7ef4bd9df: nil, __obf_5e22d6270d27491f: 0, __obf_7c17042d9b4e73cc: 0, __obf_66f73b1f1e6b88e2: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_679611dc56ff465b API, __obf_7582c70e81d83895 io.Reader, __obf_8e50793d0b2df740 int) *Iterator {
	return &Iterator{__obf_679611dc56ff465b: __obf_679611dc56ff465b.(*__obf_81639538813814ff), __obf_7582c70e81d83895: __obf_7582c70e81d83895, __obf_0b1656d7ef4bd9df: make([]byte, __obf_8e50793d0b2df740), __obf_5e22d6270d27491f: 0, __obf_7c17042d9b4e73cc: 0, __obf_66f73b1f1e6b88e2: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_679611dc56ff465b API, __obf_518a47613a744158 []byte) *Iterator {
	return &Iterator{__obf_679611dc56ff465b: __obf_679611dc56ff465b.(*__obf_81639538813814ff), __obf_7582c70e81d83895: nil, __obf_0b1656d7ef4bd9df: __obf_518a47613a744158, __obf_5e22d6270d27491f: 0, __obf_7c17042d9b4e73cc: len(__obf_518a47613a744158), __obf_66f73b1f1e6b88e2: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_679611dc56ff465b API, __obf_518a47613a744158 string) *Iterator {
	return ParseBytes(__obf_679611dc56ff465b, []byte(__obf_518a47613a744158))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_4ab56a99965952e7 *Iterator) Pool() IteratorPool {
	return __obf_4ab56a99965952e7.

		// Reset reuse iterator instance by specifying another reader
		__obf_679611dc56ff465b
}

func (__obf_4ab56a99965952e7 *Iterator) Reset(__obf_7582c70e81d83895 io.Reader) *Iterator {
	__obf_4ab56a99965952e7.__obf_7582c70e81d83895 = __obf_7582c70e81d83895
	__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = 0
	__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc = 0
	__obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2 = 0
	return __obf_4ab56a99965952e7
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_4ab56a99965952e7 *Iterator) ResetBytes(__obf_518a47613a744158 []byte) *Iterator {
	__obf_4ab56a99965952e7.__obf_7582c70e81d83895 = nil
	__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df = __obf_518a47613a744158
	__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = 0
	__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc = len(__obf_518a47613a744158)
	__obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2 = 0
	return __obf_4ab56a99965952e7
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_4ab56a99965952e7 *Iterator) WhatIsNext() ValueType {
	__obf_8afc98a56c05c14e := __obf_c47d6682a0c4b75b[__obf_4ab56a99965952e7.__obf_61df301d1f67ad73()]
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	return __obf_8afc98a56c05c14e
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_3523e5db1e354e93() bool {
	for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		switch __obf_24309b3b0ff9ba22 {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
		return false
	}
	return true
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_6a883651c29f40bf() bool {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == ',' {
		return false
	}
	if __obf_24309b3b0ff9ba22 == '}' {
		return true
	}
	__obf_4ab56a99965952e7.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_24309b3b0ff9ba22}))
	return true
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_61df301d1f67ad73() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			switch __obf_24309b3b0ff9ba22 {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
			return __obf_24309b3b0ff9ba22
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_4ab56a99965952e7 *Iterator) ReportError(__obf_4efa14fd01d4bd95 string, __obf_489aa20db03fe286 string) {
	if __obf_4ab56a99965952e7.Error != nil {
		if __obf_4ab56a99965952e7.Error != io.EOF {
			return
		}
	}
	__obf_339a495287a7c057 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 10
	if __obf_339a495287a7c057 < 0 {
		__obf_339a495287a7c057 = 0
	}
	__obf_29e749beae25e098 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f + 10
	if __obf_29e749beae25e098 > __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
		__obf_29e749beae25e098 = __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc
	}
	__obf_09284aa7b8e3e2f7 := string(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_339a495287a7c057:__obf_29e749beae25e098])
	__obf_61c6a3876741576b := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 50
	if __obf_61c6a3876741576b < 0 {
		__obf_61c6a3876741576b = 0
	}
	__obf_c4d4f8d7b4cb9412 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f + 50
	if __obf_c4d4f8d7b4cb9412 > __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
		__obf_c4d4f8d7b4cb9412 = __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc
	}
	__obf_da10893a53474db4 := string(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_61c6a3876741576b:__obf_c4d4f8d7b4cb9412])
	__obf_4ab56a99965952e7.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_4efa14fd01d4bd95, __obf_489aa20db03fe286, __obf_4ab56a99965952e7.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_5e22d6270d27491f-__obf_339a495287a7c057, __obf_09284aa7b8e3e2f7, __obf_da10893a53474db4)
}

func (__obf_4ab56a99965952e7 *Iterator) CurrentBuffer() string {
	__obf_339a495287a7c057 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 10
	if __obf_339a495287a7c057 < 0 {
		__obf_339a495287a7c057 = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_4ab56a99965952e7.__obf_5e22d6270d27491f, string(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_339a495287a7c057:__obf_4ab56a99965952e7.__obf_5e22d6270d27491f]), string(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[0:__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc]))
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_2b6cf70e8dd32b68() (__obf_59c72400ec1a6110 byte) {
	if __obf_4ab56a99965952e7.__obf_5e22d6270d27491f == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
		if __obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_59c72400ec1a6110 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f]
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f++
			return __obf_59c72400ec1a6110
		}
		return 0
	}
	__obf_59c72400ec1a6110 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f]
	__obf_4ab56a99965952e7.__obf_5e22d6270d27491f++
	return __obf_59c72400ec1a6110
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_08a0b850abb3abbd() bool {
	if __obf_4ab56a99965952e7.__obf_7582c70e81d83895 == nil {
		if __obf_4ab56a99965952e7.Error == nil {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc
			__obf_4ab56a99965952e7.
				Error = io.EOF
		}
		return false
	}
	if __obf_4ab56a99965952e7.__obf_4541c922f6acdb7d != nil {
		__obf_4ab56a99965952e7.__obf_4541c922f6acdb7d = append(__obf_4ab56a99965952e7.__obf_4541c922f6acdb7d, __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_85ab55dc89939a18:__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc]...)
		__obf_4ab56a99965952e7.__obf_85ab55dc89939a18 = 0
	}
	for {
		__obf_e40b3963a92ca4a0, __obf_fcc907dd69879566 := __obf_4ab56a99965952e7.__obf_7582c70e81d83895.Read(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df)
		if __obf_e40b3963a92ca4a0 == 0 {
			if __obf_fcc907dd69879566 != nil {
				if __obf_4ab56a99965952e7.Error == nil {
					__obf_4ab56a99965952e7.
						Error = __obf_fcc907dd69879566
				}
				return false
			}
		} else {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = 0
			__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc = __obf_e40b3963a92ca4a0
			return true
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_ba6d0be9a7ab6ae4() {
	if __obf_4ab56a99965952e7.Error != nil {
		return
	}
	__obf_4ab56a99965952e7.

		// Read read the next JSON element as generic some.
		__obf_5e22d6270d27491f--
}

func (__obf_4ab56a99965952e7 *Iterator) Read() any {
	__obf_8afc98a56c05c14e := __obf_4ab56a99965952e7.WhatIsNext()
	switch __obf_8afc98a56c05c14e {
	case StringValue:
		return __obf_4ab56a99965952e7.ReadString()
	case NumberValue:
		if __obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_fb97917857c6c8da.UseNumber {
			return json.Number(__obf_4ab56a99965952e7.__obf_0ba08d27498a0d84())
		}
		return __obf_4ab56a99965952e7.ReadFloat64()
	case NilValue:
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_4ab56a99965952e7.ReadBool()
	case ArrayValue:
		__obf_612acb1933b33fd2 := []any{}
		__obf_4ab56a99965952e7.
			ReadArrayCB(func(__obf_4ab56a99965952e7 *Iterator) bool {
				var __obf_e4d99e200ff06def any
				__obf_4ab56a99965952e7.
					ReadVal(&__obf_e4d99e200ff06def)
				__obf_612acb1933b33fd2 = append(__obf_612acb1933b33fd2, __obf_e4d99e200ff06def)
				return true
			})
		return __obf_612acb1933b33fd2
	case ObjectValue:
		__obf_eefa92392cc2442c := map[string]any{}
		__obf_4ab56a99965952e7.
			ReadMapCB(func(Iter *Iterator, __obf_cd4d02f264b18d55 string) bool {
				var __obf_e4d99e200ff06def any
				__obf_4ab56a99965952e7.
					ReadVal(&__obf_e4d99e200ff06def)
				__obf_eefa92392cc2442c[__obf_cd4d02f264b18d55] = __obf_e4d99e200ff06def
				return true
			})
		return __obf_eefa92392cc2442c
	default:
		__obf_4ab56a99965952e7.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_8afc98a56c05c14e))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_421204de653ccba7 = 10000

func (__obf_4ab56a99965952e7 *Iterator) __obf_9cff3330bd1a56ea() (__obf_8ea174385fbfb807 bool) {
	__obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2++
	if __obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2 <= __obf_421204de653ccba7 {
		return true
	}
	__obf_4ab56a99965952e7.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_6502e8959d5f85cf() (__obf_8ea174385fbfb807 bool) {
	__obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2--
	if __obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2 >= 0 {
		return true
	}
	__obf_4ab56a99965952e7.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
