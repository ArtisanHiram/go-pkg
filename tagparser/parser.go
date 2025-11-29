package __obf_db8a49503ac5ea4a

import (
	"bytes"
	"strings"
)

type Tag struct {
	Name    string
	Options map[string]string
}

func (__obf_16343ff89796b132 *Tag) HasOption(__obf_86dd6d385423a4fc string) bool {
	_, __obf_c175f5748f8957d2 := __obf_16343ff89796b132.Options[__obf_86dd6d385423a4fc]
	return __obf_c175f5748f8957d2
}

func Parse(__obf_8421f5e4eb58bbdd string) *Tag {
	__obf_5906296a9fee4dc0 := &__obf_bf5a4670145148c6{
		Parser: NewString(__obf_8421f5e4eb58bbdd),
	}
	__obf_5906296a9fee4dc0.__obf_d163bbd7b36e381b()
	return &__obf_5906296a9fee4dc0.Tag
}

type __obf_bf5a4670145148c6 struct {
	*Parser

	Tag                    Tag
	__obf_6929c54de2b1dfae bool
	__obf_55b1a8ee1757875c string
}

func (__obf_5906296a9fee4dc0 *__obf_bf5a4670145148c6) __obf_f729022363b53d6b(__obf_55b1a8ee1757875c, __obf_b7286c3c48fece5e string) {
	__obf_55b1a8ee1757875c = strings.TrimSpace(__obf_55b1a8ee1757875c)
	__obf_b7286c3c48fece5e = strings.TrimSpace(__obf_b7286c3c48fece5e)

	if !__obf_5906296a9fee4dc0.__obf_6929c54de2b1dfae {
		__obf_5906296a9fee4dc0.__obf_6929c54de2b1dfae = true
		if __obf_55b1a8ee1757875c == "" {
			__obf_5906296a9fee4dc0.
				Tag.Name = __obf_b7286c3c48fece5e
			return
		}
	}
	if __obf_5906296a9fee4dc0.Tag.Options == nil {
		__obf_5906296a9fee4dc0.
			Tag.Options = make(map[string]string)
	}
	if __obf_55b1a8ee1757875c == "" {
		__obf_5906296a9fee4dc0.
			Tag.Options[__obf_b7286c3c48fece5e] = ""
	} else {
		__obf_5906296a9fee4dc0.
			Tag.Options[__obf_55b1a8ee1757875c] = __obf_b7286c3c48fece5e
	}
}

func (__obf_5906296a9fee4dc0 *__obf_bf5a4670145148c6) __obf_d163bbd7b36e381b() {
	__obf_5906296a9fee4dc0.__obf_55b1a8ee1757875c = ""

	var __obf_19d98e7d079b0c5b []byte
	for __obf_5906296a9fee4dc0.Valid() {
		__obf_97c98900db46cf5f := __obf_5906296a9fee4dc0.Read()
		switch __obf_97c98900db46cf5f {
		case ',':
			__obf_5906296a9fee4dc0.
				Skip(' ')
			__obf_5906296a9fee4dc0.__obf_f729022363b53d6b("", string(__obf_19d98e7d079b0c5b))
			__obf_5906296a9fee4dc0.__obf_d163bbd7b36e381b()
			return
		case ':':
			__obf_5906296a9fee4dc0.__obf_55b1a8ee1757875c = string(__obf_19d98e7d079b0c5b)
			__obf_5906296a9fee4dc0.__obf_0a6b7e4a8a27fb25()
			return
		case '\'':
			__obf_5906296a9fee4dc0.__obf_aab2d178e46527be()
			return
		default:
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
		}
	}

	if len(__obf_19d98e7d079b0c5b) > 0 {
		__obf_5906296a9fee4dc0.__obf_f729022363b53d6b("", string(__obf_19d98e7d079b0c5b))
	}
}

func (__obf_5906296a9fee4dc0 *__obf_bf5a4670145148c6) __obf_0a6b7e4a8a27fb25() {
	const __obf_18592e5adf076ace = '\''
	__obf_97c98900db46cf5f := __obf_5906296a9fee4dc0.Peek()
	if __obf_97c98900db46cf5f == __obf_18592e5adf076ace {
		__obf_5906296a9fee4dc0.
			Skip(__obf_18592e5adf076ace)
		__obf_5906296a9fee4dc0.__obf_aab2d178e46527be()
		return
	}

	var __obf_19d98e7d079b0c5b []byte
	for __obf_5906296a9fee4dc0.Valid() {
		__obf_97c98900db46cf5f = __obf_5906296a9fee4dc0.Read()
		switch __obf_97c98900db46cf5f {
		case '\\':
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_5906296a9fee4dc0.Read())
		case '(':
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
			__obf_19d98e7d079b0c5b = __obf_5906296a9fee4dc0.__obf_b223aead59125e44(__obf_19d98e7d079b0c5b)
		case ',':
			__obf_5906296a9fee4dc0.
				Skip(' ')
			__obf_5906296a9fee4dc0.__obf_f729022363b53d6b(__obf_5906296a9fee4dc0.__obf_55b1a8ee1757875c, string(__obf_19d98e7d079b0c5b))
			__obf_5906296a9fee4dc0.__obf_d163bbd7b36e381b()
			return
		default:
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
		}
	}
	__obf_5906296a9fee4dc0.__obf_f729022363b53d6b(__obf_5906296a9fee4dc0.__obf_55b1a8ee1757875c, string(__obf_19d98e7d079b0c5b))
}

func (__obf_5906296a9fee4dc0 *__obf_bf5a4670145148c6) __obf_b223aead59125e44(__obf_19d98e7d079b0c5b []byte) []byte {
	var __obf_bf90ee7bbf2363d2 int
__obf_2a748f97a2ebf623:
	for __obf_5906296a9fee4dc0.Valid() {
		__obf_97c98900db46cf5f := __obf_5906296a9fee4dc0.Read()
		switch __obf_97c98900db46cf5f {
		case '\\':
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_5906296a9fee4dc0.Read())
		case '(':
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
			__obf_bf90ee7bbf2363d2++
		case ')':
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
			__obf_bf90ee7bbf2363d2--
			if __obf_bf90ee7bbf2363d2 < 0 {
				break __obf_2a748f97a2ebf623
			}
		default:
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_97c98900db46cf5f)
		}
	}
	return __obf_19d98e7d079b0c5b
}

func (__obf_5906296a9fee4dc0 *__obf_bf5a4670145148c6) __obf_aab2d178e46527be() {
	const __obf_18592e5adf076ace = '\''
	var __obf_19d98e7d079b0c5b []byte
	for __obf_5906296a9fee4dc0.Valid() {
		__obf_8cf905ffdbad40fb, __obf_c175f5748f8957d2 := __obf_5906296a9fee4dc0.ReadSep(__obf_18592e5adf076ace)
		if !__obf_c175f5748f8957d2 {
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_8cf905ffdbad40fb...)
			break
		}

		// keep the escaped single-quote, and continue until we've found the
		// one that isn't.
		if len(__obf_8cf905ffdbad40fb) > 0 && __obf_8cf905ffdbad40fb[len(__obf_8cf905ffdbad40fb)-1] == '\\' {
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_8cf905ffdbad40fb[:len(__obf_8cf905ffdbad40fb)-1]...)
			__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_18592e5adf076ace)
			continue
		}
		__obf_19d98e7d079b0c5b = append(__obf_19d98e7d079b0c5b, __obf_8cf905ffdbad40fb...)
		break
	}
	__obf_5906296a9fee4dc0.__obf_f729022363b53d6b(__obf_5906296a9fee4dc0.__obf_55b1a8ee1757875c, string(__obf_19d98e7d079b0c5b))
	if __obf_5906296a9fee4dc0.Skip(',') {
		__obf_5906296a9fee4dc0.
			Skip(' ')
	}
	__obf_5906296a9fee4dc0.__obf_d163bbd7b36e381b()
}

type Parser struct {
	__obf_19d98e7d079b0c5b []byte
	__obf_6bcc420db5909eab int
}

func New(__obf_19d98e7d079b0c5b []byte) *Parser {
	return &Parser{__obf_19d98e7d079b0c5b: __obf_19d98e7d079b0c5b}
}

func NewString(__obf_8421f5e4eb58bbdd string) *Parser {
	return New(StringToBytes(__obf_8421f5e4eb58bbdd))
}

func (__obf_5906296a9fee4dc0 *Parser) Bytes() []byte {
	return __obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab:]
}

func (__obf_5906296a9fee4dc0 *Parser) Valid() bool {
	return __obf_5906296a9fee4dc0.__obf_6bcc420db5909eab < len(__obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b)
}

func (__obf_5906296a9fee4dc0 *Parser) Read() byte {
	if __obf_5906296a9fee4dc0.Valid() {
		__obf_97c98900db46cf5f := __obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab]
		__obf_5906296a9fee4dc0.
			Advance()
		return __obf_97c98900db46cf5f
	}
	return 0
}

func (__obf_5906296a9fee4dc0 *Parser) Peek() byte {
	if __obf_5906296a9fee4dc0.Valid() {
		return __obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab]
	}
	return 0
}

func (__obf_5906296a9fee4dc0 *Parser) Advance() {
	__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab++
}

func (__obf_5906296a9fee4dc0 *Parser) Skip(__obf_9afc5992f4c56f68 byte) bool {
	if __obf_5906296a9fee4dc0.Peek() == __obf_9afc5992f4c56f68 {
		__obf_5906296a9fee4dc0.
			Advance()
		return true
	}
	return false
}

func (__obf_5906296a9fee4dc0 *Parser) SkipBytes(__obf_9afc5992f4c56f68 []byte) bool {
	if len(__obf_9afc5992f4c56f68) > len(__obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab:]) {
		return false
	}
	if !bytes.Equal(__obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab:__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab+len(__obf_9afc5992f4c56f68)], __obf_9afc5992f4c56f68) {
		return false
	}
	__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab += len(__obf_9afc5992f4c56f68)
	return true
}

func (__obf_5906296a9fee4dc0 *Parser) ReadSep(__obf_e92df56ae2f26273 byte) ([]byte, bool) {
	__obf_bd193bdcca5b398a := bytes.IndexByte(__obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab:], __obf_e92df56ae2f26273)
	if __obf_bd193bdcca5b398a == -1 {
		__obf_19d98e7d079b0c5b := __obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab:]
		__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab = len(__obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b)
		return __obf_19d98e7d079b0c5b, false
	}
	__obf_19d98e7d079b0c5b := __obf_5906296a9fee4dc0.__obf_19d98e7d079b0c5b[__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab : __obf_5906296a9fee4dc0.__obf_6bcc420db5909eab+__obf_bd193bdcca5b398a]
	__obf_5906296a9fee4dc0.__obf_6bcc420db5909eab += __obf_bd193bdcca5b398a + 1
	return __obf_19d98e7d079b0c5b, true
}

func BytesToString(__obf_19d98e7d079b0c5b []byte) string {
	return string(__obf_19d98e7d079b0c5b)
}

func StringToBytes(__obf_8421f5e4eb58bbdd string) []byte {
	return []byte(__obf_8421f5e4eb58bbdd)
}
