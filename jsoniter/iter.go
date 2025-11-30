package __obf_703d23ba520c3295

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

var __obf_6b77b261f313f65f []byte
var __obf_c43d6d20c66729ed []ValueType

func init() {
	__obf_6b77b261f313f65f = make([]byte, 256)
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < len(__obf_6b77b261f313f65f); __obf_b0a5d2bd48690f1d++ {
		__obf_6b77b261f313f65f[__obf_b0a5d2bd48690f1d] = 255
	}
	for __obf_b0a5d2bd48690f1d := '0'; __obf_b0a5d2bd48690f1d <= '9'; __obf_b0a5d2bd48690f1d++ {
		__obf_6b77b261f313f65f[__obf_b0a5d2bd48690f1d] = byte(__obf_b0a5d2bd48690f1d - '0')
	}
	for __obf_b0a5d2bd48690f1d := 'a'; __obf_b0a5d2bd48690f1d <= 'f'; __obf_b0a5d2bd48690f1d++ {
		__obf_6b77b261f313f65f[__obf_b0a5d2bd48690f1d] = byte((__obf_b0a5d2bd48690f1d - 'a') + 10)
	}
	for __obf_b0a5d2bd48690f1d := 'A'; __obf_b0a5d2bd48690f1d <= 'F'; __obf_b0a5d2bd48690f1d++ {
		__obf_6b77b261f313f65f[__obf_b0a5d2bd48690f1d] = byte((__obf_b0a5d2bd48690f1d - 'A') + 10)
	}
	__obf_c43d6d20c66729ed = make([]ValueType, 256)
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < len(__obf_c43d6d20c66729ed); __obf_b0a5d2bd48690f1d++ {
		__obf_c43d6d20c66729ed[__obf_b0a5d2bd48690f1d] = InvalidValue
	}
	__obf_c43d6d20c66729ed['"'] = StringValue
	__obf_c43d6d20c66729ed['-'] = NumberValue
	__obf_c43d6d20c66729ed['0'] = NumberValue
	__obf_c43d6d20c66729ed['1'] = NumberValue
	__obf_c43d6d20c66729ed['2'] = NumberValue
	__obf_c43d6d20c66729ed['3'] = NumberValue
	__obf_c43d6d20c66729ed['4'] = NumberValue
	__obf_c43d6d20c66729ed['5'] = NumberValue
	__obf_c43d6d20c66729ed['6'] = NumberValue
	__obf_c43d6d20c66729ed['7'] = NumberValue
	__obf_c43d6d20c66729ed['8'] = NumberValue
	__obf_c43d6d20c66729ed['9'] = NumberValue
	__obf_c43d6d20c66729ed['t'] = BoolValue
	__obf_c43d6d20c66729ed['f'] = BoolValue
	__obf_c43d6d20c66729ed['n'] = NilValue
	__obf_c43d6d20c66729ed['['] = ArrayValue
	__obf_c43d6d20c66729ed['{'] = ObjectValue
}

// Iterator is a io.Reader like object, with JSON specific read functions.
// Error is not returned as return value, but stored as Error member on this iterator instance.
type Iterator struct {
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
	__obf_3943589abf9d9fb6 io.Reader
	__obf_a065f8e0da5f5952 []byte
	__obf_da6aa1cfbd772c11 int
	__obf_687f71f0b8c00ef5 int
	__obf_5f5fdf79d12ab854 int
	__obf_2e0a81bee9ef3e76 int
	__obf_4b160b169035b16e []byte
	Error                  error
	Attachment             any // open for customized decoder
}

// NewIterator creates an empty Iterator instance
func NewIterator(__obf_25bd4f754a37b862 API) *Iterator {
	return &Iterator{__obf_25bd4f754a37b862: __obf_25bd4f754a37b862.(*__obf_ef74248d8ae9ba78), __obf_3943589abf9d9fb6: nil, __obf_a065f8e0da5f5952: nil, __obf_da6aa1cfbd772c11: 0, __obf_687f71f0b8c00ef5: 0, __obf_5f5fdf79d12ab854: 0}
}

// Parse creates an Iterator instance from io.Reader
func Parse(__obf_25bd4f754a37b862 API, __obf_3943589abf9d9fb6 io.Reader, __obf_7da20755ab0b0a83 int) *Iterator {
	return &Iterator{__obf_25bd4f754a37b862: __obf_25bd4f754a37b862.(*__obf_ef74248d8ae9ba78), __obf_3943589abf9d9fb6: __obf_3943589abf9d9fb6, __obf_a065f8e0da5f5952: make([]byte, __obf_7da20755ab0b0a83), __obf_da6aa1cfbd772c11: 0, __obf_687f71f0b8c00ef5: 0, __obf_5f5fdf79d12ab854: 0}
}

// ParseBytes creates an Iterator instance from byte array
func ParseBytes(__obf_25bd4f754a37b862 API, __obf_63efc9bf462b78dc []byte) *Iterator {
	return &Iterator{__obf_25bd4f754a37b862: __obf_25bd4f754a37b862.(*__obf_ef74248d8ae9ba78), __obf_3943589abf9d9fb6: nil, __obf_a065f8e0da5f5952: __obf_63efc9bf462b78dc, __obf_da6aa1cfbd772c11: 0, __obf_687f71f0b8c00ef5: len(__obf_63efc9bf462b78dc), __obf_5f5fdf79d12ab854: 0}
}

// ParseString creates an Iterator instance from string
func ParseString(__obf_25bd4f754a37b862 API, __obf_63efc9bf462b78dc string) *Iterator {
	return ParseBytes(__obf_25bd4f754a37b862, []byte(__obf_63efc9bf462b78dc))
}

// Pool returns a pool can provide more iterator with same configuration
func (__obf_47edab4c16a2d88d *Iterator) Pool() IteratorPool {
	return __obf_47edab4c16a2d88d.

		// Reset reuse iterator instance by specifying another reader
		__obf_25bd4f754a37b862
}

func (__obf_47edab4c16a2d88d *Iterator) Reset(__obf_3943589abf9d9fb6 io.Reader) *Iterator {
	__obf_47edab4c16a2d88d.__obf_3943589abf9d9fb6 = __obf_3943589abf9d9fb6
	__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = 0
	__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 = 0
	__obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854 = 0
	return __obf_47edab4c16a2d88d
}

// ResetBytes reuse iterator instance by specifying another byte array as input
func (__obf_47edab4c16a2d88d *Iterator) ResetBytes(__obf_63efc9bf462b78dc []byte) *Iterator {
	__obf_47edab4c16a2d88d.__obf_3943589abf9d9fb6 = nil
	__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952 = __obf_63efc9bf462b78dc
	__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = 0
	__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 = len(__obf_63efc9bf462b78dc)
	__obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854 = 0
	return __obf_47edab4c16a2d88d
}

// WhatIsNext gets ValueType of relatively next json element
func (__obf_47edab4c16a2d88d *Iterator) WhatIsNext() ValueType {
	__obf_835102d74dbf01c6 := __obf_c43d6d20c66729ed[__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()]
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	return __obf_835102d74dbf01c6
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_e2d7a927464ad20b() bool {
	for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		switch __obf_bd08f5161fff294a {
		case ' ', '\n', '\t', '\r':
			continue
		}
		__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
		return false
	}
	return true
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_ceba151f5486ca04() bool {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == ',' {
		return false
	}
	if __obf_bd08f5161fff294a == '}' {
		return true
	}
	__obf_47edab4c16a2d88d.
		ReportError("isObjectEnd", "object ended prematurely, unexpected char "+string([]byte{__obf_bd08f5161fff294a}))
	return true
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_6f7a9d6f8e846eec() byte {
	// a variation of skip whitespaces, returning the next non-whitespace token
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			switch __obf_bd08f5161fff294a {
			case ' ', '\n', '\t', '\r':
				continue
			}
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
			return __obf_bd08f5161fff294a
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			return 0
		}
	}
}

// ReportError record a error in iterator instance with current position.
func (__obf_47edab4c16a2d88d *Iterator) ReportError(__obf_8d50b99370eb3102 string, __obf_c6f61ce4298691bb string) {
	if __obf_47edab4c16a2d88d.Error != nil {
		if __obf_47edab4c16a2d88d.Error != io.EOF {
			return
		}
	}
	__obf_c7319b9a8ffd7d22 := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 10
	if __obf_c7319b9a8ffd7d22 < 0 {
		__obf_c7319b9a8ffd7d22 = 0
	}
	__obf_88648ac809c6d37e := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 + 10
	if __obf_88648ac809c6d37e > __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
		__obf_88648ac809c6d37e = __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5
	}
	__obf_7262f8966848235b := string(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_c7319b9a8ffd7d22:__obf_88648ac809c6d37e])
	__obf_b4f6d10575b8db9b := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 50
	if __obf_b4f6d10575b8db9b < 0 {
		__obf_b4f6d10575b8db9b = 0
	}
	__obf_28d4a6a51e994caf := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 + 50
	if __obf_28d4a6a51e994caf > __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
		__obf_28d4a6a51e994caf = __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5
	}
	__obf_43925118aa980ede := string(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b4f6d10575b8db9b:__obf_28d4a6a51e994caf])
	__obf_47edab4c16a2d88d.
		Error = fmt.Errorf("%s: %s, error found in #%v byte of ...|%s|..., bigger context ...|%s|...", __obf_8d50b99370eb3102, __obf_c6f61ce4298691bb, __obf_47edab4c16a2d88d.

		// CurrentBuffer gets current buffer as string for debugging purpose
		__obf_da6aa1cfbd772c11-__obf_c7319b9a8ffd7d22, __obf_7262f8966848235b, __obf_43925118aa980ede)
}

func (__obf_47edab4c16a2d88d *Iterator) CurrentBuffer() string {
	__obf_c7319b9a8ffd7d22 := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 10
	if __obf_c7319b9a8ffd7d22 < 0 {
		__obf_c7319b9a8ffd7d22 = 0
	}
	return fmt.Sprintf("parsing #%v byte, around ...|%s|..., whole buffer ...|%s|...", __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11, string(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_c7319b9a8ffd7d22:__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11]), string(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[0:__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5]))
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_2fa919905fa99cc3() (__obf_551cbec38242ce0c byte) {
	if __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
		if __obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_551cbec38242ce0c = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11]
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11++
			return __obf_551cbec38242ce0c
		}
		return 0
	}
	__obf_551cbec38242ce0c = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11]
	__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11++
	return __obf_551cbec38242ce0c
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_e147dd47ce22b671() bool {
	if __obf_47edab4c16a2d88d.__obf_3943589abf9d9fb6 == nil {
		if __obf_47edab4c16a2d88d.Error == nil {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5
			__obf_47edab4c16a2d88d.
				Error = io.EOF
		}
		return false
	}
	if __obf_47edab4c16a2d88d.__obf_4b160b169035b16e != nil {
		__obf_47edab4c16a2d88d.__obf_4b160b169035b16e = append(__obf_47edab4c16a2d88d.__obf_4b160b169035b16e, __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_2e0a81bee9ef3e76:__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5]...)
		__obf_47edab4c16a2d88d.__obf_2e0a81bee9ef3e76 = 0
	}
	for {
		__obf_716729e0f8e419ed, __obf_e56742d6e52f642e := __obf_47edab4c16a2d88d.__obf_3943589abf9d9fb6.Read(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952)
		if __obf_716729e0f8e419ed == 0 {
			if __obf_e56742d6e52f642e != nil {
				if __obf_47edab4c16a2d88d.Error == nil {
					__obf_47edab4c16a2d88d.
						Error = __obf_e56742d6e52f642e
				}
				return false
			}
		} else {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = 0
			__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 = __obf_716729e0f8e419ed
			return true
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_743b1a47c346c8a3() {
	if __obf_47edab4c16a2d88d.Error != nil {
		return
	}
	__obf_47edab4c16a2d88d.

		// Read read the next JSON element as generic some.
		__obf_da6aa1cfbd772c11--
}

func (__obf_47edab4c16a2d88d *Iterator) Read() any {
	__obf_835102d74dbf01c6 := __obf_47edab4c16a2d88d.WhatIsNext()
	switch __obf_835102d74dbf01c6 {
	case StringValue:
		return __obf_47edab4c16a2d88d.ReadString()
	case NumberValue:
		if __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea.UseNumber {
			return json.Number(__obf_47edab4c16a2d88d.__obf_4c009361bc8be406())
		}
		return __obf_47edab4c16a2d88d.ReadFloat64()
	case NilValue:
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('n', 'u', 'l', 'l')
		return nil
	case BoolValue:
		return __obf_47edab4c16a2d88d.ReadBool()
	case ArrayValue:
		__obf_f21bafabcc6ccf82 := []any{}
		__obf_47edab4c16a2d88d.
			ReadArrayCB(func(__obf_47edab4c16a2d88d *Iterator) bool {
				var __obf_146c3283557f8c5c any
				__obf_47edab4c16a2d88d.
					ReadVal(&__obf_146c3283557f8c5c)
				__obf_f21bafabcc6ccf82 = append(__obf_f21bafabcc6ccf82, __obf_146c3283557f8c5c)
				return true
			})
		return __obf_f21bafabcc6ccf82
	case ObjectValue:
		__obf_02ba706f4f104d71 := map[string]any{}
		__obf_47edab4c16a2d88d.
			ReadMapCB(func(Iter *Iterator, __obf_13f6440419533990 string) bool {
				var __obf_146c3283557f8c5c any
				__obf_47edab4c16a2d88d.
					ReadVal(&__obf_146c3283557f8c5c)
				__obf_02ba706f4f104d71[__obf_13f6440419533990] = __obf_146c3283557f8c5c
				return true
			})
		return __obf_02ba706f4f104d71
	default:
		__obf_47edab4c16a2d88d.
			ReportError("Read", fmt.Sprintf("unexpected value type: %v", __obf_835102d74dbf01c6))
		return nil
	}
}

// limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9
const __obf_c099d261a6c6591d = 10000

func (__obf_47edab4c16a2d88d *Iterator) __obf_0f044cac6de70712() (__obf_a92d797d86808a2e bool) {
	__obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854++
	if __obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854 <= __obf_c099d261a6c6591d {
		return true
	}
	__obf_47edab4c16a2d88d.
		ReportError("incrementDepth", "exceeded max depth")
	return false
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_bf8dbc480744e1b2() (__obf_a92d797d86808a2e bool) {
	__obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854--
	if __obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854 >= 0 {
		return true
	}
	__obf_47edab4c16a2d88d.
		ReportError("decrementDepth", "unexpected negative nesting")
	return false
}
