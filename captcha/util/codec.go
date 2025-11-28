package __obf_9aafb22e4ca0e7fb

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_d527642670a6cb84 = "data:image/png;base64,"
	__obf_ad10b6d5eb35b150 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_b19838e146b51f6e image.Image) (__obf_3eaf85ab33963725 []byte, __obf_c8d00ac01db0efb5 error) {
	var __obf_b69e87af3f3ce0fe bytes.Buffer
	if __obf_c8d00ac01db0efb5 = png.Encode(&__obf_b69e87af3f3ce0fe, __obf_b19838e146b51f6e); __obf_c8d00ac01db0efb5 != nil {
		return
	}
	__obf_3eaf85ab33963725 = __obf_b69e87af3f3ce0fe.Bytes()
	__obf_b69e87af3f3ce0fe.Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_b19838e146b51f6e image.Image, __obf_a3ace758376dd900 int) (__obf_3eaf85ab33963725 []byte, __obf_c8d00ac01db0efb5 error) {
	var __obf_b69e87af3f3ce0fe bytes.Buffer
	if __obf_c8d00ac01db0efb5 = jpeg.Encode(&__obf_b69e87af3f3ce0fe, __obf_b19838e146b51f6e, &jpeg.Options{Quality: __obf_a3ace758376dd900}); __obf_c8d00ac01db0efb5 != nil {
		return
	}
	__obf_3eaf85ab33963725 = __obf_b69e87af3f3ce0fe.Bytes()
	__obf_b69e87af3f3ce0fe.Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_af1f8c99779118ac []byte) (__obf_b19838e146b51f6e image.Image, __obf_c8d00ac01db0efb5 error) {
	var __obf_b69e87af3f3ce0fe bytes.Buffer
	__obf_b69e87af3f3ce0fe.Write(__obf_af1f8c99779118ac)
	__obf_b19838e146b51f6e, __obf_c8d00ac01db0efb5 = jpeg.Decode(&__obf_b69e87af3f3ce0fe)
	__obf_b69e87af3f3ce0fe.Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_af1f8c99779118ac []byte) (__obf_b19838e146b51f6e image.Image, __obf_c8d00ac01db0efb5 error) {
	var __obf_b69e87af3f3ce0fe bytes.Buffer
	__obf_b69e87af3f3ce0fe.Write(__obf_af1f8c99779118ac)
	__obf_b19838e146b51f6e, __obf_c8d00ac01db0efb5 = png.Decode(&__obf_b69e87af3f3ce0fe)
	__obf_b69e87af3f3ce0fe.Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_b19838e146b51f6e image.Image) (string, error) {
	__obf_24fee240bdda38fa, __obf_c8d00ac01db0efb5 := EncodePNGToByte(__obf_b19838e146b51f6e)
	if __obf_c8d00ac01db0efb5 != nil {
		return "", __obf_c8d00ac01db0efb5
	}
	__obf_a84aad8c19fa4cb2 := base64.StdEncoding.EncodeToString(__obf_24fee240bdda38fa)
	return __obf_d527642670a6cb84 + __obf_a84aad8c19fa4cb2, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_b19838e146b51f6e image.Image, __obf_a3ace758376dd900 int) (string, error) {
	__obf_24fee240bdda38fa, __obf_c8d00ac01db0efb5 := EncodeJPEGToByte(__obf_b19838e146b51f6e, __obf_a3ace758376dd900)
	if __obf_c8d00ac01db0efb5 != nil {
		return "", __obf_c8d00ac01db0efb5
	}
	__obf_a84aad8c19fa4cb2 := base64.StdEncoding.EncodeToString(__obf_24fee240bdda38fa)
	return __obf_ad10b6d5eb35b150 + __obf_a84aad8c19fa4cb2, nil
}
