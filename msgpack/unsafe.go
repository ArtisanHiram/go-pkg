//go:build !appengine
// +build !appengine

package __obf_de86cdc8ae98b45b

import (
	"unsafe"
)

// bytesToString converts byte slice to string.
func __obf_16417a69f4ce2476(__obf_a7fd7acf2bd4435f []byte) string {
	return *(*string)(unsafe.Pointer(&__obf_a7fd7acf2bd4435f))
}

// stringToBytes converts string to byte slice.
func __obf_26e7c5987cb4f459(__obf_a93d004caac96500 string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{__obf_a93d004caac96500, len(__obf_a93d004caac96500)},
	))
}
