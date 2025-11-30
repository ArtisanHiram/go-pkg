package __obf_33bc7bf683859fa2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_55803553715560c9, __obf_729bd81b45181227, __obf_d0b65848cf2133b8 []byte) ([]byte, error) {
	__obf_f1c6299fc5e3bf4a, __obf_19a0c091a533826e := aes.NewCipher(__obf_729bd81b45181227)
	if __obf_19a0c091a533826e != nil {
		return nil, __obf_19a0c091a533826e
	}
	__obf_999fc092b301a0ed := __obf_f1c6299fc5e3bf4a.BlockSize()
	if len(__obf_d0b65848cf2133b8) == 0 {
		__obf_d0b65848cf2133b8 = __obf_729bd81b45181227
	}
	__obf_55803553715560c9 = __obf_d3b3bc7fb56c8cac(__obf_55803553715560c9, __obf_999fc092b301a0ed)
	__obf_44d840c1eb42557f := cipher.NewCBCEncrypter(__obf_f1c6299fc5e3bf4a, __obf_d0b65848cf2133b8[:__obf_999fc092b301a0ed])
	__obf_a2afea43198859b3 := make([]byte, len(__obf_55803553715560c9))
	__obf_44d840c1eb42557f.
		CryptBlocks(__obf_a2afea43198859b3, __obf_55803553715560c9)
	return __obf_a2afea43198859b3, nil
}

func AES128Decrypt(__obf_a2afea43198859b3, __obf_729bd81b45181227, __obf_d0b65848cf2133b8 []byte) ([]byte, error) {
	__obf_f1c6299fc5e3bf4a, __obf_19a0c091a533826e := aes.NewCipher(__obf_729bd81b45181227)
	if __obf_19a0c091a533826e != nil {
		return nil, __obf_19a0c091a533826e
	}
	__obf_999fc092b301a0ed := __obf_f1c6299fc5e3bf4a.BlockSize()
	if len(__obf_d0b65848cf2133b8) == 0 {
		__obf_d0b65848cf2133b8 = __obf_729bd81b45181227
	}
	__obf_44d840c1eb42557f := cipher.NewCBCDecrypter(__obf_f1c6299fc5e3bf4a, __obf_d0b65848cf2133b8[:__obf_999fc092b301a0ed])
	__obf_55803553715560c9 := make([]byte, len(__obf_a2afea43198859b3))
	__obf_44d840c1eb42557f.
		CryptBlocks(__obf_55803553715560c9, __obf_a2afea43198859b3)
	__obf_55803553715560c9 = __obf_bb2b06b6f9c77acc(__obf_55803553715560c9)
	return __obf_55803553715560c9, nil
}

func __obf_d3b3bc7fb56c8cac(__obf_4177040155855a6a []byte, __obf_999fc092b301a0ed int) []byte {
	__obf_9f63ae145d399d0e := __obf_999fc092b301a0ed - len(__obf_4177040155855a6a)%__obf_999fc092b301a0ed
	__obf_14e7255c591f7a6a := bytes.Repeat([]byte{byte(__obf_9f63ae145d399d0e)}, __obf_9f63ae145d399d0e)
	return append(__obf_4177040155855a6a, __obf_14e7255c591f7a6a...)
}

func __obf_bb2b06b6f9c77acc(__obf_55803553715560c9 []byte) []byte {
	__obf_0a7031472e778852 := len(__obf_55803553715560c9)
	__obf_fd112f40e66df8e3 := int(__obf_55803553715560c9[__obf_0a7031472e778852-1])
	return __obf_55803553715560c9[:(__obf_0a7031472e778852 - __obf_fd112f40e66df8e3)]
}
