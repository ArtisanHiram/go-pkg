package __obf_ddeaac61ff7fc4bb

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_6a9a79a93865e754 = "data:image/png;base64,"
	__obf_96e54acce993c199 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_e730a61923e268ec image.Image) (__obf_c77942a7d5578a17 []byte, __obf_630bf6ccf46ac8b5 error) {
	var __obf_9457d53ef27a1d39 bytes.Buffer
	if __obf_630bf6ccf46ac8b5 = png.Encode(&__obf_9457d53ef27a1d39, __obf_e730a61923e268ec); __obf_630bf6ccf46ac8b5 != nil {
		return
	}
	__obf_c77942a7d5578a17 = __obf_9457d53ef27a1d39.Bytes()
	__obf_9457d53ef27a1d39.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_e730a61923e268ec image.Image, __obf_9f35251f7917e75e int) (__obf_c77942a7d5578a17 []byte, __obf_630bf6ccf46ac8b5 error) {
	var __obf_9457d53ef27a1d39 bytes.Buffer
	if __obf_630bf6ccf46ac8b5 = jpeg.Encode(&__obf_9457d53ef27a1d39, __obf_e730a61923e268ec, &jpeg.Options{Quality: __obf_9f35251f7917e75e}); __obf_630bf6ccf46ac8b5 != nil {
		return
	}
	__obf_c77942a7d5578a17 = __obf_9457d53ef27a1d39.Bytes()
	__obf_9457d53ef27a1d39.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_7723b6938513fa3a []byte) (__obf_e730a61923e268ec image.Image, __obf_630bf6ccf46ac8b5 error) {
	var __obf_9457d53ef27a1d39 bytes.Buffer
	__obf_9457d53ef27a1d39.
		Write(__obf_7723b6938513fa3a)
	__obf_e730a61923e268ec, __obf_630bf6ccf46ac8b5 = jpeg.Decode(&__obf_9457d53ef27a1d39)
	__obf_9457d53ef27a1d39.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_7723b6938513fa3a []byte) (__obf_e730a61923e268ec image.Image, __obf_630bf6ccf46ac8b5 error) {
	var __obf_9457d53ef27a1d39 bytes.Buffer
	__obf_9457d53ef27a1d39.
		Write(__obf_7723b6938513fa3a)
	__obf_e730a61923e268ec, __obf_630bf6ccf46ac8b5 = png.Decode(&__obf_9457d53ef27a1d39)
	__obf_9457d53ef27a1d39.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_e730a61923e268ec image.Image) (string, error) {
	__obf_71207d7cf5e36098, __obf_630bf6ccf46ac8b5 := EncodePNGToByte(__obf_e730a61923e268ec)
	if __obf_630bf6ccf46ac8b5 != nil {
		return "", __obf_630bf6ccf46ac8b5
	}
	__obf_b6a59d8f5ab1d2cc := base64.StdEncoding.EncodeToString(__obf_71207d7cf5e36098)
	return __obf_6a9a79a93865e754 + __obf_b6a59d8f5ab1d2cc, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_e730a61923e268ec image.Image, __obf_9f35251f7917e75e int) (string, error) {
	__obf_71207d7cf5e36098, __obf_630bf6ccf46ac8b5 := EncodeJPEGToByte(__obf_e730a61923e268ec, __obf_9f35251f7917e75e)
	if __obf_630bf6ccf46ac8b5 != nil {
		return "", __obf_630bf6ccf46ac8b5
	}
	__obf_b6a59d8f5ab1d2cc := base64.StdEncoding.EncodeToString(__obf_71207d7cf5e36098)
	return __obf_96e54acce993c199 + __obf_b6a59d8f5ab1d2cc, nil
}
