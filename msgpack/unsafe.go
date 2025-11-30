//go:build !appengine
// +build !appengine

package __obf_3e0c303610a19bd4

import (
	"unsafe"
)

// bytesToString converts byte slice to string.
func __obf_f682cfd58bdd5de3(__obf_11bcc66cde095c11 []byte) string {
	return *(*string)(unsafe.Pointer(&__obf_11bcc66cde095c11))
}

// stringToBytes converts string to byte slice.
func __obf_31018e102d310148(__obf_61027e0491b6dd3d string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{__obf_61027e0491b6dd3d, len(__obf_61027e0491b6dd3d)},
	))
}
