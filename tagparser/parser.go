package __obf_02fd394f79be5bec

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_676b9629414b33e6 *Tag) HasOption(__obf_9fb873eba167cfb6 string) bool {
	_, __obf_adc32cfdb9d8ee75 := __obf_676b9629414b33e6.Options[__obf_9fb873eba167cfb6]
	return __obf_adc32cfdb9d8ee75
}

func Parse(__obf_86c31cb7b305ff49 string) *Tag {
	__obf_563a66539a2257b2 := &__obf_01a4243139f9c42d{
		Parser: NewString(__obf_86c31cb7b305ff49),
	}
	__obf_563a66539a2257b2.__obf_05473a84061a62b6()
	return &__obf_563a66539a2257b2.Tag
}

type __obf_01a4243139f9c42d struct {
	*Parser

	Tag                    Tag
	__obf_e1fe001185082019 bool
	__obf_1295540326708cea string
}

func (__obf_563a66539a2257b2 *__obf_01a4243139f9c42d) __obf_d429c83b68a1f638(__obf_1295540326708cea, __obf_e785e853c9e11b5a string) {
	__obf_1295540326708cea = strings.TrimSpace(__obf_1295540326708cea)
	__obf_e785e853c9e11b5a = strings.TrimSpace(__obf_e785e853c9e11b5a)

	if !__obf_563a66539a2257b2.__obf_e1fe001185082019 {
		__obf_563a66539a2257b2.__obf_e1fe001185082019 = true
		if __obf_1295540326708cea == "" {
			__obf_563a66539a2257b2.
				Tag.Name = __obf_e785e853c9e11b5a
			return
		}
	}
	if __obf_563a66539a2257b2.Tag.Options == nil {
		__obf_563a66539a2257b2.
			Tag.Options = make(map[string]string)
	}
	if __obf_1295540326708cea == "" {
		__obf_563a66539a2257b2.
			Tag.Options[__obf_e785e853c9e11b5a] = ""
	} else {
		__obf_563a66539a2257b2.
			Tag.Options[__obf_1295540326708cea] = __obf_e785e853c9e11b5a
	}
}

func (__obf_563a66539a2257b2 *__obf_01a4243139f9c42d) __obf_05473a84061a62b6() {
	__obf_563a66539a2257b2.__obf_1295540326708cea = ""

	var __obf_0ab50c1778cf932b []byte
	for __obf_563a66539a2257b2.Valid() {
		__obf_c1dcf796c627d5c9 := __obf_563a66539a2257b2.Read()
		switch __obf_c1dcf796c627d5c9 {
		case ',':
			__obf_563a66539a2257b2.
				Skip(' ')
			__obf_563a66539a2257b2.__obf_d429c83b68a1f638("", string(__obf_0ab50c1778cf932b))
			__obf_563a66539a2257b2.__obf_05473a84061a62b6()
			return
		case ':':
			__obf_563a66539a2257b2.__obf_1295540326708cea = string(__obf_0ab50c1778cf932b)
			__obf_563a66539a2257b2.__obf_172ffbe513e3884f()
			return
		case '\'':
			__obf_563a66539a2257b2.__obf_6fcf0eca32384c6a()
			return
		default:
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
		}
	}

	if len(__obf_0ab50c1778cf932b) > 0 {
		__obf_563a66539a2257b2.__obf_d429c83b68a1f638("", string(__obf_0ab50c1778cf932b))
	}
}

func (__obf_563a66539a2257b2 *__obf_01a4243139f9c42d) __obf_172ffbe513e3884f() {
	const __obf_92631eb6043ce8e5 = '\''
	__obf_c1dcf796c627d5c9 := __obf_563a66539a2257b2.Peek()
	if __obf_c1dcf796c627d5c9 == __obf_92631eb6043ce8e5 {
		__obf_563a66539a2257b2.
			Skip(__obf_92631eb6043ce8e5)
		__obf_563a66539a2257b2.__obf_6fcf0eca32384c6a()
		return
	}

	var __obf_0ab50c1778cf932b []byte
	for __obf_563a66539a2257b2.Valid() {
		__obf_c1dcf796c627d5c9 = __obf_563a66539a2257b2.Read()
		switch __obf_c1dcf796c627d5c9 {
		case '\\':
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_563a66539a2257b2.Read())
		case '(':
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
			__obf_0ab50c1778cf932b = __obf_563a66539a2257b2.__obf_efc91ad0844ab148(__obf_0ab50c1778cf932b)
		case ',':
			__obf_563a66539a2257b2.
				Skip(' ')
			__obf_563a66539a2257b2.__obf_d429c83b68a1f638(__obf_563a66539a2257b2.__obf_1295540326708cea, string(__obf_0ab50c1778cf932b))
			__obf_563a66539a2257b2.__obf_05473a84061a62b6()
			return
		default:
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
		}
	}
	__obf_563a66539a2257b2.__obf_d429c83b68a1f638(__obf_563a66539a2257b2.__obf_1295540326708cea, string(__obf_0ab50c1778cf932b))
}

func (__obf_563a66539a2257b2 *__obf_01a4243139f9c42d) __obf_efc91ad0844ab148(__obf_0ab50c1778cf932b []byte) []byte {
	var __obf_1f30fb52996757bd int
__obf_99e6b31d11107105:
	for __obf_563a66539a2257b2.Valid() {
		__obf_c1dcf796c627d5c9 := __obf_563a66539a2257b2.Read()
		switch __obf_c1dcf796c627d5c9 {
		case '\\':
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_563a66539a2257b2.Read())
		case '(':
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
			__obf_1f30fb52996757bd++
		case ')':
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
			__obf_1f30fb52996757bd--
			if __obf_1f30fb52996757bd < 0 {
				break __obf_99e6b31d11107105
			}
		default:
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c1dcf796c627d5c9)
		}
	}
	return __obf_0ab50c1778cf932b
}

func (__obf_563a66539a2257b2 *__obf_01a4243139f9c42d) __obf_6fcf0eca32384c6a() {
	const __obf_92631eb6043ce8e5 = '\''
	var __obf_0ab50c1778cf932b []byte
	for __obf_563a66539a2257b2.Valid() {
		__obf_c6d82e424d0a889f, __obf_adc32cfdb9d8ee75 := __obf_563a66539a2257b2.ReadSep(__obf_92631eb6043ce8e5)
		if !__obf_adc32cfdb9d8ee75 {
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c6d82e424d0a889f...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_c6d82e424d0a889f) > 0 && __obf_c6d82e424d0a889f[len(__obf_c6d82e424d0a889f)-1] == '\\' {
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c6d82e424d0a889f[:len(__obf_c6d82e424d0a889f)-1]...)
			__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_92631eb6043ce8e5)
			continue
		}
		__obf_0ab50c1778cf932b = append(__obf_0ab50c1778cf932b, __obf_c6d82e424d0a889f...)
		break
	}
	__obf_563a66539a2257b2.__obf_d429c83b68a1f638(__obf_563a66539a2257b2.__obf_1295540326708cea, string(__obf_0ab50c1778cf932b))
	if __obf_563a66539a2257b2.Skip(',') {
		__obf_563a66539a2257b2.
			Skip(' ')
	}
	__obf_563a66539a2257b2.__obf_05473a84061a62b6()
}

type Parser struct {
	__obf_0ab50c1778cf932b []byte
	__obf_a92bc2c4a5cc778e int
}

func New(__obf_0ab50c1778cf932b []byte) *Parser {
	return &Parser{__obf_0ab50c1778cf932b: __obf_0ab50c1778cf932b}
}

func NewString(__obf_86c31cb7b305ff49 string) *Parser {
	return New(StringToBytes(__obf_86c31cb7b305ff49))
}

func (__obf_563a66539a2257b2 *Parser) Bytes() []byte {
	return __obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e:]
}

func (__obf_563a66539a2257b2 *Parser) Valid() bool {
	return __obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e < len(__obf_563a66539a2257b2.__obf_0ab50c1778cf932b)
}

func (__obf_563a66539a2257b2 *Parser) Read() byte {
	if __obf_563a66539a2257b2.Valid() {
		__obf_c1dcf796c627d5c9 := __obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e]
		__obf_563a66539a2257b2.
			Advance()
		return __obf_c1dcf796c627d5c9
	}
	return 0
}

func (__obf_563a66539a2257b2 *Parser) Peek() byte {
	if __obf_563a66539a2257b2.Valid() {
		return __obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e]
	}
	return 0
}

func (__obf_563a66539a2257b2 *Parser) Advance() {
	__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e++
}

func (__obf_563a66539a2257b2 *Parser) Skip(__obf_80b524cd896a4be2 byte) bool {
	if __obf_563a66539a2257b2.Peek() == __obf_80b524cd896a4be2 {
		__obf_563a66539a2257b2.
			Advance()
		return true
	}
	return false
}

func (__obf_563a66539a2257b2 *Parser) SkipBytes(__obf_80b524cd896a4be2 []byte) bool {
	if len(__obf_80b524cd896a4be2) > len(__obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e:]) {
		return false
	}
	if !bytes.Equal(__obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e:__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e+len(__obf_80b524cd896a4be2)], __obf_80b524cd896a4be2) {
		return false
	}
	__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e += len(__obf_80b524cd896a4be2)
	return true
}

func (__obf_563a66539a2257b2 *Parser) ReadSep(__obf_a5902335f4839e38 byte) ([]byte, bool) {
	__obf_c7a464473abcf8e9 := bytes.IndexByte(__obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e:], __obf_a5902335f4839e38)
	if __obf_c7a464473abcf8e9 == -1 {
		__obf_0ab50c1778cf932b := __obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e:]
		__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e = len(__obf_563a66539a2257b2.__obf_0ab50c1778cf932b)
		return __obf_0ab50c1778cf932b, false
	}
	__obf_0ab50c1778cf932b := __obf_563a66539a2257b2.__obf_0ab50c1778cf932b[__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e : __obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e+__obf_c7a464473abcf8e9]
	__obf_563a66539a2257b2.__obf_a92bc2c4a5cc778e += __obf_c7a464473abcf8e9 + 1
	return __obf_0ab50c1778cf932b, true
}

func BytesToString(__obf_0ab50c1778cf932b []byte) string {
	return string(__obf_0ab50c1778cf932b)
}

func StringToBytes(__obf_86c31cb7b305ff49 string) []byte {
	return []byte(__obf_86c31cb7b305ff49)
}
