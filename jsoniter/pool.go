package __obf_5b802ce8d9ba56d6

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_5d3dff745e2b805b []byte) *Iterator
	ReturnIterator(__obf_67008a6a9e5ba828 *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_3dbedd483fe5a38a io.Writer) *Stream
	ReturnStream(__obf_00fc491caa967a74 *Stream)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) BorrowStream(__obf_3dbedd483fe5a38a io.Writer) *Stream {
	__obf_00fc491caa967a74 := __obf_dca106e1cdf85392.__obf_a774d126f4306dfb.Get().(*Stream)
	__obf_00fc491caa967a74.
		Reset(__obf_3dbedd483fe5a38a)
	return __obf_00fc491caa967a74
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) ReturnStream(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.__obf_3d1a004850652c9c = nil
	__obf_00fc491caa967a74.
		Error = nil
	__obf_00fc491caa967a74.
		Attachment = nil
	__obf_dca106e1cdf85392.__obf_a774d126f4306dfb.
		Put(__obf_00fc491caa967a74)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) BorrowIterator(__obf_5d3dff745e2b805b []byte) *Iterator {
	__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.__obf_f74e8f2afdb9df78.Get().(*Iterator)
	__obf_67008a6a9e5ba828.
		ResetBytes(__obf_5d3dff745e2b805b)
	return __obf_67008a6a9e5ba828
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) ReturnIterator(__obf_67008a6a9e5ba828 *Iterator) {
	__obf_67008a6a9e5ba828.
		Error = nil
	__obf_67008a6a9e5ba828.
		Attachment = nil
	__obf_dca106e1cdf85392.__obf_f74e8f2afdb9df78.
		Put(__obf_67008a6a9e5ba828)
}
