package __obf_7a3e76871d8e36cd

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_1561272f80ffbe15 *Tag) HasOption(__obf_e970921681cdb748 string) bool {
	_, __obf_6f62606b6640bca1 := __obf_1561272f80ffbe15.Options[__obf_e970921681cdb748]
	return __obf_6f62606b6640bca1
}

func Parse(__obf_cbeee07d735d7228 string) *Tag {
	__obf_ea59cb61b72614bf := &__obf_691c267b735830b8{
		Parser: NewString(__obf_cbeee07d735d7228),
	}
	__obf_ea59cb61b72614bf.__obf_2f5e3667b42c4270()
	return &__obf_ea59cb61b72614bf.Tag
}

type __obf_691c267b735830b8 struct {
	*Parser

	Tag                    Tag
	__obf_afcde686ea56686d bool
	__obf_057359bed7c432b3 string
}

func (__obf_ea59cb61b72614bf *__obf_691c267b735830b8) __obf_c8652ea93747ae4b(__obf_057359bed7c432b3, __obf_d2df4ba851135340 string) {
	__obf_057359bed7c432b3 = strings.TrimSpace(__obf_057359bed7c432b3)
	__obf_d2df4ba851135340 = strings.TrimSpace(__obf_d2df4ba851135340)

	if !__obf_ea59cb61b72614bf.__obf_afcde686ea56686d {
		__obf_ea59cb61b72614bf.__obf_afcde686ea56686d = true
		if __obf_057359bed7c432b3 == "" {
			__obf_ea59cb61b72614bf.
				Tag.Name = __obf_d2df4ba851135340
			return
		}
	}
	if __obf_ea59cb61b72614bf.Tag.Options == nil {
		__obf_ea59cb61b72614bf.
			Tag.Options = make(map[string]string)
	}
	if __obf_057359bed7c432b3 == "" {
		__obf_ea59cb61b72614bf.
			Tag.Options[__obf_d2df4ba851135340] = ""
	} else {
		__obf_ea59cb61b72614bf.
			Tag.Options[__obf_057359bed7c432b3] = __obf_d2df4ba851135340
	}
}

func (__obf_ea59cb61b72614bf *__obf_691c267b735830b8) __obf_2f5e3667b42c4270() {
	__obf_ea59cb61b72614bf.__obf_057359bed7c432b3 = ""

	var __obf_b61eeabc6ce9f935 []byte
	for __obf_ea59cb61b72614bf.Valid() {
		__obf_19bc6ee95fd2a039 := __obf_ea59cb61b72614bf.Read()
		switch __obf_19bc6ee95fd2a039 {
		case ',':
			__obf_ea59cb61b72614bf.
				Skip(' ')
			__obf_ea59cb61b72614bf.__obf_c8652ea93747ae4b("", string(__obf_b61eeabc6ce9f935))
			__obf_ea59cb61b72614bf.__obf_2f5e3667b42c4270()
			return
		case ':':
			__obf_ea59cb61b72614bf.__obf_057359bed7c432b3 = string(__obf_b61eeabc6ce9f935)
			__obf_ea59cb61b72614bf.__obf_c99045896263f227()
			return
		case '\'':
			__obf_ea59cb61b72614bf.__obf_bd6cdfc0556ca786()
			return
		default:
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
		}
	}

	if len(__obf_b61eeabc6ce9f935) > 0 {
		__obf_ea59cb61b72614bf.__obf_c8652ea93747ae4b("", string(__obf_b61eeabc6ce9f935))
	}
}

func (__obf_ea59cb61b72614bf *__obf_691c267b735830b8) __obf_c99045896263f227() {
	const __obf_d3254a5876084250 = '\''
	__obf_19bc6ee95fd2a039 := __obf_ea59cb61b72614bf.Peek()
	if __obf_19bc6ee95fd2a039 == __obf_d3254a5876084250 {
		__obf_ea59cb61b72614bf.
			Skip(__obf_d3254a5876084250)
		__obf_ea59cb61b72614bf.__obf_bd6cdfc0556ca786()
		return
	}

	var __obf_b61eeabc6ce9f935 []byte
	for __obf_ea59cb61b72614bf.Valid() {
		__obf_19bc6ee95fd2a039 = __obf_ea59cb61b72614bf.Read()
		switch __obf_19bc6ee95fd2a039 {
		case '\\':
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_ea59cb61b72614bf.Read())
		case '(':
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
			__obf_b61eeabc6ce9f935 = __obf_ea59cb61b72614bf.__obf_c0fa4418367d9b0f(__obf_b61eeabc6ce9f935)
		case ',':
			__obf_ea59cb61b72614bf.
				Skip(' ')
			__obf_ea59cb61b72614bf.__obf_c8652ea93747ae4b(__obf_ea59cb61b72614bf.__obf_057359bed7c432b3, string(__obf_b61eeabc6ce9f935))
			__obf_ea59cb61b72614bf.__obf_2f5e3667b42c4270()
			return
		default:
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
		}
	}
	__obf_ea59cb61b72614bf.__obf_c8652ea93747ae4b(__obf_ea59cb61b72614bf.__obf_057359bed7c432b3, string(__obf_b61eeabc6ce9f935))
}

func (__obf_ea59cb61b72614bf *__obf_691c267b735830b8) __obf_c0fa4418367d9b0f(__obf_b61eeabc6ce9f935 []byte) []byte {
	var __obf_b1131d76e28fc414 int
__obf_eb3933e9b273dd71:
	for __obf_ea59cb61b72614bf.Valid() {
		__obf_19bc6ee95fd2a039 := __obf_ea59cb61b72614bf.Read()
		switch __obf_19bc6ee95fd2a039 {
		case '\\':
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_ea59cb61b72614bf.Read())
		case '(':
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
			__obf_b1131d76e28fc414++
		case ')':
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
			__obf_b1131d76e28fc414--
			if __obf_b1131d76e28fc414 < 0 {
				break __obf_eb3933e9b273dd71
			}
		default:
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_19bc6ee95fd2a039)
		}
	}
	return __obf_b61eeabc6ce9f935
}

func (__obf_ea59cb61b72614bf *__obf_691c267b735830b8) __obf_bd6cdfc0556ca786() {
	const __obf_d3254a5876084250 = '\''
	var __obf_b61eeabc6ce9f935 []byte
	for __obf_ea59cb61b72614bf.Valid() {
		__obf_1a7f9ca77442cf49, __obf_6f62606b6640bca1 := __obf_ea59cb61b72614bf.ReadSep(__obf_d3254a5876084250)
		if !__obf_6f62606b6640bca1 {
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_1a7f9ca77442cf49...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_1a7f9ca77442cf49) > 0 && __obf_1a7f9ca77442cf49[len(__obf_1a7f9ca77442cf49)-1] == '\\' {
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_1a7f9ca77442cf49[:len(__obf_1a7f9ca77442cf49)-1]...)
			__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_d3254a5876084250)
			continue
		}
		__obf_b61eeabc6ce9f935 = append(__obf_b61eeabc6ce9f935, __obf_1a7f9ca77442cf49...)
		break
	}
	__obf_ea59cb61b72614bf.__obf_c8652ea93747ae4b(__obf_ea59cb61b72614bf.__obf_057359bed7c432b3, string(__obf_b61eeabc6ce9f935))
	if __obf_ea59cb61b72614bf.Skip(',') {
		__obf_ea59cb61b72614bf.
			Skip(' ')
	}
	__obf_ea59cb61b72614bf.__obf_2f5e3667b42c4270()
}

type Parser struct {
	__obf_b61eeabc6ce9f935 []byte
	__obf_7438eb55cc32f4da int
}

func New(__obf_b61eeabc6ce9f935 []byte) *Parser {
	return &Parser{__obf_b61eeabc6ce9f935: __obf_b61eeabc6ce9f935}
}

func NewString(__obf_cbeee07d735d7228 string) *Parser {
	return New(StringToBytes(__obf_cbeee07d735d7228))
}

func (__obf_ea59cb61b72614bf *Parser) Bytes() []byte {
	return __obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da:]
}

func (__obf_ea59cb61b72614bf *Parser) Valid() bool {
	return __obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da < len(__obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935)
}

func (__obf_ea59cb61b72614bf *Parser) Read() byte {
	if __obf_ea59cb61b72614bf.Valid() {
		__obf_19bc6ee95fd2a039 := __obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da]
		__obf_ea59cb61b72614bf.
			Advance()
		return __obf_19bc6ee95fd2a039
	}
	return 0
}

func (__obf_ea59cb61b72614bf *Parser) Peek() byte {
	if __obf_ea59cb61b72614bf.Valid() {
		return __obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da]
	}
	return 0
}

func (__obf_ea59cb61b72614bf *Parser) Advance() {
	__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da++
}

func (__obf_ea59cb61b72614bf *Parser) Skip(__obf_e0707c75dba4dd9e byte) bool {
	if __obf_ea59cb61b72614bf.Peek() == __obf_e0707c75dba4dd9e {
		__obf_ea59cb61b72614bf.
			Advance()
		return true
	}
	return false
}

func (__obf_ea59cb61b72614bf *Parser) SkipBytes(__obf_e0707c75dba4dd9e []byte) bool {
	if len(__obf_e0707c75dba4dd9e) > len(__obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da:]) {
		return false
	}
	if !bytes.Equal(__obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da:__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da+len(__obf_e0707c75dba4dd9e)], __obf_e0707c75dba4dd9e) {
		return false
	}
	__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da += len(__obf_e0707c75dba4dd9e)
	return true
}

func (__obf_ea59cb61b72614bf *Parser) ReadSep(__obf_8f8aec4337ed23ef byte) ([]byte, bool) {
	__obf_29c526662e13c4a8 := bytes.IndexByte(__obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da:], __obf_8f8aec4337ed23ef)
	if __obf_29c526662e13c4a8 == -1 {
		__obf_b61eeabc6ce9f935 := __obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da:]
		__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da = len(__obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935)
		return __obf_b61eeabc6ce9f935, false
	}
	__obf_b61eeabc6ce9f935 := __obf_ea59cb61b72614bf.__obf_b61eeabc6ce9f935[__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da : __obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da+__obf_29c526662e13c4a8]
	__obf_ea59cb61b72614bf.__obf_7438eb55cc32f4da += __obf_29c526662e13c4a8 + 1
	return __obf_b61eeabc6ce9f935, true
}

func BytesToString(__obf_b61eeabc6ce9f935 []byte) string {
	return string(__obf_b61eeabc6ce9f935)
}

func StringToBytes(__obf_cbeee07d735d7228 string) []byte {
	return []byte(__obf_cbeee07d735d7228)
}
