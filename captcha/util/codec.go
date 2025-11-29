package __obf_1c2965545693a835

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_e86f261f51f8f803 = "data:image/png;base64,"
	__obf_bfcfed7c5b0e65a6 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_887e29ba67702921 image.Image) (__obf_44394f77261288e1 []byte, __obf_675a3caf504b632f error) {
	var __obf_2e288198e7e6783c bytes.Buffer
	if __obf_675a3caf504b632f = png.Encode(&__obf_2e288198e7e6783c, __obf_887e29ba67702921); __obf_675a3caf504b632f != nil {
		return
	}
	__obf_44394f77261288e1 = __obf_2e288198e7e6783c.Bytes()
	__obf_2e288198e7e6783c.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_887e29ba67702921 image.Image, __obf_c259aa6af9c2c62c int) (__obf_44394f77261288e1 []byte, __obf_675a3caf504b632f error) {
	var __obf_2e288198e7e6783c bytes.Buffer
	if __obf_675a3caf504b632f = jpeg.Encode(&__obf_2e288198e7e6783c, __obf_887e29ba67702921, &jpeg.Options{Quality: __obf_c259aa6af9c2c62c}); __obf_675a3caf504b632f != nil {
		return
	}
	__obf_44394f77261288e1 = __obf_2e288198e7e6783c.Bytes()
	__obf_2e288198e7e6783c.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_d935b3348e7c0cef []byte) (__obf_887e29ba67702921 image.Image, __obf_675a3caf504b632f error) {
	var __obf_2e288198e7e6783c bytes.Buffer
	__obf_2e288198e7e6783c.
		Write(__obf_d935b3348e7c0cef)
	__obf_887e29ba67702921, __obf_675a3caf504b632f = jpeg.Decode(&__obf_2e288198e7e6783c)
	__obf_2e288198e7e6783c.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_d935b3348e7c0cef []byte) (__obf_887e29ba67702921 image.Image, __obf_675a3caf504b632f error) {
	var __obf_2e288198e7e6783c bytes.Buffer
	__obf_2e288198e7e6783c.
		Write(__obf_d935b3348e7c0cef)
	__obf_887e29ba67702921, __obf_675a3caf504b632f = png.Decode(&__obf_2e288198e7e6783c)
	__obf_2e288198e7e6783c.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_887e29ba67702921 image.Image) (string, error) {
	__obf_2ad7ac23da962381, __obf_675a3caf504b632f := EncodePNGToByte(__obf_887e29ba67702921)
	if __obf_675a3caf504b632f != nil {
		return "", __obf_675a3caf504b632f
	}
	__obf_568f4e84b9aebd1b := base64.StdEncoding.EncodeToString(__obf_2ad7ac23da962381)
	return __obf_e86f261f51f8f803 + __obf_568f4e84b9aebd1b, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_887e29ba67702921 image.Image, __obf_c259aa6af9c2c62c int) (string, error) {
	__obf_2ad7ac23da962381, __obf_675a3caf504b632f := EncodeJPEGToByte(__obf_887e29ba67702921, __obf_c259aa6af9c2c62c)
	if __obf_675a3caf504b632f != nil {
		return "", __obf_675a3caf504b632f
	}
	__obf_568f4e84b9aebd1b := base64.StdEncoding.EncodeToString(__obf_2ad7ac23da962381)
	return __obf_bfcfed7c5b0e65a6 + __obf_568f4e84b9aebd1b, nil
}
