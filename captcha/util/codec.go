package __obf_724f8a3b7b400b23

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_4744dcfec787ae39 = "data:image/png;base64,"
	__obf_481fffbcc63cdf59 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_4aeb51bb135c8905 image.Image) (__obf_832a5d468c3a53cb []byte, __obf_043e1e14636249eb error) {
	var __obf_e5c25443cd1b87cf bytes.Buffer
	if __obf_043e1e14636249eb = png.Encode(&__obf_e5c25443cd1b87cf, __obf_4aeb51bb135c8905); __obf_043e1e14636249eb != nil {
		return
	}
	__obf_832a5d468c3a53cb = __obf_e5c25443cd1b87cf.Bytes()
	__obf_e5c25443cd1b87cf.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_4aeb51bb135c8905 image.Image, __obf_32b5c940da93f1e7 int) (__obf_832a5d468c3a53cb []byte, __obf_043e1e14636249eb error) {
	var __obf_e5c25443cd1b87cf bytes.Buffer
	if __obf_043e1e14636249eb = jpeg.Encode(&__obf_e5c25443cd1b87cf, __obf_4aeb51bb135c8905, &jpeg.Options{Quality: __obf_32b5c940da93f1e7}); __obf_043e1e14636249eb != nil {
		return
	}
	__obf_832a5d468c3a53cb = __obf_e5c25443cd1b87cf.Bytes()
	__obf_e5c25443cd1b87cf.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_a0b806d33e63dadd []byte) (__obf_4aeb51bb135c8905 image.Image, __obf_043e1e14636249eb error) {
	var __obf_e5c25443cd1b87cf bytes.Buffer
	__obf_e5c25443cd1b87cf.
		Write(__obf_a0b806d33e63dadd)
	__obf_4aeb51bb135c8905, __obf_043e1e14636249eb = jpeg.Decode(&__obf_e5c25443cd1b87cf)
	__obf_e5c25443cd1b87cf.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_a0b806d33e63dadd []byte) (__obf_4aeb51bb135c8905 image.Image, __obf_043e1e14636249eb error) {
	var __obf_e5c25443cd1b87cf bytes.Buffer
	__obf_e5c25443cd1b87cf.
		Write(__obf_a0b806d33e63dadd)
	__obf_4aeb51bb135c8905, __obf_043e1e14636249eb = png.Decode(&__obf_e5c25443cd1b87cf)
	__obf_e5c25443cd1b87cf.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_4aeb51bb135c8905 image.Image) (string, error) {
	__obf_2f105fb49def23ea, __obf_043e1e14636249eb := EncodePNGToByte(__obf_4aeb51bb135c8905)
	if __obf_043e1e14636249eb != nil {
		return "", __obf_043e1e14636249eb
	}
	__obf_1205620436b1147c := base64.StdEncoding.EncodeToString(__obf_2f105fb49def23ea)
	return __obf_4744dcfec787ae39 + __obf_1205620436b1147c, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_4aeb51bb135c8905 image.Image, __obf_32b5c940da93f1e7 int) (string, error) {
	__obf_2f105fb49def23ea, __obf_043e1e14636249eb := EncodeJPEGToByte(__obf_4aeb51bb135c8905, __obf_32b5c940da93f1e7)
	if __obf_043e1e14636249eb != nil {
		return "", __obf_043e1e14636249eb
	}
	__obf_1205620436b1147c := base64.StdEncoding.EncodeToString(__obf_2f105fb49def23ea)
	return __obf_481fffbcc63cdf59 + __obf_1205620436b1147c, nil
}
