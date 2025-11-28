package __obf_44ec487f80286103

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_ab36fc6e223b481b = "data:image/png;base64,"
	__obf_b247eda78baddd0d = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_1392ec6be47e851e image.Image) (__obf_5db072286cb88e49 []byte, __obf_9ad8a31ce70de66b error) {
	var __obf_614c03a4bd4fbf8f bytes.Buffer
	if __obf_9ad8a31ce70de66b = png.Encode(&__obf_614c03a4bd4fbf8f, __obf_1392ec6be47e851e); __obf_9ad8a31ce70de66b != nil {
		return
	}
	__obf_5db072286cb88e49 = __obf_614c03a4bd4fbf8f.Bytes()
	__obf_614c03a4bd4fbf8f.Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_1392ec6be47e851e image.Image, __obf_6979d5f6aef16474 int) (__obf_5db072286cb88e49 []byte, __obf_9ad8a31ce70de66b error) {
	var __obf_614c03a4bd4fbf8f bytes.Buffer
	if __obf_9ad8a31ce70de66b = jpeg.Encode(&__obf_614c03a4bd4fbf8f, __obf_1392ec6be47e851e, &jpeg.Options{Quality: __obf_6979d5f6aef16474}); __obf_9ad8a31ce70de66b != nil {
		return
	}
	__obf_5db072286cb88e49 = __obf_614c03a4bd4fbf8f.Bytes()
	__obf_614c03a4bd4fbf8f.Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_7571c697f623cc50 []byte) (__obf_1392ec6be47e851e image.Image, __obf_9ad8a31ce70de66b error) {
	var __obf_614c03a4bd4fbf8f bytes.Buffer
	__obf_614c03a4bd4fbf8f.Write(__obf_7571c697f623cc50)
	__obf_1392ec6be47e851e, __obf_9ad8a31ce70de66b = jpeg.Decode(&__obf_614c03a4bd4fbf8f)
	__obf_614c03a4bd4fbf8f.Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_7571c697f623cc50 []byte) (__obf_1392ec6be47e851e image.Image, __obf_9ad8a31ce70de66b error) {
	var __obf_614c03a4bd4fbf8f bytes.Buffer
	__obf_614c03a4bd4fbf8f.Write(__obf_7571c697f623cc50)
	__obf_1392ec6be47e851e, __obf_9ad8a31ce70de66b = png.Decode(&__obf_614c03a4bd4fbf8f)
	__obf_614c03a4bd4fbf8f.Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_1392ec6be47e851e image.Image) (string, error) {
	__obf_16ba35642ae15f29, __obf_9ad8a31ce70de66b := EncodePNGToByte(__obf_1392ec6be47e851e)
	if __obf_9ad8a31ce70de66b != nil {
		return "", __obf_9ad8a31ce70de66b
	}
	__obf_5d79e558162731d8 := base64.StdEncoding.EncodeToString(__obf_16ba35642ae15f29)
	return __obf_ab36fc6e223b481b + __obf_5d79e558162731d8, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_1392ec6be47e851e image.Image, __obf_6979d5f6aef16474 int) (string, error) {
	__obf_16ba35642ae15f29, __obf_9ad8a31ce70de66b := EncodeJPEGToByte(__obf_1392ec6be47e851e, __obf_6979d5f6aef16474)
	if __obf_9ad8a31ce70de66b != nil {
		return "", __obf_9ad8a31ce70de66b
	}
	__obf_5d79e558162731d8 := base64.StdEncoding.EncodeToString(__obf_16ba35642ae15f29)
	return __obf_b247eda78baddd0d + __obf_5d79e558162731d8, nil
}
