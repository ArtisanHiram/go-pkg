package __obf_0e73be9e4c3ea3fd

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const __obf_ff8c4c8658a36f7c = "data:image/png;base64,"
const __obf_08d841af39e4cf0c = "data:image/jpeg;base64,"

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_d1b36837c286905e image.Image) (__obf_95290d569d7ce0ee []byte, __obf_335fb0d6019c95ed error) {
	var __obf_6e99783d6b5fc129 bytes.Buffer
	if __obf_335fb0d6019c95ed = png.Encode(&__obf_6e99783d6b5fc129, __obf_d1b36837c286905e); __obf_335fb0d6019c95ed != nil {
		return
	}
	__obf_95290d569d7ce0ee = __obf_6e99783d6b5fc129.Bytes()
	__obf_6e99783d6b5fc129.Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_d1b36837c286905e image.Image, __obf_01f2bb93c2c9908f int) (__obf_95290d569d7ce0ee []byte, __obf_335fb0d6019c95ed error) {
	var __obf_6e99783d6b5fc129 bytes.Buffer
	if __obf_335fb0d6019c95ed = jpeg.Encode(&__obf_6e99783d6b5fc129, __obf_d1b36837c286905e, &jpeg.Options{Quality: __obf_01f2bb93c2c9908f}); __obf_335fb0d6019c95ed != nil {
		return
	}
	__obf_95290d569d7ce0ee = __obf_6e99783d6b5fc129.Bytes()
	__obf_6e99783d6b5fc129.Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_e3b1e50668ab4d43 []byte) (__obf_d1b36837c286905e image.Image, __obf_335fb0d6019c95ed error) {
	var __obf_6e99783d6b5fc129 bytes.Buffer
	__obf_6e99783d6b5fc129.Write(__obf_e3b1e50668ab4d43)
	__obf_d1b36837c286905e, __obf_335fb0d6019c95ed = jpeg.Decode(&__obf_6e99783d6b5fc129)
	__obf_6e99783d6b5fc129.Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_e3b1e50668ab4d43 []byte) (__obf_d1b36837c286905e image.Image, __obf_335fb0d6019c95ed error) {
	var __obf_6e99783d6b5fc129 bytes.Buffer
	__obf_6e99783d6b5fc129.Write(__obf_e3b1e50668ab4d43)
	__obf_d1b36837c286905e, __obf_335fb0d6019c95ed = png.Decode(&__obf_6e99783d6b5fc129)
	__obf_6e99783d6b5fc129.Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_d1b36837c286905e image.Image) (string, error) {
	__obf_0d6afd16e2597f07, __obf_335fb0d6019c95ed := EncodePNGToBase64Data(__obf_d1b36837c286905e)
	if __obf_335fb0d6019c95ed != nil {
		return "", __obf_335fb0d6019c95ed
	}

	return __obf_ff8c4c8658a36f7c + __obf_0d6afd16e2597f07, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_d1b36837c286905e image.Image, __obf_01f2bb93c2c9908f int) (string, error) {
	__obf_0d6afd16e2597f07, __obf_335fb0d6019c95ed := EncodeJPEGToBase64Data(__obf_d1b36837c286905e, __obf_01f2bb93c2c9908f)
	if __obf_335fb0d6019c95ed != nil {
		return "", __obf_335fb0d6019c95ed
	}

	return __obf_08d841af39e4cf0c + __obf_0d6afd16e2597f07, nil
}

// EncodePNGToBase64Data encodes a PNG image to Base64 data (without prefix)
func EncodePNGToBase64Data(__obf_d1b36837c286905e image.Image) (string, error) {
	__obf_afdff3e8ce12a664, __obf_335fb0d6019c95ed := EncodePNGToByte(__obf_d1b36837c286905e)
	if __obf_335fb0d6019c95ed != nil {
		return "", __obf_335fb0d6019c95ed
	}
	return base64.StdEncoding.EncodeToString(__obf_afdff3e8ce12a664), nil
}

// EncodeJPEGToBase64Data encodes a JPEG image to Base64 data (without prefix)
func EncodeJPEGToBase64Data(__obf_d1b36837c286905e image.Image, __obf_01f2bb93c2c9908f int) (string, error) {
	__obf_afdff3e8ce12a664, __obf_335fb0d6019c95ed := EncodeJPEGToByte(__obf_d1b36837c286905e, __obf_01f2bb93c2c9908f)
	if __obf_335fb0d6019c95ed != nil {
		return "", __obf_335fb0d6019c95ed
	}

	return base64.StdEncoding.EncodeToString(__obf_afdff3e8ce12a664), nil
}
