package __obf_3156656031d4185a

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_e0a41f035a5f6dd1 *Tag) HasOption(__obf_5d6ee487428fd6c6 string) bool {
	_, __obf_c22dac9c047f42f1 := __obf_e0a41f035a5f6dd1.Options[__obf_5d6ee487428fd6c6]
	return __obf_c22dac9c047f42f1
}

func Parse(__obf_0b394aa3b8f73532 string) *Tag {
	__obf_74a4c88ff661b25d := &__obf_15cc2153e1aab351{
		Parser: NewString(__obf_0b394aa3b8f73532),
	}
	__obf_74a4c88ff661b25d.__obf_ef85a6cb1ef99cb1()
	return &__obf_74a4c88ff661b25d.Tag
}

type __obf_15cc2153e1aab351 struct {
	*Parser

	Tag                    Tag
	__obf_bee576acc436c0bd bool
	__obf_27f27e9da35ead8a string
}

func (__obf_74a4c88ff661b25d *__obf_15cc2153e1aab351) __obf_b3719c54118cc5fd(__obf_27f27e9da35ead8a, __obf_ea5fc74bb9c45cc1 string) {
	__obf_27f27e9da35ead8a = strings.TrimSpace(__obf_27f27e9da35ead8a)
	__obf_ea5fc74bb9c45cc1 = strings.TrimSpace(__obf_ea5fc74bb9c45cc1)

	if !__obf_74a4c88ff661b25d.__obf_bee576acc436c0bd {
		__obf_74a4c88ff661b25d.__obf_bee576acc436c0bd = true
		if __obf_27f27e9da35ead8a == "" {
			__obf_74a4c88ff661b25d.
				Tag.Name = __obf_ea5fc74bb9c45cc1
			return
		}
	}
	if __obf_74a4c88ff661b25d.Tag.Options == nil {
		__obf_74a4c88ff661b25d.
			Tag.Options = make(map[string]string)
	}
	if __obf_27f27e9da35ead8a == "" {
		__obf_74a4c88ff661b25d.
			Tag.Options[__obf_ea5fc74bb9c45cc1] = ""
	} else {
		__obf_74a4c88ff661b25d.
			Tag.Options[__obf_27f27e9da35ead8a] = __obf_ea5fc74bb9c45cc1
	}
}

func (__obf_74a4c88ff661b25d *__obf_15cc2153e1aab351) __obf_ef85a6cb1ef99cb1() {
	__obf_74a4c88ff661b25d.__obf_27f27e9da35ead8a = ""

	var __obf_28424462e68cc1f9 []byte
	for __obf_74a4c88ff661b25d.Valid() {
		__obf_56da3659fa6540f3 := __obf_74a4c88ff661b25d.Read()
		switch __obf_56da3659fa6540f3 {
		case ',':
			__obf_74a4c88ff661b25d.
				Skip(' ')
			__obf_74a4c88ff661b25d.__obf_b3719c54118cc5fd("", string(__obf_28424462e68cc1f9))
			__obf_74a4c88ff661b25d.__obf_ef85a6cb1ef99cb1()
			return
		case ':':
			__obf_74a4c88ff661b25d.__obf_27f27e9da35ead8a = string(__obf_28424462e68cc1f9)
			__obf_74a4c88ff661b25d.__obf_45083aafe1709562()
			return
		case '\'':
			__obf_74a4c88ff661b25d.__obf_61ca79e1bf99fa01()
			return
		default:
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
		}
	}

	if len(__obf_28424462e68cc1f9) > 0 {
		__obf_74a4c88ff661b25d.__obf_b3719c54118cc5fd("", string(__obf_28424462e68cc1f9))
	}
}

func (__obf_74a4c88ff661b25d *__obf_15cc2153e1aab351) __obf_45083aafe1709562() {
	const __obf_18568b7dc112b337 = '\''
	__obf_56da3659fa6540f3 := __obf_74a4c88ff661b25d.Peek()
	if __obf_56da3659fa6540f3 == __obf_18568b7dc112b337 {
		__obf_74a4c88ff661b25d.
			Skip(__obf_18568b7dc112b337)
		__obf_74a4c88ff661b25d.__obf_61ca79e1bf99fa01()
		return
	}

	var __obf_28424462e68cc1f9 []byte
	for __obf_74a4c88ff661b25d.Valid() {
		__obf_56da3659fa6540f3 = __obf_74a4c88ff661b25d.Read()
		switch __obf_56da3659fa6540f3 {
		case '\\':
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_74a4c88ff661b25d.Read())
		case '(':
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
			__obf_28424462e68cc1f9 = __obf_74a4c88ff661b25d.__obf_e04476290c50aeba(__obf_28424462e68cc1f9)
		case ',':
			__obf_74a4c88ff661b25d.
				Skip(' ')
			__obf_74a4c88ff661b25d.__obf_b3719c54118cc5fd(__obf_74a4c88ff661b25d.__obf_27f27e9da35ead8a, string(__obf_28424462e68cc1f9))
			__obf_74a4c88ff661b25d.__obf_ef85a6cb1ef99cb1()
			return
		default:
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
		}
	}
	__obf_74a4c88ff661b25d.__obf_b3719c54118cc5fd(__obf_74a4c88ff661b25d.__obf_27f27e9da35ead8a, string(__obf_28424462e68cc1f9))
}

func (__obf_74a4c88ff661b25d *__obf_15cc2153e1aab351) __obf_e04476290c50aeba(__obf_28424462e68cc1f9 []byte) []byte {
	var __obf_33822ebf0d5c24bd int
__obf_7255f23487023773:
	for __obf_74a4c88ff661b25d.Valid() {
		__obf_56da3659fa6540f3 := __obf_74a4c88ff661b25d.Read()
		switch __obf_56da3659fa6540f3 {
		case '\\':
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_74a4c88ff661b25d.Read())
		case '(':
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
			__obf_33822ebf0d5c24bd++
		case ')':
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
			__obf_33822ebf0d5c24bd--
			if __obf_33822ebf0d5c24bd < 0 {
				break __obf_7255f23487023773
			}
		default:
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_56da3659fa6540f3)
		}
	}
	return __obf_28424462e68cc1f9
}

func (__obf_74a4c88ff661b25d *__obf_15cc2153e1aab351) __obf_61ca79e1bf99fa01() {
	const __obf_18568b7dc112b337 = '\''
	var __obf_28424462e68cc1f9 []byte
	for __obf_74a4c88ff661b25d.Valid() {
		__obf_5d9119514cf9bb70, __obf_c22dac9c047f42f1 := __obf_74a4c88ff661b25d.ReadSep(__obf_18568b7dc112b337)
		if !__obf_c22dac9c047f42f1 {
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_5d9119514cf9bb70...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_5d9119514cf9bb70) > 0 && __obf_5d9119514cf9bb70[len(__obf_5d9119514cf9bb70)-1] == '\\' {
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_5d9119514cf9bb70[:len(__obf_5d9119514cf9bb70)-1]...)
			__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_18568b7dc112b337)
			continue
		}
		__obf_28424462e68cc1f9 = append(__obf_28424462e68cc1f9, __obf_5d9119514cf9bb70...)
		break
	}
	__obf_74a4c88ff661b25d.__obf_b3719c54118cc5fd(__obf_74a4c88ff661b25d.__obf_27f27e9da35ead8a, string(__obf_28424462e68cc1f9))
	if __obf_74a4c88ff661b25d.Skip(',') {
		__obf_74a4c88ff661b25d.
			Skip(' ')
	}
	__obf_74a4c88ff661b25d.__obf_ef85a6cb1ef99cb1()
}

type Parser struct {
	__obf_28424462e68cc1f9 []byte
	__obf_978fa52110c07443 int
}

func New(__obf_28424462e68cc1f9 []byte) *Parser {
	return &Parser{__obf_28424462e68cc1f9: __obf_28424462e68cc1f9}
}

func NewString(__obf_0b394aa3b8f73532 string) *Parser {
	return New(StringToBytes(__obf_0b394aa3b8f73532))
}

func (__obf_74a4c88ff661b25d *Parser) Bytes() []byte {
	return __obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443:]
}

func (__obf_74a4c88ff661b25d *Parser) Valid() bool {
	return __obf_74a4c88ff661b25d.__obf_978fa52110c07443 < len(__obf_74a4c88ff661b25d.__obf_28424462e68cc1f9)
}

func (__obf_74a4c88ff661b25d *Parser) Read() byte {
	if __obf_74a4c88ff661b25d.Valid() {
		__obf_56da3659fa6540f3 := __obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443]
		__obf_74a4c88ff661b25d.
			Advance()
		return __obf_56da3659fa6540f3
	}
	return 0
}

func (__obf_74a4c88ff661b25d *Parser) Peek() byte {
	if __obf_74a4c88ff661b25d.Valid() {
		return __obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443]
	}
	return 0
}

func (__obf_74a4c88ff661b25d *Parser) Advance() {
	__obf_74a4c88ff661b25d.__obf_978fa52110c07443++
}

func (__obf_74a4c88ff661b25d *Parser) Skip(__obf_ad6765f77320701e byte) bool {
	if __obf_74a4c88ff661b25d.Peek() == __obf_ad6765f77320701e {
		__obf_74a4c88ff661b25d.
			Advance()
		return true
	}
	return false
}

func (__obf_74a4c88ff661b25d *Parser) SkipBytes(__obf_ad6765f77320701e []byte) bool {
	if len(__obf_ad6765f77320701e) > len(__obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443:]) {
		return false
	}
	if !bytes.Equal(__obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443:__obf_74a4c88ff661b25d.__obf_978fa52110c07443+len(__obf_ad6765f77320701e)], __obf_ad6765f77320701e) {
		return false
	}
	__obf_74a4c88ff661b25d.__obf_978fa52110c07443 += len(__obf_ad6765f77320701e)
	return true
}

func (__obf_74a4c88ff661b25d *Parser) ReadSep(__obf_44445238cb0a0eb9 byte) ([]byte, bool) {
	__obf_0cdd6162d752a954 := bytes.IndexByte(__obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443:], __obf_44445238cb0a0eb9)
	if __obf_0cdd6162d752a954 == -1 {
		__obf_28424462e68cc1f9 := __obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443:]
		__obf_74a4c88ff661b25d.__obf_978fa52110c07443 = len(__obf_74a4c88ff661b25d.__obf_28424462e68cc1f9)
		return __obf_28424462e68cc1f9, false
	}
	__obf_28424462e68cc1f9 := __obf_74a4c88ff661b25d.__obf_28424462e68cc1f9[__obf_74a4c88ff661b25d.__obf_978fa52110c07443 : __obf_74a4c88ff661b25d.__obf_978fa52110c07443+__obf_0cdd6162d752a954]
	__obf_74a4c88ff661b25d.__obf_978fa52110c07443 += __obf_0cdd6162d752a954 + 1
	return __obf_28424462e68cc1f9, true
}

func BytesToString(__obf_28424462e68cc1f9 []byte) string {
	return string(__obf_28424462e68cc1f9)
}

func StringToBytes(__obf_0b394aa3b8f73532 string) []byte {
	return []byte(__obf_0b394aa3b8f73532)
}
