package __obf_c3cd04a15c56f32f

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

var __obf_0b1a1228a30a4c87 []byte
var __obf_be756888864bba7d []ValueType

func init() {
	__obf_0b1a1228a30a4c87 = make([]byte, 256)
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < len(__obf_0b1a1228a30a4c87); __obf_28d099df85f083a8++ {
		__obf_0b1a1228a30a4c87[__obf_28d099df85f083a8] = 255
	}
	for __obf_28d099df85f083a8 := '0'; __obf_28d099df85f083a8 <= '9'; __obf_28d099df85f083a8++ {
		__obf_0b1a1228a30a4c87[__obf_28d099df85f083a8] = byte(__obf_28d099df85f083a8 - '0')
	}
	for __obf_28d099df85f083a8 := 'a'; __obf_28d099df85f083a8 <= 'f'; __obf_28d099df85f083a8++ {
		__obf_0b1a1228a30a4c87[__obf_28d099df85f083a8] = byte((__obf_28d099df85f083a8 - 'a') + 10)
	}
	for __obf_28d099df85f083a8 := 'A'; __obf_28d099df85f083a8 <= 'F'; __obf_28d099df85f083a8++ {
		__obf_0b1a1228a30a4c87[__obf_28d099df85f083a8] = byte((__obf_28d099df85f083a8 - 'A') + 10)
	}
	__obf_be756888864bba7d = make([]ValueType, 256)
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < len(__obf_be756888864bba7d); __obf_28d099df85f083a8++ {
		__obf_be756888864bba7d[__obf_28d099df85f083a8] = InvalidValue
	}
	__obf_be756888864bba7d['"'] = StringValue
	__obf_be756888864bba7d['-'] = NumberValue
	__obf_be756888864bba7d['0'] = NumberValue
	__obf_be756888864bba7d['1'] = NumberValue
	__obf_be756888864bba7d['2'] = NumberValue
	__obf_be756888864bba7d['3'] = NumberValue
	__obf_be756888864bba7d['4'] = NumberValue
	__obf_be756888864bba7d['5'] = NumberValue
	__obf_be756888864bba7d['6'] = NumberValue
	__obf_be756888864bba7d['7'] = NumberValue
	__obf_be756888864bba7d['8'] = NumberValue
	__obf_be756888864bba7d['9'] = NumberValue
	__obf_be756888864bba7d['t'] = BoolValue
	__obf_be756888864bba7d['f'] = BoolValue
	__obf_be756888864bba7d['n'] = NilValue
	__obf_be756888864bba7d['['] = ArrayValue
	__obf_be756888864bba7d['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
	__obf_a539369e10727e42 io.Reader
	__obf_ace979f6622823f3 []byte
	__obf_edd3c3885447229b int
	__obf_3a475d1c1ae2cd57 int
	__obf_2da3205c67480a39 int
	__obf_df16664751f477bb int
	__obf_4db99b2aa14dd4a3 []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_0c8065c219ccf0ab API) *Iterator {
	return &Iterator{__obf_0c8065c219ccf0ab: __obf_0c8065c219ccf0ab.(*__obf_3a25fb4c9a8342bb), __obf_a539369e10727e42: nil, __obf_ace979f6622823f3: nil, __obf_edd3c3885447229b: 0, __obf_3a475d1c1ae2cd57: 0, __obf_2da3205c67480a39: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_0c8065c219ccf0ab API, __obf_a539369e10727e42 io.Reader, __obf_5ab0791f20d1d4db int) *Iterator {
	return &Iterator{__obf_0c8065c219ccf0ab: __obf_0c8065c219ccf0ab.(*__obf_3a25fb4c9a8342bb), __obf_a539369e10727e42: __obf_a539369e10727e42, __obf_ace979f6622823f3: make([]byte, __obf_5ab0791f20d1d4db), __obf_edd3c3885447229b: 0, __obf_3a475d1c1ae2cd57: 0, __obf_2da3205c67480a39: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_0c8065c219ccf0ab API, __obf_9fb90325dbfbe010 []byte) *Iterator {
	return &Iterator{__obf_0c8065c219ccf0ab: __obf_0c8065c219ccf0ab.(*__obf_3a25fb4c9a8342bb), __obf_a539369e10727e42: nil, __obf_ace979f6622823f3: __obf_9fb90325dbfbe010, __obf_edd3c3885447229b: 0, __obf_3a475d1c1ae2cd57: len(__obf_9fb90325dbfbe010), __obf_2da3205c67480a39: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_0c8065c219ccf0ab API, __obf_9fb90325dbfbe010 string) *Iterator {
	return ParseBytes(__obf_0c8065c219ccf0ab, []byte(__obf_9fb90325dbfbe010))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_edd9fe48d39445e4 *Iterator) Pool() IteratorPool {
	return __obf_edd9fe48d39445e4.

		// Reset reuse iterator instance by specifying another reader
		__obf_0c8065c219ccf0ab
}

func (__obf_edd9fe48d39445e4 *Iterator) Reset(__obf_a539369e10727e42 io.Reader) *Iterator {
	__obf_edd9fe48d39445e4.__obf_a539369e10727e42 = __obf_a539369e10727e42
	__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = 0
	__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 = 0
	__obf_edd9fe48d39445e4.__obf_2da3205c67480a39 = 0
	return __obf_edd9fe48d39445e4
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_edd9fe48d39445e4 *Iterator) ResetBytes(__obf_9fb90325dbfbe010 []byte) *Iterator {
	__obf_edd9fe48d39445e4.__obf_a539369e10727e42 = nil
	__obf_edd9fe48d39445e4.__obf_ace979f6622823f3 = __obf_9fb90325dbfbe010
	__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = 0
	__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 = len(__obf_9fb90325dbfbe010)
	__obf_edd9fe48d39445e4.__obf_2da3205c67480a39 = 0
	return __obf_edd9fe48d39445e4
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_edd9fe48d39445e4 *Iterator) WhatIsNext() ValueType {
	__obf_64fa966d17d74b1c := __obf_be756888864bba7d[__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()]
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	return __obf_64fa966d17d74b1c
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_c0473a44bf419ac6() bool {
	for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		switch __obf_0c1bc1e511a43120 {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
		return false
	}
	return true
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_9edfbea94670bd7a() bool {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == ',' {
		return false
	}
	if __obf_0c1bc1e511a43120 == '}' {
		return true
	}
	__obf_edd9fe48d39445e4.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_0c1bc1e511a43120}))
	return true
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_a26eeeac1d6f5a11() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			switch __obf_0c1bc1e511a43120 {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
			return __obf_0c1bc1e511a43120
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_edd9fe48d39445e4 *Iterator) ReportError(__obf_04503d9ae6dbd364 string, __obf_864e3235428d31a6 string) {
	if __obf_edd9fe48d39445e4.Error != nil {
		if __obf_edd9fe48d39445e4.Error != io.EOF {
			return
		}
	}
	__obf_57d500c3170a0aef := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 10
	if __obf_57d500c3170a0aef < 0 {
		__obf_57d500c3170a0aef = 0
	}
	__obf_e3d91e323ba3c024 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b + 10
	if __obf_e3d91e323ba3c024 > __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
		__obf_e3d91e323ba3c024 = __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57
	}
	__obf_bd1c8bbb439194e6 := string(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57d500c3170a0aef:__obf_e3d91e323ba3c024])
	__obf_b3c0968b952e49bc := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 50
	if __obf_b3c0968b952e49bc < 0 {
		__obf_b3c0968b952e49bc = 0
	}
	__obf_2ea1f6762df0a644 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b + 50
	if __obf_2ea1f6762df0a644 > __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
		__obf_2ea1f6762df0a644 = __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57
	}
	__obf_904818d6d1cf3597 := string(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_b3c0968b952e49bc:__obf_2ea1f6762df0a644])
	__obf_edd9fe48d39445e4.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_04503d9ae6dbd364, __obf_864e3235428d31a6, __obf_edd9fe48d39445e4.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_edd3c3885447229b-__obf_57d500c3170a0aef, __obf_bd1c8bbb439194e6, __obf_904818d6d1cf3597)
}

func (__obf_edd9fe48d39445e4 *Iterator) CurrentBuffer() string {
	__obf_57d500c3170a0aef := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 10
	if __obf_57d500c3170a0aef < 0 {
		__obf_57d500c3170a0aef = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_edd9fe48d39445e4.__obf_edd3c3885447229b, string(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57d500c3170a0aef:__obf_edd9fe48d39445e4.__obf_edd3c3885447229b]), string(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[0:__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57]))
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_70c673c91de4f883() (__obf_316af79472661247 byte) {
	if __obf_edd9fe48d39445e4.__obf_edd3c3885447229b == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
		if __obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_316af79472661247 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b]
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b++
			return __obf_316af79472661247
		}
		return 0
	}
	__obf_316af79472661247 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b]
	__obf_edd9fe48d39445e4.__obf_edd3c3885447229b++
	return __obf_316af79472661247
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_35e704c169a88a74() bool {
	if __obf_edd9fe48d39445e4.__obf_a539369e10727e42 == nil {
		if __obf_edd9fe48d39445e4.Error == nil {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57
			__obf_edd9fe48d39445e4.
				Error = io.EOF
		}
		return false
	}
	if __obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 != nil {
		__obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 = append(__obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3, __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_df16664751f477bb:__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57]...)
		__obf_edd9fe48d39445e4.__obf_df16664751f477bb = 0
	}
	for {
		__obf_fd4464bb6b13cabd, __obf_5966ecc5edb9b17e := __obf_edd9fe48d39445e4.__obf_a539369e10727e42.Read(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3)
		if __obf_fd4464bb6b13cabd == 0 {
			if __obf_5966ecc5edb9b17e != nil {
				if __obf_edd9fe48d39445e4.Error == nil {
					__obf_edd9fe48d39445e4.
						Error = __obf_5966ecc5edb9b17e
				}
				return false
			}
		} else {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = 0
			__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 = __obf_fd4464bb6b13cabd
			return true
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_a0622aded2d86ded() {
	if __obf_edd9fe48d39445e4.Error != nil {
		return
	}
	__obf_edd9fe48d39445e4.

		// Read read the next JSON element as generic some.
		__obf_edd3c3885447229b--
}

func (__obf_edd9fe48d39445e4 *Iterator) Read() any {
	__obf_64fa966d17d74b1c := __obf_edd9fe48d39445e4.WhatIsNext()
	switch __obf_64fa966d17d74b1c {
	case StringValue:
		return __obf_edd9fe48d39445e4.ReadString()
	case NumberValue:
		if __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab.UseNumber {
			return json.Number(__obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818())
		}
		return __obf_edd9fe48d39445e4.ReadFloat64()
	case NilValue:
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_edd9fe48d39445e4.ReadBool()
	case ArrayValue:
		__obf_003fd2c5d2c6b965 := []any{}
		__obf_edd9fe48d39445e4.
			ReadArrayCB(func(__obf_edd9fe48d39445e4 *Iterator) bool {
				var __obf_991e715ce3758c78 any
				__obf_edd9fe48d39445e4.
					ReadVal(&__obf_991e715ce3758c78)
				__obf_003fd2c5d2c6b965 = append(__obf_003fd2c5d2c6b965, __obf_991e715ce3758c78)
				return true
			})
		return __obf_003fd2c5d2c6b965
	case ObjectValue:
		__obf_50acae8c0a871ba1 := map[string]any{}
		__obf_edd9fe48d39445e4.
			ReadMapCB(func(Iter *Iterator, __obf_f992c91036745ca0 string) bool {
				var __obf_991e715ce3758c78 any
				__obf_edd9fe48d39445e4.
					ReadVal(&__obf_991e715ce3758c78)
				__obf_50acae8c0a871ba1[__obf_f992c91036745ca0] = __obf_991e715ce3758c78
				return true
			})
		return __obf_50acae8c0a871ba1
	default:
		__obf_edd9fe48d39445e4.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_64fa966d17d74b1c))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_e2d5522ecc718d15 = 10000

func (__obf_edd9fe48d39445e4 *Iterator) __obf_68afb6508328afc9() (__obf_e0847f682ce51fc4 bool) {
	__obf_edd9fe48d39445e4.__obf_2da3205c67480a39++
	if __obf_edd9fe48d39445e4.__obf_2da3205c67480a39 <= __obf_e2d5522ecc718d15 {
		return true
	}
	__obf_edd9fe48d39445e4.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_0727968e0120c8a6() (__obf_e0847f682ce51fc4 bool) {
	__obf_edd9fe48d39445e4.__obf_2da3205c67480a39--
	if __obf_edd9fe48d39445e4.__obf_2da3205c67480a39 >= 0 {
		return true
	}
	__obf_edd9fe48d39445e4.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
