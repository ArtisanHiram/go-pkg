package __obf_3860403e44e18318

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_9b443580f79dcfa4 = "data:image/png;base64,"
	__obf_5be216c509a4d791 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_29c6dafe17296cb0 image.Image) (__obf_376d290df7e12311 []byte, __obf_7834a057dbd32804 error) {
	var __obf_edba0635d7aed20f bytes.Buffer
	if __obf_7834a057dbd32804 = png.Encode(&__obf_edba0635d7aed20f, __obf_29c6dafe17296cb0); __obf_7834a057dbd32804 != nil {
		return
	}
	__obf_376d290df7e12311 = __obf_edba0635d7aed20f.Bytes()
	__obf_edba0635d7aed20f.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_29c6dafe17296cb0 image.Image, __obf_66f4d7b966544de4 int) (__obf_376d290df7e12311 []byte, __obf_7834a057dbd32804 error) {
	var __obf_edba0635d7aed20f bytes.Buffer
	if __obf_7834a057dbd32804 = jpeg.Encode(&__obf_edba0635d7aed20f, __obf_29c6dafe17296cb0, &jpeg.Options{Quality: __obf_66f4d7b966544de4}); __obf_7834a057dbd32804 != nil {
		return
	}
	__obf_376d290df7e12311 = __obf_edba0635d7aed20f.Bytes()
	__obf_edba0635d7aed20f.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_3536db1209733836 []byte) (__obf_29c6dafe17296cb0 image.Image, __obf_7834a057dbd32804 error) {
	var __obf_edba0635d7aed20f bytes.Buffer
	__obf_edba0635d7aed20f.
		Write(__obf_3536db1209733836)
	__obf_29c6dafe17296cb0, __obf_7834a057dbd32804 = jpeg.Decode(&__obf_edba0635d7aed20f)
	__obf_edba0635d7aed20f.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_3536db1209733836 []byte) (__obf_29c6dafe17296cb0 image.Image, __obf_7834a057dbd32804 error) {
	var __obf_edba0635d7aed20f bytes.Buffer
	__obf_edba0635d7aed20f.
		Write(__obf_3536db1209733836)
	__obf_29c6dafe17296cb0, __obf_7834a057dbd32804 = png.Decode(&__obf_edba0635d7aed20f)
	__obf_edba0635d7aed20f.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_29c6dafe17296cb0 image.Image) (string, error) {
	__obf_1903e3c06dafd122, __obf_7834a057dbd32804 := EncodePNGToByte(__obf_29c6dafe17296cb0)
	if __obf_7834a057dbd32804 != nil {
		return "", __obf_7834a057dbd32804
	}
	__obf_61bf45b383fb54c3 := base64.StdEncoding.EncodeToString(__obf_1903e3c06dafd122)
	return __obf_9b443580f79dcfa4 + __obf_61bf45b383fb54c3, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_29c6dafe17296cb0 image.Image, __obf_66f4d7b966544de4 int) (string, error) {
	__obf_1903e3c06dafd122, __obf_7834a057dbd32804 := EncodeJPEGToByte(__obf_29c6dafe17296cb0, __obf_66f4d7b966544de4)
	if __obf_7834a057dbd32804 != nil {
		return "", __obf_7834a057dbd32804
	}
	__obf_61bf45b383fb54c3 := base64.StdEncoding.EncodeToString(__obf_1903e3c06dafd122)
	return __obf_5be216c509a4d791 + __obf_61bf45b383fb54c3, nil
}
