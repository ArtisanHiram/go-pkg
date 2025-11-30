package __obf_d314e59d3ebfecc1

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_b6c573f0abd4dd2a *Tag) HasOption(__obf_3066baa9a8a9a060 string) bool {
	_, __obf_71c903701c6b1bd8 := __obf_b6c573f0abd4dd2a.Options[__obf_3066baa9a8a9a060]
	return __obf_71c903701c6b1bd8
}

func Parse(__obf_490cf4278cae9b35 string) *Tag {
	__obf_91f289e1ef5b86ce := &__obf_79cfa003b3857669{
		Parser: NewString(__obf_490cf4278cae9b35),
	}
	__obf_91f289e1ef5b86ce.__obf_e3e02fec94a14dec()
	return &__obf_91f289e1ef5b86ce.Tag
}

type __obf_79cfa003b3857669 struct {
	*Parser

	Tag                    Tag
	__obf_1eabffdd9812c10d bool
	__obf_b4112ebf0e189519 string
}

func (__obf_91f289e1ef5b86ce *__obf_79cfa003b3857669) __obf_18a2c081351de612(__obf_b4112ebf0e189519, __obf_c2a56ca9272de076 string) {
	__obf_b4112ebf0e189519 = strings.TrimSpace(__obf_b4112ebf0e189519)
	__obf_c2a56ca9272de076 = strings.TrimSpace(__obf_c2a56ca9272de076)

	if !__obf_91f289e1ef5b86ce.__obf_1eabffdd9812c10d {
		__obf_91f289e1ef5b86ce.__obf_1eabffdd9812c10d = true
		if __obf_b4112ebf0e189519 == "" {
			__obf_91f289e1ef5b86ce.
				Tag.Name = __obf_c2a56ca9272de076
			return
		}
	}
	if __obf_91f289e1ef5b86ce.Tag.Options == nil {
		__obf_91f289e1ef5b86ce.
			Tag.Options = make(map[string]string)
	}
	if __obf_b4112ebf0e189519 == "" {
		__obf_91f289e1ef5b86ce.
			Tag.Options[__obf_c2a56ca9272de076] = ""
	} else {
		__obf_91f289e1ef5b86ce.
			Tag.Options[__obf_b4112ebf0e189519] = __obf_c2a56ca9272de076
	}
}

func (__obf_91f289e1ef5b86ce *__obf_79cfa003b3857669) __obf_e3e02fec94a14dec() {
	__obf_91f289e1ef5b86ce.__obf_b4112ebf0e189519 = ""

	var __obf_c44a57c8c8e33db7 []byte
	for __obf_91f289e1ef5b86ce.Valid() {
		__obf_8a6e9a676892e36a := __obf_91f289e1ef5b86ce.Read()
		switch __obf_8a6e9a676892e36a {
		case ',':
			__obf_91f289e1ef5b86ce.
				Skip(' ')
			__obf_91f289e1ef5b86ce.__obf_18a2c081351de612("", string(__obf_c44a57c8c8e33db7))
			__obf_91f289e1ef5b86ce.__obf_e3e02fec94a14dec()
			return
		case ':':
			__obf_91f289e1ef5b86ce.__obf_b4112ebf0e189519 = string(__obf_c44a57c8c8e33db7)
			__obf_91f289e1ef5b86ce.__obf_80a87919a2b51e9a()
			return
		case '\'':
			__obf_91f289e1ef5b86ce.__obf_9e79a8e63bd6e425()
			return
		default:
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
		}
	}

	if len(__obf_c44a57c8c8e33db7) > 0 {
		__obf_91f289e1ef5b86ce.__obf_18a2c081351de612("", string(__obf_c44a57c8c8e33db7))
	}
}

func (__obf_91f289e1ef5b86ce *__obf_79cfa003b3857669) __obf_80a87919a2b51e9a() {
	const __obf_1f77adc31dfd2927 = '\''
	__obf_8a6e9a676892e36a := __obf_91f289e1ef5b86ce.Peek()
	if __obf_8a6e9a676892e36a == __obf_1f77adc31dfd2927 {
		__obf_91f289e1ef5b86ce.
			Skip(__obf_1f77adc31dfd2927)
		__obf_91f289e1ef5b86ce.__obf_9e79a8e63bd6e425()
		return
	}

	var __obf_c44a57c8c8e33db7 []byte
	for __obf_91f289e1ef5b86ce.Valid() {
		__obf_8a6e9a676892e36a = __obf_91f289e1ef5b86ce.Read()
		switch __obf_8a6e9a676892e36a {
		case '\\':
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_91f289e1ef5b86ce.Read())
		case '(':
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
			__obf_c44a57c8c8e33db7 = __obf_91f289e1ef5b86ce.__obf_ed728079fe44bfe7(__obf_c44a57c8c8e33db7)
		case ',':
			__obf_91f289e1ef5b86ce.
				Skip(' ')
			__obf_91f289e1ef5b86ce.__obf_18a2c081351de612(__obf_91f289e1ef5b86ce.__obf_b4112ebf0e189519, string(__obf_c44a57c8c8e33db7))
			__obf_91f289e1ef5b86ce.__obf_e3e02fec94a14dec()
			return
		default:
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
		}
	}
	__obf_91f289e1ef5b86ce.__obf_18a2c081351de612(__obf_91f289e1ef5b86ce.__obf_b4112ebf0e189519, string(__obf_c44a57c8c8e33db7))
}

func (__obf_91f289e1ef5b86ce *__obf_79cfa003b3857669) __obf_ed728079fe44bfe7(__obf_c44a57c8c8e33db7 []byte) []byte {
	var __obf_f54cff3ab44e4449 int
__obf_72fac82e129bcc41:
	for __obf_91f289e1ef5b86ce.Valid() {
		__obf_8a6e9a676892e36a := __obf_91f289e1ef5b86ce.Read()
		switch __obf_8a6e9a676892e36a {
		case '\\':
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_91f289e1ef5b86ce.Read())
		case '(':
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
			__obf_f54cff3ab44e4449++
		case ')':
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
			__obf_f54cff3ab44e4449--
			if __obf_f54cff3ab44e4449 < 0 {
				break __obf_72fac82e129bcc41
			}
		default:
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_8a6e9a676892e36a)
		}
	}
	return __obf_c44a57c8c8e33db7
}

func (__obf_91f289e1ef5b86ce *__obf_79cfa003b3857669) __obf_9e79a8e63bd6e425() {
	const __obf_1f77adc31dfd2927 = '\''
	var __obf_c44a57c8c8e33db7 []byte
	for __obf_91f289e1ef5b86ce.Valid() {
		__obf_802fe445415b2775, __obf_71c903701c6b1bd8 := __obf_91f289e1ef5b86ce.ReadSep(__obf_1f77adc31dfd2927)
		if !__obf_71c903701c6b1bd8 {
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_802fe445415b2775...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_802fe445415b2775) > 0 && __obf_802fe445415b2775[len(__obf_802fe445415b2775)-1] == '\\' {
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_802fe445415b2775[:len(__obf_802fe445415b2775)-1]...)
			__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_1f77adc31dfd2927)
			continue
		}
		__obf_c44a57c8c8e33db7 = append(__obf_c44a57c8c8e33db7, __obf_802fe445415b2775...)
		break
	}
	__obf_91f289e1ef5b86ce.__obf_18a2c081351de612(__obf_91f289e1ef5b86ce.__obf_b4112ebf0e189519, string(__obf_c44a57c8c8e33db7))
	if __obf_91f289e1ef5b86ce.Skip(',') {
		__obf_91f289e1ef5b86ce.
			Skip(' ')
	}
	__obf_91f289e1ef5b86ce.__obf_e3e02fec94a14dec()
}

type Parser struct {
	__obf_c44a57c8c8e33db7 []byte
	__obf_47d1bccaa76828aa int
}

func New(__obf_c44a57c8c8e33db7 []byte) *Parser {
	return &Parser{__obf_c44a57c8c8e33db7: __obf_c44a57c8c8e33db7}
}

func NewString(__obf_490cf4278cae9b35 string) *Parser {
	return New(StringToBytes(__obf_490cf4278cae9b35))
}

func (__obf_91f289e1ef5b86ce *Parser) Bytes() []byte {
	return __obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa:]
}

func (__obf_91f289e1ef5b86ce *Parser) Valid() bool {
	return __obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa < len(__obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7)
}

func (__obf_91f289e1ef5b86ce *Parser) Read() byte {
	if __obf_91f289e1ef5b86ce.Valid() {
		__obf_8a6e9a676892e36a := __obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa]
		__obf_91f289e1ef5b86ce.
			Advance()
		return __obf_8a6e9a676892e36a
	}
	return 0
}

func (__obf_91f289e1ef5b86ce *Parser) Peek() byte {
	if __obf_91f289e1ef5b86ce.Valid() {
		return __obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa]
	}
	return 0
}

func (__obf_91f289e1ef5b86ce *Parser) Advance() {
	__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa++
}

func (__obf_91f289e1ef5b86ce *Parser) Skip(__obf_8a58dd9416cfc0d0 byte) bool {
	if __obf_91f289e1ef5b86ce.Peek() == __obf_8a58dd9416cfc0d0 {
		__obf_91f289e1ef5b86ce.
			Advance()
		return true
	}
	return false
}

func (__obf_91f289e1ef5b86ce *Parser) SkipBytes(__obf_8a58dd9416cfc0d0 []byte) bool {
	if len(__obf_8a58dd9416cfc0d0) > len(__obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa:]) {
		return false
	}
	if !bytes.Equal(__obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa:__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa+len(__obf_8a58dd9416cfc0d0)], __obf_8a58dd9416cfc0d0) {
		return false
	}
	__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa += len(__obf_8a58dd9416cfc0d0)
	return true
}

func (__obf_91f289e1ef5b86ce *Parser) ReadSep(__obf_b96ca712ae7460cb byte) ([]byte, bool) {
	__obf_c05f81861a833c64 := bytes.IndexByte(__obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa:], __obf_b96ca712ae7460cb)
	if __obf_c05f81861a833c64 == -1 {
		__obf_c44a57c8c8e33db7 := __obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa:]
		__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa = len(__obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7)
		return __obf_c44a57c8c8e33db7, false
	}
	__obf_c44a57c8c8e33db7 := __obf_91f289e1ef5b86ce.__obf_c44a57c8c8e33db7[__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa : __obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa+__obf_c05f81861a833c64]
	__obf_91f289e1ef5b86ce.__obf_47d1bccaa76828aa += __obf_c05f81861a833c64 + 1
	return __obf_c44a57c8c8e33db7, true
}

func BytesToString(__obf_c44a57c8c8e33db7 []byte) string {
	return string(__obf_c44a57c8c8e33db7)
}

func StringToBytes(__obf_490cf4278cae9b35 string) []byte {
	return []byte(__obf_490cf4278cae9b35)
}
