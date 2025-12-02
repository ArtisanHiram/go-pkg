//go:build !appengine
// +build !appengine

package __obf_a935eb7f548271a4

import (
	"unsafe"
)

// bytesToString converts byte slice to string.
func __obf_ae16b98e7d90edc4(__obf_f2ca794293605b73 []byte) string {
	return *(*string)(unsafe.Pointer(&__obf_f2ca794293605b73))
}

// stringToBytes converts string to byte slice.
func __obf_2204f386aa52f210(__obf_b62c60fba2fd9788 string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{__obf_b62c60fba2fd9788, len(__obf_b62c60fba2fd9788)},
	))
}
