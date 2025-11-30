package __obf_ae73ba41a74c0e0a

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_ec93f8fd81bbecdf = "data:image/png;base64,"
	__obf_5d20a0e64d75cc6f = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_49505236572de8ef image.Image) (__obf_8fd7dc208ce2cd86 []byte, __obf_480cd6251cefe029 error) {
	var __obf_25a33fa28fd21353 bytes.Buffer
	if __obf_480cd6251cefe029 = png.Encode(&__obf_25a33fa28fd21353, __obf_49505236572de8ef); __obf_480cd6251cefe029 != nil {
		return
	}
	__obf_8fd7dc208ce2cd86 = __obf_25a33fa28fd21353.Bytes()
	__obf_25a33fa28fd21353.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_49505236572de8ef image.Image, __obf_8bcdf485f5fd9341 int) (__obf_8fd7dc208ce2cd86 []byte, __obf_480cd6251cefe029 error) {
	var __obf_25a33fa28fd21353 bytes.Buffer
	if __obf_480cd6251cefe029 = jpeg.Encode(&__obf_25a33fa28fd21353, __obf_49505236572de8ef, &jpeg.Options{Quality: __obf_8bcdf485f5fd9341}); __obf_480cd6251cefe029 != nil {
		return
	}
	__obf_8fd7dc208ce2cd86 = __obf_25a33fa28fd21353.Bytes()
	__obf_25a33fa28fd21353.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_43ab9cb8031158e8 []byte) (__obf_49505236572de8ef image.Image, __obf_480cd6251cefe029 error) {
	var __obf_25a33fa28fd21353 bytes.Buffer
	__obf_25a33fa28fd21353.
		Write(__obf_43ab9cb8031158e8)
	__obf_49505236572de8ef, __obf_480cd6251cefe029 = jpeg.Decode(&__obf_25a33fa28fd21353)
	__obf_25a33fa28fd21353.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_43ab9cb8031158e8 []byte) (__obf_49505236572de8ef image.Image, __obf_480cd6251cefe029 error) {
	var __obf_25a33fa28fd21353 bytes.Buffer
	__obf_25a33fa28fd21353.
		Write(__obf_43ab9cb8031158e8)
	__obf_49505236572de8ef, __obf_480cd6251cefe029 = png.Decode(&__obf_25a33fa28fd21353)
	__obf_25a33fa28fd21353.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_49505236572de8ef image.Image) (string, error) {
	__obf_e827b398b3c3eafd, __obf_480cd6251cefe029 := EncodePNGToByte(__obf_49505236572de8ef)
	if __obf_480cd6251cefe029 != nil {
		return "", __obf_480cd6251cefe029
	}
	__obf_20bef7eb2615837e := base64.StdEncoding.EncodeToString(__obf_e827b398b3c3eafd)
	return __obf_ec93f8fd81bbecdf + __obf_20bef7eb2615837e, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_49505236572de8ef image.Image, __obf_8bcdf485f5fd9341 int) (string, error) {
	__obf_e827b398b3c3eafd, __obf_480cd6251cefe029 := EncodeJPEGToByte(__obf_49505236572de8ef, __obf_8bcdf485f5fd9341)
	if __obf_480cd6251cefe029 != nil {
		return "", __obf_480cd6251cefe029
	}
	__obf_20bef7eb2615837e := base64.StdEncoding.EncodeToString(__obf_e827b398b3c3eafd)
	return __obf_5d20a0e64d75cc6f + __obf_20bef7eb2615837e, nil
}
