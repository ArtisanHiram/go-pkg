package __obf_3a8eca76d39e021a

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_94bb1bb952ce2654 = "data:image/png;base64,"
	__obf_c5f0cfce7594d9c3 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_d5358d2a73ea01f6 image.Image) (__obf_1b4f28736c4bcede []byte, __obf_d188595df3b06d83 error) {
	var __obf_aa2587118f70f5b7 bytes.Buffer
	if __obf_d188595df3b06d83 = png.Encode(&__obf_aa2587118f70f5b7, __obf_d5358d2a73ea01f6); __obf_d188595df3b06d83 != nil {
		return
	}
	__obf_1b4f28736c4bcede = __obf_aa2587118f70f5b7.Bytes()
	__obf_aa2587118f70f5b7.Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_d5358d2a73ea01f6 image.Image, __obf_657e6e42f7cf2381 int) (__obf_1b4f28736c4bcede []byte, __obf_d188595df3b06d83 error) {
	var __obf_aa2587118f70f5b7 bytes.Buffer
	if __obf_d188595df3b06d83 = jpeg.Encode(&__obf_aa2587118f70f5b7, __obf_d5358d2a73ea01f6, &jpeg.Options{Quality: __obf_657e6e42f7cf2381}); __obf_d188595df3b06d83 != nil {
		return
	}
	__obf_1b4f28736c4bcede = __obf_aa2587118f70f5b7.Bytes()
	__obf_aa2587118f70f5b7.Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_c689f43cdba733b4 []byte) (__obf_d5358d2a73ea01f6 image.Image, __obf_d188595df3b06d83 error) {
	var __obf_aa2587118f70f5b7 bytes.Buffer
	__obf_aa2587118f70f5b7.Write(__obf_c689f43cdba733b4)
	__obf_d5358d2a73ea01f6, __obf_d188595df3b06d83 = jpeg.Decode(&__obf_aa2587118f70f5b7)
	__obf_aa2587118f70f5b7.Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_c689f43cdba733b4 []byte) (__obf_d5358d2a73ea01f6 image.Image, __obf_d188595df3b06d83 error) {
	var __obf_aa2587118f70f5b7 bytes.Buffer
	__obf_aa2587118f70f5b7.Write(__obf_c689f43cdba733b4)
	__obf_d5358d2a73ea01f6, __obf_d188595df3b06d83 = png.Decode(&__obf_aa2587118f70f5b7)
	__obf_aa2587118f70f5b7.Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_d5358d2a73ea01f6 image.Image) (string, error) {
	__obf_ba4465d483695888, __obf_d188595df3b06d83 := EncodePNGToByte(__obf_d5358d2a73ea01f6)
	if __obf_d188595df3b06d83 != nil {
		return "", __obf_d188595df3b06d83
	}
	__obf_fd0d65f07c08712a := base64.StdEncoding.EncodeToString(__obf_ba4465d483695888)
	return __obf_94bb1bb952ce2654 + __obf_fd0d65f07c08712a, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_d5358d2a73ea01f6 image.Image, __obf_657e6e42f7cf2381 int) (string, error) {
	__obf_ba4465d483695888, __obf_d188595df3b06d83 := EncodeJPEGToByte(__obf_d5358d2a73ea01f6, __obf_657e6e42f7cf2381)
	if __obf_d188595df3b06d83 != nil {
		return "", __obf_d188595df3b06d83
	}
	__obf_fd0d65f07c08712a := base64.StdEncoding.EncodeToString(__obf_ba4465d483695888)
	return __obf_c5f0cfce7594d9c3 + __obf_fd0d65f07c08712a, nil
}
