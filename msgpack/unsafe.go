//go:build !appengine
// +build !appengine

package __obf_3edaa49e853afa16

import (
	"unsafe"
)

// bytesToString converts byte slice to string.
func __obf_39ecaa2d03d54c22(__obf_9b4a5a04bdad2347 []byte) string {
	return *(*string)(unsafe.Pointer(&__obf_9b4a5a04bdad2347))
}

// stringToBytes converts string to byte slice.
func __obf_f0406d601699c6f7(__obf_6827ff1b59d7ccec string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{__obf_6827ff1b59d7ccec, len(__obf_6827ff1b59d7ccec)},
	))
}
