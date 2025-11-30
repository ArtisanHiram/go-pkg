package __obf_5bfb23d39a7e9f2e

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_b4e54d191964bf78 *Tag) HasOption(__obf_581c6e87e549b573 string) bool {
	_, __obf_15bf1769b56e077c := __obf_b4e54d191964bf78.Options[__obf_581c6e87e549b573]
	return __obf_15bf1769b56e077c
}

func Parse(__obf_86a9f11c82d9c3af string) *Tag {
	__obf_79bcefc0ff1d6c9c := &__obf_604e8a4b59108495{
		Parser: NewString(__obf_86a9f11c82d9c3af),
	}
	__obf_79bcefc0ff1d6c9c.__obf_c30a3e736e3358d3()
	return &__obf_79bcefc0ff1d6c9c.Tag
}

type __obf_604e8a4b59108495 struct {
	*Parser

	Tag                    Tag
	__obf_2d64713e056485c1 bool
	__obf_e5e90d376efb0736 string
}

func (__obf_79bcefc0ff1d6c9c *__obf_604e8a4b59108495) __obf_2561265d39605d14(__obf_e5e90d376efb0736, __obf_09675a318138dfc9 string) {
	__obf_e5e90d376efb0736 = strings.TrimSpace(__obf_e5e90d376efb0736)
	__obf_09675a318138dfc9 = strings.TrimSpace(__obf_09675a318138dfc9)

	if !__obf_79bcefc0ff1d6c9c.__obf_2d64713e056485c1 {
		__obf_79bcefc0ff1d6c9c.__obf_2d64713e056485c1 = true
		if __obf_e5e90d376efb0736 == "" {
			__obf_79bcefc0ff1d6c9c.
				Tag.Name = __obf_09675a318138dfc9
			return
		}
	}
	if __obf_79bcefc0ff1d6c9c.Tag.Options == nil {
		__obf_79bcefc0ff1d6c9c.
			Tag.Options = make(map[string]string)
	}
	if __obf_e5e90d376efb0736 == "" {
		__obf_79bcefc0ff1d6c9c.
			Tag.Options[__obf_09675a318138dfc9] = ""
	} else {
		__obf_79bcefc0ff1d6c9c.
			Tag.Options[__obf_e5e90d376efb0736] = __obf_09675a318138dfc9
	}
}

func (__obf_79bcefc0ff1d6c9c *__obf_604e8a4b59108495) __obf_c30a3e736e3358d3() {
	__obf_79bcefc0ff1d6c9c.__obf_e5e90d376efb0736 = ""

	var __obf_e5cf7afef54f55bd []byte
	for __obf_79bcefc0ff1d6c9c.Valid() {
		__obf_518dfda1112c9e7c := __obf_79bcefc0ff1d6c9c.Read()
		switch __obf_518dfda1112c9e7c {
		case ',':
			__obf_79bcefc0ff1d6c9c.
				Skip(' ')
			__obf_79bcefc0ff1d6c9c.__obf_2561265d39605d14("", string(__obf_e5cf7afef54f55bd))
			__obf_79bcefc0ff1d6c9c.__obf_c30a3e736e3358d3()
			return
		case ':':
			__obf_79bcefc0ff1d6c9c.__obf_e5e90d376efb0736 = string(__obf_e5cf7afef54f55bd)
			__obf_79bcefc0ff1d6c9c.__obf_40fa8c4dd601b019()
			return
		case '\'':
			__obf_79bcefc0ff1d6c9c.__obf_f1c2eb73e1454add()
			return
		default:
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
		}
	}

	if len(__obf_e5cf7afef54f55bd) > 0 {
		__obf_79bcefc0ff1d6c9c.__obf_2561265d39605d14("", string(__obf_e5cf7afef54f55bd))
	}
}

func (__obf_79bcefc0ff1d6c9c *__obf_604e8a4b59108495) __obf_40fa8c4dd601b019() {
	const __obf_d7588809f655c265 = '\''
	__obf_518dfda1112c9e7c := __obf_79bcefc0ff1d6c9c.Peek()
	if __obf_518dfda1112c9e7c == __obf_d7588809f655c265 {
		__obf_79bcefc0ff1d6c9c.
			Skip(__obf_d7588809f655c265)
		__obf_79bcefc0ff1d6c9c.__obf_f1c2eb73e1454add()
		return
	}

	var __obf_e5cf7afef54f55bd []byte
	for __obf_79bcefc0ff1d6c9c.Valid() {
		__obf_518dfda1112c9e7c = __obf_79bcefc0ff1d6c9c.Read()
		switch __obf_518dfda1112c9e7c {
		case '\\':
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_79bcefc0ff1d6c9c.Read())
		case '(':
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
			__obf_e5cf7afef54f55bd = __obf_79bcefc0ff1d6c9c.__obf_708e405427ad665b(__obf_e5cf7afef54f55bd)
		case ',':
			__obf_79bcefc0ff1d6c9c.
				Skip(' ')
			__obf_79bcefc0ff1d6c9c.__obf_2561265d39605d14(__obf_79bcefc0ff1d6c9c.__obf_e5e90d376efb0736, string(__obf_e5cf7afef54f55bd))
			__obf_79bcefc0ff1d6c9c.__obf_c30a3e736e3358d3()
			return
		default:
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
		}
	}
	__obf_79bcefc0ff1d6c9c.__obf_2561265d39605d14(__obf_79bcefc0ff1d6c9c.__obf_e5e90d376efb0736, string(__obf_e5cf7afef54f55bd))
}

func (__obf_79bcefc0ff1d6c9c *__obf_604e8a4b59108495) __obf_708e405427ad665b(__obf_e5cf7afef54f55bd []byte) []byte {
	var __obf_2c8e54475fbd489f int
__obf_4104006b5c9e81da:
	for __obf_79bcefc0ff1d6c9c.Valid() {
		__obf_518dfda1112c9e7c := __obf_79bcefc0ff1d6c9c.Read()
		switch __obf_518dfda1112c9e7c {
		case '\\':
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_79bcefc0ff1d6c9c.Read())
		case '(':
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
			__obf_2c8e54475fbd489f++
		case ')':
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
			__obf_2c8e54475fbd489f--
			if __obf_2c8e54475fbd489f < 0 {
				break __obf_4104006b5c9e81da
			}
		default:
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_518dfda1112c9e7c)
		}
	}
	return __obf_e5cf7afef54f55bd
}

func (__obf_79bcefc0ff1d6c9c *__obf_604e8a4b59108495) __obf_f1c2eb73e1454add() {
	const __obf_d7588809f655c265 = '\''
	var __obf_e5cf7afef54f55bd []byte
	for __obf_79bcefc0ff1d6c9c.Valid() {
		__obf_fee49a658d07254e, __obf_15bf1769b56e077c := __obf_79bcefc0ff1d6c9c.ReadSep(__obf_d7588809f655c265)
		if !__obf_15bf1769b56e077c {
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_fee49a658d07254e...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_fee49a658d07254e) > 0 && __obf_fee49a658d07254e[len(__obf_fee49a658d07254e)-1] == '\\' {
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_fee49a658d07254e[:len(__obf_fee49a658d07254e)-1]...)
			__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_d7588809f655c265)
			continue
		}
		__obf_e5cf7afef54f55bd = append(__obf_e5cf7afef54f55bd, __obf_fee49a658d07254e...)
		break
	}
	__obf_79bcefc0ff1d6c9c.__obf_2561265d39605d14(__obf_79bcefc0ff1d6c9c.__obf_e5e90d376efb0736, string(__obf_e5cf7afef54f55bd))
	if __obf_79bcefc0ff1d6c9c.Skip(',') {
		__obf_79bcefc0ff1d6c9c.
			Skip(' ')
	}
	__obf_79bcefc0ff1d6c9c.__obf_c30a3e736e3358d3()
}

type Parser struct {
	__obf_e5cf7afef54f55bd []byte
	__obf_0a5a8adee60ff412 int
}

func New(__obf_e5cf7afef54f55bd []byte) *Parser {
	return &Parser{__obf_e5cf7afef54f55bd: __obf_e5cf7afef54f55bd}
}

func NewString(__obf_86a9f11c82d9c3af string) *Parser {
	return New(StringToBytes(__obf_86a9f11c82d9c3af))
}

func (__obf_79bcefc0ff1d6c9c *Parser) Bytes() []byte {
	return __obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412:]
}

func (__obf_79bcefc0ff1d6c9c *Parser) Valid() bool {
	return __obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412 < len(__obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd)
}

func (__obf_79bcefc0ff1d6c9c *Parser) Read() byte {
	if __obf_79bcefc0ff1d6c9c.Valid() {
		__obf_518dfda1112c9e7c := __obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412]
		__obf_79bcefc0ff1d6c9c.
			Advance()
		return __obf_518dfda1112c9e7c
	}
	return 0
}

func (__obf_79bcefc0ff1d6c9c *Parser) Peek() byte {
	if __obf_79bcefc0ff1d6c9c.Valid() {
		return __obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412]
	}
	return 0
}

func (__obf_79bcefc0ff1d6c9c *Parser) Advance() {
	__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412++
}

func (__obf_79bcefc0ff1d6c9c *Parser) Skip(__obf_c4c5fef748ad1e3b byte) bool {
	if __obf_79bcefc0ff1d6c9c.Peek() == __obf_c4c5fef748ad1e3b {
		__obf_79bcefc0ff1d6c9c.
			Advance()
		return true
	}
	return false
}

func (__obf_79bcefc0ff1d6c9c *Parser) SkipBytes(__obf_c4c5fef748ad1e3b []byte) bool {
	if len(__obf_c4c5fef748ad1e3b) > len(__obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412:]) {
		return false
	}
	if !bytes.Equal(__obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412:__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412+len(__obf_c4c5fef748ad1e3b)], __obf_c4c5fef748ad1e3b) {
		return false
	}
	__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412 += len(__obf_c4c5fef748ad1e3b)
	return true
}

func (__obf_79bcefc0ff1d6c9c *Parser) ReadSep(__obf_646ba49041acecd3 byte) ([]byte, bool) {
	__obf_81f2ec7939c5ed76 := bytes.IndexByte(__obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412:], __obf_646ba49041acecd3)
	if __obf_81f2ec7939c5ed76 == -1 {
		__obf_e5cf7afef54f55bd := __obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412:]
		__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412 = len(__obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd)
		return __obf_e5cf7afef54f55bd, false
	}
	__obf_e5cf7afef54f55bd := __obf_79bcefc0ff1d6c9c.__obf_e5cf7afef54f55bd[__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412 : __obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412+__obf_81f2ec7939c5ed76]
	__obf_79bcefc0ff1d6c9c.__obf_0a5a8adee60ff412 += __obf_81f2ec7939c5ed76 + 1
	return __obf_e5cf7afef54f55bd, true
}

func BytesToString(__obf_e5cf7afef54f55bd []byte) string {
	return string(__obf_e5cf7afef54f55bd)
}

func StringToBytes(__obf_86a9f11c82d9c3af string) []byte {
	return []byte(__obf_86a9f11c82d9c3af)
}
