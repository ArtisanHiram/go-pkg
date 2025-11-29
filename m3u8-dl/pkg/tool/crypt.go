package __obf_61bc591a775c8985

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_7995895c59073f45, __obf_1316fa2940875f91, __obf_640e119525b967d0 []byte) ([]byte, error) {
	__obf_7b95f95cb7ab5a8c, __obf_48065909968843ec := aes.NewCipher(__obf_1316fa2940875f91)
	if __obf_48065909968843ec != nil {
		return nil, __obf_48065909968843ec
	}
	__obf_1740c6c739e3a7ae := __obf_7b95f95cb7ab5a8c.BlockSize()
	if len(__obf_640e119525b967d0) == 0 {
		__obf_640e119525b967d0 = __obf_1316fa2940875f91
	}
	__obf_7995895c59073f45 = __obf_8cd542bdb636ffa5(__obf_7995895c59073f45, __obf_1740c6c739e3a7ae)
	__obf_af9f1eddb67f71be := cipher.NewCBCEncrypter(__obf_7b95f95cb7ab5a8c, __obf_640e119525b967d0[:__obf_1740c6c739e3a7ae])
	__obf_d5b97b196b602e6f := make([]byte, len(__obf_7995895c59073f45))
	__obf_af9f1eddb67f71be.
		CryptBlocks(__obf_d5b97b196b602e6f, __obf_7995895c59073f45)
	return __obf_d5b97b196b602e6f, nil
}

func AES128Decrypt(__obf_d5b97b196b602e6f, __obf_1316fa2940875f91, __obf_640e119525b967d0 []byte) ([]byte, error) {
	__obf_7b95f95cb7ab5a8c, __obf_48065909968843ec := aes.NewCipher(__obf_1316fa2940875f91)
	if __obf_48065909968843ec != nil {
		return nil, __obf_48065909968843ec
	}
	__obf_1740c6c739e3a7ae := __obf_7b95f95cb7ab5a8c.BlockSize()
	if len(__obf_640e119525b967d0) == 0 {
		__obf_640e119525b967d0 = __obf_1316fa2940875f91
	}
	__obf_af9f1eddb67f71be := cipher.NewCBCDecrypter(__obf_7b95f95cb7ab5a8c, __obf_640e119525b967d0[:__obf_1740c6c739e3a7ae])
	__obf_7995895c59073f45 := make([]byte, len(__obf_d5b97b196b602e6f))
	__obf_af9f1eddb67f71be.
		CryptBlocks(__obf_7995895c59073f45, __obf_d5b97b196b602e6f)
	__obf_7995895c59073f45 = __obf_fa901cc34f108164(__obf_7995895c59073f45)
	return __obf_7995895c59073f45, nil
}

func __obf_8cd542bdb636ffa5(__obf_a36bc11f547855ad []byte, __obf_1740c6c739e3a7ae int) []byte {
	__obf_258ff8127137c2a9 := __obf_1740c6c739e3a7ae - len(__obf_a36bc11f547855ad)%__obf_1740c6c739e3a7ae
	__obf_9feb9d0e1e1062ce := bytes.Repeat([]byte{byte(__obf_258ff8127137c2a9)}, __obf_258ff8127137c2a9)
	return append(__obf_a36bc11f547855ad, __obf_9feb9d0e1e1062ce...)
}

func __obf_fa901cc34f108164(__obf_7995895c59073f45 []byte) []byte {
	__obf_ef6df29a3268332c := len(__obf_7995895c59073f45)
	__obf_d9dfd902d85a8105 := int(__obf_7995895c59073f45[__obf_ef6df29a3268332c-1])
	return __obf_7995895c59073f45[:(__obf_ef6df29a3268332c - __obf_d9dfd902d85a8105)]
}
