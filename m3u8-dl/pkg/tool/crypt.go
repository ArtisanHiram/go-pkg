package __obf_c37d95bc12b416d3

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_e727fcd2adaea5bb, __obf_4050616e8eb590ee, __obf_6525ad01fe41171c []byte) ([]byte, error) {
	__obf_9913d1f730b07ebd, __obf_3ce485578a57d5a7 := aes.NewCipher(__obf_4050616e8eb590ee)
	if __obf_3ce485578a57d5a7 != nil {
		return nil, __obf_3ce485578a57d5a7
	}
	__obf_0bb262c8634631a8 := __obf_9913d1f730b07ebd.BlockSize()
	if len(__obf_6525ad01fe41171c) == 0 {
		__obf_6525ad01fe41171c = __obf_4050616e8eb590ee
	}
	__obf_e727fcd2adaea5bb = __obf_9e31c4919ce7e2af(__obf_e727fcd2adaea5bb, __obf_0bb262c8634631a8)
	__obf_823ba616a3bf20fd := cipher.NewCBCEncrypter(__obf_9913d1f730b07ebd, __obf_6525ad01fe41171c[:__obf_0bb262c8634631a8])
	__obf_14e78a20790a28c7 := make([]byte, len(__obf_e727fcd2adaea5bb))
	__obf_823ba616a3bf20fd.
		CryptBlocks(__obf_14e78a20790a28c7, __obf_e727fcd2adaea5bb)
	return __obf_14e78a20790a28c7, nil
}

func AES128Decrypt(__obf_14e78a20790a28c7, __obf_4050616e8eb590ee, __obf_6525ad01fe41171c []byte) ([]byte, error) {
	__obf_9913d1f730b07ebd, __obf_3ce485578a57d5a7 := aes.NewCipher(__obf_4050616e8eb590ee)
	if __obf_3ce485578a57d5a7 != nil {
		return nil, __obf_3ce485578a57d5a7
	}
	__obf_0bb262c8634631a8 := __obf_9913d1f730b07ebd.BlockSize()
	if len(__obf_6525ad01fe41171c) == 0 {
		__obf_6525ad01fe41171c = __obf_4050616e8eb590ee
	}
	__obf_823ba616a3bf20fd := cipher.NewCBCDecrypter(__obf_9913d1f730b07ebd, __obf_6525ad01fe41171c[:__obf_0bb262c8634631a8])
	__obf_e727fcd2adaea5bb := make([]byte, len(__obf_14e78a20790a28c7))
	__obf_823ba616a3bf20fd.
		CryptBlocks(__obf_e727fcd2adaea5bb, __obf_14e78a20790a28c7)
	__obf_e727fcd2adaea5bb = __obf_2456a80a6ce1db1a(__obf_e727fcd2adaea5bb)
	return __obf_e727fcd2adaea5bb, nil
}

func __obf_9e31c4919ce7e2af(__obf_6be8f5a59a8d8531 []byte, __obf_0bb262c8634631a8 int) []byte {
	__obf_badf664a1e0db5db := __obf_0bb262c8634631a8 - len(__obf_6be8f5a59a8d8531)%__obf_0bb262c8634631a8
	__obf_40e5bc0bdc33f921 := bytes.Repeat([]byte{byte(__obf_badf664a1e0db5db)}, __obf_badf664a1e0db5db)
	return append(__obf_6be8f5a59a8d8531, __obf_40e5bc0bdc33f921...)
}

func __obf_2456a80a6ce1db1a(__obf_e727fcd2adaea5bb []byte) []byte {
	__obf_b9a5ecc0cd4d420b := len(__obf_e727fcd2adaea5bb)
	__obf_6d8669500f8dbb89 := int(__obf_e727fcd2adaea5bb[__obf_b9a5ecc0cd4d420b-1])
	return __obf_e727fcd2adaea5bb[:(__obf_b9a5ecc0cd4d420b - __obf_6d8669500f8dbb89)]
}
