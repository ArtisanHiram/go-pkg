package __obf_6f9b75fc21ef8ef6

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_bf67e815d0cca627, __obf_5c939910df183df7, __obf_1eb5b853239c04fa []byte) ([]byte, error) {
	__obf_7599ad598015b486, __obf_8eefb2cd18f480f3 := aes.NewCipher(__obf_5c939910df183df7)
	if __obf_8eefb2cd18f480f3 != nil {
		return nil, __obf_8eefb2cd18f480f3
	}
	__obf_ec8b44e20c0bb9f2 := __obf_7599ad598015b486.BlockSize()
	if len(__obf_1eb5b853239c04fa) == 0 {
		__obf_1eb5b853239c04fa = __obf_5c939910df183df7
	}
	__obf_bf67e815d0cca627 = __obf_5976e75519691575(__obf_bf67e815d0cca627, __obf_ec8b44e20c0bb9f2)
	__obf_ed05673c58508641 := cipher.NewCBCEncrypter(__obf_7599ad598015b486, __obf_1eb5b853239c04fa[:__obf_ec8b44e20c0bb9f2])
	__obf_6e6d8a85272ea949 := make([]byte, len(__obf_bf67e815d0cca627))
	__obf_ed05673c58508641.
		CryptBlocks(__obf_6e6d8a85272ea949, __obf_bf67e815d0cca627)
	return __obf_6e6d8a85272ea949, nil
}

func AES128Decrypt(__obf_6e6d8a85272ea949, __obf_5c939910df183df7, __obf_1eb5b853239c04fa []byte) ([]byte, error) {
	__obf_7599ad598015b486, __obf_8eefb2cd18f480f3 := aes.NewCipher(__obf_5c939910df183df7)
	if __obf_8eefb2cd18f480f3 != nil {
		return nil, __obf_8eefb2cd18f480f3
	}
	__obf_ec8b44e20c0bb9f2 := __obf_7599ad598015b486.BlockSize()
	if len(__obf_1eb5b853239c04fa) == 0 {
		__obf_1eb5b853239c04fa = __obf_5c939910df183df7
	}
	__obf_ed05673c58508641 := cipher.NewCBCDecrypter(__obf_7599ad598015b486, __obf_1eb5b853239c04fa[:__obf_ec8b44e20c0bb9f2])
	__obf_bf67e815d0cca627 := make([]byte, len(__obf_6e6d8a85272ea949))
	__obf_ed05673c58508641.
		CryptBlocks(__obf_bf67e815d0cca627, __obf_6e6d8a85272ea949)
	__obf_bf67e815d0cca627 = __obf_5d8ac0a0c88397dd(__obf_bf67e815d0cca627)
	return __obf_bf67e815d0cca627, nil
}

func __obf_5976e75519691575(__obf_906cf4871575d81f []byte, __obf_ec8b44e20c0bb9f2 int) []byte {
	__obf_2a831220188f7db1 := __obf_ec8b44e20c0bb9f2 - len(__obf_906cf4871575d81f)%__obf_ec8b44e20c0bb9f2
	__obf_0953fce06defcdca := bytes.Repeat([]byte{byte(__obf_2a831220188f7db1)}, __obf_2a831220188f7db1)
	return append(__obf_906cf4871575d81f, __obf_0953fce06defcdca...)
}

func __obf_5d8ac0a0c88397dd(__obf_bf67e815d0cca627 []byte) []byte {
	__obf_8fdba65d97f5157f := len(__obf_bf67e815d0cca627)
	__obf_2e97e0d727db6809 := int(__obf_bf67e815d0cca627[__obf_8fdba65d97f5157f-1])
	return __obf_bf67e815d0cca627[:(__obf_8fdba65d97f5157f - __obf_2e97e0d727db6809)]
}
