package __obf_22f65c3e757f8811

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AES128Encrypt(__obf_5b95a3df51017032, __obf_b4bfcb44095f612d, __obf_0bdd4eae9372243c []byte) ([]byte, error) {
	__obf_717b749e4f2fccb3, __obf_f749646f37e10859 := aes.NewCipher(__obf_b4bfcb44095f612d)
	if __obf_f749646f37e10859 != nil {
		return nil, __obf_f749646f37e10859
	}
	__obf_d99a06fb33740bdb := __obf_717b749e4f2fccb3.BlockSize()
	if len(__obf_0bdd4eae9372243c) == 0 {
		__obf_0bdd4eae9372243c = __obf_b4bfcb44095f612d
	}
	__obf_5b95a3df51017032 = __obf_966585cff3833afd(__obf_5b95a3df51017032, __obf_d99a06fb33740bdb)
	__obf_60dd1d9c04f0a19f := cipher.NewCBCEncrypter(__obf_717b749e4f2fccb3, __obf_0bdd4eae9372243c[:__obf_d99a06fb33740bdb])
	__obf_d4aa4807390b359d := make([]byte, len(__obf_5b95a3df51017032))
	__obf_60dd1d9c04f0a19f.
		CryptBlocks(__obf_d4aa4807390b359d, __obf_5b95a3df51017032)
	return __obf_d4aa4807390b359d, nil
}

func AES128Decrypt(__obf_d4aa4807390b359d, __obf_b4bfcb44095f612d, __obf_0bdd4eae9372243c []byte) ([]byte, error) {
	__obf_717b749e4f2fccb3, __obf_f749646f37e10859 := aes.NewCipher(__obf_b4bfcb44095f612d)
	if __obf_f749646f37e10859 != nil {
		return nil, __obf_f749646f37e10859
	}
	__obf_d99a06fb33740bdb := __obf_717b749e4f2fccb3.BlockSize()
	if len(__obf_0bdd4eae9372243c) == 0 {
		__obf_0bdd4eae9372243c = __obf_b4bfcb44095f612d
	}
	__obf_60dd1d9c04f0a19f := cipher.NewCBCDecrypter(__obf_717b749e4f2fccb3, __obf_0bdd4eae9372243c[:__obf_d99a06fb33740bdb])
	__obf_5b95a3df51017032 := make([]byte, len(__obf_d4aa4807390b359d))
	__obf_60dd1d9c04f0a19f.
		CryptBlocks(__obf_5b95a3df51017032, __obf_d4aa4807390b359d)
	__obf_5b95a3df51017032 = __obf_a9645963dbe09d5b(__obf_5b95a3df51017032)
	return __obf_5b95a3df51017032, nil
}

func __obf_966585cff3833afd(__obf_88e6dcc265c9138e []byte, __obf_d99a06fb33740bdb int) []byte {
	__obf_9e755c2114e6783f := __obf_d99a06fb33740bdb - len(__obf_88e6dcc265c9138e)%__obf_d99a06fb33740bdb
	__obf_f9eda0ba986f881b := bytes.Repeat([]byte{byte(__obf_9e755c2114e6783f)}, __obf_9e755c2114e6783f)
	return append(__obf_88e6dcc265c9138e, __obf_f9eda0ba986f881b...)
}

func __obf_a9645963dbe09d5b(__obf_5b95a3df51017032 []byte) []byte {
	__obf_7de4b2841c0603d6 := len(__obf_5b95a3df51017032)
	__obf_0225bfe1b0612179 := int(__obf_5b95a3df51017032[__obf_7de4b2841c0603d6-1])
	return __obf_5b95a3df51017032[:(__obf_7de4b2841c0603d6 - __obf_0225bfe1b0612179)]
}
