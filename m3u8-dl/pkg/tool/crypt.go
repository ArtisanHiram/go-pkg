package __obf_67ff2b45a2ee7325

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_e21fb418d3b0b2dc, __obf_c55eb7bee3cc4601, __obf_5ba6ac8a91240044 []byte) ([]byte, error) {
	__obf_3a8086263fd20509, __obf_3d6b93bb4b532726 := aes.NewCipher(__obf_c55eb7bee3cc4601)
	if __obf_3d6b93bb4b532726 != nil {
		return nil, __obf_3d6b93bb4b532726
	}
	__obf_94aa063744a072db := __obf_3a8086263fd20509.BlockSize()
	if len(__obf_5ba6ac8a91240044) == 0 {
		__obf_5ba6ac8a91240044 = __obf_c55eb7bee3cc4601
	}
	__obf_e21fb418d3b0b2dc = __obf_21e4434bb0c9b281(__obf_e21fb418d3b0b2dc, __obf_94aa063744a072db)
	__obf_7d6e7ce3d06b0599 := cipher.NewCBCEncrypter(__obf_3a8086263fd20509, __obf_5ba6ac8a91240044[:__obf_94aa063744a072db])
	__obf_e21284189ab13509 := make([]byte, len(__obf_e21fb418d3b0b2dc))
	__obf_7d6e7ce3d06b0599.
		CryptBlocks(__obf_e21284189ab13509, __obf_e21fb418d3b0b2dc)
	return __obf_e21284189ab13509, nil
}

func AES128Decrypt(__obf_e21284189ab13509, __obf_c55eb7bee3cc4601, __obf_5ba6ac8a91240044 []byte) ([]byte, error) {
	__obf_3a8086263fd20509, __obf_3d6b93bb4b532726 := aes.NewCipher(__obf_c55eb7bee3cc4601)
	if __obf_3d6b93bb4b532726 != nil {
		return nil, __obf_3d6b93bb4b532726
	}
	__obf_94aa063744a072db := __obf_3a8086263fd20509.BlockSize()
	if len(__obf_5ba6ac8a91240044) == 0 {
		__obf_5ba6ac8a91240044 = __obf_c55eb7bee3cc4601
	}
	__obf_7d6e7ce3d06b0599 := cipher.NewCBCDecrypter(__obf_3a8086263fd20509, __obf_5ba6ac8a91240044[:__obf_94aa063744a072db])
	__obf_e21fb418d3b0b2dc := make([]byte, len(__obf_e21284189ab13509))
	__obf_7d6e7ce3d06b0599.
		CryptBlocks(__obf_e21fb418d3b0b2dc, __obf_e21284189ab13509)
	__obf_e21fb418d3b0b2dc = __obf_ffeeda1d38281899(__obf_e21fb418d3b0b2dc)
	return __obf_e21fb418d3b0b2dc, nil
}

func __obf_21e4434bb0c9b281(__obf_d97e1da97b3fe7b1 []byte, __obf_94aa063744a072db int) []byte {
	__obf_6331a1100c8439e8 := __obf_94aa063744a072db - len(__obf_d97e1da97b3fe7b1)%__obf_94aa063744a072db
	__obf_46e94fbfb8813c6e := bytes.Repeat([]byte{byte(__obf_6331a1100c8439e8)}, __obf_6331a1100c8439e8)
	return append(__obf_d97e1da97b3fe7b1, __obf_46e94fbfb8813c6e...)
}

func __obf_ffeeda1d38281899(__obf_e21fb418d3b0b2dc []byte) []byte {
	__obf_ed5056e91b5682d4 := len(__obf_e21fb418d3b0b2dc)
	__obf_4c2c58ed50cc2f60 := int(__obf_e21fb418d3b0b2dc[__obf_ed5056e91b5682d4-1])
	return __obf_e21fb418d3b0b2dc[:(__obf_ed5056e91b5682d4 - __obf_4c2c58ed50cc2f60)]
}
