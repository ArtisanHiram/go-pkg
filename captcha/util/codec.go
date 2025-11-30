package __obf_611d3abd867098c0

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	__obf_e5059269513abf4a = "data:image/png;base64,"
	__obf_05b2a0c110348747 = "data:image/jpeg;base64,"
)

// EncodePNGToByte encodes a PNG image to a byte array
func EncodePNGToByte(__obf_6922282326d724ae image.Image) (__obf_44aa4aac478bf881 []byte, __obf_e9ef0dd78ff22cdd error) {
	var __obf_3bb978316365351e bytes.Buffer
	if __obf_e9ef0dd78ff22cdd = png.Encode(&__obf_3bb978316365351e, __obf_6922282326d724ae); __obf_e9ef0dd78ff22cdd != nil {
		return
	}
	__obf_44aa4aac478bf881 = __obf_3bb978316365351e.Bytes()
	__obf_3bb978316365351e.
		Reset()
	return
}

// EncodeJPEGToByte encodes a JPEG image to a byte array
func EncodeJPEGToByte(__obf_6922282326d724ae image.Image, __obf_d4c6fbd36d49b11e int) (__obf_44aa4aac478bf881 []byte, __obf_e9ef0dd78ff22cdd error) {
	var __obf_3bb978316365351e bytes.Buffer
	if __obf_e9ef0dd78ff22cdd = jpeg.Encode(&__obf_3bb978316365351e, __obf_6922282326d724ae, &jpeg.Options{Quality: __obf_d4c6fbd36d49b11e}); __obf_e9ef0dd78ff22cdd != nil {
		return
	}
	__obf_44aa4aac478bf881 = __obf_3bb978316365351e.Bytes()
	__obf_3bb978316365351e.
		Reset()
	return
}

// DecodeByteToJpeg decodes a byte array to a JPEG image
func DecodeByteToJpeg(__obf_39012e53e716527b []byte) (__obf_6922282326d724ae image.Image, __obf_e9ef0dd78ff22cdd error) {
	var __obf_3bb978316365351e bytes.Buffer
	__obf_3bb978316365351e.
		Write(__obf_39012e53e716527b)
	__obf_6922282326d724ae, __obf_e9ef0dd78ff22cdd = jpeg.Decode(&__obf_3bb978316365351e)
	__obf_3bb978316365351e.
		Reset()
	return
}

// DecodeByteToPng decodes a byte array to a PNG image
func DecodeByteToPng(__obf_39012e53e716527b []byte) (__obf_6922282326d724ae image.Image, __obf_e9ef0dd78ff22cdd error) {
	var __obf_3bb978316365351e bytes.Buffer
	__obf_3bb978316365351e.
		Write(__obf_39012e53e716527b)
	__obf_6922282326d724ae, __obf_e9ef0dd78ff22cdd = png.Decode(&__obf_3bb978316365351e)
	__obf_3bb978316365351e.
		Reset()
	return
}

// EncodePNGToBase64 encodes a PNG image to a Base64 string
func EncodePNGToBase64(__obf_6922282326d724ae image.Image) (string, error) {
	__obf_f5cbaa92f9ee0606, __obf_e9ef0dd78ff22cdd := EncodePNGToByte(__obf_6922282326d724ae)
	if __obf_e9ef0dd78ff22cdd != nil {
		return "", __obf_e9ef0dd78ff22cdd
	}
	__obf_a8c4402c7870814a := base64.StdEncoding.EncodeToString(__obf_f5cbaa92f9ee0606)
	return __obf_e5059269513abf4a + __obf_a8c4402c7870814a, nil
}

// EncodeJPEGToBase64 encodes a JPEG image to a Base64 string
func EncodeJPEGToBase64(__obf_6922282326d724ae image.Image, __obf_d4c6fbd36d49b11e int) (string, error) {
	__obf_f5cbaa92f9ee0606, __obf_e9ef0dd78ff22cdd := EncodeJPEGToByte(__obf_6922282326d724ae, __obf_d4c6fbd36d49b11e)
	if __obf_e9ef0dd78ff22cdd != nil {
		return "", __obf_e9ef0dd78ff22cdd
	}
	__obf_a8c4402c7870814a := base64.StdEncoding.EncodeToString(__obf_f5cbaa92f9ee0606)
	return __obf_05b2a0c110348747 + __obf_a8c4402c7870814a, nil
}
